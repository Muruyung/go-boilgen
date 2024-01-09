package modulegenerator

import (
	"fmt"
	"os"

	"github.com/Muruyung/go-utilities/logger"
	"github.com/dave/jennifer/jen"
	"github.com/iancoleman/strcase"
)

func modelsGenerator(dto dtoModule, isGenerate bool) error {
	if !isGenerate {
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
	var (
		file       = jen.NewFilePathName(dto.path, "models")
		upperName  = capitalize(dto.name)
		dir        = dto.name
		timeType   = "time.Time"
		modelsName = fmt.Sprintf("%sModels", upperName)
		title      = sentences(modelsName)
	)
	dto.name = lowerize(dto.name)

	file.Add(jen.Id("import").Parens(
		jen.Id("\n").
			Id(`"sort"`).Id("\n").
			Id(`"time"`).Id("\n").
			Line().
			Id(`"github.com/Muruyung/go-utilities/converter"`).Id("\n"),
	))

	var (
		generatedFields = make([]jen.Code, 0)
	)

	for _, field := range dto.arrFields {
		var (
			upperCaseField = capitalize(field)
			snakeCaseField = strcase.ToSnake(field)
		)

		generatedFields = append(
			generatedFields,
			jen.Id(upperCaseField).Id(dto.fields[field]).Tag(map[string]string{
				"json": snakeCaseField,
				"dbq":  snakeCaseField,
			}),
		)
	}

	generatedFields = append(
		generatedFields,
		jen.Id("CreatedAt").Id(timeType).Tag(map[string]string{
			"json": "created_at,omitempty",
			"dbq":  "created_at,omitempty",
		}),
		jen.Id("UpdatedAt").Id(timeType).Tag(map[string]string{
			"json": "updated_at,omitempty",
			"dbq":  "updated_at,omitempty",
		}),
		jen.Id("DeletedAt").Id("*"+timeType).Tag(map[string]string{
			"json": "deleted_at,omitempty",
			"dbq":  "deleted_at,omitempty",
		}),
	)

	file.Commentf("%s %s struct", modelsName, title)
	file.Type().Id(modelsName).Struct(
		generatedFields...,
	)
	file.Line()

	file.Commentf("GetTableName get table name of %s", title)
	file.Func().Parens(jen.Id("models").Id(modelsName)).Id("GetTableName").Params().Id("string").Block(
		jen.Return(jen.Lit(strcase.ToSnake(dto.name))),
	)
	file.Line()

	file.Commentf("GetModels get models of %s", title)
	file.Func().Parens(jen.Id("models").Id(modelsName)).Id("GetModels").Params().Id("interface{}").Block(
		jen.Return(jen.Id("models")),
	)
	file.Line()

	file.Commentf("GetModelsMap get models map of %s", title)
	file.Func().Parens(jen.Id("models").Id(modelsName)).Id("GetModelsMap").Params().Id("map[string]interface{}").Block(
		jen.Id("dataMap, _").Id(":=").Id("converter").Dot("ConvertInterfaceToMap").Parens(jen.Id("models")),
		jen.Return(jen.Id("dataMap")),
	)
	file.Line()

	file.Commentf("GetColumns get columns of %s", title)
	file.Func().Parens(jen.Id("models").Id(modelsName)).Id("GetColumns").Params().Id("[]string").Block(
		jen.Var().Parens(jen.Id("\n").
			Id("modelsMap").Id("=").Id("models").Dot("GetModelsMap()").Id("\n").
			Id("arrColumn").Id("=").Make(jen.Id("[]string"), jen.Id("0")).Id("\n"),
		),
		jen.Line(),
		jen.For(jen.Id("column").Id(":=").Range().Id("modelsMap")).Block(
			jen.Id("arrColumn").Id("=").Append(jen.Id("arrColumn"), jen.Id("column")),
		),
		jen.Id("sort").Dot("Strings").Parens(jen.Id("arrColumn")),
		jen.Line(),
		jen.Return(jen.Id("arrColumn")),
	)
	file.Line()

	file.Commentf("GetValStruct get value struct of %s", title)
	file.Func().Parens(jen.Id("models").Id(modelsName)).Id("GetValStruct").Params(jen.Id("arrColumn").Id("[]string")).Id("[]interface{}").Block(
		jen.Var().Parens(jen.Id("\n").
			Id("modelsMap").Id("=").Id("models").Dot("GetModelsMap()").Id("\n").
			Id("arrValue").Id("=").Make(jen.Id("[]interface{}"), jen.Id("0")).Id("\n"),
		),
		jen.Line(),
		jen.For(jen.Id("_, column").Id(":=").Range().Id("arrColumn")).Block(
			jen.Id("arrValue").Id("=").Append(jen.Id("arrValue"), jen.Id("modelsMap[column]")),
		),
		jen.Line(),
		jen.Return(jen.Id("[]interface{}{arrValue}")),
	)
	file.Line()

	return file.Save(dto.path + "/" + dir + ".go")
}
