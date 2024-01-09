package tes_baru_service

import (
	"context"
	"github.com/Muruyung/go-boilgen/pkg/logger"
	"github.com/Muruyung/go-boilgen/services/coba/domain/entity"
	"github.com/Muruyung/go-boilgen/services/coba/domain/service"
)

// CreateTesBaru create tes baru
func (svc *tesBaruInteractor) CreateTesBaru(ctx context.Context, dto service.DTOTesBaru) error {
	const commandName = "SVC-CREATE-TES-BARU"
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

	err = svc.repo.TesBaruRepo.Save(ctx, entityDTO)
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
		"Create tes baru success",
		nil,
	)
	return nil
}
