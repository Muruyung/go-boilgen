package tes_baru_service

import (
	"context"
	"fmt"
	"github.com/Muruyung/go-boilgen/pkg/logger"
	"github.com/Muruyung/go-boilgen/pkg/utils"
	"github.com/Muruyung/go-boilgen/services/coba/domain/entity"
)

// GetTesBaruByID get tes baru by id
func (svc *tesBaruInteractor) GetTesBaruByID(ctx context.Context, id string) (*entity.TesBaru, error) {
	const commandName = "SVC-GET-TES-BARU-BY-ID"
	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Get tes baru process...",
		nil,
	)

	var query = utils.NewQueryBuilder()
	query.AddWhere("id", "=", "string")
	res, err := svc.repo.TesBaruRepo.Get(ctx, query)
	if err != nil {
		logger.DetailLoggerError(
			ctx,
			commandName,
			fmt.Sprintf("Error get by id=%v", id),
			err,
		)
		return nil, err
	}

	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Get tes baru success",
		nil,
	)
	return res, nil
}
