package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/users/", userPage)
	http.HandleFunc("/posts/", postPage)
	http.HandleFunc("/posts/users/", listPostsPage)
	log.Fatal(http.ListenAndServe(":5000", nil))
}

var Database *mongo.Database
var MongoContext context.Context

func connectMongo() {
	var (
		client   *mongo.Client
		mongoURL = "mongodb://localhost:27017"
	)

	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURL))
	if err != nil {
		fmt.Println("ERROR: Unable to create a new client")
	}

	MongoContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(MongoContext)
	if err != nil {
		fmt.Println("ERROR: Unable to connect client with mongo context", MongoContext)
	}

	MongoContext, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err = client.Ping(MongoContext, readpref.Primary()); err != nil {
		fmt.Println("ERROR: Could not ping to Mongo DB service: ", err)
		return
	}

	fmt.Println("connected to nosql database:", mongoURL)
	Database = client.Database("instaclone")
}

func main() {
	connectMongo()
	handleRequests()
}
