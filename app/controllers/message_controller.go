package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/kautsarhasby/go-messaging-app/app/repository"
	"github.com/kautsarhasby/go-messaging-app/pkg/response"
)

func GetMessageHistory(ctx *fiber.Ctx) error {
	resp, err := repository.GetAllMessage(ctx.Context())
	if err != nil {
		errResponse := fmt.Errorf("Failed to fetch message history %v", err)
		return response.SendFailureResponse(ctx, fiber.StatusInternalServerError, errResponse.Error(), nil)
	}

	return response.SendSuccessResponse(ctx, resp)
}
