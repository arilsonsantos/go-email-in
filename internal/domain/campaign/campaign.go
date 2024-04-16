package campaign

import (
	"emailn/internal/internalerrors"
	"time"
)

type Contact struct {
	Email string `validate:"required,email"`
}

type Campaign struct {
	ID        int       //`validate:"required"`
	Name      string    `validate:"min=3,max=30"`
	CreatedAt time.Time `validate:"required" db:"created_at default:now()"`
	Content   string    `validate:"min=5,max=100"`
	Contacts  []Contact `validate:"min=1,dive"`
	Status    string
}

const (
	Pending string = "Pendent"
	Started string = "Started"
	Done    string = "Done"
)

func NewCampaign(name, content string, emails []string) (*Campaign, error) {
	contacts := make([]Contact, len(emails))

	for i, email := range emails {
		contacts[i].Email = email
	}

	campaign := &Campaign{
		//ID:        xid.New().String(),
		Name:      name,
		CreatedAt: time.Now(),
		Content:   content,
		Contacts:  contacts,
		Status:    Pending,
	}

	err := internalerrors.ValidateStruct(campaign)

	if err != nil {
		return nil, err
	}

	return campaign, nil
}
