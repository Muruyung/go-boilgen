package ngetes_usecase

import (
	"context"

	"github.com/Muruyung/go-boilgen/pkg/utils"
	"github.com/Muruyung/go-boilgen/services/coba/domain/entity"
)

// GetListNgetes get list ngetes
func (uc *ngetesInteractor) GetListNgetes(ctx context.Context, request *utils.RequestOption) ([]*entity.Ngetes, *utils.MetaResponse, error) {
	return uc.svc.NgetesSvc.GetListNgetes(ctx, request)
}
