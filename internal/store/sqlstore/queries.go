package sqlstore

const (
	insertGame       string = `INSERT INTO game (name) VALUES ($1) RETURNING id;`
	insertGamePlayer string = `INSERT INTO game_player (game_id, player_id) VALUES ($1, $2);`
	selectGame       string = `SELECT * FROM game WHERE id=$1;`
	checkGame        string = `SELECT id FROM game WHERE id=$1 LIMIT 1;`
	updateGame       string = `UPDATE game SET name=$1 WHERE id=$2;`
	selectAllGames   string = `SELECT * FROM game LIMIT $1, $2`
	deleteGame       string = `DELETE FROM game WHERE id=$1;`
)

const (
	insertPlayer        string = `INSERT INTO player (name) VALUES ($1) RETURNING id;`
	checkPlayer         string = `SELECT id FROM player WHERE id=$1 LIMIT 1;`
	selectPlayersByGame string = `SELECT player.* FROM player 
		JOIN game_player ON game_player.payer_id = player.id 
		WHERE game_player.game_id = $1;`
	selectGamesByPlayer string = `SELECT game.* FROM game 
		JOIN game_player ON game_player.game_id = game.id 
		WHERE game_player.player_id = $1;`
	selectAllPlayers string = `SELECT * FROM player LIMIT $1, $2;`
	selectPlayer     string = `SELECT * FROM player WHERE id=$1;`
	updatePlayer     string = `UPDATE player SET name=$1 WHERE id=$2;`
	deletePlayer     string = `DELETE FROM game WHERE id=$1;`
)

const (
	selectWordsByGame   string = `SELECT id, value, points, player_id FROM word WHERE game_id = $1;`
	insertWord          string = `INSERT INTO word (value, points, player_id, game_id) VALUES ($1, $2, $3, $4);`
	selectWordsByPlayer string = `SELECT id, value, points, game_id FROM word WHERE player_id = $1;`
	selectAllWords      string = `SELECT * FROM player LIMIT $1, $2;`
	updateWord          string = `UPDATE word SET value=$1, points=$2 WHERE id=$3;`
	deleteWord          string = `DELETE FROM word WHERE id=$1;`
)
