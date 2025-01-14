package model

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Picture struct {
	UserID  int64 `json:"user_id" bson:"user_id"`
	ChatID  int64 `json:"chat_id" bson:"chat_id"`
	IsMuted bool  `json:"is_muted" bson:"is_muted"`
}

func NewPicture(userID int64, chatID int64, isMuted bool) *Picture {
	return &Picture{
		UserID:  userID,
		ChatID:  chatID,
		IsMuted: isMuted,
	}
}

func GetPictureByID(db *mongo.Database, ctx context.Context, Id int64) (*Picture, error) {
	var picture *Picture
	dat, err := db.Collection("picture").FindOne(ctx, bson.M{"user_id": Id}).DecodeBytes()
	if err != nil {
		return nil, err
	}

	err = bson.Unmarshal(dat, &picture)
	return picture, err
}

func SavePicture(db *mongo.Database, ctx context.Context, picture *Picture) error {
	_, err := db.Collection("picture").UpdateOne(ctx, bson.M{"user_id": picture.UserID}, bson.D{{Key: "$set", Value: picture}}, options.Update().SetUpsert(true))
	return err
}

func DeletePictureByID(db *mongo.Database, ctx context.Context, Id int64) error {
	_, err := db.Collection("picture").DeleteOne(ctx, bson.M{"user_id": Id})
	return err
}
