package example_name_service

import (
	"context"
	"github.com/Muruyung/go-utilities/logger"
)

// ExampleCustomQueryMethod example custom query method service
func (svc *exampleNameInteractor) ExampleCustomQueryMethod(ctx context.Context, exampleParam string) (int, error) {
	const commandName = "SVC-EXAMPLE-CUSTOM-QUERY-METHOD"
	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"example custom query method process...",
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
		"example custom query method success",
		nil,
	)
	return exampleReturn, err
}
