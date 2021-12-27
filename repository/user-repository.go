package repository

import (
	"context"
	"log"
	"time"

	"github.com/horlabyc/iSub/database"
	model "github.com/horlabyc/iSub/models"
	"github.com/jinzhu/gorm"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository struct {
	database *gorm.DB
}

var UserCollection = database.OpenConnection(database.Client, "users")

func NewUserRepository(database *gorm.DB) *UserRepository {
	return &UserRepository{
		database: database,
	}
}

func CreateUser(user model.User) model.User {
	_, err := UserCollection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Fatal(err)
	}
	return user
}

func FindOneUser(match bson.M) (model.User, error) {
	var user model.User
	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return user, UserCollection.FindOne(ctx, match).Decode(&user)
}

func UpdateOneUser(filter bson.M, updateData primitive.D, options options.UpdateOptions) (*mongo.UpdateResult, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	return UserCollection.UpdateOne(
		ctx,
		filter,
		bson.D{
			{"$set", updateData},
		},
		&options,
	)
}

func (repository UserRepository) Login() []model.User {
	var allUsers []model.User
	repository.database.Find(&allUsers)
	return allUsers
}

func FindAllUsers(match primitive.D, group primitive.D, projectStage primitive.D) (*mongo.Cursor, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return UserCollection.Aggregate(ctx, mongo.Pipeline{match, group, projectStage})
}
