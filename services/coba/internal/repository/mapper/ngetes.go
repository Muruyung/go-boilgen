package mapper

import (
	"time"

	"github.com/Muruyung/go-boilgen/pkg/logger"

	"github.com/Muruyung/go-boilgen/services/coba/domain/entity"
	"github.com/Muruyung/go-boilgen/services/coba/domain/repository"
	"github.com/Muruyung/go-boilgen/services/coba/internal/repository/models"
)

type ngetesMapperInteractor struct {
	ngetesModels *models.NgetesModels
	ngetesEntity *entity.Ngetes
}

// NewNgetesMapper create new ngetes models mapper
func NewNgetesMapper(ngetesEntity *entity.Ngetes, ngetesModels *models.NgetesModels) repository.MapperCommon {
	return &ngetesMapperInteractor{
		ngetesEntity: ngetesEntity,
		ngetesModels: ngetesModels,
	}
}

// MapDomainToModels map domain to models ngetes
func (mapper *ngetesMapperInteractor) MapDomainToModels() repository.ModelsCommon {
	repoModels := models.NgetesModels{
		ID:        mapper.ngetesEntity.GetID(),
		Name:      mapper.ngetesEntity.GetName(),
		CreatedAt: mapper.ngetesEntity.GetCreatedAt(),
		UpdatedAt: time.Now(),
	}
	return repoModels
}

// MapModelsToDomain map models to domain ngetes
func (mapper *ngetesMapperInteractor) MapModelsToDomain(entityStruct interface{}) {
	domain, err := entity.NewNgetes(entity.DTONgetes{
		ID:   mapper.ngetesModels.ID,
		Name: mapper.ngetesModels.Name,
	})
	if err != nil {
		logger.Logger.Warnf("Error: %v", err)
	}

	domain, err = domain.SetCreatedAt(mapper.ngetesModels.CreatedAt)
	if err != nil {
		logger.Logger.Warnf("Error: %v", err)
	}

	domain, err = domain.SetUpdatedAt(mapper.ngetesModels.UpdatedAt)
	if err != nil {
		logger.Logger.Warnf("Error: %v", err)
	}

	domain, err = domain.SetDeletedAt(mapper.ngetesModels.DeletedAt)
	if err != nil {
		logger.Logger.Warnf("Error: %v", err)
	}

	entityDomain := entityStruct.(*entity.Ngetes)
	*entityDomain = *domain
}
