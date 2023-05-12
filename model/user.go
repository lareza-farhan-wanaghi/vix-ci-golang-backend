package model

import (
	"context"
	"self-payrol/request"
	"self-payrol/response"

	"gorm.io/gorm"
)

type (
	User struct {
		gorm.Model
		SecretID   string    `json:"secret_id"`
		Name       string    `json:"name"`
		Email      string    `json:"email"`
		Phone      string    `json:"phone"`
		Address    string    `json:"address"`
		PositionID int       `json:"position_id"`
		Position   *Position `json:"position" gorm:"foreignKey:PositionID"`
	}

	UserRepository interface {
		Create(ctx context.Context, user *User) (*User, error)
		UpdateByID(ctx context.Context, id int, user *User) (*User, error)
		FindByID(ctx context.Context, id int) (*User, error)
		Delete(ctx context.Context, id int) error
		Fetch(ctx context.Context, limit, offset int) ([]*User, error)
	}

	UserUsecase interface {
		GetByID(ctx context.Context, id int) (*response.UserResp, error)
		FetchUser(ctx context.Context, limit, offset int) ([]*response.UserResp, error)
		DestroyUser(ctx context.Context, id int) error
		EditUser(ctx context.Context, id int, req *request.UserRequest) (*response.UserResp, error)
		StoreUser(ctx context.Context, req *request.UserRequest) (*response.UserResp, error)
		WithdrawSalary(ctx context.Context, req *request.WithdrawRequest) error
	}
)
