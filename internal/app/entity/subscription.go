// Package entity contains all app entities.
package entity

import "time"

// @description Subscription object
type Subscription struct {
	// subscription uuid
	ID string `json:"id" gorm:"id;primaryKey;type:uuid"`
	// price
	Price int `json:"price" gorm:"price;not null"`
	// user uuid
	UserID string `json:"user_id" gorm:"user_id;type:uuid;not null"`
	// start date
	StartDate *time.Time `json:"start_date" gorm:"start_date;not null"`
	// end date
	EndDate *time.Time `json:"end_date" gorm:"end_date"`

	// service uuid
	ServiceID string `json:"-" gorm:"service_id;type:uuid"`
	// service name
	ServiceName string `json:"service_name" gorm:"service_name;->"`
}

func (Subscription) TableName() string {
	return "subs"
}

// @description Pagination settings for SubscriptionList result.
type SubscriptionPagination struct {
	// current page number
	Page int `json:"page" query:"page" validate:"omitempty,min=1"`
	// total pages amount
	Pages int `json:"pages"`
	// total subs
	Total int64 `json:"total"`
	// subs per page amount
	Limit int `json:"limit" query:"limit" validate:"omitempty,min=1"`
}

// Subscription list with pagination.
type SubscriptionList struct {
	Data       []Subscription          `json:"data"`
	Pagination *SubscriptionPagination `json:"pagination"`
}

// @description Subscription object variant for update it.
type SubscriptionUpdate struct {
	// subscription uuid
	ID string `json:"id" gorm:"id;primaryKey;type:uuid"`
	// price
	Price *int `json:"price" gorm:"price"`
	// user uuid
	UserID *string `json:"user_id" gorm:"user_id;type:uuid"`
	// start date
	StartDate *time.Time `json:"start_date" gorm:"start_date"`
	// end date
	EndDate *time.Time `json:"end_date" gorm:"end_date"`

	// service uuid
	ServiceID *string `json:"-" gorm:"service_id;type:uuid"`
	// service name
	ServiceName *string `json:"service_name" gorm:"service_name;->"`
}

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
	// filter fields
	Filter *SubscriptionSumFilter `json:"filter,omitempty"`
	// result
	Sum int `json:"sum"`
}
