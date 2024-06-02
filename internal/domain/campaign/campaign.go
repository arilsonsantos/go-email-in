package campaign

import (
	"emailn/internal/internalerrors"
	"time"
)

type Contact struct {
	ID         int    `sql:"id"`
	Email      string `validate:"email" db:"email" json:"email"`
	CampaignID int    `sql:"campaign_id"`
}

type Campaign struct {
	ID        int       //`validate:"required"`
	Name      string    `validate:"min=3,max=30" json:"name"`
	CreatedAt time.Time `validate:"required" db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	Content   string    `validate:"min=5,max=100" json:"content"`
	Status    string    `db:"status" json:"status"`
	Contacts  []Contact `validate:"min=1,dive" json:"contacts"`
	CreatedBy string    `validate:"email" db:"email" json:"email"`
}

const (
	Pending string = "Pendent"
	Started string = "Started"
	Done    string = "Done"
	Failed  string = "Failed"
)

func (c *Campaign) Done() {
	c.Status = Done
	c.UpdatedAt = time.Now()
}

func (c *Campaign) Started() {
	c.Status = Started
	c.UpdatedAt = time.Now()
}

func (c *Campaign) Failed() {
	c.Status = Failed
	c.UpdatedAt = time.Now()
}

func NewCampaign(name, content string, emails []string, createdBy string) (*Campaign, error) {
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
		CreatedBy: createdBy,
	}

	err := internalerrors.ValidateStruct(campaign)

	if err != nil {
		return nil, err
	}

	return campaign, nil
}
