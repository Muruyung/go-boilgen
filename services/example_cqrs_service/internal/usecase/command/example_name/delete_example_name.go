package example_name_usecase

import (
	"context"
	"fmt"
	"github.com/Muruyung/go-utilities/logger"
)

// DeleteExampleName update example name
func (uc *exampleNameInteractor) DeleteExampleName(ctx context.Context, id int) error {
	const commandName = "UC-DELETE-EXAMPLE-NAME"
	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Delete example name process...",
		nil,
	)

	err := uc.ExampleNameSvc.DeleteExampleName(ctx, id)
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
