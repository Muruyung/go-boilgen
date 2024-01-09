package tes_baru_usecase

import (
	"context"

	"github.com/Muruyung/go-boilgen/pkg/utils"
	"github.com/Muruyung/go-boilgen/services/coba/domain/entity"
)

// GetListTesBaru get list tes baru
func (uc *tesBaruInteractor) GetListTesBaru(ctx context.Context, request *utils.RequestOption) ([]*entity.TesBaru, *utils.MetaResponse, error) {
	return uc.svc.TesBaruSvc.GetListTesBaru(ctx, request)
}
