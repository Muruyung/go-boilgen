package tes_baru_repo

import (
	"github.com/Muruyung/go-boilgen/services/coba/domain/repository"
)

type mysqlTesBaruRepository struct {
	sql repository.SqlTx
}

// NewTesBaruRepository initialize new tes baru repository
func NewTesBaruRepository(db repository.SqlTx) repository.TesBaruRepository {
	return &mysqlTesBaruRepository{
		sql: db,
	}
}
