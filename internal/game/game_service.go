package game

type GameRepository interface {
	Create(*Game) (int, error)
	Get(int) (*Game, error)
	Update(*Game) error
	GetAll() ([]*Game, error)
	Delete(*Game) error
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
	g, err := s.repository.Get(gameID)
	if err != nil {
		return err
	}
	g.Players = append(g.Players, player)

	err = s.repository.Update(g)
	if err != nil {
		return err
	}
	return nil
}

func (s *GameService) ListAll() ([]*Game, error) {
	return s.repository.GetAll()
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
