package handlers

import (
	"github.com/gofiber/fiber/v3"

	"pkg.formatio/lib"
	"pkg.formatio/prisma/db"
	"pkg.formatio/services"
	"pkg.formatio/types"
)

type IUserHandler interface {
	RegisterUser(fiber.Ctx) error
	LoginUser(fiber.Ctx) error
	RefreshAccessToken(fiber.Ctx) error

	GetProfileUser(fiber.Ctx) error
	UpdateProfileUser(fiber.Ctx) error

	AuthSocialConnection(fiber.Ctx) error
}

type UserHandler struct {
	userService services.IUserService
}

// LoginUser godoc
//
// @ID			loginUser
// @Tags    auth
// @Accept  json
// @Produce json
// @Param		args					body			types.LoginUserArgs		true "Login User"
// @Success 200						{object}	types.LoginUserResult
// @Router  /auth/login/ 	[post]
func (h *UserHandler) LoginUser(c fiber.Ctx) error {
	args := types.LoginUserArgs{}

	if err := c.Bind().JSON(&args); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	user, err := h.userService.LoginUser(args)
	if err != nil {
		return lib.TranslateHandlerError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

// RegisterUser godoc
//
// @ID			registerUser
// @Tags    auth
// @Accept  json
// @Produce json
// @Param		args						body			types.RegisterUserArgs	true "Register User"
// @Success 201							{object}	types.RegisterUserResult
// @Router  /auth/register/ [post]
func (h *UserHandler) RegisterUser(c fiber.Ctx) error {
	args := types.RegisterUserArgs{}

	if err := c.Bind().JSON(&args); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	user, err := h.userService.RegisterUser(args)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

// RefreshAccessToken godoc
//
// @ID			refreshAccessToken
// @Tags    auth
// @Accept  json
// @Produce json
// @Param		args												body			types.RefreshAccessTokenArgs	true "Refresh Access Token"
// @Success 200													{object}	types.LoginUserResult
// @Router  /auth/refresh-access-token/ [post]
func (h *UserHandler) RefreshAccessToken(c fiber.Ctx) error {
	args := types.RefreshAccessTokenArgs{}

	if err := c.Bind().JSON(&args); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	tokens, err := h.userService.RefreshAccessToken(args)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(tokens)
}

// AuthSocialConnection godoc
//
// @ID			authSocialConnection
// @Tags    auth
// @Accept  json
// @Produce json
// @Param		args											body			types.Auth0UserArgs	true "Auth0 User"
// @Success 200												{object}	types.LoginUserResult
// @Router  /auth/social-connection/ 	[post]
func (u *UserHandler) AuthSocialConnection(c fiber.Ctx) error {
	args := types.Auth0UserArgs{}

	if err := c.Bind().JSON(&args); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	user, err := u.userService.AuthSocialConnection(args)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(&user)
}

// GetProfileUser godoc
//
// @ID			getProfileUser
// @Tags    user
// @Accept  json
// @Produce json
// @Success 200							{object}	db.UserModel
// @Router  /user/profile/ 	[get]
func (*UserHandler) GetProfileUser(c fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Locals[*db.UserModel](c, "user"))
}

// UpdateProfileUser godoc
//
// @ID			updateProfileUser
// @Tags    user
// @Accept  json
// @Produce json
// @Param		args						body			types.UpdateUserArgs	true "Update User"
// @Success 202							{object}	db.UserModel
// @Router  /user/profile/ 	[patch]
func (h *UserHandler) UpdateProfileUser(c fiber.Ctx) error {
	args := types.UpdateUserArgs{ID: fiber.Locals[string](c, "userId")}

	if err := c.Bind().JSON(&args); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	user, err := h.userService.UpdateUser(args)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.Status(fiber.StatusAccepted).JSON(user)
}

func NewUserHandler(userService services.IUserService) IUserHandler {
	return &UserHandler{userService: userService}
}
