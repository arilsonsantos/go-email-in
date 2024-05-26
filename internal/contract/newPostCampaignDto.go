package contract

type NewPostCampaignDto struct {
	ID      int
	Name    string
	Content string
	Emails  []string
}
