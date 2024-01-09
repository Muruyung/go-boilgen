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

// GetList get list data tes baru
func (db *mysqlTesBaruRepository) GetList(ctx context.Context, query utils.QueryBuilderInteractor) ([]*entity.TesBaru, error) {
	const commandName = "REPO-GET-LIST-TES-BARU"
	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Get list tes baru process...",
		nil,
	)

	var (
		err       error
		tableName = models.TesBaruModels{}.GetTableName()
		data      = make([]interface{}, 0)
	)

	query.AddWhere("deleted_at", "=", nil)
	stmt, val, _ := query.GetQuery(tableName, "")
	opts := &dbq.Options{
		SingleResult:   false,
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

	tesBaruModels := result.([]*models.TesBaruModels)
	if len(tesBaruModels) == 0 {
		err = fmt.Errorf("tes baru list data not found")
		logger.DetailLoggerError(
			ctx,
			commandName,
			"Data not found",
			err,
		)
		return nil, nil
	}

	tesBaru := make([]*entity.TesBaru, len(tesBaruModels))
	for key, val := range tesBaruModels {
		data = append(data, val.GetModelsMap())
		tesBaru[key] = new(entity.TesBaru)
		tesBaruMapper := mapper.NewTesBaruMapper(nil, val)
		tesBaruMapper.MapModelsToDomain(tesBaru[key])
	}

	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Get list tes baru success",
		data,
	)
	return tesBaru, nil
}
