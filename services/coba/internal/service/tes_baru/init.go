package tes_baru_service

import (
	"github.com/Muruyung/go-boilgen/services/coba/domain/repository"
	"github.com/Muruyung/go-boilgen/services/coba/domain/service"
)

type tesBaruInteractor struct {
	repo *repository.Wrapper
}

// NewTesBaruService initialize new tes baru service
func NewTesBaruService(repo *repository.Wrapper) service.TesBaruService {
	return &tesBaruInteractor{
		repo: repo,
	}
}
