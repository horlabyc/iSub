package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id        primitive.ObjectID `bson:"_id"`
	Email     string             `bson:"email" validate:"required,email"`
	Username  string             `bson:"username" validate:"required,min=2,max=100"`
	Password  string             `bson:"password" validate:"required,min=6"`
	CreatedAt string             `bson:"createdAt"`
	UpdatedAt string             `bson:"updatedAt"`
	UserId    string             `bson:"userId"`
}

type Subscription struct {
	Id               primitive.ObjectID `bson:"_id"`
	Name             string             `bson:"name" validate:"required"`
	SubscriptionType string             `bson:"subscriptionType"`
	Status           string             `bson:"status"`
	LastRenewalDate  time.Time          `bson:"lastRenewalDate" validate:"required" `
	NextRenewalDate  time.Time          `bson:nextRenewalDate validate:"required"`
	Cost             string             `bson:"cost"`
	Currency         string             `bson:"currency"`
	UserID           primitive.ObjectID `bson:"_id"`
}
