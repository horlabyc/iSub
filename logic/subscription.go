package logic

import (
	"strings"
	"time"

	model "github.com/horlabyc/iSub/models"
	"github.com/horlabyc/iSub/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	SUBSCRIPTION_STATUS_PENDING = "pending"
)

func CreateSub(payload model.Subscription, userId string) model.Subscription {
	var sub = payload
	sub.Id = primitive.NewObjectID()
	sub.Status = SUBSCRIPTION_STATUS_PENDING
	sub.Currency = strings.ToUpper(payload.Currency)
	sub.UserId = userId
	history := model.SubscriptionHistory{
		Description: "Create new subscription",
		Date:        time.Now().String(),
	}
	sub.History = append(sub.History, history)
	sub = repository.CreateSubscription(sub)
	return sub
}
