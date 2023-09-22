package modulegenerator

import (
	"fmt"
	"os"

	"github.com/Muruyung/go-utilities/logger"
)

func modelsGenerator(dto dtoModule, isAll, isOnly bool) error {
	if !isAll && !isOnly {
		return nil
	}

	dto.path += "repository" + dto.sep + "models" + dto.sep
	var err error

	if _, err = os.Stat(dto.path); os.IsNotExist(err) {
		err = os.MkdirAll(dto.path, 0777)
		if err != nil {
			return err
		}
	}

	if _, err = os.Stat(dto.path + dto.name + ".go"); os.IsNotExist(err) {
		err = generateRepoModels(dto)
		if err != nil {
			logger.Logger.Error(fmt.Sprintf(defaultErr, err))
			return err
		} else {
			logger.Logger.Info("repository models created")
		}
	} else {
		logger.Logger.Warn("repository models already created")
	}

	return nil
}

func generateRepoModels(dto dtoModule) error {
	// var (
	// 	file      = jen.NewFilePathName(path, "models")
	// 	upperName = capitalize(name)
	// 	title     = sentences(name)
	// 	dir       = name
	// 	timeType  = "time.Time"
	// )
	// name = lowerize(name)

	// file.Add(jen.Id("import").Parens(
	// 	jen.Id("\n").
	// 		Id(`"sort"`).Id("\n").
	// 		Line().
	// 		Id(`"github.com/Muruyung/go-utilities/converter"`).Id("\n"),
	// ))

	// var (
	// 	generatedFields = make([]jen.Code, 0)
	// )

	// for field, fieldType := range dto.fields {
	// 	var upperCaseField = capitalize(field)

	// 	generatedFields = append(
	// 		generatedFields,
	// 		jen.Id(upperCaseField).Id(fieldType).Tag(map[string]string{
	// 			"json": field,
	// 			"dbq":  field,
	// 		}),
	// 	)
	// }

	// generatedFields = append(
	// 	generatedFields,
	// 	jen.Id("CreatedAt").Id(timeType).Tag(map[string]string{
	// 		"json": "created_at",
	// 		"dbq":  "created_at",
	// 	}),
	// 	jen.Id("UpdatedAt").Id(timeType).Tag(map[string]string{
	// 		"json": "updated_at",
	// 		"dbq":  "updated_at",
	// 	}),
	// 	jen.Id("DeletedAt").Id("*"+timeType).Tag(map[string]string{
	// 		"json": "deleted_at",
	// 		"dbq":  "deleted_at",
	// 	}),
	// )

	return nil
}
