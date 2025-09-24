package domain

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"

	"github.com/typelate/example-sortable/internal/database"
	"github.com/typelate/example-sortable/internal/hypertext"
)

func (svc *Service) Lists(ctx context.Context) []database.List {
	var result []database.List
	_ = svc.tx.ReadOnly(ctx, func(db database.ReadOnlyQuerier) error {
		lists, err := db.Lists(ctx)
		result = lists
		return err
	})
	return result
}

func (svc *Service) GetList(ctx context.Context, id int32) hypertext.ListData {
	var data hypertext.ListData
	data.Err = svc.tx.ReadOnly(ctx, func(db database.ReadOnlyQuerier) error {
		list, err := db.ListByID(ctx, id)
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				return fmt.Errorf("list %d not found", id)
			}
			return err
		}
		data.List = list
		tasks, err := db.TasksByListID(ctx, id)
		if err != nil {
			return err
		}
		data.Tasks = tasks
		return nil
	})
	return data
}

func (svc *Service) UpdateList(ctx context.Context, listID int32, form hypertext.UpdateListForm) hypertext.ListData {
	svc.log.Debug("updating list", "list_id", listID, "task_ids", form.TaskIDs)
	var data hypertext.ListData
	data.Err = svc.tx.UpdatePriorityList(ctx, func(db database.TaskPriorityUpdater) error {
		list, err := db.ListByID(ctx, listID)
		if err != nil {
			return err
		}
		data.List = list
		for index, taskID := range form.TaskIDs {
			if err := db.SetTaskPriority(ctx, database.SetTaskPriorityParams{
				Priority: int32(len(form.TaskIDs) - index),
				ID:       taskID,
			}); err != nil {
				return fmt.Errorf("failed to update task: %w", err)
			}
		}
		return nil
	})
	err := svc.tx.ReadOnly(ctx, func(db database.ReadOnlyQuerier) error {
		tasks, err := db.TasksByListID(ctx, listID)
		data.Tasks = tasks
		return err
	})
	if data.Err != nil {
		data.Err = errors.Join(data.Err, err)
	}
	return data
}
