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

// TODO: tuliskan validasi untuk CompanyRequest dengan rule semua field required

func (req CompanyRequest) Validate() error {
	return validation.ValidateStruct(
		&req,
		validation.Field(&req.Name, validation.Required),
		validation.Field(&req.Balance, validation.Required),
		validation.Field(&req.Address, validation.Required),
	)
}

func (req TopupCompanyBalance) Validate() error {
	return validation.ValidateStruct(
		&req,
		validation.Field(&req.Balance, validation.Required),
	)
}
