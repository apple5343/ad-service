CREATE TABLE IF NOT EXISTS clients (
    client_id UUID PRIMARY KEY,
    login VARCHAR(255) NOT NULL,
    age INT NOT NULL,
    location VARCHAR(255) NOT NULL,
    gender VARCHAR(255) NOT NULL
)