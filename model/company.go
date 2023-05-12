package model

import (
	"context"
	"self-payrol/request"
	"self-payrol/response"

	"gorm.io/gorm"
)

type (
	Company struct {
		gorm.Model
		Name    string `json:"name"`
		Address string `json:"address"`
		Balance int    `json:"balance"`
	}

	CompanyRepository interface {
		Get(ctx context.Context) (*Company, error)
		CreateOrUpdate(ctx context.Context, Company *Company) (*Company, error)
		AddBalance(ctx context.Context, balance int) (*Company, error)
		DebitBalance(ctx context.Context, amount int, note string) error
	}

	CompanyUsecase interface {
		GetCompanyInfo(ctx context.Context) (*response.CompanyResp, int, error)
		CreateOrUpdateCompany(ctx context.Context, req request.CompanyRequest) (*response.CompanyResp, int, error)
		TopupBalance(ctx context.Context, req request.TopupCompanyBalance) (*response.CompanyResp, int, error)
	}
)
