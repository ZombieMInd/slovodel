package game

type Game struct {
	ID      int64    `json:"id"`
	Name    string   `json:"name"`
	Players []Player `json:"players"`
	Words   []Word   `json:"words"`
}

type Word struct {
	ID    int64  `json:"id"`
	Text  string `json:"text"`
	Value string `json:"value"`
}

type Player struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
