CREATE TABLE study_sessions (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL,

    session_date DATE NOT NULL,

    notes TEXT,
    reflection TEXT,

    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP NULL,

    CONSTRAINT fk_study_sessions_user
        FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE
);

CREATE UNIQUE INDEX idx_study_sessions_user_date
ON study_sessions(user_id, session_date);

CREATE INDEX idx_study_sessions_user_id
ON study_sessions(user_id);

CREATE TABLE study_session_items (
    id UUID PRIMARY KEY,

    study_session_id UUID NOT NULL,

    item_type VARCHAR(50) NOT NULL,

    item_id UUID NOT NULL,

    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP NULL,

    CONSTRAINT fk_study_session_items_session
        FOREIGN KEY (study_session_id)
        REFERENCES study_sessions(id)
        ON DELETE CASCADE
);

CREATE INDEX idx_study_session_items_session_id
ON study_session_items(study_session_id);

CREATE INDEX idx_study_session_items_item
ON study_session_items(item_type, item_id);