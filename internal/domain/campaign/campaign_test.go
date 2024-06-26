package campaign

import (
	"github.com/jaswdr/faker"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	name      = "My campaign"
	content   = "Body of the campaign"
	contacts  = []string{"test1@email.com", "test2@email.com"}
	createdBy = "teste@email.com"
)

func Test_NewCampaing_CreateCampaign(t *testing.T) {
	assertions := assert.New(t)

	campaign, _ := NewCampaign(name, content, contacts, createdBy)

	println(campaign.CreatedBy)
	assertions.Equal(name, campaign.Name)
	assertions.Equal(content, campaign.Content)
	assertions.Equal(len(contacts), len(campaign.Contacts))
	assertions.Equal(contacts[0], campaign.Contacts[0].Email)
	assertions.Equal(Pending, campaign.Status)
	assertions.Equal(createdBy, campaign.CreatedBy)
}

func Test_NewCampaign_ID_IsNotNIL(t *testing.T) {
	assertions := assert.New(t)
	campaign, _ := NewCampaign(name, content, contacts, createdBy)

	assertions.NotNil(campaign.ID)
}

func Test_NewCampaign_CreateAt_ShouldBeNow(t *testing.T) {
	assertions := assert.New(t)
	now := time.Now().Add(-time.Minute)

	campaign, _ := NewCampaign(name, content, contacts, createdBy)

	assertions.Greater(campaign.CreatedAt, now)
}

func Test_NewCampaign_EmptyName(t *testing.T) {
	assertions := assert.New(t)
	contacts := []string{"email@email.com"}

	_, err := NewCampaign("", content, contacts, createdBy)
	assertions.NotNil(err)
	assert.Equal(t, "name is less than the minimum 3", err.Error())
}

func Test_NewCampaign_EmptyContent(t *testing.T) {
	assertions := assert.New(t)
	contacts := []string{"email@email.com"}

	_, err := NewCampaign(name, "", contacts, createdBy)
	assertions.NotNil(err)
	assert.Equal(t, "content is less than the minimum 5", err.Error())
}

func Test_NewCampaign_ContentWithMoreThan100Characters(t *testing.T) {
	assertions := assert.New(t)
	//create a new faker for testing
	fake := faker.New()

	_, err := NewCampaign(fake.Lorem().Text(100), "", contacts, createdBy)
	assertions.NotNil(err)
	assert.Equal(t, "name is greater than the maximum 30", err.Error())
}

func Test_NewCampaign_EmptyContact(t *testing.T) {
	assertions := assert.New(t)

	_, err := NewCampaign(name, content, []string{}, createdBy)
	assertions.NotNil(err)
	assert.Equal(t, "contacts is less than the minimum 1", err.Error())
}

func Test_NewCampaign_InvalidEmail(t *testing.T) {
	assertions := assert.New(t)
	contacts := []string{"email"}

	_, err := NewCampaign(name, content, contacts, createdBy)
	assertions.NotNil(err)
	assert.Equal(t, "email is not a valid email", err.Error())
}

func Test_NewCampaign_EmptyCreatedBy(t *testing.T) {
	assertions := assert.New(t)
	contacts := []string{"teste@email.com"}

	_, err := NewCampaign(name, "teste@email.com", contacts, "")
	assertions.NotNil(err)
	assert.Equal(t, "createdby is not a valid email", err.Error())
}
