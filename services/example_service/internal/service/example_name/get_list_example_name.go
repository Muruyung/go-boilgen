package example_name_service

import (
	"context"

	"github.com/Muruyung/go-boilgen/pkg/utils"
	"github.com/Muruyung/go-boilgen/services/example_service/domain/entity"
	"github.com/Muruyung/go-utilities/logger"
)

// GetListExampleName get list example name
func (svc *exampleNameInteractor) GetListExampleName(ctx context.Context, request *utils.RequestOption) ([]*entity.ExampleName, *utils.MetaResponse, error) {
	const commandName = "SVC-GET-LIST-EXAMPLE-NAME"
	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Get list example name process...",
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

	res, err := svc.repo.ExampleNameRepo.GetList(ctx, query)
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
		totalCount, err := svc.repo.ExampleNameRepo.GetCount(ctx, queryPagination)
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
		"Get list example name success",
		nil,
	)
	return res, metaRes, nil
}
