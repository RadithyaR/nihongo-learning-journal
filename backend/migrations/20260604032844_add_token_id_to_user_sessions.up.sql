ALTER TABLE user_sessions
ADD COLUMN token_id UUID NOT NULL DEFAULT gen_random_uuid();

CREATE UNIQUE INDEX idx_user_sessions_token_id
ON user_sessions(token_id);