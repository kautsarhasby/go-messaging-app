package repository

import (
	"context"

	"github.com/kautsarhasby/go-messaging-app/app/models"
	"github.com/kautsarhasby/go-messaging-app/pkg/database"
)

func InsertUser(ctx context.Context, user *models.User) error {
	return database.DB.WithContext(ctx).Create(user).Error
}

func InsertUserSession(ctx context.Context, session *models.UserSession) error {
	return database.DB.WithContext(ctx).Create(session).Error
}

func GetUserByUsername(ctx context.Context, username string) (models.User, error) {
	var response models.User
	err := database.DB.WithContext(ctx).Where("username = ?", username).Last(&response).Error
	return response, err
}
