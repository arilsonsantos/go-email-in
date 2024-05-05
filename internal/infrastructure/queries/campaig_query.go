package queries

const (
	INSERT_CAMPAIGN_NAME = "INSERT INTO go.campaign (name, created_at) VALUES (:name, now()) RETURNING id"
	SELECT_ID_NAME_BY_ID = "SELECT id, name FROM go.campaign WHERE id = :id"
	SELECT_ALL           = `select cp.id, 
                                cp.name, 
                                cp.created_at as createdAt, 
                                cp.content, cp.status, 
                                ct.id IDEmail, 
                                ct.email 
                            from go.campaign cp join go.contact ct on ct.campaign_id = cp.id
                            order by cp.id`
)
