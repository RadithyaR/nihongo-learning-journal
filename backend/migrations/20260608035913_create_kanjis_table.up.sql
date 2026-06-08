CREATE TABLE kanjis (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    user_id UUID NOT NULL,

    character VARCHAR(10) NOT NULL,

    meaning TEXT,

    onyomi VARCHAR(255),

    kunyomi VARCHAR(255),

    jlpt_level VARCHAR(10),

    favourite BOOLEAN NOT NULL DEFAULT FALSE,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP NULL,

    CONSTRAINT fk_kanjis_user
        FOREIGN KEY(user_id)
        REFERENCES users(id)
        ON DELETE CASCADE,

    CONSTRAINT uq_user_kanji
        UNIQUE(user_id, character)
);