package command

import (
	"context"
	"time"
)

// DTOExampleName dto for example name usecase
type DTOExampleName struct {
	Name      string
	Status    int
	IsActive  bool
	TestID    int64
	StartDate time.Time
}

// ExampleNameUseCase example name command usecase template
type ExampleNameUseCase interface {
	ExampleCustomCreateCommandDefault(ctx context.Context) error
	CreateExampleName(ctx context.Context, dto DTOExampleName) error
	UpdateExampleName(ctx context.Context, id int, dto DTOExampleName) error
	DeleteExampleName(ctx context.Context, id int) error
}
