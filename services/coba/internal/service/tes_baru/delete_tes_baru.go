package tes_baru_service

import (
	"context"
	"fmt"
	"github.com/Muruyung/go-boilgen/pkg/logger"
)

// DeleteTesBaru update tes baru
func (svc *tesBaruInteractor) DeleteTesBaru(ctx context.Context, id string) error {
	const commandName = "SVC-DELETE-TES-BARU"
	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Delete tes baru process...",
		nil,
	)

	err := svc.repo.TesBaruRepo.Delete(ctx, id)
	if err != nil {
		logger.DetailLoggerError(
			ctx,
			commandName,
			fmt.Sprintf("Error delete by id=%v", id),
			err,
		)
		return err
	}

	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Delete tes baru success",
		nil,
	)
	return nil
}
