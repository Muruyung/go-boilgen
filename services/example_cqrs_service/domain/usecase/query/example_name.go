package query

import (
	"context"
	"github.com/Muruyung/go-boilgen/services/example_cqrs_service/domain/entity"
	goutils "github.com/Muruyung/go-utilities"
)

// ExampleNameUseCase example name query usecase template
type ExampleNameUseCase interface {
	ExampleCustomQueryMethod(ctx context.Context, exampleParam string) (int, error)
	GetExampleNameByID(ctx context.Context, id int) (*entity.ExampleName, error)
	GetListExampleName(ctx context.Context, request *goutils.RequestOption) ([]*entity.ExampleName, *goutils.MetaResponse, error)
}
