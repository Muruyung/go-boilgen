package service

import (
	"context"
	"github.com/Muruyung/go-boilgen/services/example_service/domain/entity"
	goutils "github.com/Muruyung/go-utilities"
	"time"
)

// DTOExampleName dto for example name service
type DTOExampleName struct {
	Status    int
	IsActive  bool
	TestID    int64
	StartDate time.Time
	Name      string
}

// ExampleNameService example name service template
type ExampleNameService interface {
	ExampleCustomMethodDefault(ctx context.Context) error
	ExampleCustomMethod(ctx context.Context, exampleParam string) (int, error)
	GetExampleNameByID(ctx context.Context, id int) (*entity.ExampleName, error)
	GetListExampleName(ctx context.Context, request *goutils.RequestOption) ([]*entity.ExampleName, *goutils.MetaResponse, error)
	CreateExampleName(ctx context.Context, dto DTOExampleName) error
	UpdateExampleName(ctx context.Context, id int, dto DTOExampleName) error
	DeleteExampleName(ctx context.Context, id int) error
}
