package router

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/kautsarhasby/go-messaging-app/app/repository"
	"github.com/kautsarhasby/go-messaging-app/pkg/response"
	"github.com/kautsarhasby/go-messaging-app/pkg/tokens"
)

func MiddlewareValidateAuth(ctx *fiber.Ctx) error {
	token := ctx.Get("Authorization")
	if token == "" {
		fmt.Println("token authorization empty")
		return response.SendFailureResponse(ctx, fiber.StatusUnauthorized, "Unauthorized", nil)
	}

	claim, err := tokens.ValidateToken(ctx.Context(), token)
	if err != nil {
		fmt.Println("Invalid Token")
		return response.SendFailureResponse(ctx, fiber.StatusUnauthorized, "Unauthorized", nil)
	}

	if _, err := repository.GetUserSessionByToken(ctx.Context(), token); err != nil {
		fmt.Println("Failed to get user session on DB")
		return response.SendFailureResponse(ctx, fiber.StatusUnauthorized, "Unauthorized", nil)
	}

	ctx.Locals("username", claim.Username)
	ctx.Locals("fullname", claim.Fullname)

	return ctx.Next()
}

func MiddlewareRefreshAuth(ctx *fiber.Ctx) error {
	token := ctx.Get("Authorization")
	if token == "" {
		fmt.Println("token authorization empty")
		return response.SendFailureResponse(ctx, fiber.StatusUnauthorized, "Unauthorized", nil)
	}

	claim, err := tokens.ValidateToken(ctx.Context(), token)
	if err != nil {
		fmt.Println("Invalid Token")
		return response.SendFailureResponse(ctx, fiber.StatusUnauthorized, "Unauthorized", nil)
	}

	ctx.Locals("username", claim.Username)
	ctx.Locals("fullname", claim.Fullname)

	return ctx.Next()
}
