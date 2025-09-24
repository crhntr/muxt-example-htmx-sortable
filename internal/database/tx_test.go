package database_test

import (
	"context"
	"testing"

	"github.com/crhntr/transaction"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/require"

	"github.com/typelate/example-sortable/internal/database"
	"github.com/typelate/example-sortable/internal/database/internal/fake"
)

func TestTransactions_ReadOnly(t *testing.T) {
	t.Run("options", func(t *testing.T) {
		var (
			ctx = context.TODO()
			m   = new(fake.TransactionManager)
			tx  = new(fake.Tx)
		)
		m.CallCalls(func(ctx context.Context, options pgx.TxOptions, t transaction.Func) error { return t(ctx, tx) })
		txs := database.NewTransactionsWithCaller(m)

		_ = txs.ReadOnly(ctx, func(querier database.ReadOnlyQuerier) error { return nil })

		require.Equal(t, 1, m.CallCallCount())
		_, o, _ := m.CallArgsForCall(0)
		require.Equal(t, pgx.ReadOnly, o.AccessMode)
	})
}

func TestTransactions_UpdatePriorityList(t *testing.T) {
	t.Run("options", func(t *testing.T) {
		var (
			ctx = context.TODO()
			m   = new(fake.TransactionManager)
			tx  = new(fake.Tx)
		)
		m.CallCalls(func(ctx context.Context, options pgx.TxOptions, t transaction.Func) error { return t(ctx, tx) })
		txs := database.NewTransactionsWithCaller(m)

		_ = txs.UpdatePriorityList(ctx, func(updater database.TaskPriorityUpdater) error { return nil })

		require.Equal(t, 1, m.CallCallCount())
		_, o, _ := m.CallArgsForCall(0)
		require.Equal(t, pgx.ReadWrite, o.AccessMode)
	})
}
