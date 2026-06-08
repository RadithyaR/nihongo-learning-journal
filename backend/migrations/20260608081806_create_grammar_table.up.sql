CREATE TABLE grammars (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    user_id UUID NOT NULL,

    pattern VARCHAR(255) NOT NULL,

    meaning TEXT,

    example TEXT,

    note TEXT,

    favourite BOOLEAN NOT NULL DEFAULT FALSE,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),

    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),

    deleted_at TIMESTAMP NULL,

    CONSTRAINT fk_grammars_user
        FOREIGN KEY(user_id)
        REFERENCES users(id)
        ON DELETE CASCADE
);

CREATE INDEX idx_grammars_user_id
ON grammars(user_id);

CREATE INDEX idx_grammars_deleted_at
ON grammars(deleted_at);