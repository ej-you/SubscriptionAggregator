package v1

import (
	"errors"
	"fmt"
	"time"

	"SubscriptionAggregator/internal/pkg/utils"
)

// inPathUUID is input data with UUID in path.
type inPathUUID struct {
	// uuid
	ID string `path:"id" validate:"required,uuid4"`
}

// @description inSubsCreate is body input data with subs data.
type inSubsCreate struct {
	// service name
	ServiceName string `json:"service_name" validate:"required,max=100" maxLength:"100" example:"Yandex Plus"`
	// price
	Price int `json:"price" validate:"required" example:"400"`
	// user uuid
	UserID string `json:"user_id" validate:"required,uuid4" example:"60601fee-2bf1-4721-ae6f-7636e79a0cba"`
	// start date
	StartDate string `json:"start_date" validate:"required" example:"07-2025"`
	// end date
	EndDate *string `json:"end_date,omitempty" validate:"omitempty" example:"08-2025"`

	// string start date parsed into time.Time
	StartDateParsed *time.Time `json:"-"`
	// string end date parsed into time.Time
	EndDateParsed *time.Time `json:"-"`
}

// ParseDates parses given string dates into StartDateParsed and EndDateParsed fields.
// It returns parsing error if it occurs.
func (c *inSubsCreate) ParseDates() (err error) {
	c.StartDateParsed, c.EndDateParsed, err = parseDates(&c.StartDate, c.EndDate)
	return err // err OR nil
}

// @description inSubsUpdate is body input data with optional subs data.
type inSubsUpdate struct {
	// service name
	ServiceName *string `json:"service_name,omitempty" validate:"omitempty,max=100" maxLength:"100" example:"Yandex Plus"`
	// price
	Price *int `json:"price,omitempty" validate:"omitempty" example:"400"`
	// user uuid
	UserID *string `json:"user_id,omitempty" validate:"omitempty,uuid4" example:"60601fee-2bf1-4721-ae6f-7636e79a0cba"`
	// start date
	StartDate *string `json:"start_date,omitempty" validate:"omitempty" example:"07-2025"`
	// end date
	EndDate *string `json:"end_date,omitempty" validate:"omitempty" example:"08-2025"`

	// string start date parsed into time.Time
	StartDateParsed *time.Time `json:"-"`
	// string end date parsed into time.Time
	EndDateParsed *time.Time `json:"-"`
}

// ParseDates parses given string dates into StartDateParsed and EndDateParsed fields.
// It returns parsing error if it occurs.
func (c *inSubsUpdate) ParseDates() (err error) {
	c.StartDateParsed, c.EndDateParsed, err = parseDates(c.StartDate, c.EndDate)
	return err // err OR nil
}

// @description inSubSumFilter is query-params with user ans service.
type inSubSumFilter struct {
	// service name
	ServiceName string `query:"service_name,omitempty" validate:"omitempty,max=100"`
	// user uuid
	UserID string `query:"user_id,omitempty" validate:"omitempty,uuid4"`
	// start date
	StartDate *string `query:"start_date,omitempty" validate:"omitempty"`
	// end date
	EndDate *string `query:"end_date,omitempty" validate:"omitempty"`

	// string start date parsed into time.Time
	StartDateParsed *time.Time `json:"-"`
	// string end date parsed into time.Time
	EndDateParsed *time.Time `json:"-"`
}

// ParseDates parses given string dates into StartDateParsed and EndDateParsed fields.
// It returns parsing error if it occurs.
func (c *inSubSumFilter) ParseDates() (err error) {
	c.StartDateParsed, c.EndDateParsed, err = parseDates(c.StartDate, c.EndDate)
	return err // err OR nil
}

// parseDates parses given start and end string dates into time.Time structs.
// It returns parsing error if it occurs. Also it checks that end date is after startd date
// if both start and end dates is not nil.
func parseDates(startStr, endStr *string) (startDate, endDate *time.Time, err error) {
	// parse start date if it is presented
	if startStr != nil {
		parsedStart, err := utils.ParseDate(*startStr)
		if err != nil {
			return startDate, endDate, fmt.Errorf("parse start date: %w", err)
		}
		startDate = &parsedStart
	}
	// parse end date if it is presented
	if endStr != nil {
		parsedEnd, err := utils.ParseDate(*endStr)
		if err != nil {
			return startDate, endDate, fmt.Errorf("parse end date: %w", err)
		}
		endDate = &parsedEnd
	}
	if startDate != nil && endDate != nil {
		// if end date after start date
		if startDate.After(*endDate) {
			return startDate, endDate, errors.New("end date after start date")
		}
	}
	return startDate, endDate, nil
}
