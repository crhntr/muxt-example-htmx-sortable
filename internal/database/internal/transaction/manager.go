package transaction

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type Func func(ctx context.Context, tx pgx.Tx) error

type TxBeginner interface {
	BeginTx(ctx context.Context, options pgx.TxOptions) (pgx.Tx, error)
}

type Manager struct {
	conn TxBeginner
}

func NewManager(conn TxBeginner) *Manager {
	return &Manager{conn: conn}
}

func (t Manager) Call(ctx context.Context, options pgx.TxOptions, f Func) error {
	var recoverErr error
	return errors.Join(func() error {
		tx, err := t.conn.BeginTx(ctx, options)
		if err != nil {
			return err
		}
		defer func() {
			if r := recover(); r != nil {
				err, ok := r.(error)
				if !ok {
					err = fmt.Errorf("recovered from panic: %v", r)
				}
				recoverErr = errors.Join(err, tx.Rollback(ctx))
			}
		}()
		if err := f(ctx, tx); err != nil {
			return errors.Join(err, tx.Rollback(ctx))
		}
		return tx.Commit(ctx)
	}(), recoverErr)
}
