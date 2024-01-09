package command

import (
	"context"
)

// DTONgetes dto for ngetes usecase
type DTONgetes struct {
	Name string
}

// NgetesUseCase ngetes command usecase template
type NgetesUseCase interface {
	CreateNgetes(ctx context.Context, dto DTONgetes) error
	UpdateNgetes(ctx context.Context, id string, dto DTONgetes) error
}
