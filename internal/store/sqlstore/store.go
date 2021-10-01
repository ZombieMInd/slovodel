package sqlstore

import (
	"context"
	"database/sql"

	"github.com/ZombieMInd/slovodel/internal/game"
)

type DB interface {
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
}

type Store struct {
	db               DB
	gameRepository   game.GameRepository
	playerRepository game.PlayerRepository
	wordRepository   game.WordRepository
}

func New(db DB) *Store {
	return &Store{db: db}
}

func (s *Store) Game() game.GameRepository {
	if s.gameRepository == nil {
		s.gameRepository = &GameRepository{store: s}
	}

	return s.gameRepository
}

func (s *Store) Player() game.PlayerRepository {
	if s.playerRepository == nil {
		s.playerRepository = &PlayerRepository{store: s}
	}

	return s.playerRepository
}

func (s *Store) Word() game.WordRepository {
	if s.wordRepository == nil {
		s.wordRepository = &WordRepository{store: s}
	}

	return s.wordRepository
}
