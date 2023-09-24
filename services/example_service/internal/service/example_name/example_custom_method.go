package example_name

import (
	"context"
	goutils "github.com/Muruyung/go-utilities"
	"github.com/Muruyung/go-utilities/logger"
)

// ExampleCustomMethod example custom method service
func (svc *exampleNameInteractor) ExampleCustomMethod(ctx context.Context, exampleParam string) (int, error) {
	const commandName = "SVC-EXAMPLE-CUSTOM-METHOD"
	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Get list example custom method process...",
		nil,
	)

	var query = goutils.NewQueryBuilder()
	exampeReturn, err := svc.repo.ExampleNameRepo.ExampleCustomMethod(ctx, query)
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
		"Get list example custom method success",
		nil,
	)
	return exampeReturn, err
}
