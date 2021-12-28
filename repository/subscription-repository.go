package repository

import (
	"context"
	"log"

	"github.com/horlabyc/iSub/database"
	model "github.com/horlabyc/iSub/models"
)

var SubscriptionCollection = database.OpenConnection(database.Client, "subscription")

func CreateSubscription(sub model.Subscription) model.Subscription {
	_, err := SubscriptionCollection.InsertOne(context.TODO(), sub)
	if err != nil {
		log.Fatal(err)
	}
	return sub
}
