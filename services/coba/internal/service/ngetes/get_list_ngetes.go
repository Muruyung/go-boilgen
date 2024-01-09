package ngetes_service

import (
	"context"
	"github.com/Muruyung/go-boilgen/pkg/logger"
	"github.com/Muruyung/go-boilgen/pkg/utils"
	"github.com/Muruyung/go-boilgen/services/coba/domain/entity"
)

// GetListNgetes get list ngetes
func (svc *ngetesInteractor) GetListNgetes(ctx context.Context, request *utils.RequestOption) ([]*entity.Ngetes, *utils.MetaResponse, error) {
	const commandName = "SVC-GET-LIST-NGETES"
	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Get list ngetes process...",
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

	res, err := svc.repo.NgetesRepo.GetList(ctx, query)
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
		totalCount, err := svc.repo.NgetesRepo.GetCount(ctx, queryPagination)
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
		"Get list ngetes success",
		nil,
	)
	return res, metaRes, nil
}
