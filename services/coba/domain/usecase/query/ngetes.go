package query

import (
	"context"
	"github.com/Muruyung/go-boilgen/pkg/utils"
	"github.com/Muruyung/go-boilgen/services/coba/domain/entity"
)

// NgetesUseCase ngetes query usecase template
type NgetesUseCase interface {
	GetNgetesByID(ctx context.Context, id string) (*entity.Ngetes, error)
	GetListNgetes(ctx context.Context, request *utils.RequestOption) ([]*entity.Ngetes, *utils.MetaResponse, error)
}
