package configs

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoInstance struct {
	Client *mongo.Client
	DB     *mongo.Database
}

var MI MongoInstance

func ConnectDB() {
	dbName := fmt.Sprintf("%s-backend", AppEnv())
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(EnvMongoURI()).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Fatalf(err.Error())
	}
	// defer func() {
	// 	if err = client.Disconnect(context.TODO()); err != nil {
	// 		panic(err)
	// 	}
	// }()
	var result bson.M
	if err := client.Database(dbName).RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Decode(&result); err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	MI = MongoInstance{
		Client: client,
		DB:     client.Database(dbName),
	}
}

func StoreRequestInDb(request any, collection string) (string, error) {
	coll := MI.DB.Collection(collection)
	insertResult, err := coll.InsertOne(context.TODO(), request)
	if err != nil {
		log.Printf("there was an error recording the request \n%v\n", err)
		return "", err
	}
	uuid, ok := insertResult.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", err
	}
	return uuid.Hex(), nil
}

func UpdateRequestInDb(ID string, request any, collection string) (string, error) {
	coll := MI.DB.Collection(collection)
	filter := bson.M{"_id": ID}
	update := bson.M{"$set": request}
	_, err := coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Printf("there was an error updating the request \n%v\n", err)
		return "", err
	}
	return ID, nil
}
