package example_name_usecase

import (
	"context"
	"github.com/Muruyung/go-utilities/logger"
)

// ExampleCustomQueryMethod example custom query method use case
func (uc *exampleNameInteractor) ExampleCustomQueryMethod(ctx context.Context, exampleParam string) (int, error) {
	const commandName = "UC-EXAMPLE-CUSTOM-QUERY-METHOD"
	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"example custom query method process...",
		nil,
	)

	exampleReturn, err := uc.ExampleNameSvc.ExampleCustomQueryMethod(ctx, exampleParam)
	if err != nil {
		logger.DetailLoggerError(
			ctx,
			commandName,
			"Error example custom query method",
			err,
		)
		return exampleReturn, err
	}

	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"example custom query method success",
		nil,
	)
	return exampleReturn, err
}
