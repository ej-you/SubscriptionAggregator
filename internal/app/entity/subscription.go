// Package entity contains all app entities.
package entity

import "time"

// Subscription model.
// @description Subscription object
type Subscription struct {
	// subscription uuid
	ID string `json:"id" gorm:"id;primaryKey;type:uuid"`
	// service name
	ServiceName string `json:"service_name" gorm:"service_name;not null"`
	// price
	Price int `json:"price" gorm:"price;not null"`
	// user uuid
	UserID string `json:"user_id" gorm:"user_id;not null"`
	// start date
	StartDate *time.Time `json:"start_date" gorm:"start_date;not null"`
	// end date
	EndDate *time.Time `json:"end_date,omitempty" gorm:"end_date"`
}

func (Subscription) TableName() string {
	return "subs"
}

// Subscription list.
type SubscriptionList []Subscription

// @description Filter for SubscriptionSum result.
type SubscriptionSumFilter struct {
	// service name
	ServiceName string `json:"service_name,omitempty"`
	// user uuid
	UserID string `json:"user_id,omitempty"`
	// start date
	StartDate *time.Time `json:"start_date,omitempty"`
	// end date
	EndDate *time.Time `json:"end_date,omitempty"`
}

// @description Sum of subs prices filtered by Filter.
type SubscriptionSum struct {
	Filter *SubscriptionSumFilter `json:"filter,omitempty"`
	Sum    int                    `json:"sum"`
}
