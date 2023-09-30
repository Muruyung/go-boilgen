package example_name_service

import (
	"context"
	"github.com/Muruyung/go-utilities/logger"
)

// ExampleCustomMethodDefault example custom method default service
func (svc *exampleNameInteractor) ExampleCustomMethodDefault(ctx context.Context) error {
	const commandName = "SVC-EXAMPLE-CUSTOM-METHOD-DEFAULT"
	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"example custom method default process...",
		nil,
	)

	var (
		err error
	)

	// TODO: Implement code here

	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"example custom method default success",
		nil,
	)
	return err
}
