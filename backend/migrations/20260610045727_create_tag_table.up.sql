CREATE TABLE tags (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    user_id UUID NOT NULL,

    name VARCHAR(100) NOT NULL,

    color VARCHAR(20),

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP NULL,

    CONSTRAINT fk_tags_user
        FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE
);

CREATE INDEX idx_tags_user_id
ON tags(user_id);

CREATE INDEX idx_tags_deleted_at
ON tags(deleted_at);