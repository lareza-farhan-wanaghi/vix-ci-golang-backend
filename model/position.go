package model

import (
	"context"
	"self-payrol/request"
	"self-payrol/response"

	"gorm.io/gorm"
)

type (
	Position struct {
		gorm.Model
		Name   string `json:"name"`
		Salary int    `json:"salary"`
	}

	PositionRepository interface {
		Create(ctx context.Context, Position *Position) (*Position, error)
		UpdateByID(ctx context.Context, id int, Position *Position) (*Position, error)
		FindByID(ctx context.Context, id int) (*Position, error)
		Delete(ctx context.Context, id int) error
		Fetch(ctx context.Context, limit, offset int) ([]*Position, error)
	}

	PositionUsecase interface {
		GetByID(ctx context.Context, id int) (*response.PositionResp, error)
		FetchPosition(ctx context.Context, limit, offset int) ([]*response.PositionResp, error)
		DestroyPosition(ctx context.Context, id int) error
		EditPosition(ctx context.Context, id int, req *request.PositionRequest) (*response.PositionResp, error)
		StorePosition(ctx context.Context, req *request.PositionRequest) (*response.PositionResp, error)
	}
)
