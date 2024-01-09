package ngetes_service

import (
	"github.com/Muruyung/go-boilgen/services/coba/domain/repository"
	"github.com/Muruyung/go-boilgen/services/coba/domain/service"
)

type ngetesInteractor struct {
	repo *repository.Wrapper
}

// NewNgetesService initialize new ngetes service
func NewNgetesService(repo *repository.Wrapper) service.NgetesService {
	return &ngetesInteractor{
		repo: repo,
	}
}
