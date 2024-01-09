package command

import (
	"context"
)

// DTOTesBaru dto for tes baru usecase
type DTOTesBaru struct {
	Apacing int
	Test    bool
}

// TesBaruUseCase tes baru command usecase template
type TesBaruUseCase interface {
	CreateTesBaru(ctx context.Context, dto DTOTesBaru) error
	UpdateTesBaru(ctx context.Context, id string, dto DTOTesBaru) error
	DeleteTesBaru(ctx context.Context, id string) error
}
