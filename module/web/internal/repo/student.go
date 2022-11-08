package repo

import (
	"context"
	"github.com/dimayaschhu/vocabulary/pkg/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type StudentRepo struct {
	client       *mongo.Client
	nameDatabase string
}

func NewStudentRepo(client *mongo.Client, dbConfig db.Config) *StudentRepo {
	return &StudentRepo{client: client, nameDatabase: dbConfig.GetNameDB()}
}

func (w *StudentRepo) CreateStudent(name string, lesson, chatId int) error {
	coll := w.client.Database(w.nameDatabase).Collection("students")
	doc := bson.D{{"name", name}, {"lesson", lesson}, {"chatId", chatId}}

	_, err := coll.InsertOne(context.TODO(), doc)

	return err
}

func (w *StudentRepo) ExistStudent(name string) bool {
	coll := w.client.Database(w.nameDatabase).Collection("students")

	quantity, _ := coll.CountDocuments(context.TODO(), bson.D{{"name", name}})

	return quantity > 0
}
