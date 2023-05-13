package request

import validation "github.com/go-ozzo/ozzo-validation"

type (
	CompanyRequest struct {
		Name    string `json:"name" validate:"required"`
		Balance int    `json:"balance" validate:"required"`
		Address string `json:"address" validate:"required"`
	}

	TopupCompanyBalance struct {
		Balance int `json:"balance" validate:"required"`
	}
)

// Validate validates the company request data
func (req CompanyRequest) Validate() error {
	return validation.ValidateStruct(
		&req,
		validation.Field(&req.Name, validation.Required),
		validation.Field(&req.Balance, validation.Required),
		validation.Field(&req.Address, validation.Required),
	)
}

// Validate validates the top up request data
func (req TopupCompanyBalance) Validate() error {
	return validation.ValidateStruct(
		&req,
		validation.Field(&req.Balance, validation.Required),
	)
}
