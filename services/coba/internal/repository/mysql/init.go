package repository_mysql

import (
	"github.com/Muruyung/go-boilgen/services/coba/domain/repository"
	ngetes_repo "github.com/Muruyung/go-boilgen/services/coba/internal/repository/mysql/ngetes"
	tes_baru_repo "github.com/Muruyung/go-boilgen/services/coba/internal/repository/mysql/tes_baru"
)

func Init(db repository.SqlTx) *repository.Wrapper {
	return &repository.Wrapper{
		TesBaruRepo: tes_baru_repo.NewTesBaruRepository(db),
		NgetesRepo:  ngetes_repo.NewNgetesRepository(db),
		SqlTx:       db,
	}
}
