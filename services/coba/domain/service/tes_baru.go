package service

import (
	"context"
	"github.com/Muruyung/go-boilgen/pkg/utils"
	"github.com/Muruyung/go-boilgen/services/coba/domain/entity"
)

// DTOTesBaru dto for tes baru service
type DTOTesBaru struct {
	Apacing int
	Test    bool
}

// TesBaruService tes baru service template
type TesBaruService interface {
	GetTesBaruByID(ctx context.Context, id string) (*entity.TesBaru, error)
	GetListTesBaru(ctx context.Context, request *utils.RequestOption) ([]*entity.TesBaru, *utils.MetaResponse, error)
	CreateTesBaru(ctx context.Context, dto DTOTesBaru) error
	UpdateTesBaru(ctx context.Context, id string, dto DTOTesBaru) error
	DeleteTesBaru(ctx context.Context, id string) error
}
