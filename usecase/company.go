package usecase

import (
	"context"
	"net/http"
	"self-payrol/helper"
	"self-payrol/model"
	"self-payrol/request"
	"self-payrol/response"
)

type companyUsecase struct {
	companyRepo model.CompanyRepository
}

// NewCompanyUsecase returns the usecase implementation of the company group path
func NewCompanyUsecase(repo model.CompanyRepository) model.CompanyUsecase {
	return &companyUsecase{companyRepo: repo}
}

// GetCompanyInfo handles the usecase of the path that gets the company data
func (c *companyUsecase) GetCompanyInfo(ctx context.Context) (*response.CompanyResp, int, error) {
	company, err := c.companyRepo.Get(ctx)
	if err != nil {
		return nil, http.StatusNotFound, err
	}

	resp := helper.NewCompanyResp(company)
	return resp, http.StatusOK, err
}

// GetCompanyInfo handles the usecase of the path that creates or updates the company data
func (c *companyUsecase) CreateOrUpdateCompany(ctx context.Context, req request.CompanyRequest) (*response.CompanyResp, int, error) {
	company, err := c.companyRepo.CreateOrUpdate(ctx, &model.Company{
		Name:    req.Name,
		Address: req.Address,
		Balance: req.Balance,
	})
	if err != nil {
		return nil, http.StatusUnprocessableEntity, err
	}

	resp := helper.NewCompanyResp(company)
	return resp, http.StatusOK, nil

}

// GetCompanyInfo handles the usecase of the path that top ups the balance of the company data
func (c *companyUsecase) TopupBalance(ctx context.Context, req request.TopupCompanyBalance) (*response.CompanyResp, int, error) {
	company, err := c.companyRepo.AddBalance(ctx, req.Balance)
	if err != nil {
		return nil, http.StatusUnprocessableEntity, err
	}

	resp := helper.NewCompanyResp(company)
	return resp, http.StatusOK, nil
}
