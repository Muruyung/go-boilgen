package modulegenerator

import (
	"fmt"
	"os"

	"github.com/Muruyung/go-utilities/logger"
	"github.com/dave/jennifer/jen"
)

func entityGenerator(dto dtoModule) error {
	dto.path += "entity" + dto.sep
	var err error

	if _, err = os.Stat(dto.path); os.IsNotExist(err) {
		err = os.MkdirAll(dto.path, 0777)
		if err != nil {
			return err
		}
	}

	if _, err = os.Stat(dto.path + dto.name + ".go"); os.IsNotExist(err) {
		err = generateEntity(dto)
		if err != nil {
			return err
		} else {
			logger.Logger.Info("entity created")
		}
	} else {
		logger.Logger.Warn("entity already created")
	}

	return err
}

func generateEntity(dto dtoModule) error {
	var (
		file      = jen.NewFilePathName(dto.path, "entity")
		upperName = capitalize(dto.name)
		title     = sentences(dto.name)
		dir       = dto.name
		timeType  = "time.Time"
	)
	dto.name = lowerize(dto.name)

	file.Add(jen.Id("import").Parens(
		jen.Id("\n").
			Id(`"time"`).Id("\n").
			Id(fmt.Sprintf(`"%s/pkg/logger"`, projectName)).Id("\n"),
	))

	var (
		generatedFields    = make([]jen.Code, 0)
		generatedDtoFields = make([]jen.Code, 0)
		generatedInit      = make([]jen.Code, 0)
	)

	for _, field := range dto.arrFields {
		var (
			lowerCaseField = lowerize(field)
			upperCaseField = capitalize(field)
		)

		if lowerCaseField == "ctx" {
			continue
		}

		generatedFields = append(
			generatedFields,
			jen.Id(lowerCaseField).Id(dto.fields[field]),
		)

		generatedDtoFields = append(
			generatedDtoFields,
			jen.Id(upperCaseField).Id(dto.fields[field]),
		)

		generatedInit = append(
			generatedInit,
			jen.Id(lowerCaseField).Op(":").Id(fmt.Sprintf("dto.%s", upperCaseField)).Id(","),
		)
	}

	generatedFields = append(
		generatedFields,
		jen.Id("createdAt").Id(timeType),
		jen.Id("updatedAt").Id(timeType),
		jen.Id("deletedAt").Id("*"+timeType),
	)

	file.Commentf("%s %s entity", upperName, title)
	file.Type().Id(upperName).Struct(
		generatedFields...,
	)

	var (
		dtoName    = fmt.Sprintf("DTO%s", upperName)
		funcName   = fmt.Sprintf("New%s", upperName)
		entityName = fmt.Sprintf("*%s", upperName)
	)

	file.Commentf("%s dto for %s entity", dtoName, title)
	file.Type().Id(dtoName).Struct(
		generatedDtoFields...,
	)

	file.Commentf("%s build new entity %s", funcName, title)
	file.Func().Id(funcName).Params(jen.Id("dto").Id(dtoName)).
		Parens(jen.List(jen.Id(entityName), jen.Id("error"))).Block(
		jen.Id(dto.name).Id(":=").Id("&"+upperName).Block(
			generatedInit...,
		),
		jen.Line(),
		jen.Id("err").Id(":=").Id(dto.name).Dot("validate").Call(),
		jen.If(jen.Id("err").Op("!=").Nil()).Block(
			jen.Id("logger").Dot("Logger").Dot("Error").Call(jen.Id("err")),
			jen.Return(jen.Nil(), jen.Id("err")),
		),
		jen.Return(jen.Id(dto.name), jen.Nil()),
	)

	file.Line()

	file.Func().Params(jen.Id("data").Id(entityName)).
		Id("validate").Params().Id("error").Block(
		jen.Return(jen.Nil()),
	)

	file.Line()

	for _, field := range dto.arrFields {
		var (
			lowerCaseField = lowerize(field)
			upperCaseField = capitalize(field)
			getterFuncName = fmt.Sprintf("Get%s", upperCaseField)
			setterFuncName = fmt.Sprintf("Set%s", upperCaseField)
		)

		if lowerCaseField == "ctx" {
			continue
		}

		file.Commentf("%s get %s value", getterFuncName, lowerCaseField)
		file.Func().Params(jen.Id("data").Id(entityName)).
			Id(getterFuncName).Params().Id(dto.fields[field]).
			Block(
				jen.Return(jen.Id("data").Dot(lowerCaseField)),
			)

		file.Commentf("%s set %s value", setterFuncName, lowerCaseField)
		file.Func().Params(jen.Id("data").Id(entityName)).
			Id(setterFuncName).Params(jen.Id(lowerCaseField).Id(dto.fields[field])).
			Parens(jen.List(jen.Id(entityName), jen.Error())).
			Block(
				jen.Id("data").Dot(lowerCaseField).Op("=").Id(lowerCaseField),
				jen.Id("err").Id(":=").Id("data").Dot("validate").Call(),
				jen.If(jen.Id("err").Op("!=").Nil()).Block(
					jen.Id("logger").Dot("Logger").Dot("Error").Call(jen.Id("err")),
					jen.Return(jen.Nil(), jen.Id("err")),
				),
				jen.Return(jen.Id("data"), jen.Nil()),
			)

		file.Line()
	}

	file.Comment("GetCreatedAt get createdAt value")
	file.Func().Params(jen.Id("data").Id(entityName)).
		Id("GetCreatedAt").Params().Id(timeType).
		Block(
			jen.Return(jen.Id("data").Dot("createdAt")),
		)

	file.Comment("SetCreatedAt set createdAt value")
	file.Func().Params(jen.Id("data").Id(entityName)).
		Id("SetCreatedAt").Params(jen.Id("date").Id("time.Time")).
		Parens(jen.List(jen.Id(entityName), jen.Error())).
		Block(
			jen.Id("data").Dot("createdAt").Op("=").Id("date"),
			jen.Id("err").Id(":=").Id("data").Dot("validate").Call(),
			jen.If(jen.Id("err").Op("!=").Nil()).Block(
				jen.Id("logger").Dot("Logger").Dot("Error").Call(jen.Id("err")),
				jen.Return(jen.Nil(), jen.Id("err")),
			),
			jen.Return(jen.Id("data"), jen.Nil()),
		)

	file.Comment("GetUpdatedAt get updatedAt value")
	file.Func().Params(jen.Id("data").Id(entityName)).
		Id("GetUpdatedAt").Params().Id(timeType).
		Block(
			jen.Return(jen.Id("data").Dot("updatedAt")),
		)

	file.Comment("SetUpdatedAt set updatedAt value")
	file.Func().Params(jen.Id("data").Id(entityName)).
		Id("SetUpdatedAt").Params(jen.Id("date").Id("time.Time")).
		Parens(jen.List(jen.Id(entityName), jen.Error())).
		Block(
			jen.Id("data").Dot("updatedAt").Op("=").Id("date"),
			jen.Id("err").Id(":=").Id("data").Dot("validate").Call(),
			jen.If(jen.Id("err").Op("!=").Nil()).Block(
				jen.Id("logger").Dot("Logger").Dot("Error").Call(jen.Id("err")),
				jen.Return(jen.Nil(), jen.Id("err")),
			),
			jen.Return(jen.Id("data"), jen.Nil()),
		)

	file.Comment("GetDeletedAt get deletedAt value")
	file.Func().Params(jen.Id("data").Id(entityName)).
		Id("GetDeletedAt").Params().Id("*" + timeType).
		Block(
			jen.Return(jen.Id("data").Dot("deletedAt")),
		)

	file.Comment("SetDeletedAt set deletedAt value")
	file.Func().Params(jen.Id("data").Id(entityName)).
		Id("SetDeletedAt").Params(jen.Id("date").Id("*time.Time")).
		Parens(jen.List(jen.Id(entityName), jen.Error())).
		Block(
			jen.Id("data").Dot("deletedAt").Op("=").Id("date"),
			jen.Id("err").Id(":=").Id("data").Dot("validate").Call(),
			jen.If(jen.Id("err").Op("!=").Nil()).Block(
				jen.Id("logger").Dot("Logger").Dot("Error").Call(jen.Id("err")),
				jen.Return(jen.Nil(), jen.Id("err")),
			),
			jen.Return(jen.Id("data"), jen.Nil()),
		)

	return file.Save(dto.path + "/" + dir + ".go")
}
