CREATE TABLE vocabularies (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    word VARCHAR(255) NOT NULL,
    reading VARCHAR(255),
    source VARCHAR(255),
    note TEXT,
    status VARCHAR(20) NOT NULL DEFAULT 'NEW',
    favourite BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP NULL,
    CONSTRAINT fk_vocabularies_user
        FOREIGN KEY(user_id)
        REFERENCES users(id)
        ON DELETE CASCADE,
    CONSTRAINT uq_user_word
        UNIQUE(user_id, word)
);

CREATE INDEX idx_vocabularies_user_id
ON vocabularies(user_id);

CREATE INDEX idx_vocabularies_status
ON vocabularies(status);

CREATE INDEX idx_vocabularies_deleted_at
ON vocabularies(deleted_at);