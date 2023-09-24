package example_name

import (
	"github.com/Muruyung/go-boilgen/services/example_service/domain/repository"
	"github.com/Muruyung/go-boilgen/services/example_service/domain/service"
)

type exampleNameInteractor struct {
	repo *repository.Wrapper
}

// NewExampleNameService initialize new example name service
func NewExampleNameService(repo *repository.Wrapper) service.ExampleNameService {
	return &exampleNameInteractor{
		repo: repo,
	}
}
