package example_name_usecase

import (
	"context"
	"fmt"
	"github.com/Muruyung/go-boilgen/services/example_cqrs_service/domain/entity"
	"github.com/Muruyung/go-utilities/logger"
)

// GetExampleNameByID get example name by id
func (uc *exampleNameInteractor) GetExampleNameByID(ctx context.Context, id int) (*entity.ExampleName, error) {
	const commandName = "UC-GET-EXAMPLE-NAME-BY-ID"
	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Get example name process...",
		nil,
	)

	res, err := uc.ExampleNameSvc.GetExampleNameByID(ctx, id)
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
		"Get example name success",
		nil,
	)
	return res, nil
}
