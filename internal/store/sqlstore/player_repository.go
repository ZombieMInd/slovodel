package sqlstore

import (
	"context"
	"database/sql"

	"github.com/ZombieMInd/slovodel/internal/game"
)

type PlayerRepository struct {
	store *Store
}

func (r *PlayerRepository) Create(p *game.Player) (int, error) {
	ctx := context.TODO()

	tx, err := r.store.db.BeginTx(ctx, nil)
	if err != nil {
		return -1, err
	}
	defer tx.Rollback()

	res, err := tx.ExecContext(ctx, insertPlayer,
		p.Name,
	)
	if err != nil {
		return -1, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return -1, err
	}

	if err = tx.Commit(); err != nil {
		return -1, err
	}

	return int(id), nil
}

func (r *PlayerRepository) GetAll(offset, limit int) ([]*game.Player, error) {
	ctx := context.TODO()
	players := []*game.Player{}

	tx, err := r.store.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	rows, err := tx.QueryContext(ctx, selectAllPlayers,
		offset,
		limit,
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		p := game.Player{}
		if err := rows.Scan(&p.ID, &p.Name); err != nil {
			return nil, err
		}

		players = append(players, &p)
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return players, nil
}

func (r *PlayerRepository) Get(id int) (*game.Player, error) {
	ctx := context.TODO()
	p := game.Player{}

	tx, err := r.store.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	rows, err := tx.QueryContext(ctx, selectPlayer,
		id,
	)
	if err != nil {
		return nil, err
	}

	err = rows.Scan(&p.Name)
	if err != nil {
		return nil, err
	}

	p.Words, err = r.selectPlayerWords(ctx, tx, id)
	if err != nil {
		return nil, err
	}

	p.Games, err = r.selectPlayerGames(ctx, tx, id)
	if err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return &p, nil
}

func (r *PlayerRepository) Update(p *game.Player) error {
	ctx := context.TODO()

	tx, err := r.store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.ExecContext(ctx, checkPlayer,
		p.ID,
	)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, updatePlayer,
		p.Name,
		p.ID,
	)
	if err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *PlayerRepository) Delete(p *game.Player) error {
	ctx := context.TODO()

	tx, err := r.store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.ExecContext(ctx, deletePlayer,
		p.ID,
	)
	if err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *PlayerRepository) selectPlayerWords(ctx context.Context, tx *sql.Tx, id int) ([]*game.Word, error) {
	wordsRows, err := tx.QueryContext(ctx, selectWordsByPlayer,
		id,
	)
	if err != nil {
		return nil, err
	}

	defer wordsRows.Close()
	words := []*game.Word{}

	for wordsRows.Next() {
		w := game.Word{}
		if err := wordsRows.Scan(&w.ID, &w.Value, &w.Points, &w.Game.ID); err != nil {
			return words, err
		}
		w.Player.ID = id

		words = append(words, &w)
	}

	return words, nil
}

func (r *PlayerRepository) selectPlayerGames(ctx context.Context, tx *sql.Tx, id int) ([]*game.Game, error) {
	gameRows, err := tx.QueryContext(ctx, selectGamesByPlayer,
		id,
	)
	if err != nil {
		return nil, err
	}

	defer gameRows.Close()
	games := []*game.Game{}

	for gameRows.Next() {
		g := game.Game{}
		if err := gameRows.Scan(&g.ID, &g.Name); err != nil {
			return nil, err
		}

		games = append(games, &g)
	}

	return games, nil
}
