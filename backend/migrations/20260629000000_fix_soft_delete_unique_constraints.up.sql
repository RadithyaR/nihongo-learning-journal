DROP INDEX IF EXISTS uq_user_word;

CREATE UNIQUE INDEX uq_user_word
ON vocabularies(user_id, word)
WHERE deleted_at IS NULL;

DROP INDEX IF EXISTS uq_user_kanji;

CREATE UNIQUE INDEX uq_user_kanji
ON kanjis(user_id, character)
WHERE deleted_at IS NULL;
