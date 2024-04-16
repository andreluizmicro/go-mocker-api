package application

import (
	"github.com/andreluizmicro/go-mocker-api/internal/domain/contracts"
	"github.com/andreluizmicro/go-mocker-api/internal/domain/entity"
)

type MockService struct {
	mockRepository contracts.MockRepositoryInterface
}

func NewMockService(mockRepository contracts.MockRepositoryInterface) *MockService {
	return &MockService{
		mockRepository: mockRepository,
	}
}

func (s *MockService) Create(payload []byte) (*string, error) {
	mock := entity.NewMock(payload)
	sessiontId, err := s.mockRepository.Create(*mock)
	if err != nil {
		return nil, err
	}
	return sessiontId, nil
}

func (s *MockService) Find(sessionId string) ([]byte, error) {
	data, err := s.mockRepository.Find(sessionId)
	if err != nil {
		return nil, err
	}
	return data, nil
}
