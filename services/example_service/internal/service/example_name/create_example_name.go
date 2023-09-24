package example_name

import (
	"context"
	"github.com/Muruyung/go-boilgen/services/example_service/domain/entity"
	"github.com/Muruyung/go-boilgen/services/example_service/domain/service"
	"github.com/Muruyung/go-utilities/logger"
)

// CreateExampleName create example name
func (svc *exampleNameInteractor) CreateExampleName(ctx context.Context, dto service.DTOExampleName) error {
	const commandName = "SVC-CREATE-EXAMPLE-NAME"
	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Create example name process...",
		nil,
	)

	entityDTO, err := entity.NewExampleName(entity.DTOExampleName{
		Name:      dto.Name,
		Status:    dto.Status,
		IsActive:  dto.IsActive,
		TestID:    dto.TestID,
		StartDate: dto.StartDate,
	})
	if err != nil {
		logger.DetailLoggerError(
			ctx,
			commandName,
			"Error generate entity",
			err,
		)
		return err
	}

	err = svc.repo.ExampleNameRepo.Save(ctx, entityDTO)
	if err != nil {
		logger.DetailLoggerError(
			ctx,
			commandName,
			"Error create",
			err,
		)
		return err
	}

	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Create example name success",
		nil,
	)
	return nil
}
