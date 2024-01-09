package tes_baru_usecase

import (
	"github.com/Muruyung/go-boilgen/services/coba/domain/service"
	"github.com/Muruyung/go-boilgen/services/coba/domain/usecase/query"
)

type tesBaruInteractor struct {
	svc *service.Wrapper
}

// NewTesBaruUseCase initialize new tes baru use case
func NewTesBaruUseCase(svc *service.Wrapper) query.TesBaruUseCase {
	return &tesBaruInteractor{
		svc: svc,
	}
}
