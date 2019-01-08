package main

import (
	"context"
	"fmt"
	"github.com/mongodb/mongo-go-driver/mongo"
)

func main() {
	var (
		client     *mongo.Client
		err        error
		database   *mongo.Database
		collection *mongo.Collection
	)
	//1、 建立连接
	if client, err = mongo.Connect(context.TODO(), "mongodb://127.0.0.1:27017"); err != nil {
		fmt.Println(err)
		return
	}
	//2、 选择数据库my_db
	database = client.Database("my_db")
	//3、 选择表my_collection
	database.Collection("my_collection")
	collection = collection
}
