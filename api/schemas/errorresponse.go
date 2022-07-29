package schemas

// Error Response for Validators
type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}