package service

import (
	"github.com/Muruyung/go-boilgen/services/coba/domain/repository"
	"github.com/Muruyung/go-boilgen/services/coba/domain/service"
	ngetes_svc "github.com/Muruyung/go-boilgen/services/coba/internal/service/ngetes"
	tes_baru_svc "github.com/Muruyung/go-boilgen/services/coba/internal/service/tes_baru"
)

func Init(repo *repository.Wrapper, svcTx service.SvcTx) *service.Wrapper {
	return &service.Wrapper{
		TesBaruSvc: tes_baru_svc.NewTesBaruService(repo),
		NgetesSvc:  ngetes_svc.NewNgetesService(repo),
		SvcTx:      svcTx,
	}
}
