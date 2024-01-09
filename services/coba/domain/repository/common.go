package repository

import (
	"context"
	"database/sql"

	"github.com/Muruyung/go-boilgen/pkg/database"
)

// MapperCommon template for common mapper models
type MapperCommon interface {
	MapDomainToModels() ModelsCommon
	MapModelsToDomain(entityStruct interface{})
}

// ModelsCommon template for common models repository
type ModelsCommon interface {
	GetTableName() string
	GetModels() interface{}
	GetModelsMap() map[string]interface{}
	GetColumns() []string
	GetValStruct(arrColumn []string) []interface{}
}

// SqlTx template for common transaction repository
type SqlTx interface {
	BeginTx(ctx context.Context, operation func(ctx context.Context, repo *Wrapper) error) error
	Session() *database.TX
	Wrapper() *Wrapper
	DB() *sql.DB
}
