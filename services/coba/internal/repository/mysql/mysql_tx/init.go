package mysqltx

import (
	"database/sql"

	"github.com/Muruyung/go-boilgen/pkg/database"
	"github.com/Muruyung/go-boilgen/services/coba/domain/repository"
	repository_mysql "github.com/Muruyung/go-boilgen/services/coba/internal/repository/mysql"
)

type mysqlTxRepository struct {
	db      *sql.DB
	tx      *database.TX
	wrapper *repository.Wrapper
}

// NewMysqlTx initialize new sqltx repository
func NewMysqlTx(db *sql.DB) repository.SqlTx {
	tx := &mysqlTxRepository{
		db: db,
	}
	tx.wrapper = repository_mysql.Init(tx)
	return tx
}
