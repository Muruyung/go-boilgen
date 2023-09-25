package example_name

import (
	"context"
	"github.com/Muruyung/go-boilgen/services/example_service/domain/entity"
	"github.com/Muruyung/go-boilgen/services/example_service/domain/service"
	"github.com/Muruyung/go-utilities/logger"
)

// UpdateExampleName update example name
func (svc *exampleNameInteractor) UpdateExampleName(ctx context.Context, id int, dto service.DTOExampleName) error {
	const commandName = "SVC-UPDATE-EXAMPLE-NAME"
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

	err = svc.repo.ExampleNameRepo.Update(ctx, id, entityDTO)
	if err != nil {
		logger.DetailLoggerError(
			ctx,
			commandName,
			"Error update",
			err,
		)
		return err
	}

	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Update example name success",
		nil,
	)
	return nil
}
