package ngetes_usecase

import (
	"context"

	"github.com/Muruyung/go-boilgen/services/coba/domain/entity"
)

// GetNgetesByID get ngetes by id
func (uc *ngetesInteractor) GetNgetesByID(ctx context.Context, id string) (*entity.Ngetes, error) {
	return uc.svc.NgetesSvc.GetNgetesByID(ctx, id)
}
