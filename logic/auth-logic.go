package logic

import (
	"log"
	"time"

	helper "github.com/horlabyc/iSub/helpers"
	"github.com/horlabyc/iSub/repository"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	model "github.com/horlabyc/iSub/models"
)

func RegisterUser(payload model.User) (model.User, error) {
	var user = payload
	var _, err = repository.FindOne(bson.M{"email": payload.Email})
	if err != mongo.ErrNoDocuments {
		return model.User{}, errors.New("user already exists")
	}
	_, err = repository.FindOne(bson.M{"username": payload.Username})
	if err != mongo.ErrNoDocuments {
		return model.User{}, errors.New("user already exists")
	}
	hashedPassword, hashError := helper.HashPassword(payload.Password)
	if hashError != nil {
		log.Panic(hashError.Error())
	}
	user.Password = hashedPassword
	time := time.Now().String()
	user.CreatedAt = time
	user.UpdatedAt = time
	user.Id = primitive.NewObjectID()
	userId := user.Id.Hex()
	user.UserId = userId
	// token, refreshToken := helper.GenerateToken(&user.Username, &user.Email, &user.Id)
	user = repository.CreateUser(user)
	return user, nil
}

func LoginUser(payload model.User) (model.User, string, string, error) {
	var existingUser, err = repository.FindOne(bson.M{"email": payload.Email})
	if err == mongo.ErrNoDocuments {
		return model.User{}, "", "", errors.New("email or password is incorrect")
	}
	passwordIsValid, errorMsg := helper.VerifyPassword(existingUser.Password, payload.Password)
	if !passwordIsValid {
		return model.User{}, "", "", errors.New(errorMsg)
	}
	token, refreshToken := helper.GenerateToken(&existingUser.Username, &existingUser.Email, &existingUser.Id)
	// UpdateUserTokens(token, refreshToken, existingUser.UserId)
	// updatedUser, e := repository.FindOne(bson.M{"userid": existingUser.UserId})
	log.Println(existingUser.Email, existingUser.Username)
	return existingUser, token, refreshToken, nil
}

func UpdateUserTokens(signedToken string, signedRefreshToken string, userId string) {
	var updateData primitive.D
	updateData = append(updateData, bson.E{"token", signedToken})
	updateData = append(updateData, bson.E{"refreshToken", signedRefreshToken})
	time := time.Now().String()
	updateData = append(updateData, bson.E{"updatedAt", time})
	upsert := true
	options := options.UpdateOptions{
		Upsert: &upsert,
	}
	filter := bson.M{"userid": userId}
	_, err := repository.UpdateOne(filter, updateData, options)
	if err != nil {
		log.Panic(err)
		return
	}
	return
}
