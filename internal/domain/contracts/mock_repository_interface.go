package contracts

import "github.com/andreluizmicro/go-mocker-api/internal/domain/entity"

type MockRepositoryInterface interface {
	Create(entity.Mock) (*string, error)
	Find(sessionId string) ([]byte, error)
}
