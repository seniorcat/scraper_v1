-- +goose Up
-- +goose StatementBegin
CREATE TABLE "content"."recept" (
    id  INTEGER PRIMARY KEY,
    name TEXT,
    description TEXT,
    cooking_time TEXT,
    number_servings TEXT,
    image_src TEXT);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "content"."recept";
-- +goose StatementEnd