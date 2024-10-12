package mongoapi

import (
	"context"
	"fmt"
	"log"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func newMongoClient() *mongo.Client {
	mongoTestClient, err := mongo.Connect(context.Background(),
		options.Client().ApplyURI("mongodb://admin:1234@127.0.0.1:27017/?directConnection=true&serverSelectionTimeoutMS=2000"))

	if err != nil {
		log.Fatal("error while connection mongodb", err)
	}

	log.Println("mongodb successfully connected.")

	err = mongoTestClient.Ping(context.Background(), readpref.Primary())

	if err != nil {
		log.Fatal("ping failed", err)
	}

	log.Println("Ping success")

	return mongoTestClient
}

func TestMongoOperations(t *testing.T) {
	mongoTestClient := newMongoClient()
	defer mongoTestClient.Disconnect(context.Background())

	//dummy data
	//card1 := uuid.New().String()
	//card2 := uuid.New().String()

	//connect to collection
	coll := mongoTestClient.Database("cards").Collection("cards")

	cardRepo := CardRepo{MongoCollection: coll}

	// Insert card 1 data

	t.Run("Insert Card 1", func(t *testing.T) {
		card := Card{
			CardId: "1234",
			Front:  "Dummy1",
			Back:   "Hello World",
			Any:    123456789,
		}
		fmt.Println("Test")
		result, err := cardRepo.InsertCard(&card)

		if err != nil {
			t.Fatal("insert 1 operation failed", err)
		}

		t.Log("Insert 1 successful", result)
	})
}
