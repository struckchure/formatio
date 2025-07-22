package services_test

import (
	"testing"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"pkg.formatio/lib"
	dao_mocks "pkg.formatio/mocks/dao"
	lib_mocks "pkg.formatio/mocks/lib"
	"pkg.formatio/prisma/db"
	"pkg.formatio/services"
	"pkg.formatio/types"
)

type UserServiceSuite struct {
	suite.Suite

	userService       services.IUserService
	mockHasher        *lib_mocks.MockIHasher
	mockAuth0         *lib_mocks.MockIAuth0
	mockJwt           *lib_mocks.MockIJwt
	mockConnectionDao *dao_mocks.MockISocialConnectionDao
	mockUserDao       *dao_mocks.MockIUserDao
}

func (s *UserServiceSuite) SetupTest() {
	s.mockHasher = new(lib_mocks.MockIHasher)
	s.mockAuth0 = new(lib_mocks.MockIAuth0)
	s.mockJwt = new(lib_mocks.MockIJwt)
	s.mockConnectionDao = new(dao_mocks.MockISocialConnectionDao)
	s.mockUserDao = new(dao_mocks.MockIUserDao)

	s.userService = services.NewUserService(s.mockHasher, s.mockAuth0, s.mockJwt, s.mockConnectionDao, s.mockUserDao)
}

func (s *UserServiceSuite) TearDownTest() {
	s.mockHasher.AssertExpectations(s.T())
	s.mockAuth0.AssertExpectations(s.T())
	s.mockJwt.AssertExpectations(s.T())
	s.mockConnectionDao.AssertExpectations(s.T())
	s.mockUserDao.AssertExpectations(s.T())
}

func (s *UserServiceSuite) TestLoginUser_SuccessfulLogin() {
	email := "test@example.com"
	password := "password123"
	hashedPassword := "hashedPassword123"
	userId := "user123"

	mockUser := &db.UserModel{
		InnerUser: db.InnerUser{
			ID:       userId,
			Email:    &email,
			Password: &hashedPassword,
		},
	}

	s.mockUserDao.On("GetUser", types.GetUserArgs{Email: &email}).Return(mockUser, nil)
	s.mockHasher.On("PasswordIsCorrect", hashedPassword, password).Return(true)
	s.mockJwt.On("GenerateJWT", userId).Return(&lib.AuthTokens{
		AccessToken:  "access_token",
		RefreshToken: "refresh_token",
	}, nil)

	result, err := s.userService.LoginUser(types.LoginUserArgs{
		Email:    email,
		Password: password,
	})

	s.NoError(err)
	s.NotNil(result)
	s.Equal("access_token", result.Tokens.AccessToken)
	s.Equal("refresh_token", result.Tokens.RefreshToken)
}

func (s *UserServiceSuite) TestLoginUser_InvalidCredentials_UserNotFound() {
	email := "nonexistent@example.com"
	password := "password123"

	s.mockUserDao.On("GetUser", types.GetUserArgs{Email: &email}).Return(nil, fiber.NewError(fiber.StatusNotFound, "user not found"))

	result, err := s.userService.LoginUser(types.LoginUserArgs{
		Email:    email,
		Password: password,
	})

	s.Error(err)
	s.Nil(result)
	s.Equal(fiber.StatusBadRequest, err.(*fiber.Error).Code)
	s.Equal("invalid credentials", err.(*fiber.Error).Message)
}

func (s *UserServiceSuite) TestLoginUser_InvalidCredentials_IncorrectPassword() {
	email := "test@example.com"
	password := "wrongpassword"
	hashedPassword := "hashedPassword123"
	userId := "user123"

	mockUser := &db.UserModel{
		InnerUser: db.InnerUser{
			ID:       userId,
			Password: &hashedPassword,
		},
	}

	s.mockUserDao.On("GetUser", types.GetUserArgs{Email: &email}).Return(mockUser, nil)
	s.mockHasher.On("PasswordIsCorrect", hashedPassword, password).Return(false)

	result, err := s.userService.LoginUser(types.LoginUserArgs{
		Email:    email,
		Password: password,
	})

	s.Error(err)
	s.Nil(result)
	s.Equal(fiber.StatusBadRequest, err.(*fiber.Error).Code)
	s.Equal("invalid credentials", err.(*fiber.Error).Message)
}

func (s *UserServiceSuite) TestRegisterUser_SuccessfulRegistration() {
	email := "newuser@example.com"
	password := "password123"
	hashedPassword := "hashedPassword123"
	userId := "user123"
	firstName := "John"
	lastName := "Doe"

	mockUser := &db.UserModel{
		InnerUser: db.InnerUser{
			ID:       userId,
			Email:    &email,
			Password: &hashedPassword,
		},
	}

	s.mockHasher.On("HashPassword", password).Return(hashedPassword)
	s.mockUserDao.On("CreateUser", mock.Anything).Return(mockUser, nil)
	s.mockJwt.On("GenerateJWT", userId).Return(&lib.AuthTokens{
		AccessToken:  "access_token",
		RefreshToken: "refresh_token",
	}, nil)

	result, err := s.userService.RegisterUser(types.RegisterUserArgs{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  password,
	})

	s.NoError(err)
	s.NotNil(result)
	s.Equal("access_token", result.Tokens.AccessToken)
	s.Equal("refresh_token", result.Tokens.RefreshToken)
}

func (s *UserServiceSuite) TestGetProfileUser_SuccessfulGetProfile() {
	userId := "user123"
	email := "test@example.com"

	mockUser := &db.UserModel{
		InnerUser: db.InnerUser{
			ID:    userId,
			Email: &email,
		},
	}

	s.mockUserDao.On("GetUser", types.GetUserArgs{ID: &userId}).Return(mockUser, nil)

	result, err := s.userService.GetProfileUser(types.GetUserArgs{ID: &userId})

	s.NoError(err)
	s.NotNil(result)
	s.Equal(userId, result.ID)
}

func (s *UserServiceSuite) TestUpdateUser_SuccessfulUpdate() {
	userId := "user123"
	email := "newemail@example.com"
	firstName := "NewFirstName"
	lastName := "NewLastName"
	password := "newpassword"
	hashedPassword := "hashedPassword123"

	mockUser := &db.UserModel{
		InnerUser: db.InnerUser{
			ID:    userId,
			Email: &email,
		},
	}

	s.mockHasher.On("HashPassword", password).Return(hashedPassword)
	s.mockUserDao.On("UpdateUser", mock.Anything).Return(mockUser, nil)

	result, err := s.userService.UpdateUser(types.UpdateUserArgs{
		ID:        userId,
		FirstName: &firstName,
		LastName:  &lastName,
		Email:     &email,
		Password:  &password,
	})

	resultEmail, _ := result.Email()

	s.NoError(err)
	s.NotNil(result)
	s.Equal(userId, result.ID)
	s.Equal(email, resultEmail)
}

func (s *UserServiceSuite) TestRefreshAccessToken_SuccessfulTokenRefresh() {
	refreshToken := "valid_refresh_token"
	userId := "user123"

	claims := &jwt.Token{
		Claims: jwt.MapClaims{
			"sub": userId,
		},
	}

	s.mockJwt.On("VerifyJWT", refreshToken, lib.REFRESH_TOKEN_TYPE).Return(claims, nil)
	s.mockJwt.On("GenerateJWT", userId).Return(&lib.AuthTokens{
		AccessToken:  "new_access_token",
		RefreshToken: "new_refresh_token",
	}, nil)

	result, err := s.userService.RefreshAccessToken(types.RefreshAccessTokenArgs{RefreshToken: refreshToken})

	s.NoError(err)
	s.NotNil(result)
	s.Equal("new_access_token", result.Tokens.AccessToken)
	s.Equal("new_refresh_token", result.Tokens.RefreshToken)
}

func (s *UserServiceSuite) TestAuthSocialConnection_SuccessfulSocialConnection() {
	userId := "user123"
	email := "socialuser@example.com"

	claims := &lib.Auth0TokenClaims{
		RegisteredClaims: struct {
			Iss string   "json:\"iss\""
			Sub string   "json:\"sub\""
			Aud []string "json:\"aud\""
			Exp int64    "json:\"exp\""
			Iat int64    "json:\"iat\""
		}{
			Sub: "social|user",
		},
		CustomClaims: lib.CustomClaims{
			Auth0UserInfo: lib.Auth0UserInfo{
				GivenName:  "John",
				FamilyName: "Doe",
				Email:      email,
			},
		},
	}

	s.mockAuth0.On("GetTokenClaims", mock.Anything).Return(claims, nil)
	s.mockConnectionDao.On("ListConnections", mock.Anything).Return([]db.SocialConnectionModel{}, nil)
	s.mockUserDao.On("GetUser", types.GetUserArgs{Email: &email}).Return(nil, nil)
	s.mockUserDao.On("CreateUser", mock.Anything).Return(&db.UserModel{InnerUser: db.InnerUser{ID: userId}}, nil)
	s.mockConnectionDao.On("CreateConnection", mock.Anything).Return(&db.SocialConnectionModel{}, nil)
	s.mockJwt.On("GenerateJWT", userId).Return(&lib.AuthTokens{
		AccessToken:  "access_token",
		RefreshToken: "refresh_token",
	}, nil)

	result, err := s.userService.AuthSocialConnection(types.Auth0UserArgs{
		Token: "social_token",
	})

	s.NoError(err)
	s.NotNil(result)
	s.Equal("access_token", result.Tokens.AccessToken)
	s.Equal("refresh_token", result.Tokens.RefreshToken)
}

func TestUserService(t *testing.T) {
	suite.Run(t, new(UserServiceSuite))
}
