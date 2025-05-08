package tx

import (
	"context"

	"cinema/pkg/ent"
	"cinema/pkg/logging"

	"go.uber.org/zap"
)

type Tx interface {
	Client() *ent.Client
	OnRollback(f ent.RollbackHook)
	OnCommit(f ent.CommitHook)
}

func WithTransaction(ctx context.Context, client *ent.Client, fn Fn) error {
	tx, err := client.Tx(ctx)
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			// a panic occurred, rollback and repanic
			if err := tx.Rollback(); err != nil {
				logging.Logger(ctx).Error("can not rollback transaction", zap.Error(err))
			}
			panic(p)
		} else if err != nil {
			// something went wrong, rollback
			if err := tx.Rollback(); err != nil {
				logging.Logger(ctx).Error("can not rollback transaction", zap.Error(err))
			}
		} else {
			// all good, commit
			if err := tx.Commit(); err != nil {
				logging.Logger(ctx).Error("can not commit transaction", zap.Error(err))
			}
		}
	}()
	err = fn(ctx, tx)

	return err
}

type Fn func(context.Context, Tx) error
