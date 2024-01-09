package repository

import (
	"context"
	"github.com/Muruyung/go-boilgen/pkg/utils"
	"github.com/Muruyung/go-boilgen/services/coba/domain/entity"
)

// TesBaruRepository tes baru repository template
type TesBaruRepository interface {
	Get(ctx context.Context, query utils.QueryBuilderInteractor) (*entity.TesBaru, error)
	GetList(ctx context.Context, query utils.QueryBuilderInteractor) ([]*entity.TesBaru, error)
	GetCount(ctx context.Context, query utils.QueryBuilderInteractor) (int, error)
	Save(ctx context.Context, data *entity.TesBaru) error
	Update(ctx context.Context, id string, data *entity.TesBaru) error
	Delete(ctx context.Context, id string) error
}
