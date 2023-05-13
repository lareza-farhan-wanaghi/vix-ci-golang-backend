package delivery

import (
	"self-payrol/helper"
	"self-payrol/model"
	"self-payrol/request"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/labstack/echo/v4"
)

type companyDelivery struct {
	companyUsecase model.CompanyUsecase
}

type CompanyDelivery interface {
	Mount(group *echo.Group)
}

// NewCompanyDelivery returns the delivery implementation of the company group path
func NewCompanyDelivery(companyUsecase model.CompanyUsecase) CompanyDelivery {
	return &companyDelivery{companyUsecase: companyUsecase}
}

// Mount mounts the available paths of the company group path
func (comp *companyDelivery) Mount(group *echo.Group) {
	group.GET("", comp.GetDetailCompanyHandler)
	group.POST("", comp.UpdateOrCreateCompanyHandler)
	group.POST("/topup", comp.TopupBalanceHandler)

}

// GetDetailCompanyHandler handles the delivery of the path that gets the company detail
func (comp *companyDelivery) GetDetailCompanyHandler(e echo.Context) error {
	ctx := e.Request().Context()

	info, i, err := comp.companyUsecase.GetCompanyInfo(ctx)
	if err != nil {
		return helper.ResponseErrorJson(e, i, err)
	}

	return helper.ResponseSuccessJson(e, "success", info)

}

// UpdateOrCreateCompanyHandler handles the delivery of the path that inserts or updates the company
func (comp *companyDelivery) UpdateOrCreateCompanyHandler(e echo.Context) error {
	ctx := e.Request().Context()

	var req request.CompanyRequest

	if err := e.Bind(&req); err != nil {
		return helper.ResponseValidationErrorJson(e, "Error binding struct", err.Error())
	}

	if err := req.Validate(); err != nil {
		errVal := err.(validation.Errors)
		return helper.ResponseValidationErrorJson(e, "Error validation", errVal)
	}

	company, i, err := comp.companyUsecase.CreateOrUpdateCompany(ctx, req)
	if err != nil {
		return helper.ResponseErrorJson(e, i, err)
	}

	return helper.ResponseSuccessJson(e, "success", company)
}

// TopupBalanceHandler handles the delivery of the path that top ups the balance of the company
func (comp *companyDelivery) TopupBalanceHandler(e echo.Context) error {
	ctx := e.Request().Context()

	var req request.TopupCompanyBalance

	if err := e.Bind(&req); err != nil {
		return helper.ResponseValidationErrorJson(e, "Error binding struct", err.Error())
	}

	if err := req.Validate(); err != nil {
		errVal := err.(validation.Errors)
		return helper.ResponseValidationErrorJson(e, "Error validation", errVal)
	}

	company, i, err := comp.companyUsecase.TopupBalance(ctx, req)
	if err != nil {
		return helper.ResponseErrorJson(e, i, err)
	}

	return helper.ResponseSuccessJson(e, "success", company)
}
