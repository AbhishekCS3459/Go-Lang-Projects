package db

import (
	"context"
	"fmt"
	"time"

	"github.com/AbhishekCS3459/HR_MS/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mg = &models.Mg

func ConnectDB(MONGOURI string, DB_NAME string) {
	// logic to connect to db
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(MONGOURI).SetServerAPIOptions(serverAPI)
	// create a new client and connect to server
	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		panic(err)
	}
	// send a ping to make sure the connection is established
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Second)
	defer cancel()
	err = client.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}
	// set the db instance
	mg.Client = client
	mg.Db = client.Database(DB_NAME)
	fmt.Print("Connected to MongoDB!")
}
