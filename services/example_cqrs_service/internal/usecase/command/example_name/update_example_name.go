package example_name_usecase

import (
	"context"
	"fmt"
	"github.com/Muruyung/go-boilgen/services/example_cqrs_service/domain/service"
	"github.com/Muruyung/go-boilgen/services/example_cqrs_service/domain/usecase/command"
	"github.com/Muruyung/go-utilities/logger"
)

// UpdateExampleName update example name
func (uc *exampleNameInteractor) UpdateExampleName(ctx context.Context, id int, dto command.DTOExampleName) error {
	const commandName = "UC-UPDATE-EXAMPLE-NAME"
	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Update example name process...",
		nil,
	)

	err := uc.ExampleNameSvc.UpdateExampleName(ctx, id, service.DTOExampleName{
		TestID:    dto.TestID,
		StartDate: dto.StartDate,
		Name:      dto.Name,
		Status:    dto.Status,
		IsActive:  dto.IsActive,
	})
	if err != nil {
		logger.DetailLoggerError(
			ctx,
			commandName,
			fmt.Sprintf("Error update by id=%v", id),
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
