package campaign

import "context"

type Repository interface {
	Save(ctx context.Context, campaign *Campaign) (int, error)
	Get() (*[]Campaign, error)
	GetBy(id int) (*Campaign, error)
	Update(campaign *Campaign) error
}
