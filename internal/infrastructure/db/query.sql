-- name: CreateContact :one
INSERT INTO Contact (email) VALUES (?) RETURNING *;
-- name: CreateCampaign :one
INSERT INTO Campaign (id, name, created_at, content, contact_id, status) VALUES (?, ?, ?, ?, ?, ?) RETURNING *;
-- name: GetCampaignsWithContacts :many
select c.id, c.name, c.created_at, c.content, c.status, co.email from Campaign c join Contact co on c.contact_id = co.id;
