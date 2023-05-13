package usecase

import (
	"context"
	"self-payrol/helper"
	"self-payrol/model"
	"self-payrol/request"
	"self-payrol/response"
)

type positionUsecase struct {
	positionRepository model.PositionRepository
}

// NewPositionUsecase returns the usecase implementation of the position group path
func NewPositionUsecase(position model.PositionRepository) model.PositionUsecase {
	return &positionUsecase{positionRepository: position}
}

// GetByID handles the usecase of the path that gets position data by its id
func (p *positionUsecase) GetByID(ctx context.Context, id int) (*response.PositionResp, error) {
	position, err := p.positionRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	resp := helper.NewPositionResp(position)
	return resp, nil
}

// FetchPosition handles the usecase of the path that gets all position data
func (p *positionUsecase) FetchPosition(ctx context.Context, limit, offset int) ([]*response.PositionResp, error) {

	positions, err := p.positionRepository.Fetch(ctx, limit, offset)
	if err != nil {
		return nil, err
	}
	resps := []*response.PositionResp{}

	for _, position := range positions {
		resp := helper.NewPositionResp(position)
		resps = append(resps, resp)
	}
	return resps, nil
}

// DestroyPosition handles the usecase of the path that deletes position data
func (p *positionUsecase) DestroyPosition(ctx context.Context, id int) error {
	err := p.positionRepository.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

// EditPosition handles the usecase of the path that updates position data
func (p *positionUsecase) EditPosition(ctx context.Context, id int, req *request.PositionRequest) (*response.PositionResp, error) {
	_, err := p.positionRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	position, err := p.positionRepository.UpdateByID(ctx, id, &model.Position{
		Name:   req.Name,
		Salary: req.Salary,
	})
	if err != nil {
		return nil, err
	}

	resp := helper.NewPositionResp(position)

	return resp, nil
}

// StorePosition handles the usecase of the path that inserts position data
func (p *positionUsecase) StorePosition(ctx context.Context, req *request.PositionRequest) (*response.PositionResp, error) {
	newPosition := &model.Position{
		Name:   req.Name,
		Salary: req.Salary,
	}

	position, err := p.positionRepository.Create(ctx, newPosition)
	if err != nil {
		return nil, err
	}

	resp := helper.NewPositionResp(position)

	return resp, nil
}
