package usecase

import (
	repository "Test/Repository"
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
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
	EnableCors(&w)
	w.Header().Add("Content-Type", "application/json")

	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	var card repository.Card

	err := json.NewDecoder(r.Body).Decode(&card)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("invalid body", err)
		res.Error = err.Error()
		return
	}

	// assign new CardId
	card.CardId = uuid.NewString()

	repo := repository.CardRepo{MongoCollection: svg.MongoCollection}

	// insert Card
	insertID, err := repo.InsertCard(&card)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("insert error", err)
		return
	}

	res.Data = card.CardId
	w.WriteHeader(http.StatusOK)

	log.Println("Card inserted with id", insertID, card)
}
func (svg *CardService) GetCardByID(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
	w.Header().Add("Content-Type", "application/json")

	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	// get CardId
	cId := mux.Vars(r)["id"]
	log.Println("CardId", cId)

	repo := repository.CardRepo{MongoCollection: svg.MongoCollection}

	card, err := repo.FindCardByID(cId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("error", err)
		res.Error = err.Error()
		return

	}

	res.Data = card
	w.WriteHeader(http.StatusOK)
}
func (svg *CardService) GetAllCard(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
	w.Header().Add("Content-Type", "application/json")

	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	repo := repository.CardRepo{MongoCollection: svg.MongoCollection}

	card, err := repo.FindAllCard()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("error", err)
		res.Error = err.Error()
		return

	}

	res.Data = card
	w.WriteHeader(http.StatusOK)
}
func (svg *CardService) UpdateCardByID(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	w.Header().Add("Content-Type", "application/json")

	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	// get CardId
	cId := mux.Vars(r)["id"]
	log.Println("CardId", cId)

	if cId == "" {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("invalid cardId")
		res.Error = "invalid cardId"
		return
	}

	var card repository.Card

	err := json.NewDecoder(r.Body).Decode(&card)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("invalid body", err)
		res.Error = err.Error()
		return
	}

	card.CardId = cId

	repo := repository.CardRepo{MongoCollection: svg.MongoCollection}
	count, err := repo.UpdateCardByID(cId, &card)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("error ", err)
		res.Error = err.Error()
		return
	}

	res.Data = count
	w.WriteHeader(http.StatusOK)
}
func (svg *CardService) DeleteCardByID(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
	w.Header().Add("Content-Type", "application/json")

	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	// get CardId
	cId := mux.Vars(r)["id"]
	log.Println("CardId", cId)

	repo := repository.CardRepo{MongoCollection: svg.MongoCollection}

	count, err := repo.DeleteCardByID(cId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("error ", err)
		res.Error = err.Error()
		return
	}

	res.Data = count
	w.WriteHeader(http.StatusOK)
}
func (svg *CardService) DeleteAllCard(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
	w.Header().Add("Content-Type", "application/json")

	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	repo := repository.CardRepo{MongoCollection: svg.MongoCollection}

	count, err := repo.DeleteAllCard()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("error ", err)
		res.Error = err.Error()
		return
	}

	res.Data = count
	w.WriteHeader(http.StatusOK)
}

func EnableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Headers", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST,GET,OPTION,PUT,DELETE")
}
