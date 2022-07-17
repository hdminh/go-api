package database

import (
	"go_code/models"
	"go_code/configs"

    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
	"go.mongodb.org/mongo-driver/bson"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")

func FindUserByUsername(ctx context.Context, username string) (*models.User, error) {
	var user models.User
	if err := userCollection.FindOne(ctx, bson.M{"username": username}).Decode(&user); err != nil {
		 return nil, err
	}
	return &user, nil
}

func FindUserById(ctx context.Context, userId string) (*models.User, error) {
	objId, _ := primitive.ObjectIDFromHex(userId)
	var user models.User
	if err := userCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&user); err != nil {
		 return nil, err
	}
	return &user, nil
}