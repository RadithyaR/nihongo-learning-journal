CREATE TABLE srs_records (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    user_id UUID NOT NULL REFERENCES users(id),

    item_type VARCHAR(20) NOT NULL,
    item_id UUID NOT NULL,

    ease_factor NUMERIC(4,2) NOT NULL DEFAULT 2.5,

    interval_days INTEGER NOT NULL DEFAULT 0,

    review_count INTEGER NOT NULL DEFAULT 0,

    last_reviewed_at TIMESTAMP NULL,

    next_review_at TIMESTAMP NOT NULL,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);

CREATE INDEX idx_srs_records_user_id
ON srs_records(user_id);

CREATE INDEX idx_srs_records_next_review_at
ON srs_records(next_review_at);

CREATE UNIQUE INDEX idx_srs_records_unique_item
ON srs_records(
	user_id,
	item_type,
	item_id
);