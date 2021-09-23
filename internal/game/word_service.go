package game

type WordRepository interface {
	Create(*Word) (int, error)
	GetAll() ([]*Word, error)
	Get(int) (*Word, error)
	Update(*Word) error
	Delete(*Word) error
}

type WordService struct {
	repository WordRepository
}

func NewWordService(r WordRepository) *WordService {
	return &WordService{repository: r}
}

func (s *WordService) Create(w *Word) (int, error) {
	return s.repository.Create(w)
}

func (s *WordService) ListAll() ([]*Word, error) {
	return s.repository.GetAll()
}

func (s *WordService) Edit(w *Word) error {
	_, err := s.repository.Get(w.ID)
	if err != nil {
		return err
	}

	return s.repository.Update(w)
}

func (s *WordService) Delete(id int) error {
	w, err := s.repository.Get(id)
	if err != nil {
		return err
	}

	return s.repository.Delete(w)
}
