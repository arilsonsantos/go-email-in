package contract

type NewCampaignDto struct {
	ID      int
	Name    string
	Content string
	Emails  []string
}
