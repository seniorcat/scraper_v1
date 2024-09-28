-- +goose Up
-- +goose StatementBegin

CREATE TABLE "content"."category" (
    slug TEXT PRIMARY KEY,
    name TEXT,
    href TEXT,
    parent_slug TEXT
);
ALTER TABLE ONLY "content"."category"
    ADD CONSTRAINT category_fk_parent_slug FOREIGN KEY (parent_slug) 
    REFERENCES "content"."category"(slug) DEFERRABLE INITIALLY DEFERRED;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "content"."category";
-- +goose StatementEnd
