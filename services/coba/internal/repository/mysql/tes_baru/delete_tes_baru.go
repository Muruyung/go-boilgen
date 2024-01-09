package tes_baru_repo

import (
	"context"
	"fmt"
	"time"

	"github.com/Muruyung/go-boilgen/pkg/logger"
	"github.com/Muruyung/go-utilities/converter"
	"github.com/rocketlaunchr/dbq/v2"

	"github.com/Muruyung/go-boilgen/services/coba/internal/repository/models"
)

// Delete delete data tes baru
func (db *mysqlTesBaruRepository) Delete(ctx context.Context, id string) error {
	const commandName = "REPO-DELETE-TES-BARU"
	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Delete tes baru process...",
		nil,
	)

	var (
		err       error
		tableName = models.TesBaruModels{}.GetTableName()
		sqlDB     dbq.ExecContexter
	)

	if db.sql.Session().UseTx {
		sqlDB = db.sql.Session().Tx
	} else {
		sqlDB = db.sql.DB()
	}

	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	stmt := fmt.Sprintf(`UPDATE %s SET deleted_at = ? WHERE id = ?`, tableName)
	_, err = dbq.E(ctx, sqlDB, stmt, nil, converter.ConvertDateToString(time.Now()), id)
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
		"Delete tes baru success",
		nil,
	)

	return nil
}
