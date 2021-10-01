package sqlstore

import (
	"context"
	"database/sql"

	"github.com/ZombieMInd/slovodel/internal/game"
)

type GameRepository struct {
	store *Store
}

func (r *GameRepository) Create(g *game.Game) (int, error) {
	ctx := context.TODO()

	tx, err := r.store.db.BeginTx(ctx, nil)
	if err != nil {
		return -1, err
	}
	defer tx.Rollback()

	res, err := tx.ExecContext(ctx, insertGame,
		g.Name,
	)
	if err != nil {
		return -1, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return -1, err
	}

	for _, p := range g.Players {
		CheckPlayer(ctx, tx, p.ID)
		_, err = tx.ExecContext(ctx, insertGamePlayer,
			g.ID,
			p.ID,
		)
		if err != nil {
			return -1, err
		}
	}

	if err = tx.Commit(); err != nil {
		return -1, err
	}

	return int(id), nil
}

func (r *GameRepository) Get(id int) (*game.Game, error) {
	ctx := context.TODO()
	g := game.Game{}

	tx, err := r.store.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	res, err := tx.QueryContext(ctx, selectGame,
		id,
	)
	if err != nil {
		return nil, err
	}
	err = res.Scan(&g.ID, &g.Name)
	if err != nil {
		return nil, err
	}

	g.Players, err = r.selectGamePlayers(ctx, tx, id)
	if err != nil {
		return nil, err
	}

	g.Words, err = r.selectGameWords(ctx, tx, id)
	if err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return &g, nil
}

func (r *GameRepository) Update(g *game.Game) error {
	ctx := context.TODO()

	tx, err := r.store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.ExecContext(ctx, checkGame,
		g.ID,
	)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, updateGame,
		g.Name,
		g.ID,
	)
	if err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *GameRepository) AddWord(g *game.Game, w *game.Word) error {
	ctx := context.TODO()

	tx, err := r.store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.ExecContext(ctx, checkGame,
		g.ID,
	)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, insertWord,
		w.Value,
		w.Points,
		w.Player.ID,
		w.Game.ID,
	)
	if err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *GameRepository) AddPlayer(g *game.Game, p *game.Player) error {
	ctx := context.TODO()

	tx, err := r.store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.ExecContext(ctx, checkGame,
		g.ID,
	)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, checkPlayer,
		p.ID,
	)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, insertGamePlayer,
		g.ID,
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

func (r *GameRepository) GetAll(offset, limit int) ([]*game.Game, error) {
	ctx := context.TODO()
	games := []*game.Game{}

	tx, err := r.store.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	rows, err := tx.QueryContext(ctx, selectAllGames,
		offset,
		limit,
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		g := game.Game{}
		if err := rows.Scan(&g.ID, &g.Name); err != nil {
			return nil, err
		}

		games = append(games, &g)
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return games, nil
}

func (r *GameRepository) Delete(g *game.Game) error {
	ctx := context.TODO()

	tx, err := r.store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.ExecContext(ctx, deleteGame,
		g.ID,
	)
	if err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func CheckPlayer(ctx context.Context, tx *sql.Tx, id int) error {
	res, err := tx.QueryContext(ctx, checkPlayer, id)
	if err != nil {
		return err
	}
	resultPlayer := struct {
		ID int
	}{}
	res.Scan(&resultPlayer)
	if resultPlayer.ID != id {
		return err
	}
	return nil
}

func (r *GameRepository) selectGamePlayers(ctx context.Context, tx *sql.Tx, id int) ([]*game.Player, error) {
	playersRows, err := tx.QueryContext(ctx, selectPlayersByGame,
		id,
	)
	if err != nil {
		return nil, err
	}

	defer playersRows.Close()
	players := []*game.Player{}

	for playersRows.Next() {
		p := game.Player{}
		if err := playersRows.Scan(&p.ID, &p.Name); err != nil {
			return players, err
		}

		players = append(players, &p)
	}

	return players, nil
}

func (r *GameRepository) selectGameWords(ctx context.Context, tx *sql.Tx, id int) ([]*game.Word, error) {
	wordsRows, err := tx.QueryContext(ctx, selectWordsByGame,
		id,
	)
	if err != nil {
		return nil, err
	}

	defer wordsRows.Close()
	words := []*game.Word{}

	for wordsRows.Next() {
		w := game.Word{}
		if err := wordsRows.Scan(&w.ID, &w.Value, &w.Points, &w.Player.ID); err != nil {
			return words, err
		}
		w.Game.ID = id

		words = append(words, &w)
	}

	return words, nil
}
