DROP INDEX IF EXISTS uq_user_word;

ALTER TABLE vocabularies
ADD CONSTRAINT uq_user_word UNIQUE(user_id, word);

DROP INDEX IF EXISTS uq_user_kanji;

ALTER TABLE kanjis
ADD CONSTRAINT uq_user_kanji UNIQUE(user_id, character);
