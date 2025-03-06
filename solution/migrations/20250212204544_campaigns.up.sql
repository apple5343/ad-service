CREATE TABLE IF NOT EXISTS campaigns (
    campaign_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    advertiser_id UUID NOT NULL REFERENCES advertisers(advertiser_id) ON DELETE CASCADE,
    impressions_limit INT NOT NULL,
    clicks_limit INT NOT NULL,
    cost_per_impression FLOAT NOT NULL,
    cost_per_click FLOAT NOT NULL,
    ad_title VARCHAR(255) NOT NULL,
    ad_text TEXT NOT NULL,
    start_date INT NOT NULL,
    end_date INT NOT NULL,
    gender VARCHAR(255),
    age_from INT,
    age_to INT,
    active BOOLEAN NOT NULL,
    location VARCHAR(255),
    impression_count INT NOT NULL DEFAULT 0,
    click_count INT NOT NULL DEFAULT 0,
    image_url VARCHAR(512),
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_campaigns_active ON campaigns(active);

