package tes_baru_usecase

import (
	"context"

	"github.com/Muruyung/go-boilgen/services/coba/domain/service"
	"github.com/Muruyung/go-boilgen/services/coba/domain/usecase/command"
)

// CreateTesBaru create tes baru
func (uc *tesBaruInteractor) CreateTesBaru(ctx context.Context, dto command.DTOTesBaru) error {
	return uc.tx.BeginTx(ctx, func(ctx context.Context, svc *service.Wrapper) error {
		return svc.TesBaruSvc.CreateTesBaru(ctx, service.DTOTesBaru{
			Apacing: dto.Apacing,
			Test:    dto.Test,
		})
	})
}
