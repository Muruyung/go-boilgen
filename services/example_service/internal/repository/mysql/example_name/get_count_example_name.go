package example_name_repo

import (
	"context"
	"fmt"

	goutils "github.com/Muruyung/go-utilities"
	"github.com/Muruyung/go-utilities/logger"
	"github.com/rocketlaunchr/dbq/v2"

	"github.com/Muruyung/go-boilgen/services/example_service/internal/repository/models"
)

// GetCount get count data example name
func (db *mysqlExampleNameRepository) GetCount(ctx context.Context, query goutils.QueryBuilderInteractor) (int, error) {
	const commandName = "REPO-GET-COUNT-EXAMPLE-NAME"
	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Get count example name process...",
		0,
	)

	var (
		err       error
		tableName = models.ExampleNameModels{}.GetTableName()
		count     map[string]int
	)

	query.AddCount("id", "count")
	stmt, val, _ := query.GetQuery(tableName, "")
	opts := &dbq.Options{
		SingleResult:   true,
		ConcreteStruct: map[string]int{},
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
		return 0, err
	}

	if result != nil {
		count = result.(map[string]int)
	} else {
		err = fmt.Errorf("example name data not found")
		logger.DetailLoggerError(
			ctx,
			commandName,
			"Data not found",
			err,
		)
		return 0, nil
	}

	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Get count example name success",
		count,
	)
	return count["count"], nil
}
