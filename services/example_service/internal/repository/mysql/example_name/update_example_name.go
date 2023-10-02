package example_name_repo

import (
	"context"
	"fmt"
	"time"

	"github.com/Muruyung/go-utilities/logger"
	"github.com/rocketlaunchr/dbq/v2"

	"github.com/Muruyung/go-boilgen/services/example_service/domain/entity"
	"github.com/Muruyung/go-boilgen/services/example_service/internal/repository/mapper"
	"github.com/Muruyung/go-boilgen/services/example_service/internal/repository/models"
)

// Update update data example name
func (db *mysqlExampleNameRepository) Update(ctx context.Context, id int, data *entity.ExampleName) error {
	const commandName = "REPO-UPDATE-EXAMPLE-NAME"
	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Update example name process...",
		nil,
	)

	var (
		err               error
		tableName         = models.ExampleNameModels{}.GetTableName()
		exampleNameMapper = mapper.NewExampleNameMapper(data, nil).MapDomainToModels()
		exampleNameModels = exampleNameMapper.GetModelsMap()
		arrColumn         = exampleNameMapper.GetColumns()
		values            = make([]interface{}, 0)
	)
	exampleNameModels["updated_at"] = time.Now()
	arrColumn = append(arrColumn, "updated_at")
	lastIndex := len(arrColumn) - 1

	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	err = dbq.Tx(ctx, db.mysql, func(tx interface{}, Q dbq.QFn, E dbq.EFn, txCommit dbq.TxCommit) {
		stmt := fmt.Sprintf(`UPDATE %s SET`, tableName)
		for key, val := range arrColumn {
			if exampleNameModels[val] != nil {
				stmt = fmt.Sprintf(`%s %s = ?`, stmt, val)
				values = append(values, exampleNameModels[val])
			}

			if key < lastIndex && exampleNameModels[val] != nil {
				stmt += `, `
			} else if key == lastIndex {
				stmt = fmt.Sprintf(`%s WHERE id = ?`, stmt)
			}
		}
		values = append(values, id)

		_, err = E(ctx, stmt, nil, values...)
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
			"Update example name success",
			nil,
		)
	})

	return err
}
