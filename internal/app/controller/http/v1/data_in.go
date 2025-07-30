package v1

import (
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
	c.StartDateParsed, err = parseDate(&c.StartDate)
	if err != nil {
		return fmt.Errorf("parse start date: %w", err)
	}
	c.EndDateParsed, err = parseDate(c.EndDate)
	if err != nil {
		return fmt.Errorf("parse end date: %w", err)
	}
	return nil
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
func (u *inSubsUpdate) ParseDates() (err error) {
	u.StartDateParsed, err = parseDate(u.StartDate)
	if err != nil {
		return fmt.Errorf("parse start date: %w", err)
	}
	u.EndDateParsed, err = parseDate(u.EndDate)
	if err != nil {
		return fmt.Errorf("parse end date: %w", err)
	}
	return nil
}

// @description inSubSumFilter is query-params with user ans service.
type inSubSumFilter struct {
	// service name
	ServiceName string `query:"service_name,omitempty" validate:"omitempty,max=100"`
	// user uuid
	UserID string `query:"user_id,omitempty" validate:"omitempty,uuid4"`
	// start date
	StartDate *string `query:"start_date,omitempty" validate:"required"`
	// end date
	EndDate *string `query:"end_date,omitempty" validate:"required"`

	// string start date parsed into time.Time
	StartDateParsed *time.Time `json:"-"`
	// string end date parsed into time.Time
	EndDateParsed *time.Time `json:"-"`
}

// ParseDates parses given string dates into StartDateParsed and EndDateParsed fields.
// It returns parsing error if it occurs.
func (f *inSubSumFilter) ParseDates() (err error) {
	f.StartDateParsed, err = parseDate(f.StartDate)
	if err != nil {
		return fmt.Errorf("parse start date: %w", err)
	}
	f.EndDateParsed, err = parseDate(f.EndDate)
	if err != nil {
		return fmt.Errorf("parse end date: %w", err)
	}
	return nil
}

// parseDate parses given string date into time.Time struct.
// Also it returns parsing error if it occurs.
func parseDate(strDate *string) (*time.Time, error) {
	// parse start date if it is presented
	if strDate == nil {
		return nil, nil
	}
	parsedStart, err := utils.ParseDate(*strDate)
	if err != nil {
		return nil, err
	}
	return &parsedStart, nil
}
