package db

import (
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	config "landate/config"
)

/*
MongoDBClientConnection initializes and establishes
a connection to the MongoDB database.
@return {*mongo.Client} - The connected MongoDB client.
*/
func mongoDBClientConnection() *mongo.Client {

	URI := config.GetEnvConfig("MONGO_URI")
	clientOption := options.Client().ApplyURI(URI)
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal("Unable to connect to MONGODB")
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB")
	return client
}

// MongoDB represents the MongoDB client used for database operations.
var MongoClient *mongo.Client = mongoDBClientConnection()

/*
GetCollection retrieves a MongoDB collection from the connected client.
@param {client} - The connected MongoDB client.
@param {collectionName} - The name of the collection to retrieve.
@return {*mongo.Collection} - The retrieved MongoDB collection.
*/
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("documentDB").Collection(collectionName)
	return collection
}
