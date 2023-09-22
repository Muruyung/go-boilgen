package modulegenerator

import (
	"fmt"
	"os"

	"github.com/dave/jennifer/jen"

	"github.com/Muruyung/go-utilities/logger"
)

func entityGenerator(path, sep, name string, fields map[string]string) error {
	path += "entity" + sep
	var err error

	if _, err = os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, 0777)
		if err != nil {
			return err
		}
	}

	if _, err = os.Stat(path + name + ".go"); os.IsNotExist(err) {
		err = generateEntity(name, path, fields)
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

func generateEntity(name, path string, fields map[string]string) error {
	var (
		file      = jen.NewFilePathName(path, "entity")
		upperName = capitalize(name)
		title     = sentences(name)
		dir       = name
		timeType  = "time.Time"
	)
	name = lowerize(name)

	file.Add(jen.Id("import").Parens(
		jen.Id("\n").
			Id(`"time"`).Id("\n").
			Id(`"github.com/Muruyung/go-utilities/logger"`).Id("\n"),
	))

	var (
		generatedFields    = make([]jen.Code, 0)
		generatedDtoFields = make([]jen.Code, 0)
		generatedInit      = make([]jen.Code, 0)
	)

	for field, fieldType := range fields {
		var (
			lowerCaseField = lowerize(field)
			upperCaseField = capitalize(field)
		)

		if lowerCaseField == "ctx" {
			continue
		}

		generatedFields = append(
			generatedFields,
			jen.Id(lowerCaseField).Id(fieldType),
		)

		generatedDtoFields = append(
			generatedDtoFields,
			jen.Id(upperCaseField).Id(fieldType),
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
		jen.Id(name).Id(":=").Id("&"+upperName).Block(
			generatedInit...,
		),
		jen.Line(),
		jen.Id("err").Id(":=").Id(name).Dot("validate").Call(),
		jen.If(jen.Id("err").Op("!=").Nil()).Block(
			jen.Id("logger").Dot("Logger").Dot("Error").Call(jen.Id("err")),
			jen.Return(jen.Nil(), jen.Id("err")),
		),
		jen.Return(jen.Id(name), jen.Nil()),
	)

	file.Line()

	file.Func().Params(jen.Id("strc").Id(entityName)).
		Id("validate").Params().Id("error").Block(
		jen.Return(jen.Nil()),
	)

	file.Line()

	for field, fieldType := range fields {
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
		file.Func().Params(jen.Id("strc").Id(entityName)).
			Id(getterFuncName).Params().Id(fieldType).
			Block(
				jen.Return(jen.Id("strc").Dot(lowerCaseField)),
			)

		file.Commentf("%s set %s value", setterFuncName, lowerCaseField)
		file.Func().Params(jen.Id("strc").Id(entityName)).
			Id(setterFuncName).Params(jen.Id(lowerCaseField).Id(fieldType)).
			Parens(jen.List(jen.Id(entityName), jen.Error())).
			Block(
				jen.Id("strc").Dot(lowerCaseField).Op("=").Id(lowerCaseField),
				jen.Id("err").Id(":=").Id("strc").Dot("validate").Call(),
				jen.If(jen.Id("err").Op("!=").Nil()).Block(
					jen.Id("logger").Dot("Logger").Dot("Error").Call(jen.Id("err")),
					jen.Return(jen.Nil(), jen.Id("err")),
				),
				jen.Return(jen.Id("strc"), jen.Nil()),
			)

		file.Line()
	}

	file.Comment("GetCreatedAt get createdAt value")
	file.Func().Params(jen.Id("strc").Id(entityName)).
		Id("GetCreatedAt").Params().Id(timeType).
		Block(
			jen.Return(jen.Id("strc").Dot("createdAt")),
		)

	file.Comment("GetUpdatedAt get updatedAt value")
	file.Func().Params(jen.Id("strc").Id(entityName)).
		Id("GetUpdatedAt").Params().Id(timeType).
		Block(
			jen.Return(jen.Id("strc").Dot("updatedAt")),
		)

	file.Comment("GetDeletedAt get deletedAt value")
	file.Func().Params(jen.Id("strc").Id(entityName)).
		Id("GetDeletedAt").Params().Id("*" + timeType).
		Block(
			jen.Return(jen.Id("strc").Dot("deletedAt")),
		)

	return file.Save(path + "/" + dir + ".go")
}
