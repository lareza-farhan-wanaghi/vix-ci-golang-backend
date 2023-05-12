package usecase

import (
	"context"
	"errors"
	"fmt"
	"self-payrol/helper"
	"self-payrol/model"
	"self-payrol/request"
	"self-payrol/response"
	"time"

	"gorm.io/gorm"
)

type userUsecase struct {
	userRepository  model.UserRepository
	positionRepo    model.PositionRepository
	companyRepo     model.CompanyRepository
	transactionRepo model.TransactionRepository
}

func NewUserUsecase(user model.UserRepository, post model.PositionRepository, company model.CompanyRepository, transaction model.TransactionRepository) model.UserUsecase {
	return &userUsecase{
		userRepository:  user,
		positionRepo:    post,
		companyRepo:     company,
		transactionRepo: transaction,
	}
}

func (p *userUsecase) WithdrawSalary(ctx context.Context, req *request.WithdrawRequest) error {
	user, err := p.userRepository.FindByID(ctx, req.ID)
	if err != nil {
		return err
	}

	if err = helper.ValidatePassword(user.SecretID, req.SecretID); err != nil {
		return errors.New("secret id not valid")
	}

	notes := fmt.Sprintf("%s(%d) %s ", user.Name, user.ID, " withdraw salary")

	now := time.Now()
	currentYear, currentMonth, _ := now.Date()
	currentLocation := now.Location()

	firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1)

	_, err = p.transactionRepo.FindByNoteAndBeweenDates(ctx, notes, firstOfMonth, lastOfMonth)
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	if err == nil {
		return errors.New("you have withdrawn in this month")
	}

	err = p.companyRepo.DebitBalance(ctx, user.Position.Salary, notes)
	if err != nil {
		return err
	}

	return nil
}

func (p *userUsecase) GetByID(ctx context.Context, id int) (*response.UserResp, error) {
	user, err := p.userRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	resp := helper.NewUserResp(user)
	return resp, nil
}

func (p *userUsecase) FetchUser(ctx context.Context, limit, offset int) ([]*response.UserResp, error) {

	users, err := p.userRepository.Fetch(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	resps := []*response.UserResp{}
	for _, user := range users {
		resp := helper.NewUserResp(user)
		resps = append(resps, resp)
	}
	return resps, nil

}

func (p *userUsecase) DestroyUser(ctx context.Context, id int) error {
	err := p.userRepository.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (p *userUsecase) EditUser(ctx context.Context, id int, req *request.UserRequest) (*response.UserResp, error) {
	_, err := p.userRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	secretHash, err := helper.HashPassword(req.SecretID)
	if err != nil {
		return nil, errors.New("secret not valid ")
	}

	user, err := p.userRepository.UpdateByID(ctx, id, &model.User{
		SecretID:   secretHash,
		Name:       req.Name,
		Email:      req.Email,
		Phone:      req.Phone,
		Address:    req.Address,
		PositionID: req.PositionID,
	})
	if err != nil {
		return nil, err
	}

	resp := helper.NewUserResp(user)

	return resp, nil
}

func (p *userUsecase) StoreUser(ctx context.Context, req *request.UserRequest) (*response.UserResp, error) {
	secretHash, err := helper.HashPassword(req.SecretID)
	if err != nil {
		return nil, errors.New("secret not valid ")
	}

	newUser := &model.User{
		SecretID:   secretHash,
		Name:       req.Name,
		Email:      req.Email,
		Phone:      req.Phone,
		Address:    req.Address,
		PositionID: req.PositionID,
	}

	_, err = p.positionRepo.FindByID(ctx, req.PositionID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("position id not valid ")
		}

		return nil, err
	}

	user, err := p.userRepository.Create(ctx, newUser)
	if err != nil {
		return nil, err
	}

	resp := helper.NewUserResp(user)
	return resp, nil
}
