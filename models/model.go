package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

// subscription_type_enum
// status_enum

// User has many CreditCards, UserID is the foreign key
type User struct {
	gorm.Model
	Email         string `gorm:"Not Null" json:email`
	Username      string `gorm:"Not Null" json:username`
	Password      string `gorm:"Not Null" json:password`
	Subscriptions []Subscription
}

type Subscription struct {
	gorm.Model
	Name             string    `gorm:"Not Null" json:name`
	SubscriptionType string    `sql:"type:subscription_type_enum Not Null" json:"subscription_type"`
	Status           string    `gorm:"type:status_enum Not Null" json:"status"`
	LastRenewalDate  time.Time `gorm:"Not Null" json:last_renewal_date`
	NextRenewalDate  time.Time `gorm:"Not Null" json:next_renewal_date`
	Cost             string    `gorm:"Not Null" json:cost`
	Currency         string    `gorm:"Not Null" json:currency`
	UserID           uint
}
