package mapper

import (
	"time"

	"github.com/Muruyung/go-boilgen/pkg/logger"

	"github.com/Muruyung/go-boilgen/services/coba/domain/entity"
	"github.com/Muruyung/go-boilgen/services/coba/domain/repository"
	"github.com/Muruyung/go-boilgen/services/coba/internal/repository/models"
)

type tesBaruMapperInteractor struct {
	tesBaruModels *models.TesBaruModels
	tesBaruEntity *entity.TesBaru
}

// NewTesBaruMapper create new tes baru models mapper
func NewTesBaruMapper(tesBaruEntity *entity.TesBaru, tesBaruModels *models.TesBaruModels) repository.MapperCommon {
	return &tesBaruMapperInteractor{
		tesBaruEntity: tesBaruEntity,
		tesBaruModels: tesBaruModels,
	}
}

// MapDomainToModels map domain to models tes baru
func (mapper *tesBaruMapperInteractor) MapDomainToModels() repository.ModelsCommon {
	repoModels := models.TesBaruModels{
		ID:        mapper.tesBaruEntity.GetID(),
		Apacing:   mapper.tesBaruEntity.GetApacing(),
		Test:      mapper.tesBaruEntity.GetTest(),
		CreatedAt: mapper.tesBaruEntity.GetCreatedAt(),
		UpdatedAt: time.Now(),
	}
	return repoModels
}

// MapModelsToDomain map models to domain tes baru
func (mapper *tesBaruMapperInteractor) MapModelsToDomain(entityStruct interface{}) {
	domain, err := entity.NewTesBaru(entity.DTOTesBaru{
		ID:      mapper.tesBaruModels.ID,
		Apacing: mapper.tesBaruModels.Apacing,
		Test:    mapper.tesBaruModels.Test,
	})
	if err != nil {
		logger.Logger.Warnf("Error: %v", err)
	}

	domain, err = domain.SetCreatedAt(mapper.tesBaruModels.CreatedAt)
	if err != nil {
		logger.Logger.Warnf("Error: %v", err)
	}

	domain, err = domain.SetUpdatedAt(mapper.tesBaruModels.UpdatedAt)
	if err != nil {
		logger.Logger.Warnf("Error: %v", err)
	}

	domain, err = domain.SetDeletedAt(mapper.tesBaruModels.DeletedAt)
	if err != nil {
		logger.Logger.Warnf("Error: %v", err)
	}

	entityDomain := entityStruct.(*entity.TesBaru)
	*entityDomain = *domain
}
