package sqlstore

import (
	"context"

	"github.com/ZombieMInd/slovodel/internal/game"
)

type WordRepository struct {
	store *Store
}

func (r *WordRepository) Update(w *game.Word) error {
	ctx := context.TODO()

	tx, err := r.store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.ExecContext(ctx, updateWord,
		w.Value,
		w.Points,
		w.ID,
	)
	if err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *WordRepository) Delete(id int) error {
	ctx := context.TODO()

	tx, err := r.store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.ExecContext(ctx, deletePlayer,
		id,
	)
	if err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}
