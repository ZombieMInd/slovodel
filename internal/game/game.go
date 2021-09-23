package game

type Game struct {
	ID      int             `json:"id"`
	Name    string          `json:"name"`
	Players []*Player       `json:"players"`
	Words   []*Word         `json:"words"`
	Result  []*PlayerResult `json:"result"`
}

type Player struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Words []*Word `json:"words"`
	Games []*Game `json:"games"`
}

type Word struct {
	ID     int     `json:"id"`
	Value  string  `json:"value"`
	Points int     `json:"points"`
	Game   *Game   `json:"game"`
	Player *Player `json:"player"`
}

type PlayerResult struct {
	Player      *Player `json:"player"`
	TotalPoints int     `json:"total_points"`
}
