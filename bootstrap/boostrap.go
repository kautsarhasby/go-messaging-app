package bootstrap

import (
	"github.com/gofiber/contrib/v3/monitor"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/gofiber/template/html/v2"

	"github.com/kautsarhasby/go-messaging-app/pkg/database"
	"github.com/kautsarhasby/go-messaging-app/pkg/env"
	"github.com/kautsarhasby/go-messaging-app/pkg/router"
)

func NewApplication() *fiber.App {
	env.SetupEnvFile()
	database.SetupDatabase()
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{Views: engine})
	app.Use(recover.New())
	app.Use(logger.New())
	app.Get("/dashboard", monitor.New())
	router.InstallRouter(app)

	return app
}
