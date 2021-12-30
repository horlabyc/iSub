package repository

import (
	"context"
	"log"
	"time"

	"github.com/horlabyc/iSub/database"
	model "github.com/horlabyc/iSub/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var SubscriptionCollection = database.OpenConnection(database.Client, "subscription")

func CreateSubscription(sub model.Subscription) model.Subscription {
	_, err := SubscriptionCollection.InsertOne(context.TODO(), sub)
	if err != nil {
		log.Fatal(err)
	}
	return sub
}

func FindOneSub(match bson.M) (model.Subscription, error) {
	var sub model.Subscription
	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return sub, SubscriptionCollection.FindOne(ctx, match).Decode(&sub)
}

func UpdateOneSub(filter bson.M, updateData primitive.D, options options.UpdateOptions) (*mongo.UpdateResult, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	return SubscriptionCollection.UpdateOne(
		ctx,
		filter,
		bson.D{
			{"$set", updateData},
		},
		&options,
	)
}

func FindAllSubs(match primitive.D, group primitive.D, projectStage primitive.D) (*mongo.Cursor, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return SubscriptionCollection.Aggregate(ctx, mongo.Pipeline{match, group, projectStage})
}
