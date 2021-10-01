package game

type PlayerRepository interface {
	Create(*Player) (int, error)
	GetAll(offset, limit int) ([]*Player, error)
	Get(int) (*Player, error)
	Update(*Player) error
	Delete(*Player) error
}

type PlayerService struct {
	repository PlayerRepository
}

func NewPlayerService(r PlayerRepository) *PlayerService {
	return &PlayerService{repository: r}
}

func (s *PlayerService) Create(name string) (int, error) {
	return s.repository.Create(&Player{Name: name})
}

func (s *PlayerService) ListAll(offset, limit int) ([]*Player, error) {
	return s.repository.GetAll(offset, limit)
}

func (s *PlayerService) Get(id int) (*Player, error) {
	return s.repository.Get(id)
}

func (s *PlayerService) Edit(p *Player) error {
	_, err := s.repository.Get(p.ID)
	if err != nil {
		return err
	}
	return s.repository.Update(p)
}

func (s *PlayerService) Delete(id int) error {
	p, err := s.repository.Get(id)
	if err != nil {
		return err
	}
	return s.repository.Delete(p)
}
