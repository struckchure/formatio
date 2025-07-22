package middlewares

import (
	"strings"

	"github.com/gofiber/fiber/v3"
	"pkg.formatio/prisma/db"
	"pkg.formatio/services"
	"pkg.formatio/types"
)

type IUserMiddleware interface {
	Use(c fiber.Ctx) error
}

type UserMiddleware struct {
	userService services.IUserService
}

func (m *UserMiddleware) Use(c fiber.Ctx) error {
	userId := fiber.Locals[string](c, "userId")
	if userId == "" {
		return c.
			Status(fiber.StatusUnauthorized).
			JSON(fiber.Map{"message": "authentication failed"})
	}

	user, err := m.userService.GetProfileUser(types.GetUserArgs{ID: &userId})
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(err)
	}

	fiber.Locals(c, "user", user)

	return c.Next()
}

func UserRolesMiddleware(userRoles []string) fiber.Handler {
	return func(c fiber.Ctx) error {
		var isNotAllowed bool

		user := fiber.Locals[*db.UserModel](c, "user")
		if user.Roles != nil {
			isNotAllowed = !strings.Contains(
				strings.Join(user.Roles, ","), strings.Join(userRoles, ","))
		}

		if isNotAllowed {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"message": "access denied"})
		}

		return c.Next()
	}
}

func NewUserMiddleware(userService services.IUserService) IUserMiddleware {
	return &UserMiddleware{userService: userService}
}
