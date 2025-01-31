


-- +goose Up
-- +goose StatementBegin
INSERT INTO list (name) VALUES('kitchen'), ('car');
INSERT INTO plan (instructions) VALUES('clean sink'), ('sweep floors'), ('empty trunk'), ('clean windshield'), ('stare into fridge');
INSERT into task (list_id, plan_id, priority) VALUES
  (1, 1, 3),
  (1, 2, 2),
  (1, 5, 1),
  (2, 3, 2),
  (2, 4, 1);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DELETE FROM list WHERE id IN (1, 2);
DELETE FROM plan WHERE id IN (1, 2, 3, 4, 5);
DELETE FROM task WHERE id IN (1, 2, 3, 4, 5);
-- +goose StatementEnd