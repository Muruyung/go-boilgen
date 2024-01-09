package ngetes_service

import (
	"context"
	"github.com/Muruyung/go-boilgen/pkg/logger"
	"github.com/Muruyung/go-boilgen/services/coba/domain/entity"
	"github.com/Muruyung/go-boilgen/services/coba/domain/service"
)

// UpdateNgetes update ngetes
func (svc *ngetesInteractor) UpdateNgetes(ctx context.Context, id string, dto service.DTONgetes) error {
	const commandName = "SVC-UPDATE-NGETES"
	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Create ngetes process...",
		nil,
	)

	entityDTO, err := entity.NewNgetes(entity.DTONgetes{
		Name: dto.Name,
	})
	if err != nil {
		logger.DetailLoggerError(
			ctx,
			commandName,
			"Error generate entity",
			err,
		)
		return err
	}

	err = svc.repo.NgetesRepo.Update(ctx, id, entityDTO)
	if err != nil {
		logger.DetailLoggerError(
			ctx,
			commandName,
			"Error update",
			err,
		)
		return err
	}

	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Update ngetes success",
		nil,
	)
	return nil
}
