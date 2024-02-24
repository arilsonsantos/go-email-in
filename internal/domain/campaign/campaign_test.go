package campaign

import (
	"github.com/jaswdr/faker"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	name     = "My campaign"
	content  = "Body of the campaign"
	contacts = []string{"test1@email.com", "test2@email.com"}
)

func Test_NewCampaing_CreateCampaign(t *testing.T) {
	assertions := assert.New(t)

	campaing, _ := NewCampaign(name, content, contacts)

	assertions.Equal(campaing.Name, name)
	assertions.Equal(campaing.Content, content)
	assertions.Equal(len(campaing.Contacts), len(contacts))
	assertions.Equal(campaing.Contacts[0].Email, contacts[0])
	assertions.Equal(campaing.Status, Pending)
}

func Test_NewCampaign_ID_IsNotNIL(t *testing.T) {
	assertions := assert.New(t)
	campaign, _ := NewCampaign(name, content, contacts)

	assertions.NotNil(campaign.ID)
}

func Test_NewCampaign_CreateAt_ShouldBeNow(t *testing.T) {
	assertions := assert.New(t)
	now := time.Now().Add(-time.Minute)

	campaign, _ := NewCampaign(name, content, contacts)

	assertions.Greater(campaign.CreatedAt, now)
}

func Test_NewCampaign_EmptyName(t *testing.T) {
	assertions := assert.New(t)
	contacts := []string{"email@email.com"}

	_, err := NewCampaign("", content, contacts)
	assertions.NotNil(err)
	assert.Equal(t, "name is less than the minimum 3", err.Error())
}

func Test_NewCampaign_EmptyContent(t *testing.T) {
	assertions := assert.New(t)
	contacts := []string{"email@email.com"}

	_, err := NewCampaign(name, "", contacts)
	assertions.NotNil(err)
	assert.Equal(t, "content is less than the minimum 5", err.Error())
}

func Test_NewCampaign_ContentWithMoreThan100Characters(t *testing.T) {
	assertions := assert.New(t)
	//create a new faker for testing
	fake := faker.New()

	_, err := NewCampaign(fake.Lorem().Text(100), "", contacts)
	assertions.NotNil(err)
	assert.Equal(t, "name is greater than the maximum 30", err.Error())
}

func Test_NewCampaign_EmptyContact(t *testing.T) {
	assertions := assert.New(t)

	_, err := NewCampaign(name, content, []string{})
	assertions.NotNil(err)
	assert.Equal(t, "contacts is less than the minimum 1", err.Error())
}

func Test_NewCampaign_InvalidEmail(t *testing.T) {
	assertions := assert.New(t)
	contacts := []string{"email"}

	_, err := NewCampaign(name, content, contacts)
	assertions.NotNil(err)
	assert.Equal(t, "email is not a valid email", err.Error())
}
