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

// Get get single data ngetes
func (db *mysqlNgetesRepository) Get(ctx context.Context, query utils.QueryBuilderInteractor) (*entity.Ngetes, error) {
	const commandName = "REPO-GET-NGETES"
	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Get ngetes process...",
		nil,
	)

	var (
		err         error
		tableName   = models.NgetesModels{}.GetTableName()
		ngetes      = &entity.Ngetes{}
		ngetesModel *models.NgetesModels
	)

	query.AddPagination(utils.NewPagination(1, 1))
	query.AddWhere("deleted_at", "=", nil)
	stmt, val, _ := query.GetQuery(tableName, "")
	opts := &dbq.Options{
		SingleResult:   true,
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

	if result != nil {
		ngetesModel = result.(*models.NgetesModels)
		ngetesMapper := mapper.NewNgetesMapper(nil, ngetesModel)
		ngetesMapper.MapModelsToDomain(ngetes)
	} else {
		err = fmt.Errorf("ngetes data not found")
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
		"Get ngetes success",
		ngetesModel.GetModelsMap(),
	)
	return ngetes, nil
}
