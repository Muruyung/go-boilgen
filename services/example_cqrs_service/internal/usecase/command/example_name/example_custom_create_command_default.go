package example_name_usecase

import (
	"context"
	"github.com/Muruyung/go-utilities/logger"
)

// ExampleCustomCreateCommandDefault example custom create command default use case
func (uc *exampleNameInteractor) ExampleCustomCreateCommandDefault(ctx context.Context) error {
	const commandName = "UC-EXAMPLE-CUSTOM-CREATE-COMMAND-DEFAULT"
	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"example custom create command default process...",
		nil,
	)

	err := uc.ExampleNameSvc.ExampleCustomCreateCommandDefault(ctx)
	if err != nil {
		logger.DetailLoggerError(
			ctx,
			commandName,
			"Error example custom create command default",
			err,
		)
		return err
	}

	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"example custom create command default success",
		nil,
	)
	return err
}
