CREATE TABLE IF NOT EXISTS scores (
    client_id UUID NOT NULL REFERENCES clients(client_id) ON DELETE CASCADE,
    advertiser_id UUID NOT NULL REFERENCES advertisers(advertiser_id) ON DELETE CASCADE,
    score INT NOT NULL,
    UNIQUE (client_id, advertiser_id)
)