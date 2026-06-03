CREATE TABLE vocabulary_meanings (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    vocabulary_id UUID NOT NULL,

    meaning TEXT NOT NULL,

    order_number INTEGER NOT NULL,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),

    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_vocabulary_meanings_vocabulary
        FOREIGN KEY(vocabulary_id)
        REFERENCES vocabularies(id)
        ON DELETE CASCADE
);

CREATE INDEX idx_vocabulary_meanings_vocabulary_id
ON vocabulary_meanings(vocabulary_id);