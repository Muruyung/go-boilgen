package example_name_service

import (
	"context"
	"fmt"
	"github.com/Muruyung/go-utilities/logger"
)

// DeleteExampleName update example name
func (svc *exampleNameInteractor) DeleteExampleName(ctx context.Context, id int) error {
	const commandName = "SVC-DELETE-EXAMPLE-NAME"
	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Delete example name process...",
		nil,
	)

	err := svc.repo.ExampleNameRepo.Delete(ctx, id)
	if err != nil {
		logger.DetailLoggerError(
			ctx,
			commandName,
			fmt.Sprintf("Error delete by id=%v", id),
			err,
		)
		return err
	}

	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Delete example name success",
		nil,
	)
	return nil
}
