CREATE TABLE word (
    id INT PRIMARY KEY,
    value VARCHAR(64),
    points INT,
    game_id INT FORIGN KEY REFERENCES game(id),
    player_id INT FORIGN KEY REFERENCES player(id)
)