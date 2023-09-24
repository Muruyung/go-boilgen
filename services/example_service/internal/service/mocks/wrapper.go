package mocks

import "github.com/golang/mock/gomock"

type Wrapper struct {
	*MockExampleNameUseCase
}

func Init(ctrl *gomock.Controller) Wrapper {
	return Wrapper{
		MockExampleNameUseCase: NewMockExampleNameUseCase(ctrl),
	}
}
