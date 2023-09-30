package example_name_repo

import (
	"github.com/Muruyung/go-boilgen/services/example_service/domain/repository"

	"database/sql"
)

type mysqlExampleNameRepository struct {
	mysql *sql.DB
}

// NewExampleNameRepository initialize new example name repository
func NewExampleNameRepository(db *sql.DB) repository.ExampleNameRepository {
	return &mysqlExampleNameRepository{
		mysql: db,
	}
}
