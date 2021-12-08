package logic

import (
	"errors"
	"log"
	"time"

	helper "github.com/horlabyc/iSub/helpers"
	"github.com/horlabyc/iSub/repository"

	model "github.com/horlabyc/iSub/models"
)

func RegisterUser(repository *repository.UserRepository, payload model.User) (model.User, string, string, error) {
	var user = payload
	var existingUser = repository.Find(map[string]interface{}{"email": payload.Email}).RowsAffected
	if existingUser != 0 {
		return model.User{}, "", "", errors.New("user already exists")
	}
	hashedPassword, hashError := helper.HashPassword(payload.Password)
	if hashError != nil {
		log.Panic(hashError.Error())
	}
	user.Password = hashedPassword
	time := time.Now()
	user.CreatedAt = time
	user.UpdatedAt = time
	token, refreshToken := helper.GenerateToken(&user.Username, &user.Email, &user.ID)
	user = repository.CreateUser(user)
	return user, token, refreshToken, nil
}
