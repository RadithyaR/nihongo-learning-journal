DROP TABLE IF EXISTS review_logs;

CREATE TABLE review_logs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    user_id UUID NOT NULL,

    item_type VARCHAR(20) NOT NULL,

    item_id UUID NOT NULL,

    rating VARCHAR(20) NOT NULL,

    reviewed_at TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_review_logs_user
        FOREIGN KEY(user_id)
        REFERENCES users(id)
        ON DELETE CASCADE
);

CREATE INDEX idx_review_logs_user_id
ON review_logs(user_id);

CREATE INDEX idx_review_logs_item_type
ON review_logs(item_type);

CREATE INDEX idx_review_logs_item_id
ON review_logs(item_id);

CREATE INDEX idx_review_logs_reviewed_at
ON review_logs(reviewed_at);