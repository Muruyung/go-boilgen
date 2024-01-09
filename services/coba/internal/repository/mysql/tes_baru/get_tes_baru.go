package tes_baru_repo

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

// Get get single data tes baru
func (db *mysqlTesBaruRepository) Get(ctx context.Context, query utils.QueryBuilderInteractor) (*entity.TesBaru, error) {
	const commandName = "REPO-GET-TES-BARU"
	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Get tes baru process...",
		nil,
	)

	var (
		err          error
		tableName    = models.TesBaruModels{}.GetTableName()
		tesBaru      = &entity.TesBaru{}
		tesBaruModel *models.TesBaruModels
	)

	query.AddPagination(utils.NewPagination(1, 1))
	query.AddWhere("deleted_at", "=", nil)
	stmt, val, _ := query.GetQuery(tableName, "")
	opts := &dbq.Options{
		SingleResult:   true,
		ConcreteStruct: models.TesBaruModels{},
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
		tesBaruModel = result.(*models.TesBaruModels)
		tesBaruMapper := mapper.NewTesBaruMapper(nil, tesBaruModel)
		tesBaruMapper.MapModelsToDomain(tesBaru)
	} else {
		err = fmt.Errorf("tes baru data not found")
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
		"Get tes baru success",
		tesBaruModel.GetModelsMap(),
	)
	return tesBaru, nil
}
