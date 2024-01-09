package repository

type Wrapper struct {
	TesBaruRepo TesBaruRepository
	NgetesRepo  NgetesRepository
	SqlTx
}
