package seed

import (
	"self-payrol/helper"
	"self-payrol/model"
)

var UserSeed = []model.User{
	{
		SecretID:   helper.UnsafeHashPassword("123456"),
		Name:       "Abdul Jamal",
		Email:      "abdul@example.com",
		Phone:      "08123456789",
		Address:    "Random street no.21",
		PositionID: 1,
	},
}
