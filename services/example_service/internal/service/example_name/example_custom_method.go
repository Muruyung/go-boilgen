package example_name_service

import (
	"context"
	"github.com/Muruyung/go-utilities/logger"
)

// ExampleCustomMethod example custom method service
func (svc *exampleNameInteractor) ExampleCustomMethod(ctx context.Context, exampleParam string) (int, error) {
	const commandName = "SVC-EXAMPLE-CUSTOM-METHOD"
	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"example custom method process...",
		nil,
	)

	var (
		exampleReturn int
		err           error
	)

	// TODO: Implement code here

	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"example custom method success",
		nil,
	)
	return exampleReturn, err
}
