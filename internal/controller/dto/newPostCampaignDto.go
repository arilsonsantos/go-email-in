package dto

type NewPostCampaignDto struct {
	ID        int
	Name      string
	Content   string
	Emails    []string
	CreatedBy string
}
