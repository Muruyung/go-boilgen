package database

import "database/sql"

type TX struct {
	*sql.Tx
	UseTx bool
}

// InitSqlTx initialize sql transaction
func InitSqlTx(db *sql.DB) *TX {
	tx, _ := db.Begin()
	return &TX{
		Tx:    tx,
		UseTx: false,
	}
}
