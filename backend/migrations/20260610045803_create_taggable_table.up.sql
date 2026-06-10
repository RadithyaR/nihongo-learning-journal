CREATE TABLE taggables (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    tag_id UUID NOT NULL,

    item_type VARCHAR(20) NOT NULL,

    item_id UUID NOT NULL,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP NULL,

    CONSTRAINT fk_taggables_tag
        FOREIGN KEY (tag_id)
        REFERENCES tags(id)
        ON DELETE CASCADE
);

CREATE INDEX idx_taggables_tag_id
ON taggables(tag_id);

CREATE INDEX idx_taggables_item
ON taggables(item_type, item_id);

CREATE INDEX idx_taggables_deleted_at
ON taggables(deleted_at);

CREATE UNIQUE INDEX uq_taggables_unique
ON taggables(
    tag_id,
    item_type,
    item_id
);