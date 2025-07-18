package v1

import "time"

// inPathUUID is input data with UUID in path.
type inPathUUID struct {
	// uuid
	ID string `path:"id" validate:"required,uuid4"`
}

// @description inSubs is body input data with subs data.
type inSubs struct {
	// service name
	ServiceName string `json:"service_name" validate:"required,max=100" maxLength:"100" example:"Yandex Plus"`
	// price
	Price int `json:"price" validate:"required" example:"400"`
	// user uuid
	UserID string `json:"user_id" validate:"required,uuid4" example:"60601fee-2bf1-4721-ae6f-7636e79a0cba"`
	// start date
	StartDate time.Time `json:"start_date" validate:"required" example:"07-2025"`
	// end date
	EndDate *time.Time `json:"end_date,omitempty" validate:"omitempty" example:"08-2025"`
}

// @description inSubSumFilter is query-params with user ans service.
type inSubSumFilter struct {
	// service name
	ServiceName string `query:"service_name" validate:"max=100,omitempty" maxLength:"100" example:"Yandex Plus"`
	// user uuid
	UserID string `query:"user_id" validate:"uuid4,omitempty" example:"60601fee-2bf1-4721-ae6f-7636e79a0cba"`
	// start date
	StartDate *time.Time `query:"start_date,omitempty" validate:"omitempty" example:"07-2025"`
	// end date
	EndDate *time.Time `query:"end_date,omitempty" validate:"omitempty" example:"08-2025"`
}
