-- +goose Up
CREATE INDEX good_category_idx on Good (id_goodCategory);
-- +goose Down
DROP INDEX good_category_idx;
