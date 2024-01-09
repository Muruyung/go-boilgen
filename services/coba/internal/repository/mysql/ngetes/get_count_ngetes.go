package ngetes_repo

import (
	"context"
	"fmt"

	"github.com/Muruyung/go-boilgen/pkg/logger"
	"github.com/Muruyung/go-boilgen/pkg/utils"
	"github.com/rocketlaunchr/dbq/v2"

	"github.com/Muruyung/go-boilgen/services/coba/internal/repository/models"
)

// GetCount get count data ngetes
func (db *mysqlNgetesRepository) GetCount(ctx context.Context, query utils.QueryBuilderInteractor) (int, error) {
	const commandName = "REPO-GET-COUNT-NGETES"
	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Get count ngetes process...",
		0,
	)

	var (
		err       error
		tableName = models.NgetesModels{}.GetTableName()
		count     *map[string]int
	)

	query.AddCount("id", "count")
	query.AddWhere("deleted_at", "=", nil)
	stmt, val, _ := query.GetQuery(tableName, "")
	opts := &dbq.Options{
		SingleResult:   true,
		ConcreteStruct: map[string]int{},
		DecoderConfig:  dbq.StdTimeConversionConfig(),
	}

	result, err := dbq.Q(ctx, db.sql.DB(), stmt, opts, val...)
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
		count = result.(*map[string]int)
	} else {
		err = fmt.Errorf("ngetes data not found")
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
		"Get count ngetes success",
		count,
	)
	return (*count)["count"], nil
}
