package repo

import (
	"context"
	"github.com/dimayaschhu/vocabulary/pkg/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type WordRepo struct {
	client       *mongo.Client
	nameDatabase string
}

func NewWordRepo(client *mongo.Client, dbConfig db.Config) *WordRepo {
	return &WordRepo{client: client, nameDatabase: dbConfig.GetNameDB()}
}

func (w *WordRepo) CreateWord(name, translate string, lesson int) error {
	coll := w.client.Database(w.nameDatabase).Collection("words")
	doc := bson.D{{"name", name}, {"translate", translate}, {"lesson", lesson}}

	_, err := coll.InsertOne(context.TODO(), doc)

	return err
}

func (w *WordRepo) ExistStudent(name string) bool {
	coll := w.client.Database(w.nameDatabase).Collection("words")

	quantity, _ := coll.CountDocuments(context.TODO(), bson.D{{"name", name}})

	return quantity > 0
}
