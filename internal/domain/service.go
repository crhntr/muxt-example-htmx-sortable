package domain

import (
	"context"
	"log/slog"

	"github.com/crhntr/muxt-example-htmx-sortable/internal/database"
)

type (
	TransactionManager interface {
		ReadOnly(ctx context.Context, f database.ReadOnlyFunc) error
		UpdatePriorityList(ctx context.Context, f database.TaskPriorityUpdateFunc) error
	}
)

type Service struct {
	log *slog.Logger
	tx  TransactionManager
}

func New(log *slog.Logger, tx TransactionManager) *Service {
	return &Service{
		log: log,
		tx:  tx,
	}
}
