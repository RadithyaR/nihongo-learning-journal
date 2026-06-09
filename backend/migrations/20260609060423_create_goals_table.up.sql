CREATE TABLE goals (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    user_id UUID NOT NULL,

    title VARCHAR(255) NOT NULL,

    description TEXT NULL,

    goal_type VARCHAR(20) NULL,

    target_level VARCHAR(10) NULL,

    target_count INTEGER NULL,

    target_date DATE NOT NULL,

    status VARCHAR(20) NOT NULL DEFAULT 'IN_PROGRESS',

    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP NULL,

    CONSTRAINT fk_goals_user
        FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE
);

CREATE INDEX idx_goals_user_id
ON goals(user_id);

CREATE INDEX idx_goals_status
ON goals(status);

CREATE INDEX idx_goals_target_date
ON goals(target_date);