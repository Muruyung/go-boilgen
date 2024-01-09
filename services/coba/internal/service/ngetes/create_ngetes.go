package ngetes_service

import (
	"context"
	"github.com/Muruyung/go-boilgen/pkg/logger"
	"github.com/Muruyung/go-boilgen/services/coba/domain/entity"
	"github.com/Muruyung/go-boilgen/services/coba/domain/service"
)

// CreateNgetes create ngetes
func (svc *ngetesInteractor) CreateNgetes(ctx context.Context, dto service.DTONgetes) error {
	const commandName = "SVC-CREATE-NGETES"
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

	err = svc.repo.NgetesRepo.Save(ctx, entityDTO)
	if err != nil {
		logger.DetailLoggerError(
			ctx,
			commandName,
			"Error create",
			err,
		)
		return err
	}

	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Create ngetes success",
		nil,
	)
	return nil
}
