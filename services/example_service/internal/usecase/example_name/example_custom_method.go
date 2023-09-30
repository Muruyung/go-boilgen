package example_name_usecase

import (
	"context"
	"github.com/Muruyung/go-utilities/logger"
)

// ExampleCustomMethod example custom method use case
func (uc *exampleNameInteractor) ExampleCustomMethod(ctx context.Context, exampleParam string) (int, error) {
	const commandName = "UC-EXAMPLE-CUSTOM-METHOD"
	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"example custom method process...",
		nil,
	)

	exampeReturn, err := uc.ExampleNameSvc.ExampleCustomMethod(ctx, exampleParam)
	if err != nil {
		logger.DetailLoggerError(
			ctx,
			commandName,
			"Error example custom method",
			err,
		)
		return exampeReturn, err
	}

	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"example custom method success",
		nil,
	)
	return exampeReturn, err
}
