package tes_baru_usecase

import (
	"context"

	"github.com/Muruyung/go-boilgen/services/coba/domain/service"
	"github.com/Muruyung/go-boilgen/services/coba/domain/usecase/command"
)

// UpdateTesBaru update tes baru
func (uc *tesBaruInteractor) UpdateTesBaru(ctx context.Context, id string, dto command.DTOTesBaru) error {
	return uc.tx.BeginTx(ctx, func(ctx context.Context, svc *service.Wrapper) error {
		return svc.TesBaruSvc.UpdateTesBaru(ctx, id, service.DTOTesBaru{
			Apacing: dto.Apacing,
			Test:    dto.Test,
		})
	})
}
