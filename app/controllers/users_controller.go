package controllers

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/kautsarhasby/go-messaging-app/app/models"
	"github.com/kautsarhasby/go-messaging-app/app/repository"
	"github.com/kautsarhasby/go-messaging-app/pkg/response"
	"github.com/kautsarhasby/go-messaging-app/pkg/tokens"
	"golang.org/x/crypto/bcrypt"
)

func Register(ctx *fiber.Ctx) error {
	request := new(models.RegisterRequest)
	err := ctx.BodyParser(request)
	if err != nil {
		errResponse := fmt.Errorf("Failed to parse request  %v", err)
		return response.SendFailureResponse(ctx, fiber.StatusBadRequest, errResponse.Error(), nil)

	}

	err = request.Validate()
	if err != nil {
		errResponse := fmt.Errorf("Failed to validate request: %v", err)
		return response.SendFailureResponse(ctx, fiber.StatusBadRequest, errResponse.Error(), nil)
	}

	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		errResponse := fmt.Errorf("Failed to Insert hash password : %v", err)
		return response.SendFailureResponse(ctx, fiber.StatusInternalServerError, errResponse.Error(), nil)
	}
	request.Password = string(password)

	user := &models.User{
		Username: request.Username,
		Fullname: request.Fullname,
		Password: request.Password,
	}

	err = repository.InsertUser(ctx.Context(), user)
	if err != nil {
		errResponse := fmt.Errorf("Failed to insert new user: %v", err)
		return response.SendFailureResponse(ctx, fiber.StatusInternalServerError, errResponse.Error(), nil)
	}
	res := user
	res.Password = ""

	return response.SendSuccessResponse(ctx, res)
}

func Login(ctx *fiber.Ctx) error {
	request := new(models.LoginRequest)
	now := time.Now()
	err := ctx.BodyParser(request)
	if err != nil {
		errResponse := fmt.Errorf("Failed to parse request  %v", err)
		return response.SendFailureResponse(ctx, fiber.StatusBadRequest, errResponse.Error(), nil)
	}

	user, err := repository.GetUserByUsername(ctx.Context(), request.Username)
	if err != nil {
		errResponse := fmt.Errorf("Username or Password invalid  %v", err)
		return response.SendFailureResponse(ctx, fiber.StatusBadRequest, errResponse.Error(), nil)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		errResponse := fmt.Errorf("Username or Password invalid  %v", err)
		return response.SendFailureResponse(ctx, fiber.StatusBadRequest, errResponse.Error(), nil)
	}
	token, err := tokens.GenerateToken(ctx.Context(), user.Username, user.Fullname, "token")
	if err != nil {
		errResponse := fmt.Errorf("Failed to generate token  %v", err)
		return response.SendFailureResponse(ctx, fiber.StatusBadRequest, errResponse.Error(), nil)
	}

	refreshToken, err := tokens.GenerateToken(ctx.Context(), user.Username, user.Fullname, "refreshToken")
	if err != nil {
		errResponse := fmt.Errorf("Failed to generate token  %v", err)
		return response.SendFailureResponse(ctx, fiber.StatusBadRequest, errResponse.Error(), nil)
	}

	userSession := &models.UserSession{
		UserID:              int(user.ID),
		Token:               token,
		RefreshToken:        refreshToken,
		TokenExpired:        now.Add(tokens.TokenType["token"]),
		RefreshTokenExpired: now.Add(tokens.TokenType["refreshToken"]),
	}
	if err := repository.InsertUserSession(ctx.Context(), userSession); err != nil {
		errResponse := fmt.Errorf("Failed to insert user session  %v", err)
		return response.SendFailureResponse(ctx, fiber.StatusInternalServerError, errResponse.Error(), nil)
	}

	resp := models.LoginResponse{
		Username:     user.Username,
		Fullname:     user.Fullname,
		Token:        token,
		RefreshToken: refreshToken,
	}

	return response.SendSuccessResponse(ctx, resp)
}

func Logout(ctx *fiber.Ctx) error {
	token := ctx.Get("Authorization")
	if err := repository.DeleteUserSessionByToken(ctx.Context(), token); err != nil {
		errResponse := fmt.Errorf("Failed logout %v", err)
		return response.SendFailureResponse(ctx, fiber.StatusInternalServerError, errResponse.Error(), nil)
	}
	return response.SendSuccessResponse(ctx, nil)
}

func RefreshToken(ctx *fiber.Ctx) error {
	now := time.Now()
	refresh_token := ctx.Get("Authorization")
	username := ctx.Locals("username").(string)
	fullname := ctx.Locals("fullname").(string)
	token, err := tokens.GenerateToken(ctx.Context(), username, fullname, "token")
	if err != nil {
		errResponse := fmt.Errorf("Failed to generate token", err)
		return response.SendFailureResponse(ctx, fiber.StatusInternalServerError, errResponse.Error(), nil)
	}
	if err := repository.UpdateUserSessionByToken(ctx.Context(), token, refresh_token, now.Add(tokens.TokenType["token"])); err != nil {
		errResponse := fmt.Errorf("Failed to update token %v", err)
		return response.SendFailureResponse(ctx, fiber.StatusInternalServerError, errResponse.Error(), nil)
	}
	return response.SendSuccessResponse(ctx, fiber.Map{"token": token})
}
