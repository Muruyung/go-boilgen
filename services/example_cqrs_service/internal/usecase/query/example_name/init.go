package example_name_usecase

import (
	"github.com/Muruyung/go-boilgen/services/example_cqrs_service/domain/service"
	"github.com/Muruyung/go-boilgen/services/example_cqrs_service/domain/usecase/query"
)

type exampleNameInteractor struct {
	*service.Wrapper
}

// NewExampleNameUseCase initialize new example name use case
func NewExampleNameUseCase(svc *service.Wrapper) query.ExampleNameUseCase {
	return &exampleNameInteractor{
		Wrapper: svc,
	}
}
