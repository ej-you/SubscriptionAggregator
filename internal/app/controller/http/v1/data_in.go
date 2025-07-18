package v1

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
}

// @description inSubsUpdate is body input data with optional subs data.
type inSubsUpdate struct {
	// service name
	ServiceName string `json:"service_name,omitempty" validate:"omitempty,max=100" maxLength:"100" example:"Yandex Plus"`
	// price
	Price int `json:"price,omitempty" validate:"omitempty" example:"400"`
	// user uuid
	UserID string `json:"user_id,omitempty" validate:"omitempty,uuid4" example:"60601fee-2bf1-4721-ae6f-7636e79a0cba"`
	// start date
	StartDate *string `json:"start_date,omitempty" validate:"omitempty" example:"07-2025"`
	// end date
	EndDate *string `json:"end_date,omitempty" validate:"omitempty" example:"08-2025"`
}

// @description inSubSumFilter is query-params with user ans service.
type inSubSumFilter struct {
	// service name
	ServiceName string `query:"service_name,omitempty" validate:"omitempty,max=100" maxLength:"100" example:"Yandex Plus"`
	// user uuid
	UserID string `query:"user_id,omitempty" validate:"omitempty,uuid4" example:"60601fee-2bf1-4721-ae6f-7636e79a0cba"`
	// start date
	StartDate *string `query:"start_date,omitempty" validate:"omitempty" example:"07-2025"`
	// end date
	EndDate *string `query:"end_date,omitempty" validate:"omitempty" example:"08-2025"`
}
