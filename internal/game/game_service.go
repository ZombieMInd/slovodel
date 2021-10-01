package game

type GameRepository interface {
	Create(*Game) (int, error)
	Get(int) (*Game, error)
	Update(*Game) error
	GetAll(offset, limit int) ([]*Game, error)
	Delete(*Game) error
	AddWord(*Game, *Word) error
	AddPlayer(*Game, *Player) error
}

type GameService struct {
	repository GameRepository
}

func NewGameService(r GameRepository) *GameService {
	return &GameService{repository: r}
}

func (s *GameService) Create(name string) (id int, err error) {
	return s.repository.Create(&Game{Name: name})
}

func (s *GameService) CreateWithPlayers(g *Game) (id int, err error) {
	return s.repository.Create(g)
}

func (s *GameService) AddPlayer(gameID int, player *Player) error {
	err := s.repository.AddPlayer(&Game{ID: gameID}, player)
	if err != nil {
		return err
	}
	return nil
}

func (s *GameService) AddWord(gameID int, word *Word) error {
	err := s.repository.AddWord(&Game{ID: gameID}, word)
	if err != nil {
		return err
	}
	return nil
}

func (s *GameService) ListAll(offset, limit int) ([]*Game, error) {
	return s.repository.GetAll(offset, limit)
}

func (s *GameService) Get(id int) (*Game, error) {
	g, err := s.repository.Get(id)
	if err != nil {
		return nil, err
	}

	g.Result, err = s.getResult(g)
	if err != nil {
		return nil, err
	}
	return g, nil
}

func (s *GameService) GetResultFromID(id int) ([]*PlayerResult, error) {
	g, err := s.repository.Get(id)
	if err != nil {
		return nil, err
	}

	return s.getResult(g)
}

func (s *GameService) getResult(g *Game) ([]*PlayerResult, error) {
	var result []*PlayerResult

	for _, player := range g.Players {
		playerRes := &PlayerResult{Player: player, TotalPoints: 0}
		for _, word := range g.Words {
			if word.Player == player {
				playerRes.TotalPoints += word.Points
			}
		}
		result = append(result, playerRes)
	}

	return result, nil
}

func (s *GameService) Edit(g *Game) error {
	_, err := s.repository.Get(g.ID)
	if err != nil {
		return err
	}

	return s.repository.Update(g)
}

func (s *GameService) Delete(id int) error {
	g, err := s.repository.Get(id)
	if err != nil {
		return err
	}

	return s.repository.Delete(g)
}
