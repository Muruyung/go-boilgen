package mysqltx

import "github.com/Muruyung/go-boilgen/services/coba/domain/repository"

// Wrapper get repository wrapper
func (db *mysqlTxRepository) Wrapper() *repository.Wrapper {
	return db.wrapper
}
