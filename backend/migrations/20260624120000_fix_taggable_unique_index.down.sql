DROP INDEX IF EXISTS uq_taggables_unique;

CREATE UNIQUE INDEX uq_taggables_unique
ON taggables(tag_id, item_type, item_id);
