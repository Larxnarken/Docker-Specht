package repository

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Card struct {
	CardId string `json:"card_id" bson:"card_id"`
	Front  string `json:"front" bson:"front"`
	Back   string `json:"back" bson:"back"`
}

type CardRepo struct {
	MongoCollection *mongo.Collection
}

func (r *CardRepo) InsertCard(card *Card) (interface{}, error) {
	result, err := r.MongoCollection.InsertOne(context.Background(), card)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *CardRepo) FindCardByID(cId string) (*Card, error) {
	var card Card

	err := r.MongoCollection.FindOne(context.Background(),
		bson.D{{Key: "card_id", Value: cId}}).Decode(&card)

	if err != nil {
		return nil, err
	}

	return &card, nil
}

func (r *CardRepo) FindAllCard() ([]Card, error) {
	results, err := r.MongoCollection.Find(context.Background(), bson.D{})

	if err != nil {
		return nil, err
	}

	var cards []Card
	err = results.All(context.Background(), &cards)

	if err != nil {
		return nil, fmt.Errorf("results decode error %s", err.Error())
	}

	return cards, nil
}

func (r *CardRepo) UpdateCardByID(cId string, newCard *Card) (int64, error) {

	result, err := r.MongoCollection.UpdateOne(context.Background(),
		bson.D{{Key: "card_id", Value: cId}},
		bson.D{{Key: "$set", Value: newCard}})

	if err != nil {
		return 0, err
	}

	return result.ModifiedCount, nil
}

func (r *CardRepo) DeleteCardByID(cId string) (int64, error) {
	result, err := r.MongoCollection.DeleteOne(context.Background(),
		bson.D{{Key: "card_id", Value: cId}})

	if err != nil {
		return 0, err
	}

	return result.DeletedCount, nil
}

func (r *CardRepo) DeleteAllCard() (int64, error) {
	result, err := r.MongoCollection.DeleteMany(context.Background(), bson.D{})

	if err != nil {
		return 0, err
	}

	return result.DeletedCount, nil
}
