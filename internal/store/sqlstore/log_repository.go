package sqlstore

import (
	"context"
	"fmt"

	"github.com/ZombieMInd/slovodel/internal/logger"
	"github.com/google/uuid"
)

type LogRepository struct {
	store *Store
}

func (r *LogRepository) Save(l *logger.LogRequest) error {
	ctx := context.Background()

	tx, err := r.store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.ExecContext(ctx, InsertLog,
		l.UUID,
		l.IP,
		l.UserUUID,
		l.Timestamp,
	)
	if err != nil {
		fmt.Println(err)
		return err
	}

	for _, e := range l.Events {
		_, err = tx.ExecContext(ctx, InsertEvent,
			uuid.New(),
			e.EventType,
			e.EventText,
			l.UUID,
		)
		if err != nil {
			return err
		}
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}
