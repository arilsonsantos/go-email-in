package contract

type NewCampaignResponseDto struct {
	ID       int
	Name     string
	Content  string
	Status   string
	Contacts []NewContactDto
}

type NewContactDto struct {
	Id    int
	Email string
}
