package tes_baru_usecase

import (
	"context"

	"github.com/Muruyung/go-boilgen/services/coba/domain/service"
)

// DeleteTesBaru update tes baru
func (uc *tesBaruInteractor) DeleteTesBaru(ctx context.Context, id string) error {
	return uc.tx.BeginTx(ctx, func(ctx context.Context, svc *service.Wrapper) error {
		return svc.TesBaruSvc.DeleteTesBaru(ctx, id)
	})
}
