package repositories

import (
	"context"
	"time"

	"github.com/cassiusbessa/micro-service-update-service/entities"
	"github.com/cassiusbessa/micro-service-update-service/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func MongoConnection() (*mongo.Client, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var err error
	client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	return client, cancel
}

func UpdateService(db string, id string, service entities.Service) (bool, error) {
	collection := client.Database(db).Collection("company")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	mongoId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		cancel()
		return false, err
	}

	filter := bson.M{
		"services._id": mongoId,
	}
	arrayFilter := bson.M{"elem._id": mongoId}
	update := bson.M{
		"$set": bson.M{
			"services.$[elem]": service,
		},
	}
	updateOpts := options.Update().SetArrayFilters(options.ArrayFilters{
		Filters: []interface{}{arrayFilter},
	})

	result, err := collection.UpdateOne(ctx, filter, update, updateOpts)
	if err != nil {
		cancel()
		return false, err
	}
	if result.MatchedCount == 0 {
		cancel()
		return false, nil
	}

	defer cancel()
	return true, nil
}

func SaveError(db string, customErr errors.CustomError) {
	collection := client.Database(db).Collection("errors")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	_, err := collection.InsertOne(ctx, customErr)
	if err != nil {
		cancel()
	}
	defer cancel()
}
