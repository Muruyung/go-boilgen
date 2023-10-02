package example_name_usecase

import (
	"context"
	"github.com/Muruyung/go-boilgen/services/example_cqrs_service/domain/service"
	"github.com/Muruyung/go-boilgen/services/example_cqrs_service/domain/usecase/command"
	"github.com/Muruyung/go-utilities/logger"
)

// CreateExampleName create example name
func (uc *exampleNameInteractor) CreateExampleName(ctx context.Context, dto command.DTOExampleName) error {
	const commandName = "UC-CREATE-EXAMPLE-NAME"
	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Create example name process...",
		nil,
	)

	err := uc.ExampleNameSvc.CreateExampleName(ctx, service.DTOExampleName{
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
