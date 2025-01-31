-- +goose Up
-- +goose StatementBegin
CREATE TABLE list (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE plan (
    id SERIAL PRIMARY KEY,
    instructions TEXT NOT NULL
);

CREATE TABLE task (
  id SERIAL PRIMARY KEY,
  list_id SERIAL NOT NULL REFERENCES list(id),
  plan_id INTEGER NOT NULL REFERENCES plan(id),
  priority INTEGER NOT NULL DEFAULT 0,
  CONSTRAINT unique_list_priority UNIQUE (list_id, priority) DEFERRABLE INITIALLY IMMEDIATE
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS task;
DROP TABLE IF EXISTS plan;
DROP TABLE IF EXISTS list;
-- +goose StatementEnd
