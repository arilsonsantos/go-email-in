package dto

type NewGetCampaignDto struct {
	ID       int
	Name     string
	Content  string
	Status   string
	Contacts []NewGetContactDto
}

type NewGetContactDto struct {
	Id    int
	Email string
}
