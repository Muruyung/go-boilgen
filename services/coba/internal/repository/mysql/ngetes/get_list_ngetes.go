package ngetes_repo

import (
	"context"
	"fmt"

	"github.com/Muruyung/go-boilgen/pkg/logger"
	"github.com/Muruyung/go-boilgen/pkg/utils"
	"github.com/rocketlaunchr/dbq/v2"

	"github.com/Muruyung/go-boilgen/services/coba/domain/entity"
	"github.com/Muruyung/go-boilgen/services/coba/internal/repository/mapper"
	"github.com/Muruyung/go-boilgen/services/coba/internal/repository/models"
)

// GetList get list data ngetes
func (db *mysqlNgetesRepository) GetList(ctx context.Context, query utils.QueryBuilderInteractor) ([]*entity.Ngetes, error) {
	const commandName = "REPO-GET-LIST-NGETES"
	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Get list ngetes process...",
		nil,
	)

	var (
		err       error
		tableName = models.NgetesModels{}.GetTableName()
		data      = make([]interface{}, 0)
	)

	query.AddWhere("deleted_at", "=", nil)
	stmt, val, _ := query.GetQuery(tableName, "")
	opts := &dbq.Options{
		SingleResult:   false,
		ConcreteStruct: models.NgetesModels{},
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
		return nil, err
	}

	ngetesModels := result.([]*models.NgetesModels)
	if len(ngetesModels) == 0 {
		err = fmt.Errorf("ngetes list data not found")
		logger.DetailLoggerError(
			ctx,
			commandName,
			"Data not found",
			err,
		)
		return nil, nil
	}

	ngetes := make([]*entity.Ngetes, len(ngetesModels))
	for key, val := range ngetesModels {
		data = append(data, val.GetModelsMap())
		ngetes[key] = new(entity.Ngetes)
		ngetesMapper := mapper.NewNgetesMapper(nil, val)
		ngetesMapper.MapModelsToDomain(ngetes[key])
	}

	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Get list ngetes success",
		data,
	)
	return ngetes, nil
}
