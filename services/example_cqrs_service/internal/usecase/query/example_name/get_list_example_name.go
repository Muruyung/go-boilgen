package example_name_usecase

import (
	"context"

	"github.com/Muruyung/go-boilgen/pkg/utils"
	"github.com/Muruyung/go-boilgen/services/example_cqrs_service/domain/entity"
	"github.com/Muruyung/go-utilities/logger"
)

// GetListExampleName get list example name
func (uc *exampleNameInteractor) GetListExampleName(ctx context.Context, request *utils.RequestOption) ([]*entity.ExampleName, *utils.MetaResponse, error) {
	const commandName = "UC-GET-LIST-EXAMPLE-NAME"
	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Get list example name process...",
		nil,
	)

	res, metaRes, err := uc.ExampleNameSvc.GetListExampleName(ctx, request)
	if err != nil {
		logger.DetailLoggerError(
			ctx,
			commandName,
			"Error get list",
			err,
		)
		return nil, nil, err
	}

	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Get list example name success",
		nil,
	)
	return res, metaRes, nil
}
