package repository

import (
	"context"
	"github.com/Muruyung/go-boilgen/services/example_service/domain/entity"
	goutils "github.com/Muruyung/go-utilities"
)

// ExampleNameRepository example name repository wrapper
type ExampleNameRepository interface {
	ExampleCustomMethod(ctx context.Context, query goutils.QueryBuilderInteractor) (int, error)
	Get(ctx context.Context, query goutils.QueryBuilderInteractor) (*entity.ExampleName, error)
	GetList(ctx context.Context, query goutils.QueryBuilderInteractor) ([]*entity.ExampleName, error)
	GetCount(ctx context.Context, query goutils.QueryBuilderInteractor) (int, error)
	Save(ctx context.Context, data *entity.ExampleName) error
	Update(ctx context.Context, id int, data *entity.ExampleName) error
	Delete(ctx context.Context, id int) error
}
