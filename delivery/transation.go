package delivery

import (
	"self-payrol/helper"
	"self-payrol/model"
	"strconv"

	"github.com/labstack/echo/v4"
)

type transactionDelivery struct {
	transactionUsecase model.TransactionUsecase
}

type TransactionDelivery interface {
	Mount(group *echo.Group)
}

// NewTransactionDelivery returns the delivery implementation of the transaction group path
func NewTransactionDelivery(transactionUsecase model.TransactionUsecase) TransactionDelivery {
	return &transactionDelivery{transactionUsecase: transactionUsecase}
}

// Mount mounts the available paths of the transaction group path
func (p *transactionDelivery) Mount(group *echo.Group) {
	group.GET("", p.FetchTransactionHandler)
}

// FetchTransactionHandler handles the delivery of the path that gets all transaction data
func (p *transactionDelivery) FetchTransactionHandler(c echo.Context) error {
	ctx := c.Request().Context()

	limit := c.QueryParam("limit")
	offset := c.QueryParam("skip")

	limitInt, _ := strconv.Atoi(limit)
	offsetInt, _ := strconv.Atoi(offset)

	transactions, i, err := p.transactionUsecase.Fetch(ctx, limitInt, offsetInt)
	if err != nil {
		return helper.ResponseErrorJson(c, i, err)
	}

	return helper.ResponseSuccessJson(c, "success", transactions)
}
