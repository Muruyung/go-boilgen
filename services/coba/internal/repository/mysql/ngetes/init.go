package ngetes_repo

import (
	"github.com/Muruyung/go-boilgen/services/coba/domain/repository"
)

type mysqlNgetesRepository struct {
	sql repository.SqlTx
}

// NewNgetesRepository initialize new ngetes repository
func NewNgetesRepository(db repository.SqlTx) repository.NgetesRepository {
	return &mysqlNgetesRepository{
		sql: db,
	}
}
