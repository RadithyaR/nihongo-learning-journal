CREATE TABLE review_logs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    user_id UUID NOT NULL,

    vocabulary_id UUID NOT NULL,

    rating VARCHAR(20) NOT NULL,

    reviewed_at TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_review_logs_user
        FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE,

    CONSTRAINT fk_review_logs_vocabulary
        FOREIGN KEY (vocabulary_id)
        REFERENCES vocabularies(id)
        ON DELETE CASCADE
);

CREATE INDEX idx_review_logs_user_id
ON review_logs(user_id);

CREATE INDEX idx_review_logs_vocabulary_id
ON review_logs(vocabulary_id);

CREATE INDEX idx_review_logs_reviewed_at
ON review_logs(reviewed_at);