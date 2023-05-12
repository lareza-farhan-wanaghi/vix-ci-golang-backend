package usecase

import (
	"context"
	"net/http"
	"self-payrol/helper"
	"self-payrol/model"
	"self-payrol/response"
)

type transactionUsecase struct {
	transactionRepository model.TransactionRepository
}

func NewTransactionUsecase(transaction model.TransactionRepository) model.TransactionUsecase {
	return &transactionUsecase{transactionRepository: transaction}
}

func (t *transactionUsecase) Fetch(ctx context.Context, limit, offset int) ([]*response.TransactionResp, int, error) {
	transations, err := t.transactionRepository.Fetch(ctx, limit, offset)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	resps := []*response.TransactionResp{}
	for _, transaction := range transations {
		resp := helper.NewTransactionResp(transaction)
		resps = append(resps, resp)
	}

	return resps, http.StatusOK, err
}
