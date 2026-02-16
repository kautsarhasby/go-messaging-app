package repository

import (
	"context"
	"fmt"

	"github.com/kautsarhasby/go-messaging-app/app/models"
	"github.com/kautsarhasby/go-messaging-app/pkg/database"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func InsertMessage(ctx context.Context, message models.MessagePayload) error {
	_, err := database.MongoDB.InsertOne(ctx, message)
	return err
}

func GetAllMessage(ctx context.Context) ([]models.MessagePayload, error) {
	var (
		err      error
		response []models.MessagePayload
	)
	cursor, err := database.MongoDB.Find(ctx, bson.D{})
	if err != nil {
		return nil, fmt.Errorf("failed to insert new message : %v", err)
	}

	for cursor.Next(ctx) {
		payload := models.MessagePayload{}
		err := cursor.Decode(&payload)
		if err != nil {
			return nil, fmt.Errorf("failed to decode new message : %v", err)
		}
		response = append(response, payload)

	}
	return response, nil
}
