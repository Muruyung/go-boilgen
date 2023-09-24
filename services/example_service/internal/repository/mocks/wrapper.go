package mocks

import "github.com/golang/mock/gomock"

type Wrapper struct {
	*MockExampleNameRepository
}

func Init(ctrl *gomock.Controller) Wrapper {
	return Wrapper{
		MockExampleNameRepository: NewMockExampleNameRepository(ctrl),
	}
}
