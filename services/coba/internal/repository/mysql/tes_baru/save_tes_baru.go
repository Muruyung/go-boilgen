package tes_baru_repo

import (
	"context"
	"time"

	"github.com/Muruyung/go-boilgen/pkg/logger"
	"github.com/rocketlaunchr/dbq/v2"

	"github.com/Muruyung/go-boilgen/services/coba/domain/entity"
	"github.com/Muruyung/go-boilgen/services/coba/internal/repository/mapper"
	"github.com/Muruyung/go-boilgen/services/coba/internal/repository/models"
)

// Save create data tes baru
func (db *mysqlTesBaruRepository) Save(ctx context.Context, data *entity.TesBaru) error {
	const commandName = "REPO-SAVE-TES-BARU"
	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Save tes baru process...",
		nil,
	)

	data, _ = data.SetCreatedAt(time.Now())
	var (
		err           error
		tableName     = models.TesBaruModels{}.GetTableName()
		tesBaruMapper = mapper.NewTesBaruMapper(data, nil).MapDomainToModels()
		arrColumn     = tesBaruMapper.GetColumns()
		arrValue      = tesBaruMapper.GetValStruct(arrColumn)
		sqlDB         dbq.ExecContexter
	)

	if db.sql.Session().UseTx {
		sqlDB = db.sql.Session().Tx
	} else {
		sqlDB = db.sql.DB()
	}

	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	stmt := dbq.INSERTStmt(tableName, arrColumn, len(arrValue), dbq.MySQL)
	_, err = dbq.E(ctx, sqlDB, stmt, nil, arrValue)
	if err != nil {
		logger.DetailLoggerError(
			ctx,
			commandName,
			"Error execute query",
			err,
		)
		return err
	}

	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Save tes baru success",
		nil,
	)

	return nil
}
