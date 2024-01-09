package tes_baru_repo

import (
	"context"
	"fmt"
	"time"

	"github.com/Muruyung/go-boilgen/pkg/logger"
	"github.com/rocketlaunchr/dbq/v2"

	"github.com/Muruyung/go-boilgen/services/coba/domain/entity"
	"github.com/Muruyung/go-boilgen/services/coba/internal/repository/mapper"
	"github.com/Muruyung/go-boilgen/services/coba/internal/repository/models"
)

// Update update data tes baru
func (db *mysqlTesBaruRepository) Update(ctx context.Context, id string, data *entity.TesBaru) error {
	const commandName = "REPO-UPDATE-TES-BARU"
	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Update tes baru process...",
		nil,
	)

	var (
		err           error
		tableName     = models.TesBaruModels{}.GetTableName()
		tesBaruMapper = mapper.NewTesBaruMapper(data, nil).MapDomainToModels()
		tesBaruModels = tesBaruMapper.GetModelsMap()
		arrColumn     = tesBaruMapper.GetColumns()
		values        = make([]interface{}, 0)
		lastIndex     = len(arrColumn) - 1
		sqlDB         dbq.ExecContexter
	)

	if db.sql.Session().UseTx {
		sqlDB = db.sql.Session().Tx
	} else {
		sqlDB = db.sql.DB()
	}

	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	stmt := fmt.Sprintf(`UPDATE %s SET`, tableName)
	for key, val := range arrColumn {
		if tesBaruModels[val] != nil {
			stmt = fmt.Sprintf(`%s %s = ?`, stmt, val)
			values = append(values, tesBaruModels[val])
		}

		if key < lastIndex && tesBaruModels[val] != nil {
			stmt += `, `
		} else if key == lastIndex {
			stmt = fmt.Sprintf(`%s WHERE id = ?`, stmt)
		}
	}
	values = append(values, id)

	_, err = dbq.E(ctx, sqlDB, stmt, nil, values...)
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
		"Update tes baru success",
		nil,
	)

	return nil
}
