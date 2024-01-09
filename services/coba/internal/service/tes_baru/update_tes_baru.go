package tes_baru_service

import (
	"context"
	"github.com/Muruyung/go-boilgen/pkg/logger"
	"github.com/Muruyung/go-boilgen/services/coba/domain/entity"
	"github.com/Muruyung/go-boilgen/services/coba/domain/service"
)

// UpdateTesBaru update tes baru
func (svc *tesBaruInteractor) UpdateTesBaru(ctx context.Context, id string, dto service.DTOTesBaru) error {
	const commandName = "SVC-UPDATE-TES-BARU"
	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Create tes baru process...",
		nil,
	)

	entityDTO, err := entity.NewTesBaru(entity.DTOTesBaru{
		Apacing: dto.Apacing,
		Test:    dto.Test,
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

	err = svc.repo.TesBaruRepo.Update(ctx, id, entityDTO)
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
		"Update tes baru success",
		nil,
	)
	return nil
}
