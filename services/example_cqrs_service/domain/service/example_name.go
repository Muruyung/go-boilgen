package service

import (
	"context"
	"github.com/Muruyung/go-boilgen/services/example_cqrs_service/domain/entity"
	goutils "github.com/Muruyung/go-utilities"
	"time"
)

// DTOExampleName dto for example name service
type DTOExampleName struct {
	StartDate time.Time
	Name      string
	Status    int
	IsActive  bool
	TestID    int64
}

// ExampleNameService example name service template
type ExampleNameService interface {
	ExampleCustomQueryMethod(ctx context.Context, exampleParam string) (int, error)
	ExampleCustomCreateCommandDefault(ctx context.Context) error
	GetExampleNameByID(ctx context.Context, id int) (*entity.ExampleName, error)
	GetListExampleName(ctx context.Context, request *goutils.RequestOption) ([]*entity.ExampleName, *goutils.MetaResponse, error)
	CreateExampleName(ctx context.Context, dto DTOExampleName) error
	UpdateExampleName(ctx context.Context, id int, dto DTOExampleName) error
	DeleteExampleName(ctx context.Context, id int) error
}
