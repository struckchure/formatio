package app

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/idempotency"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"

	_ "pkg.formatio/docs"
)

func NewApp() *fiber.App {
	app := fiber.New(fiber.Config{
		StructValidator:          NewStructValidator(),
		EnableSplittingOnParsers: true,
	})

	app.Use(recover.New())
	app.Use(idempotency.New())
	app.Use(logger.New(logger.Config{TimeFormat: "02/01/2006 15:04:05"}))
	app.Use(cors.New(
		cors.Config{
			AllowOrigins: []string{"*"},
			AllowHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization"},
		},
	))

	return app
}
