package mapper

import (
	"time"

	"github.com/Muruyung/go-utilities/logger"

	"github.com/Muruyung/go-boilgen/services/example_service/domain/entity"
	"github.com/Muruyung/go-boilgen/services/example_service/domain/repository"
	"github.com/Muruyung/go-boilgen/services/example_service/internal/repository/models"
)

type exampleNameMapperInteractor struct {
	exampleNameModels *models.ExampleNameModels
	exampleNameEntity *entity.ExampleName
}

// NewExampleNameMapper create new example name models mapper
func NewExampleNameMapper(exampleNameEntity *entity.ExampleName, exampleNameModels *models.ExampleNameModels) repository.MapperCommon {
	return &exampleNameMapperInteractor{
		exampleNameEntity: exampleNameEntity,
		exampleNameModels: exampleNameModels,
	}
}

// MapDomainToModels map domain to models example name
func (mapper *exampleNameMapperInteractor) MapDomainToModels() repository.ModelsCommon {
	repoModels := models.ExampleNameModels{
		ID:        mapper.exampleNameEntity.GetID(),
		Name:      mapper.exampleNameEntity.GetName(),
		Status:    mapper.exampleNameEntity.GetStatus(),
		IsActive:  mapper.exampleNameEntity.GetIsActive(),
		TestID:    mapper.exampleNameEntity.GetTestID(),
		StartDate: mapper.exampleNameEntity.GetStartDate(),
		CreatedAt: mapper.exampleNameEntity.GetCreatedAt(),
		UpdatedAt: time.Now(),
	}
	return repoModels
}

// MapModelsToDomain map models to domain example name
func (mapper *exampleNameMapperInteractor) MapModelsToDomain(entityStruct interface{}) {
	domain, err := entity.NewExampleName(entity.DTOExampleName{
		ID:        mapper.exampleNameModels.ID,
		Name:      mapper.exampleNameModels.Name,
		Status:    mapper.exampleNameModels.Status,
		IsActive:  mapper.exampleNameModels.IsActive,
		TestID:    mapper.exampleNameModels.TestID,
		StartDate: mapper.exampleNameModels.StartDate,
	})
	if err != nil {
		logger.Logger.Warnf("Error: %v", err)
	}

	domain, err = domain.SetCreatedAt(mapper.exampleNameModels.CreatedAt)
	if err != nil {
		logger.Logger.Warnf("Error: %v", err)
	}

	domain, err = domain.SetUpdatedAt(mapper.exampleNameModels.UpdatedAt)
	if err != nil {
		logger.Logger.Warnf("Error: %v", err)
	}

	domain, err = domain.SetDeletedAt(mapper.exampleNameModels.DeletedAt)
	if err != nil {
		logger.Logger.Warnf("Error: %v", err)
	}

	entityDomain := entityStruct.(*entity.ExampleName)
	*entityDomain = *domain
}
