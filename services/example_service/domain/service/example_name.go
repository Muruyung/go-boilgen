package service

import (
	"context"
	"time"

	"github.com/Muruyung/go-boilgen/pkg/utils"
	"github.com/Muruyung/go-boilgen/services/example_service/domain/entity"
)

// DTOExampleName dto for example name service
type DTOExampleName struct {
	Name      string
	Status    int
	IsActive  bool
	TestID    int64
	StartDate time.Time
}

// ExampleNameService example name service template
type ExampleNameService interface {
	ExampleCustomMethod(ctx context.Context, exampleParam string) (int, error)
	ExampleCustomMethodDefault(ctx context.Context) error
	GetExampleNameByID(ctx context.Context, id int) (*entity.ExampleName, error)
	GetListExampleName(ctx context.Context, request *utils.RequestOption) ([]*entity.ExampleName, *utils.MetaResponse, error)
	CreateExampleName(ctx context.Context, dto DTOExampleName) error
	UpdateExampleName(ctx context.Context, id int, dto DTOExampleName) error
	DeleteExampleName(ctx context.Context, id int) error
}
