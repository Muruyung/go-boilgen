package example_name_repo

import (
	"context"
	"fmt"

	"github.com/Muruyung/go-utilities/logger"
	"github.com/rocketlaunchr/dbq/v2"

	"github.com/Muruyung/go-boilgen/pkg/utils"
	"github.com/Muruyung/go-boilgen/services/example_service/domain/entity"
	"github.com/Muruyung/go-boilgen/services/example_service/internal/repository/mapper"
	"github.com/Muruyung/go-boilgen/services/example_service/internal/repository/models"
)

// GetList get list data example name
func (db *mysqlExampleNameRepository) GetList(ctx context.Context, query utils.QueryBuilderInteractor) ([]*entity.ExampleName, error) {
	const commandName = "REPO-GET-LIST-EXAMPLE-NAME"
	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Get list example name process...",
		nil,
	)

	var (
		err       error
		tableName = models.ExampleNameModels{}.GetTableName()
		data      = make([]interface{}, 0)
	)

	query.AddWhere("deleted_at", "=", nil)
	stmt, val, _ := query.GetQuery(tableName, "")
	opts := &dbq.Options{
		SingleResult:   false,
		ConcreteStruct: models.ExampleNameModels{},
		DecoderConfig:  dbq.StdTimeConversionConfig(),
	}

	result, err := dbq.Q(ctx, db.mysql, stmt, opts, val...)
	if err != nil {
		logger.DetailLoggerError(
			ctx,
			commandName,
			"Error execute query",
			err,
		)
		return nil, err
	}

	exampleNameModels := result.([]interface{})
	if len(exampleNameModels) == 0 {
		err = fmt.Errorf("example name list data not found")
		logger.DetailLoggerError(
			ctx,
			commandName,
			"Data not found",
			err,
		)
		return nil, nil
	}

	exampleName := make([]*entity.ExampleName, len(exampleNameModels))
	for key, val := range exampleNameModels {
		exampleNameModel := val.(*models.ExampleNameModels)
		data = append(data, exampleNameModel.GetModelsMap())
		exampleName[key] = &entity.ExampleName{}
		exampleNameMapper := mapper.NewExampleNameMapper(nil, exampleNameModel)
		exampleNameMapper.MapModelsToDomain(exampleName[key])
	}

	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Get list example name success",
		data,
	)
	return exampleName, nil
}
