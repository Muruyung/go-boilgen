package example_name_usecase

import (
	"context"
	"fmt"
	"github.com/Muruyung/go-boilgen/services/example_service/domain/service"
	"github.com/Muruyung/go-boilgen/services/example_service/domain/usecase"
	"github.com/Muruyung/go-utilities/logger"
)

// UpdateExampleName update example name
func (uc *exampleNameInteractor) UpdateExampleName(ctx context.Context, id int, dto usecase.DTOExampleName) error {
	const commandName = "UC-UPDATE-EXAMPLE-NAME"
	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Update example name process...",
		nil,
	)

	err := uc.ExampleNameSvc.UpdateExampleName(ctx, id, service.DTOExampleName{
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
