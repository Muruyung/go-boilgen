package example_name

import (
	"github.com/Muruyung/go-boilgen/services/example_service/domain/service"
	"github.com/Muruyung/go-boilgen/services/example_service/domain/usecase"
)

type exampleNameInteractor struct {
	*service.Wrapper
}

// NewExampleNameUseCase initialize new example name use case
func NewExampleNameUseCase(svc *service.Wrapper) usecase.ExampleNameUseCase {
	return &exampleNameInteractor{
		Wrapper: svc,
	}
}
