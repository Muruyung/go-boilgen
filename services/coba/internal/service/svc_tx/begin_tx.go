package svctx

import (
	"context"

	"github.com/Muruyung/go-boilgen/services/coba/domain/repository"
	"github.com/Muruyung/go-boilgen/services/coba/domain/service"
	svc "github.com/Muruyung/go-boilgen/services/coba/internal/service"
)

// BeginTx begin service transaction
func (svcTx *svcTxInteractor) BeginTx(ctx context.Context, operation func(ctx context.Context, svc *service.Wrapper) error) error {
	return svcTx.repo.BeginTx(ctx, func(ctx context.Context, repo *repository.Wrapper) error {
		return operation(ctx, svc.Init(repo, svcTx))
	})
}
