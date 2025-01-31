-- name: Lists :many
-- interface: ReadOnlyQuerier
SELECT * FROM list;

-- name: TasksByListID :many
-- interface: ReadOnlyQuerier
SELECT task.*, plan.instructions FROM task JOIN plan ON plan.id = plan_id WHERE list_id = sqlc.arg(id) ORDER BY priority DESC;

-- name: ListByID :one
-- interface: ReadOnlyQuerier TaskPriorityUpdater
SELECT * FROM list WHERE id = sqlc.arg(id);

-- name: SetTaskPriority :exec
-- interface: TaskPriorityUpdater
UPDATE task SET priority = sqlc.arg(priority) WHERE id = sqlc.arg(id);
