package main

import (
	"context"
	"github.com/dimayaschhu/vocabulary/migrations"
	"github.com/dimayaschhu/vocabulary/pkg/db"
	"github.com/dimayaschhu/vocabulary/pkg/di"
	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/fx"
	"gorm.io/gorm"
	"os"
	"strings"
)

func NewMigrationsCommand() *cobra.Command {
	migrationsCmd := &cobra.Command{
		Use: "migrations",
	}

	migrateCmd := &cobra.Command{
		Use:   "migrate",
		Short: "Execute migrations",
		RunE: func(cmd *cobra.Command, args []string) error {

			fxOptions := di.AppProviders()
			fxOptions = append(fxOptions,
				fx.Invoke(func(db *gorm.DB) {
					if err := migrations.Create(db); err != nil {
						println(err.Error())
					}
				}))

			app := fx.New(fxOptions...)

			return app.Start(context.Background())
		},
	}

	createMigrationCmd := &cobra.Command{
		Use:   "create",
		Short: "Create new migration",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			fxOptions := di.AppProviders()
			fxOptions = append(fxOptions,
				fx.Invoke(func(client *mongo.Client, dbConfig db.Config) {
					dat, err := os.ReadFile("./data/dat")
					if err != nil {
						println(err.Error())
					}
					for _, words := range strings.Split(string(dat), "|") {
						word := strings.Split(words, "-")
						coll := client.Database(dbConfig.GetNameDB()).Collection("word")
						doc := bson.D{{"name", word[0]}, {"translate", word[1]}, {"lesson", 3}}

						_, err := coll.InsertOne(context.TODO(), doc)
						if err != nil {
							panic(err.Error())
						}
					}

				}))

			app := fx.New(fxOptions...)

			return app.Start(context.Background())

		},
	}

	migrationsCmd.AddCommand(migrateCmd)
	migrationsCmd.AddCommand(createMigrationCmd)

	return migrationsCmd
}

type Word struct {
	Id        int `gorm:"primaryKey;autoIncrement"`
	Name      string
	Translate string
	Lesson    int
}
