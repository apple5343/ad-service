CREATE TABLE IF NOT EXISTS clicks (
    campaign_id UUID NOT NULL REFERENCES campaigns(campaign_id) ON DELETE CASCADE,
    client_id UUID NOT NULL REFERENCES clients(client_id) ON DELETE CASCADE,
    cost FLOAT NOT NULL,
    date INT NOT NULL,
    UNIQUE (campaign_id, client_id)
)