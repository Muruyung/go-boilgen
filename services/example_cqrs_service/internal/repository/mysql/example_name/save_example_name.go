package example_name_repo

import (
	"context"
	"time"

	"github.com/Muruyung/go-utilities/logger"
	"github.com/rocketlaunchr/dbq/v2"

	"github.com/Muruyung/go-boilgen/services/example_cqrs_service/domain/entity"
	"github.com/Muruyung/go-boilgen/services/example_cqrs_service/internal/repository/mapper"
	"github.com/Muruyung/go-boilgen/services/example_cqrs_service/internal/repository/models"
)

// Save create data example name
func (db *mysqlExampleNameRepository) Save(ctx context.Context, data *entity.ExampleName) error {
	const commandName = "REPO-SAVE-EXAMPLE-NAME"
	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Save example name process...",
		nil,
	)

	data, _ = data.SetCreatedAt(time.Now())
	var (
		err               error
		tableName         = models.ExampleNameModels{}.GetTableName()
		exampleNameMapper = mapper.NewExampleNameMapper(data, nil).MapDomainToModels()
		arrColumn         = exampleNameMapper.GetColumns()
		arrValue          = exampleNameMapper.GetValStruct(arrColumn)
	)

	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	err = dbq.Tx(ctx, db.mysql, func(tx interface{}, Q dbq.QFn, E dbq.EFn, txCommit dbq.TxCommit) {
		stmt := dbq.INSERTStmt(tableName, arrColumn, len(arrValue), dbq.MySQL)
		_, err = E(ctx, stmt, nil, arrValue)
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
			"Save example name success",
			nil,
		)
	})

	return err
}
