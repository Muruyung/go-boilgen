package ngetes_usecase

import (
	"github.com/Muruyung/go-boilgen/services/coba/domain/service"
	"github.com/Muruyung/go-boilgen/services/coba/domain/usecase/command"
)

type ngetesInteractor struct {
	tx service.SvcTx
}

// NewNgetesUseCase initialize new ngetes use case
func NewNgetesUseCase(svc *service.Wrapper) command.NgetesUseCase {
	return &ngetesInteractor{
		tx: svc.SvcTx,
	}
}
