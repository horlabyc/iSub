package model

import (
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
	Id               primitive.ObjectID    `bson:"_id"`
	Name             string                `bson:"name"`
	SubscriptionType string                `bson:"subscriptionType"`
	Status           string                `bson:"status"`
	LastRenewalDate  string                `bson:"lastRenewalDate"`
	NextRenewalDate  string                `bson:nextRenewalDate`
	Cost             string                `bson:"cost"`
	Currency         string                `bson:"currency"`
	UserId           string                `bson:"userId"`
	History          []SubscriptionHistory `bson:"history"`
}

type SubscriptionHistory struct {
	Description string `bson:"description"`
	Date        string `bson:"date"`
}
