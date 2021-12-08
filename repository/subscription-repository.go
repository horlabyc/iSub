package repository

import "github.com/jinzhu/gorm"

type SubscriptionRepository struct {
	database *gorm.DB
}

func NewSubscriptionRepository(database *gorm.DB) *SubscriptionRepository {
	return &SubscriptionRepository{
		database: database,
	}
}
