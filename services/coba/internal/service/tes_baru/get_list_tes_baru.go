package tes_baru_service

import (
	"context"
	"github.com/Muruyung/go-boilgen/pkg/logger"
	"github.com/Muruyung/go-boilgen/pkg/utils"
	"github.com/Muruyung/go-boilgen/services/coba/domain/entity"
)

// GetListTesBaru get list tes baru
func (svc *tesBaruInteractor) GetListTesBaru(ctx context.Context, request *utils.RequestOption) ([]*entity.TesBaru, *utils.MetaResponse, error) {
	const commandName = "SVC-GET-LIST-TES-BARU"
	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Get list tes baru process...",
		nil,
	)

	var (
		query           = utils.NewQueryBuilder()
		queryPagination = utils.NewQueryBuilder()
		metaRes         *utils.MetaResponse
		page            int
		limit           int
	)

	if request != nil {
		query, page, limit = request.SetPaginationWithSort(query)
	}

	res, err := svc.repo.TesBaruRepo.GetList(ctx, query)
	if err != nil {
		logger.DetailLoggerError(
			ctx,
			commandName,
			"Error get list",
			err,
		)
		return nil, nil, err
	}

	if request != nil && request.GetPagination() != nil {
		totalCount, err := svc.repo.TesBaruRepo.GetCount(ctx, queryPagination)
		if err != nil {
			logger.DetailLoggerError(
				ctx,
				commandName,
				"Error get total count list",

				err,
			)
			return nil, nil, err
		}

		var meta = utils.MapMetaResponse(totalCount, len(res), page, limit)
		metaRes = &meta
	}

	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Get list tes baru success",
		nil,
	)
	return res, metaRes, nil
}
