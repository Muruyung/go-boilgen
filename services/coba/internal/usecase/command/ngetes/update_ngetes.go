package ngetes_usecase

import (
	"context"

	"github.com/Muruyung/go-boilgen/services/coba/domain/service"
	"github.com/Muruyung/go-boilgen/services/coba/domain/usecase/command"
)

// UpdateNgetes update ngetes
func (uc *ngetesInteractor) UpdateNgetes(ctx context.Context, id string, dto command.DTONgetes) error {
	return uc.tx.BeginTx(ctx, func(ctx context.Context, svc *service.Wrapper) error {
		return svc.NgetesSvc.UpdateNgetes(ctx, id, service.DTONgetes{
			Name: dto.Name,
		})
	})
}
