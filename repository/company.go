package repository

import (
	"context"
	"errors"
	"self-payrol/config"
	"self-payrol/model"

	"gorm.io/gorm"
)

type companyRepository struct {
	Cfg config.Config
}

// NewCompanyRepository returns the repository of the company model
func NewCompanyRepository(cfg config.Config) model.CompanyRepository {
	return &companyRepository{Cfg: cfg}
}

// Get gets company data
func (c *companyRepository) Get(ctx context.Context) (*model.Company, error) {
	company := new(model.Company)

	if err := c.Cfg.Database().WithContext(ctx).First(company).Error; err != nil {
		return nil, err
	}

	return company, nil
}

// CreateOrUpdate creates company data or updates the existing one
func (c *companyRepository) CreateOrUpdate(ctx context.Context, company *model.Company) (*model.Company, error) {
	companyModel := new(model.Company)

	if err := c.Cfg.Database().WithContext(ctx).Debug().First(&companyModel).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if err := c.Cfg.Database().WithContext(ctx).Create(&company).Find(companyModel).Error; err != nil {
				return nil, err
			}

			return companyModel, nil
		}
		return nil, err
	}

	if err := c.Cfg.Database().WithContext(ctx).
		Model(&model.Company{}).Where("id = ?", companyModel.ID).Updates(company).Find(companyModel).Error; err != nil {
		return nil, err
	}

	return companyModel, nil
}

// DebitBalance decreases the company's balance
func (c *companyRepository) DebitBalance(ctx context.Context, amount int, note string) error {
	company, err := c.Get(ctx)
	if err != nil {
		return errors.New("company data not found")
	}

	if company.Balance < amount {
		return errors.New("insufficient company's balance")
	}
	company.Balance -= amount

	if err := c.Cfg.Database().WithContext(ctx).Model(company).Updates(company).Find(company).Error; err != nil {
		return err
	}

	if err := c.Cfg.Database().WithContext(ctx).Create(&model.Transaction{
		Amount: amount,
		Note:   note,
		Type:   model.TransactionTypeDebit,
	}).Error; err != nil {
		return err
	}

	return nil
}

// AddBalance increases the company's balance
func (c *companyRepository) AddBalance(ctx context.Context, balance int) (*model.Company, error) {
	company, err := c.Get(ctx)
	if err != nil {
		return nil, errors.New("company data not found")
	}

	company.Balance += balance

	if err := c.Cfg.Database().WithContext(ctx).Model(company).Updates(company).Find(company).Error; err != nil {
		return nil, err

	}

	if err := c.Cfg.Database().WithContext(ctx).Create(&model.Transaction{
		Amount: balance,
		Note:   "Topup balance company",
		Type:   model.TransactionsTypeCredit,
	}).Error; err != nil {
		return nil, err
	}

	return company, nil
}
