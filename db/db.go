package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	ctx context.Context
	cancel context.CancelFunc
)

var (
	clientConn *mongo.Client
)

func Init() {
	uri := "mongodb://2land:2land1234@10.0.16.192:27017"
	var dbErr error
	clientConn, dbErr = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	if dbErr != nil {
		panic(dbErr.Error())
	}

	connectErr := clientConn.Ping(context.TODO(), nil)
	if connectErr != nil {
		panic(connectErr.Error())
	}

	fmt.Println("Contected to MongoDb !!!")
}

func GetConnection() *mongo.Client {
	return clientConn
}