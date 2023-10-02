package example_name_repo

import (
	"context"
	"fmt"

	goutils "github.com/Muruyung/go-utilities"
	"github.com/Muruyung/go-utilities/logger"
	"github.com/rocketlaunchr/dbq/v2"

	"github.com/Muruyung/go-boilgen/services/example_service/domain/entity"
	"github.com/Muruyung/go-boilgen/services/example_service/internal/repository/mapper"
	"github.com/Muruyung/go-boilgen/services/example_service/internal/repository/models"
)

// Get get single data example name
func (db *mysqlExampleNameRepository) Get(ctx context.Context, query goutils.QueryBuilderInteractor) (*entity.ExampleName, error) {
	const commandName = "REPO-GET-EXAMPLE-NAME"
	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Get example name process...",
		nil,
	)

	var (
		err              error
		tableName        = models.ExampleNameModels{}.GetTableName()
		exampleName      = &entity.ExampleName{}
		exampleNameModel *models.ExampleNameModels
	)

	query.AddPagination(goutils.NewPagination(1, 1))
	query.AddWhere("deleted_at", "!=", nil)
	stmt, val, _ := query.GetQuery(tableName, "")
	opts := &dbq.Options{
		SingleResult:   true,
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

	if result != nil {
		exampleNameModel = result.(*models.ExampleNameModels)
		exampleNameMapper := mapper.NewExampleNameMapper(nil, exampleNameModel)
		exampleNameMapper.MapModelsToDomain(exampleName)
	} else {
		err = fmt.Errorf("example name data not found")
		logger.DetailLoggerError(
			ctx,
			commandName,
			"Data not found",
			err,
		)
		return nil, nil
	}

	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Get example name success",
		exampleNameModel.GetModelsMap(),
	)
	return exampleName, nil
}
