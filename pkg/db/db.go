package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewDB() (*mongo.Client, error) {
	return mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27018"))
}

type Config interface {
	GetNameDB() string
}

type ConfigTest struct {
}

func NewConfigTest() Config {
	return &ConfigTest{}
}

func (c *ConfigTest) GetNameDB() string {
	return "vocabularyTest"
}

type ConfigProd struct{}

func NewConfigProd() Config {
	return &ConfigProd{}
}

func (c *ConfigProd) GetNameDB() string {
	return "vocabulary"
}
