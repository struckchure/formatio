package services

import (
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/samber/lo"

	"pkg.formatio/dao"
	"pkg.formatio/lib"
	"pkg.formatio/prisma/db"
	"pkg.formatio/types"
)

type IUserService interface {
	RegisterUser(types.RegisterUserArgs) (*types.RegisterUserResult, error)
	LoginUser(types.LoginUserArgs) (*types.LoginUserResult, error)
	RefreshAccessToken(types.RefreshAccessTokenArgs) (*types.LoginUserResult, error)

	GetProfileUser(types.GetUserArgs) (*db.UserModel, error)
	UpdateUser(types.UpdateUserArgs) (*db.UserModel, error)

	AuthSocialConnection(types.Auth0UserArgs) (*types.LoginUserResult, error)
}

type UserService struct {
	hasher lib.IHasher
	auth0  lib.IAuth0
	jwt    lib.IJwt

	connectionDAO dao.ISocialConnectionDao
	userDAO       dao.IUserDao
}

func (u *UserService) GetProfileUser(args types.GetUserArgs) (*db.UserModel, error) {
	user, err := u.userDAO.GetUser(types.GetUserArgs{ID: args.ID})

	userWithoutPassword := lib.RemoveField(user, lib.RemoveFieldOptionFunc("password"))

	return userWithoutPassword, err
}

func (u *UserService) LoginUser(args types.LoginUserArgs) (*types.LoginUserResult, error) {
	user, err := u.userDAO.GetUser(types.GetUserArgs{Email: &args.Email})
	if err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, "invalid credentials")
	}

	userPassword, _ := user.Password()
	passwordIsNotCorrect := !u.hasher.PasswordIsCorrect(userPassword, args.Password)
	if passwordIsNotCorrect {
		return nil, fiber.NewError(fiber.StatusBadRequest, "invalid credentials")
	}

	tokens, err := u.jwt.GenerateJWT(user.ID)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return &types.LoginUserResult{Tokens: *tokens}, nil
}

func (u *UserService) RefreshAccessToken(args types.RefreshAccessTokenArgs) (*types.LoginUserResult, error) {
	claims, err := u.jwt.VerifyJWT(args.RefreshToken, lib.REFRESH_TOKEN_TYPE)
	if err != nil {
		return nil, err
	}

	userId, err := claims.Claims.GetSubject()
	if err != nil {
		return nil, err
	}

	tokens, err := u.jwt.GenerateJWT(userId)
	if err != nil {
		return nil, err
	}

	return &types.LoginUserResult{Tokens: *tokens}, nil
}

func (u *UserService) RegisterUser(args types.RegisterUserArgs) (*types.RegisterUserResult, error) {
	user, err := u.userDAO.CreateUser(types.CreateUserArgs{
		FirstName: args.FirstName,
		LastName:  args.LastName,
		Email:     args.Email,
	})
	if args.Password != "" {
		args.Password = u.hasher.HashPassword(args.Password)
	}

	if err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	tokens, err := u.jwt.GenerateJWT(user.ID)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	user = lib.RemoveField(user, lib.RemoveFieldOptionFunc("password"))

	return &types.RegisterUserResult{UserModel: *user, Tokens: *tokens}, nil
}

func (u *UserService) UpdateUser(args types.UpdateUserArgs) (*db.UserModel, error) {
	if len(*args.Password) > 0 {
		args.Password = lo.ToPtr(u.hasher.HashPassword(*args.Password))
	}

	user, err := u.userDAO.UpdateUser(types.UpdateUserArgs{
		ID:        args.ID,
		FirstName: args.FirstName,
		LastName:  args.LastName,
		Email:     args.Email,
		Password:  args.Password,
	})
	if err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	user = lib.RemoveField(user, lib.RemoveFieldOptionFunc("password"))

	return user, nil
}

func getFullName(claims lib.Auth0TokenClaims) (firstName string, lastName string) {
	if claims.CustomClaims.GivenName != "" {
		firstName = claims.CustomClaims.GivenName
	}

	if claims.CustomClaims.FamilyName != "" {
		lastName = claims.CustomClaims.FamilyName
	}

	if claims.CustomClaims.Name != "" && firstName == "" || lastName == "" {
		firstName = strings.Split(claims.CustomClaims.Name, " ")[0]
		lastName = strings.Join(strings.Split(claims.CustomClaims.Name, " ")[1:], " ")
	}

	return firstName, lastName
}

func (u *UserService) AuthSocialConnection(args types.Auth0UserArgs) (*types.LoginUserResult, error) {
	claims, err := u.auth0.GetTokenClaims(args.Token)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	firstName, lastName := getFullName(*claims)

	var userId string
	var tokens *lib.AuthTokens

	connections, _ := u.connectionDAO.ListConnections(
		types.ListConnectionsArgs{
			UserId:         claims.RegisteredClaims.Sub,
			ConnectionType: strings.Split(claims.RegisteredClaims.Sub, "|")[0],
		})
	connectionExists := len(connections) > 0
	if connectionExists {
		userId = connections[0].UserID
	} else {
		user, _ := u.userDAO.GetUser(types.GetUserArgs{Email: &claims.CustomClaims.Email})
		if user == nil {
			newUser, err := u.RegisterUser(types.RegisterUserArgs{
				FirstName: firstName,
				LastName:  lastName,
				Email:     claims.CustomClaims.Email,
			})
			if err != nil {
				return nil, fiber.NewError(fiber.StatusBadRequest, err.Error())
			}

			userId = newUser.ID
		} else {
			userId = user.ID
		}

		_, err = u.connectionDAO.CreateConnection(types.CreateConnectionArgs{
			UserId:         userId,
			ConnectionId:   strings.Split(claims.RegisteredClaims.Sub, "|")[0],
			ConnectionType: strings.Split(claims.RegisteredClaims.Sub, "|")[1],
		})

		if err != nil {
			return nil, fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
	}

	tokens, err = u.jwt.GenerateJWT(userId)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return &types.LoginUserResult{Tokens: *tokens}, nil
}

func NewUserService(
	hasher lib.IHasher,
	auth0 lib.IAuth0,
	jwt lib.IJwt,

	connectionDAO dao.ISocialConnectionDao,
	userDAO dao.IUserDao,
) IUserService {
	return &UserService{
		hasher: hasher,
		auth0:  auth0,
		jwt:    jwt,

		connectionDAO: connectionDAO,
		userDAO:       userDAO,
	}
}
