CREATE SCHEMA go;

CREATE SEQUENCE go.campaign_id_seq;

CREATE TABLE IF NOT EXISTS go.campaign (
   id INT PRIMARY KEY DEFAULT nextval('go.campaign_id_seq'),
   name TEXT,
   created_at TIMESTAMP,
   content TEXT,
    email TEXT,
   contact_ID INT,
   status TEXT
);


CREATE SEQUENCE go.contact_id_seq;

CREATE TABLE IF NOT EXISTS go.contact
(
    id    INT  primary key default nextval('go.contact_id_seq'),
    email TEXT,
    campaign_id INT,
    FOREIGN KEY (campaign_id) REFERENCES go.campaign(id)
);
