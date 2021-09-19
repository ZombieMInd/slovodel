package game

type MockGameRepository struct {
	games []*Game
}

func NewMockGameRepository() *MockGameRepository {
	return &MockGameRepository{games: make([]*Game, 0)}
}

func (r *MockGameRepository) Create(g *Game) error {
	r.games = append(r.games, g)
	return nil
}

func (r *MockGameRepository) Update(game *Game) error {
	for _, g := range r.games {
		if g.ID == game.ID {
			g = game
			return nil
		}
	}
	return ErrGameNotFound
}

func (r *MockGameRepository) Get(id int64) (*Game, error) {
	for _, g := range r.games {
		if g.ID == id {
			return g, nil
		}
	}
	return nil, ErrGameNotFound
}

func (r *MockGameRepository) GetByName(name string) (*Game, error) {
	for _, g := range r.games {
		if g.Name == name {
			return g, nil
		}
	}
	return nil, ErrGameNotFound
}

func (r *MockGameRepository) Delete(name string) error {
	for i, g := range r.games {
		if g.Name == name {
			r.games[i] = r.games[len(r.games)-1]
			r.games = r.games[:len(r.games)-1]
			return nil
		}
	}
	return ErrGameNotFound
}

func (r *MockGameRepository) GetAll() ([]*Game, error) {
	return r.games, nil
}
