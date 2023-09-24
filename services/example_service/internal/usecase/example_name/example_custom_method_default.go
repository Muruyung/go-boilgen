package example_name

import (
	"context"
	"github.com/Muruyung/go-utilities/logger"
)

// ExampleCustomMethodDefault example custom method default use case
func (uc *exampleNameInteractor) ExampleCustomMethodDefault(ctx context.Context) error {
	const commandName = "UC-EXAMPLE-CUSTOM-METHOD-DEFAULT"
	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Get list example custom method default process...",
		nil,
	)

	err := uc.ExampleNameSvc.ExampleCustomMethodDefault(ctx)
	if err != nil {
		logger.DetailLoggerError(
			ctx,
			commandName,
			"Error example custom method default",
			err,
		)
		return err
	}

	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Get list example custom method default success",
		nil,
	)
	return err
}
