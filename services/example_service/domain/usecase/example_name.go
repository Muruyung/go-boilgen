package usecase

import (
	"context"
	"github.com/Muruyung/go-boilgen/services/example_service/domain/entity"
	goutils "github.com/Muruyung/go-utilities"
	"time"
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
	GetExampleNameByID(ctx context.Context, id int) (*entity.ExampleName, error)
	GetListExampleName(ctx context.Context, request *goutils.RequestOption) ([]*entity.ExampleName, *goutils.MetaResponse, error)
	CreateExampleName(ctx context.Context, dto DTOExampleName) error
	UpdateExampleName(ctx context.Context, id int, dto DTOExampleName) error
	DeleteExampleName(ctx context.Context, id int) error
}
