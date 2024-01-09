package tes_baru_usecase

import (
	"context"

	"github.com/Muruyung/go-boilgen/services/coba/domain/entity"
)

// GetTesBaruByID get tes baru by id
func (uc *tesBaruInteractor) GetTesBaruByID(ctx context.Context, id string) (*entity.TesBaru, error) {
	return uc.svc.TesBaruSvc.GetTesBaruByID(ctx, id)
}
