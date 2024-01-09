package svctx

import (
	"github.com/Muruyung/go-boilgen/services/coba/domain/repository"
	"github.com/Muruyung/go-boilgen/services/coba/domain/service"
)

type svcTxInteractor struct {
	repo *repository.Wrapper
}

// NewSvcTx initialize new service transaction
func NewSvcTx(repo *repository.Wrapper) service.SvcTx {
	return &svcTxInteractor{
		repo: repo,
	}
}
