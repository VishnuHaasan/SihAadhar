package config

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/vishnuhaasan/PushNotifications/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var Database *mongo.Database
var mongoClient *mongo.Client

func SetCollections() {
	models.InitUser(Database.Collection("users"))
	models.InitRequest(Database.Collection("requests"))
}

func Connect() (*mongo.Database, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		log.Fatal(err)
	}
	mongoClient = client
	return client.Database("main"), cancel
}

func ConnectDB() context.CancelFunc {
	db, cancel := Connect()
	Database = db
	SetCollections()
	return cancel
}

func PingDb() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := mongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
}
