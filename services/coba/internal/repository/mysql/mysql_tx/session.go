package mysqltx

import "github.com/Muruyung/go-boilgen/pkg/database"

// Session get tx session
func (db *mysqlTxRepository) Session() *database.TX {
	if db.tx.UseTx {
		return db.tx
	}

	return &database.TX{
		UseTx: false,
	}
}
