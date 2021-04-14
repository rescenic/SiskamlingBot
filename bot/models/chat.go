package models

import (
	"SiskamlingBot/bot/helpers/database"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Chat struct {
	ChatID    int64  `json:"chat_id" bson:"chat_id" `
	ChatType  string `json:"chat_type" bson:"chat_type" `
	ChatLink  string `json:"chat_link" bson:"chat_link" `
	ChatTitle string `json:"chat_title" bson:"chat_title" `
}

func GetChatByID(ctx context.Context, Id int) (*Chat, error) {
	var chat Chat
	dat, err := database.Mongo.Collection("chat").FindOne(ctx, bson.M{"chat_id": Id}).DecodeBytes()
	if err != nil {
		return nil, err
	}

	err = bson.Unmarshal(dat, chat)
	return &chat, err
}

func SaveChat(ctx context.Context, chat Chat) error {
	_, err := database.Mongo.Collection("chat").UpdateOne(ctx, bson.M{"chat_id": chat.ChatID}, bson.D{{"$set", chat}}, options.Update().SetUpsert(true))
	return err
}

func DeleteChatByID(ctx context.Context, Id int) error {
	_, err := database.Mongo.Collection("chat").DeleteOne(ctx, bson.M{"chat_id": Id})
	return err
}
