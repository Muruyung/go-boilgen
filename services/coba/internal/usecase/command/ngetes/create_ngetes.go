package ngetes_usecase

import (
	"context"

	"github.com/Muruyung/go-boilgen/services/coba/domain/service"
	"github.com/Muruyung/go-boilgen/services/coba/domain/usecase/command"
)

// CreateNgetes create ngetes
func (uc *ngetesInteractor) CreateNgetes(ctx context.Context, dto command.DTONgetes) error {
	return uc.tx.BeginTx(ctx, func(ctx context.Context, svc *service.Wrapper) error {
		return svc.NgetesSvc.CreateNgetes(ctx, service.DTONgetes{
			Name: dto.Name,
		})
	})
}
