package entity

type Mock struct {
	Payload []byte
}

func NewMock(payload []byte) *Mock {
	return &Mock{
		Payload: payload,
	}
}
