package ngetes_repo

import (
	"context"
	"time"

	"github.com/Muruyung/go-boilgen/pkg/logger"
	"github.com/rocketlaunchr/dbq/v2"

	"github.com/Muruyung/go-boilgen/services/coba/domain/entity"
	"github.com/Muruyung/go-boilgen/services/coba/internal/repository/mapper"
	"github.com/Muruyung/go-boilgen/services/coba/internal/repository/models"
)

// Save create data ngetes
func (db *mysqlNgetesRepository) Save(ctx context.Context, data *entity.Ngetes) error {
	const commandName = "REPO-SAVE-NGETES"
	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Save ngetes process...",
		nil,
	)

	data, _ = data.SetCreatedAt(time.Now())
	var (
		err          error
		tableName    = models.NgetesModels{}.GetTableName()
		ngetesMapper = mapper.NewNgetesMapper(data, nil).MapDomainToModels()
		arrColumn    = ngetesMapper.GetColumns()
		arrValue     = ngetesMapper.GetValStruct(arrColumn)
		sqlDB        dbq.ExecContexter
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
		"Save ngetes success",
		nil,
	)

	return nil
}
