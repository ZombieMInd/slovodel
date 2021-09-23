package store

import (
	"fmt"

	"github.com/ZombieMInd/slovodel/internal/logger"
)

type MockLogRepository struct{}

func NewMockLogRepository() *MockLogRepository {
	return &MockLogRepository{}
}

func (r *MockLogRepository) Save(*logger.LogRequest) error {
	fmt.Println("Success!")
	return nil
}
