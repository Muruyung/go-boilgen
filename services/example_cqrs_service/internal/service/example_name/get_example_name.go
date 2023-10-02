package example_name_service

import (
	"context"
	"fmt"
	"github.com/Muruyung/go-boilgen/services/example_cqrs_service/domain/entity"
	goutils "github.com/Muruyung/go-utilities"
	"github.com/Muruyung/go-utilities/logger"
)

// GetExampleNameByID get example name by id
func (svc *exampleNameInteractor) GetExampleNameByID(ctx context.Context, id int) (*entity.ExampleName, error) {
	const commandName = "SVC-GET-EXAMPLE-NAME-BY-ID"
	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Get example name process...",
		nil,
	)

	var query = goutils.NewQueryBuilder()
	query.AddWhere("id", "=", "int")
	res, err := svc.repo.ExampleNameRepo.Get(ctx, query)
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
