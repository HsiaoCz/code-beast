package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CatFact struct {
	Fact   string `bson:"fact" json:"fact"`
	Length string `bson:"length" json:"length"`
}

type Store interface {
	GetAll() ([]*CatFact, error)
	Put(*CatFact) error
}

type MongoStore struct {
	client     *mongo.Client
	database   string
	collection string
}

func NewMongoStore() (*MongoStore, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return nil, err
	}

	return &MongoStore{
		client:     client,
		database:   "catfact",
		collection: "facts",
	}, nil
}

func (ms *MongoStore) GetAll() ([]*CatFact, error) {
	coll := ms.client.Database(ms.database).Collection(ms.collection)
	query := bson.M{}

	cursor, err := coll.Find(context.Background(), query)
	if err != nil {
		return nil, err
	}
	result := []*CatFact{}
	if err := cursor.All(context.Background(), &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (ms *MongoStore) Put(catfact *CatFact) error {
	coll := ms.client.Database(ms.database).Collection(ms.collection)
	_, err := coll.InsertOne(context.Background(), catfact)
	return err
}

type Server struct {
	store Store
}

func NewServer(store Store) *Server {
	return &Server{
		store: store,
	}
}

func (s *Server) handleGetAllFacts(w http.ResponseWriter, r *http.Request) {
	result, err := s.store.GetAll()
	if err != nil {
		log.Fatal(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

type CatFactWorker struct {
	store       Store
	apiEndpoint string
}

func NewCatFactWorker(store Store, apiEndpoint string) *CatFactWorker {
	return &CatFactWorker{
		store:       store,
		apiEndpoint: apiEndpoint,
	}
}

func (cfw *CatFactWorker) start() error {
	ticker := time.NewTicker(2 * time.Second)
	for {
		resp, err := http.Get(cfw.apiEndpoint)
		if err != nil {
			return err
		}
		var catfact CatFact
		if err := json.NewDecoder(resp.Body).Decode(&catfact); err != nil {
			return err
		}
		if err := cfw.store.Put(&catfact); err != nil {
			return err
		}
		<-ticker.C
	}
}

func main() {
	mongoStore, err := NewMongoStore()
	if err != nil {
		log.Fatal(err)
	}
	worker := NewCatFactWorker(mongoStore, "https://catfact.ninja/fact")
	go worker.start()

	server := NewServer(mongoStore)

	http.HandleFunc("/facts", server.handleGetAllFacts)
	http.ListenAndServe(":8991", nil)
}
