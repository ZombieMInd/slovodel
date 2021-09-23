package store

import (
	"github.com/ZombieMInd/slovodel/internal/logger"
)

type Store interface {
	Log() logger.LogRepository
}
