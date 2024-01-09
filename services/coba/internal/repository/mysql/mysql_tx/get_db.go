package mysqltx

import "database/sql"

// DB get sql db engine
func (db *mysqlTxRepository) DB() *sql.DB {
	return db.db
}
