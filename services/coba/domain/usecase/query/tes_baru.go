package query

import (
	"context"
	"github.com/Muruyung/go-boilgen/pkg/utils"
	"github.com/Muruyung/go-boilgen/services/coba/domain/entity"
)

// TesBaruUseCase tes baru query usecase template
type TesBaruUseCase interface {
	GetTesBaruByID(ctx context.Context, id string) (*entity.TesBaru, error)
	GetListTesBaru(ctx context.Context, request *utils.RequestOption) ([]*entity.TesBaru, *utils.MetaResponse, error)
}
