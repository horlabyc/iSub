package logic

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	helper "github.com/horlabyc/iSub/helpers"
	model "github.com/horlabyc/iSub/models"
	"github.com/horlabyc/iSub/repository"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	SUBSCRIPTION_STATUS_PENDING = "pending"
	SUBSCRIPTION_STATUS_ACTIVE  = "active"
	SUBSCRIPTION_STATUS_CANCEL  = "canceled"
)

func CreateSub(payload model.Subscription, userId string) model.Subscription {
	var sub = payload
	sub.Id = primitive.NewObjectID()
	sub.Status = SUBSCRIPTION_STATUS_PENDING
	sub.Currency = strings.ToUpper(payload.Currency)
	sub.UserId = userId
	t := time.Now().String()
	sub.CreatedAt = t
	history := model.SubscriptionHistory{
		Description: "Create new subscription",
		Date:        time.Now().String(),
	}
	sub.History = append(sub.History, history)
	sub = repository.CreateSubscription(sub)
	return sub
}

func ActivateSub(userId string, subId string) (model.Subscription, error) {
	var subObjectId, error = helper.ConvertToObjectId(subId)
	if error != nil {
		return model.Subscription{}, errors.New("Invalid subscription id provided")
	}
	var existingSub, err = repository.FindOneSub(bson.M{"_id": subObjectId})
	if err == mongo.ErrNoDocuments {
		return model.Subscription{}, errors.New("invalid subscription Id provided")
	}
	if existingSub.Status == SUBSCRIPTION_STATUS_ACTIVE {
		return model.Subscription{}, errors.New("Subscription already activated")
	}
	var nextRenewalDate string
	switch existingSub.SubscriptionType {
	case "annually":
		nextRenewalDate = time.Now().AddDate(1, 0, 0).String()
	case "monthly":
		nextRenewalDate = time.Now().AddDate(0, 1, 0).String()
	case "weekly":
		nextRenewalDate = time.Now().AddDate(0, 0, 7).String()
	case "daily":
		nextRenewalDate = time.Now().AddDate(0, 0, 1).String()
	}
	var history = existingSub.History
	history = append(history, model.SubscriptionHistory{
		Description: "Activate subscription",
		Date:        time.Now().String(),
	})
	var updateData primitive.D
	updateData = append(updateData, bson.E{"status", SUBSCRIPTION_STATUS_ACTIVE})
	updateData = append(updateData, bson.E{"lastRenewalDate", time.Now().String()})
	updateData = append(updateData, bson.E{"nextRenewalDate", nextRenewalDate})
	updateData = append(updateData, bson.E{"history", history})
	upsert := true
	options := options.UpdateOptions{
		Upsert: &upsert,
	}
	filter := bson.M{"_id": subObjectId}
	_, e := repository.UpdateOneSub(filter, updateData, options)
	if e != nil {
		log.Panic(e)
		return model.Subscription{}, errors.New("error occured while activating the subscription")
	}
	updatedSub, err := repository.FindOneSub(bson.M{"_id": subObjectId})
	return updatedSub, nil
}

func CancelSub(userId string, subId string) (model.Subscription, error) {
	var subObjectId, error = helper.ConvertToObjectId(subId)
	if error != nil {
		return model.Subscription{}, errors.New("Invalid subscription id provided")
	}
	var existingSub, err = repository.FindOneSub(bson.M{"_id": subObjectId})
	if err == mongo.ErrNoDocuments {
		return model.Subscription{}, errors.New("invalid subscription Id provided")
	}
	if existingSub.Status != SUBSCRIPTION_STATUS_ACTIVE {
		return model.Subscription{}, errors.New("Subscription is not active")
	}
	var history = existingSub.History
	history = append(history, model.SubscriptionHistory{
		Description: "Cancel subscription",
		Date:        time.Now().String(),
	})
	var updateData primitive.D
	updateData = append(updateData, bson.E{"status", SUBSCRIPTION_STATUS_CANCEL})
	updateData = append(updateData, bson.E{"history", history})
	upsert := true
	options := options.UpdateOptions{
		Upsert: &upsert,
	}
	filter := bson.M{"_id": subObjectId}
	_, e := repository.UpdateOneSub(filter, updateData, options)
	if e != nil {
		log.Panic(e)
		return model.Subscription{}, errors.New("error occured while canceling the subscription")
	}
	updatedSub, err := repository.FindOneSub(bson.M{"_id": subObjectId})
	return updatedSub, nil
}

func GetSub(userId string, subId string) (model.Subscription, error) {
	var subObjectId, error = helper.ConvertToObjectId(subId)
	if error != nil {
		return model.Subscription{}, errors.New("Invalid subscription id provided")
	}
	return repository.FindOneSub(bson.M{"userId": userId, "_id": subObjectId})
}

func GetAllSub(page int, limit int, skip int, userId string) ([]bson.M, error) {
	fmt.Println(page, skip, limit, userId)
	match := bson.D{{"$match", bson.D{{"userId", userId}}}}
	group := bson.D{{"$group", bson.D{
		{"_id", bson.D{{"_id", "null"}}},
		{"total", bson.D{{"$sum", 1}}},
		{"data", bson.D{{"$push", "$$ROOT"}}},
	}}}
	projectStage := bson.D{
		{"$project", bson.D{
			{"_id", 0},
			{"totalCount", 1},
			{"subscriptions", bson.D{{"$slice", []interface{}{"$data", skip, limit}}}},
		}},
	}
	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var allSubs []bson.M
	result, err := repository.FindAllSubs(match, group, projectStage)
	if err = result.All(ctx, &allSubs); err != nil {
		return []bson.M{}, err
	}
	return allSubs, err
}
