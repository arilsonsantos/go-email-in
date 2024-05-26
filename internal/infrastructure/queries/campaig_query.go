package queries

const (
	INSERT_CAMPAIGN = "INSERT INTO go.campaign (id, name, created_at, content, status) VALUES (nextval('go.campaign_id_seq'::regclass), $1, $2, $3, $4) RETURNING id"
	INSERT_CONTACT  = "INSERT INTO go.contact (id, email, campaign_id) VALUES (nextval('go.contact_id_seq'::regclass),$1, $2) RETURNING id"

	SELECT_BY_ID = `SELECT 
                                cp.id, 
                                cp.name, 
                                cp.created_at as createdAt, 
                                cp.content, 
                                cp.status, 
                                ct.id IDEmail, 
                                ct.email 
                            FROM 
                                go.campaign cp join go.contact ct on ct.campaign_id = cp.id
                            WHERE 
                                cp.id = $1`
	SELECT_ALL = `select cp.id, 
                                cp.name, 
                                cp.created_at as createdAt, 
                                cp.content, 
                                cp.status, 
                                ct.id IDEmail, 
                                ct.email 
                            from go.campaign cp join go.contact ct on ct.campaign_id = cp.id
                            order by cp.id`
)
