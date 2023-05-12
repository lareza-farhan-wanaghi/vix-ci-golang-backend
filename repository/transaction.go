package repository

import (
	"context"
	"self-payrol/config"
	"self-payrol/model"
	"time"
)

type transactionRepository struct {
	Cfg config.Config
}

func NewTransactionRepository(cfg config.Config) model.TransactionRepository {
	return &transactionRepository{Cfg: cfg}
}

func (t *transactionRepository) Fetch(ctx context.Context, limit, offset int) ([]*model.Transaction, error) {
	var data []*model.Transaction

	if err := t.Cfg.Database().WithContext(ctx).
		Limit(limit).Offset(offset).Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func (t *transactionRepository) FindByNoteAndBeweenDates(ctx context.Context, note string, startDate time.Time, endDate time.Time) (*model.Transaction, error) {
	var data *model.Transaction

	if err := t.Cfg.Database().WithContext(ctx).
		Where("note = ?", note).Where("created_at BETWEEN ? AND ?", startDate, endDate).First(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}
