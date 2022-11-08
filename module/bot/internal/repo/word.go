package repo

import (
	"context"
	"github.com/dimayaschhu/vocabulary/module/bot/internal/model"
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

func (w *WordRepo) GetWordsByLimit(limit int) ([]model.Word, error) {
	words := make([]model.Word, 0, 3)
	coll := w.client.Database(w.nameDatabase).Collection("word")
	filter := []bson.D{bson.D{{"$sample", bson.D{{"size", limit}}}}}
	r, err := coll.Aggregate(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	results := []bson.M{}
	if err := r.All(context.TODO(), &results); err != nil {
		return nil, err
	}
	for _, r := range results {
		word := model.Word{Name: r["name"].(string), Translate: r["translate"].(string)}
		words = append(words, word)
	}

	return words, nil
}
