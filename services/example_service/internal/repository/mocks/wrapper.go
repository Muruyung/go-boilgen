package mocks

import "github.com/golang/mock/gomock"

type Wrapper struct {
	*MockExampleNameRepository
	*MockModelsCommon
}

func Init(ctrl *gomock.Controller) Wrapper {
	return Wrapper{
		MockModelsCommon:          NewMockModelsCommon(ctrl),
		MockExampleNameRepository: NewMockExampleNameRepository(ctrl),
	}
}
