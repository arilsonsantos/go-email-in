package mocks

import (
	"context"
	"emailn/internal/contract"
	"fmt"
	"github.com/stretchr/testify/mock"
)

type CampaignServiceMock struct {
	context.Context
	mock.Mock
}

func (r *CampaignServiceMock) GetCampaigns() (*[]contract.NewGetCampaignDto, error) {
	args := r.Called()
	fmt.Printf("test")
	return nil, args.Error(1)
}

func (r *CampaignServiceMock) GetBy(id int) (*contract.NewGetCampaignDto, error) {
	args := r.Called(id)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*contract.NewGetCampaignDto), nil
}

func (r *CampaignServiceMock) CreateCampaign(ctx context.Context, dto contract.NewPostCampaignDto) (int, error) {
	args := r.Called(dto)
	return args.Get(0).(int), args.Error(1)
}
