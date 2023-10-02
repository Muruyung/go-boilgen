package example_name_service

import (
	"context"
	"github.com/Muruyung/go-utilities/logger"
)

// ExampleCustomCreateCommandDefault example custom create command default service
func (svc *exampleNameInteractor) ExampleCustomCreateCommandDefault(ctx context.Context) error {
	const commandName = "SVC-EXAMPLE-CUSTOM-CREATE-COMMAND-DEFAULT"
	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"example custom create command default process...",
		nil,
	)

	var (
		err error
	)

	// TODO: Implement code here

	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"example custom create command default success",
		nil,
	)
	return err
}
