package ngetes_service

import (
	"context"
	"fmt"
	"github.com/Muruyung/go-boilgen/pkg/logger"
	"github.com/Muruyung/go-boilgen/pkg/utils"
	"github.com/Muruyung/go-boilgen/services/coba/domain/entity"
)

// GetNgetesByID get ngetes by id
func (svc *ngetesInteractor) GetNgetesByID(ctx context.Context, id string) (*entity.Ngetes, error) {
	const commandName = "SVC-GET-NGETES-BY-ID"
	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Get ngetes process...",
		nil,
	)

	var query = utils.NewQueryBuilder()
	query.AddWhere("id", "=", "string")
	res, err := svc.repo.NgetesRepo.Get(ctx, query)
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
		"Get ngetes success",
		nil,
	)
	return res, nil
}
