package usecase

import (
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

type CardService struct {
	MongoCollection *mongo.Collection
}

type Response struct {
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

func (svg *CardService) CreateCard(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
}
func (svg *CardService) GetCardByID(w http.ResponseWriter, r *http.Request)    {}
func (svg *CardService) GetAllCard(w http.ResponseWriter, r *http.Request)     {}
func (svg *CardService) UpdateCardByID(w http.ResponseWriter, r *http.Request) {}
func (svg *CardService) DeleteCardByID(w http.ResponseWriter, r *http.Request) {}
func (svg *CardService) DeleteAllCard(w http.ResponseWriter, r *http.Request)  {}
