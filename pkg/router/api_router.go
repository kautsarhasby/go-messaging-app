package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/kautsarhasby/go-messaging-app/app/controllers"
)

type ApiRouter struct{}

func (h ApiRouter) InstallRouter(app *fiber.App) {
	api := app.Group("/api", limiter.New())
	api.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Hello from api",
		})
	})

	usersGroup := app.Group("/users")
	usersV1Group := usersGroup.Group("/v1")
	usersV1Group.Post("/register", controllers.Register)
	usersV1Group.Post("/login", controllers.Login)
	usersV1Group.Delete("/logout", MiddlewareValidateAuth, controllers.Logout)
	usersV1Group.Put("/refresh-token", MiddlewareRefreshAuth, controllers.RefreshToken)

	messagesGroup := app.Group("/messages")
	messagesV1Group := messagesGroup.Group("/v1")
	messagesV1Group.Get("/history", MiddlewareValidateAuth, controllers.GetMessageHistory)
}

func NewApiRouter() *ApiRouter {
	return &ApiRouter{}
}
