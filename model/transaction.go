package model

import (
	"context"
	"self-payrol/response"
	"time"

	"gorm.io/gorm"
)

const (
	TransactionTypeDebit   = "debit"
	TransactionsTypeCredit = "credit"
)

type (
	Transaction struct {
		gorm.Model
		Amount int    `json:"amount"`
		Note   string `json:"note"`
		Type   string `json:"type"`
	}

	TransactionRepository interface {
		Fetch(ctx context.Context, limit, offset int) ([]*Transaction, error)
		FindByNoteAndBeweenDates(ctx context.Context, note string, startDate time.Time, endDate time.Time) (*Transaction, error)
	}

	TransactionUsecase interface {
		Fetch(ctx context.Context, limit, offset int) ([]*response.TransactionResp, int, error)
	}
)
