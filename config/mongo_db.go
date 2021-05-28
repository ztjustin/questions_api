package config

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	host = "localhost"
	port = 27017
)

func ClientDB() *mongo.Client {
	ClientOpts := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%d", host, port))
	Client, err := mongo.Connect(context.TODO(), ClientOpts)
	if err != nil {
		log.Fatal(err)
	}

	//TODO: Check the connections
	err = Client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to DB")

	return Client

}
