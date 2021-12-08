package repository

import (
	model "github.com/horlabyc/iSub/models"
	"github.com/jinzhu/gorm"
)

type UserRepository struct {
	database *gorm.DB
}

func NewUserRepository(database *gorm.DB) *UserRepository {
	return &UserRepository{
		database: database,
	}
}

func (repository UserRepository) CreateUser(user model.User) model.User {
	// var newUser model.User
	repository.database.Create(&user)
	return user
}

func (repository UserRepository) Find(match map[string]interface{}) *gorm.DB {
	var user model.User
	return repository.database.Where(match).Find(&user)
}

func (repository UserRepository) Login() []model.User {
	var allUsers []model.User
	repository.database.Find(&allUsers)
	return allUsers
}

func (repository UserRepository) FindAll() []model.User {
	var allUsers []model.User
	repository.database.Find(&allUsers)
	return allUsers
}
