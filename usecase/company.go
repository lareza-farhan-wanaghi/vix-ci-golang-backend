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

func NewCompanyUsecase(repo model.CompanyRepository) model.CompanyUsecase {
	return &companyUsecase{companyRepo: repo}
}

func (c *companyUsecase) GetCompanyInfo(ctx context.Context) (*response.CompanyResp, int, error) {
	company, err := c.companyRepo.Get(ctx)
	if err != nil {
		return nil, http.StatusNotFound, err
	}

	resp := helper.NewCompanyResp(company)
	return resp, http.StatusOK, err
}

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

func (c *companyUsecase) TopupBalance(ctx context.Context, req request.TopupCompanyBalance) (*response.CompanyResp, int, error) {
	company, err := c.companyRepo.AddBalance(ctx, req.Balance)
	if err != nil {
		return nil, http.StatusUnprocessableEntity, err
	}

	resp := helper.NewCompanyResp(company)
	return resp, http.StatusOK, nil
}
