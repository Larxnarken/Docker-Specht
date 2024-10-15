package main

import (
	usecase "Test/Usecase"
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var mongoClient *mongo.Client

func init() {
	// load .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("env load error", err)
	}

	log.Println("env file loaded")

	// create Mongo client
	mongoClient, err = mongo.Connect(context.Background(), options.Client().ApplyURI(os.Getenv("MONGO_URI")))

	if err != nil {
		log.Fatal("Connection error", err)
	}

	err = mongoClient.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Ping failed", err)
	}

	log.Println("Mongo connected")

}

func main() {
	// Close Mongo Connection
	defer mongoClient.Disconnect(context.Background())

	coll := mongoClient.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("COLLECTION_NAME"))

	// Create Card Service
	CardService := usecase.CardService{MongoCollection: coll}

	r := mux.NewRouter()

	r.HandleFunc("/health", healthHandler).Methods(http.MethodGet)

	r.HandleFunc("/card", CardService.CreateCard).Methods(http.MethodPost)

	r.HandleFunc("/card/{id}", CardService.GetCardByID).Methods(http.MethodGet)
	r.HandleFunc("/card", CardService.GetAllCard).Methods(http.MethodGet)

	r.HandleFunc("/card/{id}", CardService.UpdateCardByID).Methods(http.MethodPut)

	r.HandleFunc("/card/{id}", CardService.DeleteCardByID).Methods(http.MethodDelete)
	r.HandleFunc("/card", CardService.DeleteAllCard).Methods(http.MethodDelete)

	log.Println("service running on 4444")
	http.ListenAndServe(":4444", r)

}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("running..."))
}
