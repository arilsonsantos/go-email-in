package queries

const (
	INSERT_CAMPAIGN_NAME = "INSERT INTO go.campaign (name, created_at) VALUES (:name, now()) RETURNING id"
	SELECT_ID_NAME_BY_ID = "SELECT id, name FROM go.campaign WHERE id = :id"
	SELECT_ID_NAME       = "SELECT id, name FROM go.campaign"
)
