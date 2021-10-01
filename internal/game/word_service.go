package game

type WordRepository interface {
	Update(*Word) error
	Delete(int) error
}

type WordService struct {
	repository WordRepository
}

func NewWordService(r WordRepository) *WordService {
	return &WordService{repository: r}
}

func (s *WordService) Edit(w *Word) error {
	return s.repository.Update(w)
}

func (s *WordService) Delete(id int) error {
	return s.repository.Delete(id)
}
