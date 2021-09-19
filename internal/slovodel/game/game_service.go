package game

type GameRepository interface {
	Create(*Game) error
	Update(*Game) error
	Get(int64) (*Game, error)
	GetByName(string) (*Game, error)
	Delete(string) error
	GetAll() ([]*Game, error)
}

type GameService struct {
	repository GameRepository
}

func NewGameService(r GameRepository) *GameService {
	return &GameService{repository: r}
}

func (s *GameService) Create(g *Game) error {
	return s.repository.Create(g)
}

func (s *GameService) Update(g *Game) error {
	return s.repository.Update(g)
}

func (s *GameService) Get(id int64) (*Game, error) {
	return s.repository.Get(id)
}

func (s *GameService) GetByName(name string) (*Game, error) {
	return s.repository.GetByName(name)
}

func (s *GameService) Delete(name string) error {
	return s.repository.Delete(name)
}

func (s *GameService) GetAll() ([]*Game, error) {
	return s.repository.GetAll()
}
