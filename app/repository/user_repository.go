package repository

import (
	"context"
	"time"

	"github.com/kautsarhasby/go-messaging-app/app/models"
	"github.com/kautsarhasby/go-messaging-app/pkg/database"
)

func InsertUser(ctx context.Context, user *models.User) error {
	return database.DB.WithContext(ctx).Create(user).Error
}

func InsertUserSession(ctx context.Context, session *models.UserSession) error {
	return database.DB.WithContext(ctx).Create(session).Error
}

func GetUserSessionByToken(ctx context.Context, token string) (models.UserSession, error) {
	var response models.UserSession
	err := database.DB.WithContext(ctx).Where("token = ?", token).Last(&response).Error
	return response, err
}

func UpdateUserSessionByToken(ctx context.Context, token, refreshToken string, tokenExpired time.Time) error {
	return database.DB.WithContext(ctx).Exec("UPDATE user_sessions SET token = ? ,token_expired = ? WHERE refresh_token = ?", token, tokenExpired, refreshToken).Error
}

func DeleteUserSessionByToken(ctx context.Context, token string) error {
	return database.DB.WithContext(ctx).Exec("DELETE FROM user_sessions WHERE token = ?", token).Error
}

func GetUserByUsername(ctx context.Context, username string) (models.User, error) {
	var response models.User
	err := database.DB.WithContext(ctx).Where("username = ?", username).Last(&response).Error
	return response, err
}
