package domain_test

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/typelate/example-sortable/internal/database"
	"github.com/typelate/example-sortable/internal/domain"
	"github.com/typelate/example-sortable/internal/domain/internal/fake"
	"github.com/typelate/example-sortable/internal/hypertext"
)

func TestService_UpdateList(t *testing.T) {
	t.Run("reorder succeeds", func(t *testing.T) {
		qr := new(fake.TaskPriorityUpdater)
		tx := new(fake.TransactionManager)
		tx.UpdatePriorityListCalls(func(ctx context.Context, f database.TaskPriorityUpdateFunc) error {
			return f(qr)
		})
		svc := domain.New(discard(), tx)
		ctx := t.Context()

		result := svc.UpdateList(ctx, 32, hypertext.UpdateListForm{
			TaskIDs: []int32{1, 2, 3, 4},
		})

		require.NoError(t, result.Err)
		require.Equal(t, 1, tx.ReadOnlyCallCount())
		require.Equal(t, 4, qr.SetTaskPriorityCallCount())
		var gotParams []database.SetTaskPriorityParams
		for i := 0; i < qr.SetTaskPriorityCallCount(); i++ {
			_, ps := qr.SetTaskPriorityArgsForCall(i)
			gotParams = append(gotParams, ps)
		}
		assert.Equal(t, []database.SetTaskPriorityParams{
			{ID: 1, Priority: 4},
			{ID: 2, Priority: 3},
			{ID: 3, Priority: 2},
			{ID: 4, Priority: 1},
		}, gotParams)
	})
	t.Run("get list fails", func(t *testing.T) {
		qr := new(fake.TaskPriorityUpdater)
		tx := new(fake.TransactionManager)
		tx.UpdatePriorityListCalls(func(ctx context.Context, f database.TaskPriorityUpdateFunc) error {
			return f(qr)
		})
		svc := domain.New(discard(), tx)
		ctx := t.Context()
		qr.ListByIDReturns(database.List{}, errors.New("banana"))

		result := svc.UpdateList(ctx, 32, hypertext.UpdateListForm{
			TaskIDs: []int32{1, 2, 3, 4},
		})

		require.EqualError(t, result.Err, "banana")
	})
	t.Run("update task priority fails", func(t *testing.T) {
		qr := new(fake.TaskPriorityUpdater)
		roQuerier := new(fake.ReadOnlyQuerier)
		tx := new(fake.TransactionManager)
		tx.UpdatePriorityListCalls(func(ctx context.Context, f database.TaskPriorityUpdateFunc) error { return f(qr) })
		tx.ReadOnlyCalls(func(ctx context.Context, f database.ReadOnlyFunc) error { return f(roQuerier) })
		svc := domain.New(discard(), tx)
		ctx := t.Context()
		qr.ListByIDReturns(database.List{
			ID:   1,
			Name: "peach",
		}, nil)
		qr.SetTaskPriorityReturns(errors.New("banana"))
		roQuerier.TasksByListIDReturns([]database.TasksByListIDRow{
			{ID: 101, ListID: 2, PlanID: 3, Priority: 1, Instructions: "peach"},
			{ID: 102, ListID: 2, PlanID: 4, Priority: 2, Instructions: "pear"},
		}, nil)

		data := svc.UpdateList(ctx, 32, hypertext.UpdateListForm{
			TaskIDs: []int32{1, 2, 3, 4},
		})

		require.EqualError(t, data.Err, "failed to update task: banana")
		require.Equal(t, database.List{
			ID:   1,
			Name: "peach",
		}, data.List)
		require.Equal(t, 1, qr.ListByIDCallCount())
		require.Equal(t, 1, roQuerier.TasksByListIDCallCount())
		require.Equal(t, []database.TasksByListIDRow{
			{ID: 101, ListID: 2, PlanID: 3, Priority: 1, Instructions: "peach"},
			{ID: 102, ListID: 2, PlanID: 4, Priority: 2, Instructions: "pear"},
		}, data.Tasks)
	})
	t.Run("list rows fails", func(t *testing.T) {
		qr := new(fake.TaskPriorityUpdater)
		roQuerier := new(fake.ReadOnlyQuerier)
		tx := new(fake.TransactionManager)
		tx.UpdatePriorityListCalls(func(ctx context.Context, f database.TaskPriorityUpdateFunc) error { return f(qr) })
		tx.ReadOnlyCalls(func(ctx context.Context, f database.ReadOnlyFunc) error { return f(roQuerier) })
		svc := domain.New(discard(), tx)
		ctx := t.Context()
		qr.ListByIDReturns(database.List{
			ID:   1,
			Name: "peach",
		}, nil)

		data := svc.UpdateList(ctx, 32, hypertext.UpdateListForm{
			TaskIDs: []int32{1, 2, 3, 4},
		})

		require.Equal(t, 1, qr.ListByIDCallCount())
		require.Equal(t, database.List{
			ID:   1,
			Name: "peach",
		}, data.List)
		require.Equal(t, 1, qr.ListByIDCallCount())
		require.Equal(t, 1, roQuerier.TasksByListIDCallCount())
		require.Empty(t, data.Tasks)
	})
}

func TestService_Lists(t *testing.T) {
	t.Run("query succeeds", func(t *testing.T) {
		qr := new(fake.ReadOnlyQuerier)
		tx := new(fake.TransactionManager)
		svc := domain.New(discard(), tx)

		qr.ListsReturns([]database.List{
			{ID: 1, Name: "peach"},
			{ID: 2, Name: "pear"},
		}, nil)
		tx.ReadOnlyCalls(func(_ context.Context, f database.ReadOnlyFunc) error {
			return f(qr)
		})

		ctx := t.Context()

		data := svc.Lists(ctx)

		require.NotEmpty(t, data)
	})
	t.Run("query fails", func(t *testing.T) {
		qr := new(fake.ReadOnlyQuerier)
		tx := new(fake.TransactionManager)
		tx.ReadOnlyCalls(func(_ context.Context, f database.ReadOnlyFunc) error {
			return f(qr)
		})
		qr.ListsReturns(nil, errors.New("banana"))

		svc := domain.New(discard(), tx)

		ctx := t.Context()

		data := svc.Lists(ctx)

		require.Empty(t, data)
	})
}

func TestService_GetList(t *testing.T) {
	t.Run("queries succeed", func(t *testing.T) {
		qr := new(fake.ReadOnlyQuerier)
		tx := new(fake.TransactionManager)
		tx.ReadOnlyCalls(func(_ context.Context, f database.ReadOnlyFunc) error {
			return f(qr)
		})
		qr.ListByIDReturns(database.List{ID: 50, Name: "peach"}, nil)
		qr.TasksByListIDReturns([]database.TasksByListIDRow{
			{ID: 101, ListID: 2, PlanID: 3, Priority: 1, Instructions: "eat"},
		}, nil)

		svc := domain.New(discard(), tx)

		ctx := t.Context()

		data := svc.GetList(ctx, 42)

		require.NotZero(t, qr.ListByIDCallCount())
		_, listID := qr.ListByIDArgsForCall(0)
		require.EqualValues(t, 42, listID)

		require.Equal(t, hypertext.ListData{
			Err: nil,
			List: database.List{
				ID:   50,
				Name: "peach",
			},
			Tasks: []database.TasksByListIDRow{
				{ID: 101, ListID: 2, PlanID: 3, Priority: 1, Instructions: "eat"},
			},
		}, data)
	})
	t.Run("list query fails", func(t *testing.T) {
		qr := new(fake.ReadOnlyQuerier)
		tx := new(fake.TransactionManager)
		tx.ReadOnlyCalls(func(_ context.Context, f database.ReadOnlyFunc) error {
			return f(qr)
		})
		qr.TasksByListIDReturns(nil, errors.New("banana"))

		svc := domain.New(discard(), tx)

		ctx := t.Context()

		data := svc.GetList(ctx, 42)

		require.Equal(t, hypertext.ListData{
			Err:   errors.New("banana"),
			List:  database.List{},
			Tasks: nil,
		}, data)
	})
	t.Run("list tasks fails", func(t *testing.T) {
		qr := new(fake.ReadOnlyQuerier)
		tx := new(fake.TransactionManager)
		tx.ReadOnlyCalls(func(_ context.Context, f database.ReadOnlyFunc) error {
			return f(qr)
		})
		qr.ListByIDReturns(database.List{ID: 50, Name: "peach"}, nil)
		qr.TasksByListIDReturns(nil, errors.New("banana"))

		svc := domain.New(discard(), tx)

		ctx := t.Context()

		data := svc.GetList(ctx, 42)

		require.Equal(t, hypertext.ListData{
			Err:   errors.New("banana"),
			List:  database.List{ID: 50, Name: "peach"},
			Tasks: nil,
		}, data)
	})
	t.Run("query tasks fails with rows not found", func(t *testing.T) {
		qr := new(fake.ReadOnlyQuerier)
		tx := new(fake.TransactionManager)
		tx.ReadOnlyCalls(func(_ context.Context, f database.ReadOnlyFunc) error {
			return f(qr)
		})
		qr.TasksByListIDReturns(nil, fmt.Errorf("banana: %w", pgx.ErrNoRows))

		svc := domain.New(discard(), tx)

		ctx := t.Context()

		data := svc.GetList(ctx, 42)

		require.Equal(t, hypertext.ListData{
			Err:   fmt.Errorf("banana: %w", pgx.ErrNoRows),
			List:  database.List{},
			Tasks: nil,
		}, data)
	})
	t.Run("query list fails with rows not found", func(t *testing.T) {
		qr := new(fake.ReadOnlyQuerier)
		tx := new(fake.TransactionManager)
		tx.ReadOnlyCalls(func(_ context.Context, f database.ReadOnlyFunc) error {
			return f(qr)
		})
		qr.ListByIDReturns(database.List{}, fmt.Errorf("list %d not found", 42))
		qr.TasksByListIDReturns(nil, nil)

		svc := domain.New(discard(), tx)

		ctx := t.Context()

		data := svc.GetList(ctx, 42)

		require.Equal(t, hypertext.ListData{
			Err:   errors.New("list 42 not found"),
			List:  database.List{},
			Tasks: nil,
		}, data)
	})
}

func discard() *slog.Logger {
	return slog.New(slog.NewTextHandler(io.Discard, &slog.HandlerOptions{
		AddSource:   false,
		Level:       slog.LevelDebug,
		ReplaceAttr: nil,
	}))
}
