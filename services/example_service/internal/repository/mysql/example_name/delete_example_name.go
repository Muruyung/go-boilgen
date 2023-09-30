package example_name_repo

import (
	"context"
	"fmt"
	"time"

	"github.com/Muruyung/go-utilities/converter"
	"github.com/Muruyung/go-utilities/logger"
	"github.com/rocketlaunchr/dbq/v2"

	"github.com/Muruyung/go-boilgen/services/example_service/internal/repository/models"
)

// Delete delete data example name
func (db *mysqlExampleNameRepository) Delete(ctx context.Context, id int) error {
	const commandName = "REPO-DELETE-EXAMPLE-NAME"
	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Delete example name process...",
		nil,
	)

	var (
		err       error
		tableName = models.ExampleNameModels{}.GetTableName()
	)

	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	err = dbq.Tx(ctx, db.mysql, func(tx interface{}, Q dbq.QFn, E dbq.EFn, txCommit dbq.TxCommit) {
		stmt := fmt.Sprintf(`UPDATE %s SET deleted_at = ? WHERE id = ?`, tableName)

		_, err = E(ctx, stmt, nil, converter.ConvertDateToString(time.Now()), id)
		if err != nil {
			logger.DetailLoggerError(
				ctx,
				commandName,
				"Error execute query",
				err,
			)
			return
		}

		err = txCommit()
		if err != nil {
			logger.DetailLoggerError(
				ctx,
				commandName,
				"Failed commit query",
				err,
			)
			return
		}

		logger.DetailLoggerInfo(
			ctx,
			commandName,
			"Delete example name success",
			nil,
		)
	})

	return err
}
