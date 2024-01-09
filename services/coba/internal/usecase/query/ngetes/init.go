package ngetes_usecase

import (
	"github.com/Muruyung/go-boilgen/services/coba/domain/service"
	"github.com/Muruyung/go-boilgen/services/coba/domain/usecase/query"
)

type ngetesInteractor struct {
	svc *service.Wrapper
}

// NewNgetesUseCase initialize new ngetes use case
func NewNgetesUseCase(svc *service.Wrapper) query.NgetesUseCase {
	return &ngetesInteractor{
		svc: svc,
	}
}
