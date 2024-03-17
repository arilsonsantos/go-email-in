-- Tabela Contact
CREATE TABLE IF NOT EXISTS Contact (
       ID TEXT PRIMARY KEY,
       Email TEXT
);

-- Tabela Campaign
CREATE TABLE IF NOT EXISTS Campaign (
        ID TEXT PRIMARY KEY,
        Name TEXT,
        Created_At DATETIME,
        Content TEXT,
        Contact_ID TEXT,
        Status TEXT,
        FOREIGN KEY (ContactID) REFERENCES Contact(ID)
);
