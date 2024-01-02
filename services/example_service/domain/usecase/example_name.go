package usecase

import (
	"context"
	"time"

	"github.com/Muruyung/go-boilgen/pkg/utils"
	"github.com/Muruyung/go-boilgen/services/example_service/domain/entity"
)

// DTOExampleName dto for example name usecase
type DTOExampleName struct {
	StartDate time.Time
	Name      string
	Status    int
	IsActive  bool
	TestID    int64
}

// ExampleNameUseCase example name usecase template
type ExampleNameUseCase interface {
	ExampleCustomMethod(ctx context.Context, exampleParam string) (int, error)
	ExampleCustomMethodDefault(ctx context.Context) error
	GetExampleNameByID(ctx context.Context, id int) (*entity.ExampleName, error)
	GetListExampleName(ctx context.Context, request *utils.RequestOption) ([]*entity.ExampleName, *utils.MetaResponse, error)
	CreateExampleName(ctx context.Context, dto DTOExampleName) error
	UpdateExampleName(ctx context.Context, id int, dto DTOExampleName) error
	DeleteExampleName(ctx context.Context, id int) error
}
