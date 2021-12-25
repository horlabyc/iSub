package database

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	helper "github.com/horlabyc/iSub/helpers"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var clientInstance *mongo.Client
var mongoOnce sync.Once

func GetMongoClient() *mongo.Client {
	DB_USER := helper.LoadEnv("DB_USER")
	DB_PASSWORD := helper.LoadEnv("DB_PASSWORD")
	DB_NAME := helper.LoadEnv("DB_NAME")
	DB_HOST := helper.LoadEnv("DB_HOST")
	ConnectionString := "mongodb+srv://" + DB_USER + ":" + DB_PASSWORD + "@" + DB_HOST + "/" + DB_NAME
	clientOption := options.Client().ApplyURI(ConnectionString)
	client, err := mongo.NewClient(clientOption)
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to database")
	// mongoOnce.Do(func() {
	// })
	return client
}

var Client *mongo.Client = GetMongoClient()

func OpenConnection(client *mongo.Client, collectionName string) *mongo.Collection {
	DB_NAME := helper.LoadEnv("DB_NAME")
	fmt.Println(DB_NAME)
	err := client.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		log.Fatal("Ping error", err)
	}
	collection := client.Database(DB_NAME).Collection(collectionName)
	return collection
}

// func ConnectDB() {
// 	var err error
// 	p := helper.LoadEnv("DB_PORT")
// 	dbPort, err := strconv.ParseUint(p, 10, 32)
// 	configData := fmt.Sprintf(
// 		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
// 		helper.LoadEnv("DB_HOST"),
// 		dbPort,
// 		helper.LoadEnv("DB_USER"),
// 		helper.LoadEnv("DB_PASSWORD"),
// 		helper.LoadEnv("DB_NAME"),
// 	)
// 	DB, err = gorm.Open("postgres", configData)
// 	if err != nil {
// 		fmt.Println(
// 			err.Error(),
// 		)
// 		panic("connection to database failed")
// 	}
// 	fmt.Println("Connection Opened to Database")
// }
