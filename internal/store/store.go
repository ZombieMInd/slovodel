package store

import "github.com/ZombieMInd/slovodel/internal/game"

type Store interface {
	Game() game.GameRepository
	Player() game.PlayerRepository
	Word() game.WordRepository
}
