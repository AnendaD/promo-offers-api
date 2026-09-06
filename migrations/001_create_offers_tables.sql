CREATE TABLE IF NOT EXISTS merchants (
    id VARCHAR(36) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS offers (
    id VARCHAR(36) PRIMARY KEY,
    merchant_id VARCHAR(36) NOT NULL REFERENCES merchants(id),
    type VARCHAR(20) NOT NULL,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    discount INT NOT NULL,
    minimum_cost INT DEFAULT 0,
    max_discount INT,
    category VARCHAR(100),
    payment_method VARCHAR(50),
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    starts_at TIMESTAMP NOT NULL,
    ends_at TIMESTAMP NOT NULL,
    INDEX idx_merchant_id (merchant_id),
    INDEX idx_active_period (is_active, starts_at, ends_at)
);

CREATE TABLE IF NOT EXISTS offer_incompatibilities (
    id VARCHAR(36) PRIMARY KEY,
    offer_id_1 VARCHAR(36) NOT NULL REFERENCES offers(id),
    offer_id_2 VARCHAR(36) NOT NULL REFERENCES offers(id),
    UNIQUE KEY unique_incompatibility (offer_id_1, offer_id_2)
);

CREATE TABLE IF NOT EXISTS user_segments (
    id VARCHAR(36) PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS user_segment_offers (
    user_segment_id VARCHAR(36) NOT NULL REFERENCES user_segments(id),
    offer_id VARCHAR(36) NOT NULL REFERENCES offers(id),
    PRIMARY KEY (user_segment_id, offer_id)
);
