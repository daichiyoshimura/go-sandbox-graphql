package db

import (
	"context"
	"fmt"
	"sandbox-gql/ent"
)

type TxFunc[T any] func(tx *ent.Tx) (T, error)

func RunInTransaction[T any](ctx context.Context, client *ent.Client, fn TxFunc[T]) (T, error) {
	var zero T
	tx, err := client.Tx(ctx)
	if err != nil {
		return zero, err
	}

	result, err := fn(tx)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return zero, fmt.Errorf("failed to rollback transaction: %v, original error: %v", rollbackErr, err)
		}
		return zero, err
	}

	if err := tx.Commit(); err != nil {
		return zero, fmt.Errorf("failed to commit transaction: %v", err)
	}

	return result, nil
}
