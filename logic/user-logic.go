package logic

import (
	model "github.com/horlabyc/iSub/models"
	"github.com/horlabyc/iSub/repository"
	"go.mongodb.org/mongo-driver/bson"
)

func GetUser(userId string) (model.User, error) {
	return repository.FindOneUser(bson.M{"userId": userId})
}
