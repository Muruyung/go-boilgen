package service

import (
	"context"
	"github.com/Muruyung/go-boilgen/pkg/utils"
	"github.com/Muruyung/go-boilgen/services/coba/domain/entity"
)

// DTONgetes dto for ngetes service
type DTONgetes struct {
	Name string
}

// NgetesService ngetes service template
type NgetesService interface {
	GetNgetesByID(ctx context.Context, id string) (*entity.Ngetes, error)
	GetListNgetes(ctx context.Context, request *utils.RequestOption) ([]*entity.Ngetes, *utils.MetaResponse, error)
	CreateNgetes(ctx context.Context, dto DTONgetes) error
	UpdateNgetes(ctx context.Context, id string, dto DTONgetes) error
}
