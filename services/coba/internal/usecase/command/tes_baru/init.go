package tes_baru_usecase

import (
	"github.com/Muruyung/go-boilgen/services/coba/domain/service"
	"github.com/Muruyung/go-boilgen/services/coba/domain/usecase/command"
)

type tesBaruInteractor struct {
	tx service.SvcTx
}

// NewTesBaruUseCase initialize new tes baru use case
func NewTesBaruUseCase(svc *service.Wrapper) command.TesBaruUseCase {
	return &tesBaruInteractor{
		tx: svc.SvcTx,
	}
}
