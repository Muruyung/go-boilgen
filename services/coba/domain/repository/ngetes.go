package repository

import (
	"context"
	"github.com/Muruyung/go-boilgen/pkg/utils"
	"github.com/Muruyung/go-boilgen/services/coba/domain/entity"
)

// NgetesRepository ngetes repository template
type NgetesRepository interface {
	Get(ctx context.Context, query utils.QueryBuilderInteractor) (*entity.Ngetes, error)
	GetList(ctx context.Context, query utils.QueryBuilderInteractor) ([]*entity.Ngetes, error)
	GetCount(ctx context.Context, query utils.QueryBuilderInteractor) (int, error)
	Save(ctx context.Context, data *entity.Ngetes) error
	Update(ctx context.Context, id string, data *entity.Ngetes) error
}
