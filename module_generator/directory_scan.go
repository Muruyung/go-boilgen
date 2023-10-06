package modulegenerator

import (
	"os"

	"github.com/Muruyung/go-utilities/logger"
)

func directoryScan(path string) ([]string, error) {
	var listName = make([]string, 0)
	dir, err := os.Open(path)
	if err != nil {
		logger.Logger.Error(defaultErr, err)
		return nil, err
	}

	defer dir.Close()
	files, err := dir.Readdir(-1)
	if err != nil {
		logger.Logger.Error(defaultErr, err)
		return nil, err
	}

	for _, file := range files {
		listName = append(listName, file.Name())
	}
	return listName, nil
}
