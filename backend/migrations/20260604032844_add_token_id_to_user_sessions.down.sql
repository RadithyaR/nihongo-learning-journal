DROP INDEX IF EXISTS idx_user_sessions_token_id;

ALTER TABLE user_sessions
DROP COLUMN token_id;