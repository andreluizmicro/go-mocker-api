package repository

import (
	"fmt"
	"os"

	"github.com/andreluizmicro/go-mocker-api/internal/domain/entity"
	"github.com/andreluizmicro/go-mocker-api/internal/domain/errors"
)

type MockRepository struct {
	sessionId string
}

func NewMockRepository(sessionId string) *MockRepository {
	return &MockRepository{
		sessionId: sessionId,
	}
}

func (r *MockRepository) Create(mock entity.Mock) (*string, error) {
	_ = os.Mkdir(fmt.Sprintf("./tmp/%s/%s", "mock", r.sessionId), 0777)

	path := fmt.Sprintf("./tmp/mock/%s/%s", r.sessionId, "mock.json")
	file, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	file.Write(mock.Payload)

	return &r.sessionId, nil
}

func (r *MockRepository) Find(sessionId string) ([]byte, error) {
	file, err := os.ReadFile(fmt.Sprintf("./tmp/mock/%s/mock.json", sessionId))
	if err != nil {
		fmt.Println(err)
		return nil, errors.ErrMockNotFound
	}

	return file, err
}
