package example_name

import (
	"context"
	goutils "github.com/Muruyung/go-utilities"
	"github.com/Muruyung/go-utilities/logger"
)

// ExampleCustomMethodDefault example custom method default service
func (svc *exampleNameInteractor) ExampleCustomMethodDefault(ctx context.Context) error {
	const commandName = "SVC-EXAMPLE-CUSTOM-METHOD-DEFAULT"
	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Get list example custom method default process...",
		nil,
	)

	var query = goutils.NewQueryBuilder()
	err := svc.repo.ExampleNameRepo.ExampleCustomMethodDefault(ctx, query)
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
