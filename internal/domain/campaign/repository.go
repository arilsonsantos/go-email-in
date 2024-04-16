package campaign

type Repository interface {
	Save(campaign *Campaign) (int, error)
	Get() ([]Campaign, error)
	GetBy(id int) (*Campaign, error)
}
