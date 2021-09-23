package redisstore

import (
	"encoding/json"
	"fmt"

	"github.com/ZombieMInd/slovodel/internal/logger"
)

type LogRepository struct {
	store *Store
}

func (r *LogRepository) Save(l *logger.LogRequest) error {
	key := fmt.Sprintf("%s:%s", l.UserUUID, l.UUID.String())

	jsonLog, err := json.Marshal(l)
	if err != nil {
		return err
	}

	err = r.store.client.Set(key, jsonLog, 0).Err()
	if err != nil {
		return err
	}
	return nil
}
