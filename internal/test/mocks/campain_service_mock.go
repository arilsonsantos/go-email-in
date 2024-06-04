package mocks

import (
	"context"
	contract2 "emailn/internal/controller/dto"
	"fmt"
	"github.com/stretchr/testify/mock"
)

type CampaignServiceMock struct {
	context.Context
	mock.Mock
}

func (r *CampaignServiceMock) GetCampaigns() (*[]contract2.NewGetCampaignDto, error) {
	args := r.Called()
	fmt.Printf("test")
	return nil, args.Error(1)
}

func (r *CampaignServiceMock) GetBy(id int) (*contract2.NewGetCampaignDto, error) {
	args := r.Called(id)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*contract2.NewGetCampaignDto), nil
}

func (r *CampaignServiceMock) CreateCampaign(ctx context.Context, dto contract2.NewPostCampaignDto) (int, error) {
	args := r.Called(dto)
	return args.Get(0).(int), args.Error(1)
}

func (r *CampaignServiceMock) Start(id int) error {
	args := r.Called(id)
	return args.Error(0)
}
