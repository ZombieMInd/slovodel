CREATE TABLE game_player (
    id INT PRIMARY KEY,
    game_id INT FORIGN KEY REFERENCES game(id),
    player_id INT FORIGN KEY REFERENCES player(id)
)