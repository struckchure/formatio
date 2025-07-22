package middlewares

import (
	"strings"

	"github.com/gofiber/fiber/v3"
	"pkg.formatio/lib"
)

type IJwtMiddleware interface {
	Use(c fiber.Ctx) error
}

type JwtMiddleware struct {
	jwt lib.IJwt
}

type headersArgs struct {
	Authorization string `json:"authorization" header:"authorization"`
}

func (m *JwtMiddleware) Use(c fiber.Ctx) error {
	headers := new(headersArgs)
	if err := c.Bind().Header(headers); err != nil {
		return c.
			Status(fiber.StatusBadRequest).
			JSON(err)
	}

	jwtToken := strings.SplitN(headers.Authorization, " ", 2)
	verifiedJwtToken, err := m.jwt.VerifyJWT(jwtToken[1], lib.ACCESS_TOKEN_TYPE)
	if err != nil {
		return c.
			Status(fiber.StatusUnauthorized).
			JSON(fiber.Map{"message": err.Error()})
	}

	userId, err := verifiedJwtToken.Claims.GetSubject()
	if err != nil {
		return c.
			Status(fiber.StatusUnauthorized).
			JSON(fiber.Map{"message": err.Error()})
	}

	fiber.Locals(c, "userId", userId)

	return c.Next()
}

func NewJwtMiddleware(jwt lib.IJwt) IJwtMiddleware {
	return &JwtMiddleware{jwt: jwt}
}
