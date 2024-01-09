package modulegenerator

import (
	"fmt"
	"os"

	"github.com/Muruyung/go-utilities/logger"
	"github.com/dave/jennifer/jen"
)

func mapperGenerator(dto dtoModule, domainPath string, isGenerate bool) error {
	if !isGenerate {
		return nil
	}
	var err error

	domainPath += "repository" + dto.sep
	if _, err = os.Stat(domainPath); os.IsNotExist(err) {
		err = os.MkdirAll(domainPath, 0777)
		if err != nil {
			return err
		}
	}

	dto.path += "repository" + dto.sep + "mapper" + dto.sep
	if _, err = os.Stat(dto.path); os.IsNotExist(err) {
		err = os.MkdirAll(dto.path, 0777)
		if err != nil {
			return err
		}
	}

	if _, err = os.Stat(dto.path + dto.name + ".go"); os.IsNotExist(err) {
		err = generateRepoMapper(dto)
		if err != nil {
			logger.Logger.Error(fmt.Sprintf(defaultErr, err))
			return err
		} else {
			logger.Logger.Info("repository mapper created")
		}
	} else {
		logger.Logger.Warn("repository mapper already created")
	}

	return nil
}

func generateRepoMapper(dto dtoModule) error {
	var (
		file      = jen.NewFilePathName(dto.path, "mapper")
		upperName = capitalize(dto.name)
		dir       = dto.name
		title     = sentences(upperName)
		modelVar  = fmt.Sprintf("models.%sModels", upperName)
		entityVar = fmt.Sprintf("entity.%s", upperName)
	)
	dto.name = lowerize(dto.name)
	var (
		interactorName = fmt.Sprintf("%sMapperInteractor", dto.name)
		modelsName     = fmt.Sprintf("%sModels", dto.name)
		entityName     = fmt.Sprintf("%sEntity", dto.name)
	)

	file.Add(jen.Id("import").Parens(
		jen.Id(`"time"`).Id("\n").
			Line().
			Id(fmt.Sprintf(`"%s/pkg/logger"`, projectName)).Id("\n").
			Line().
			Id(fmt.Sprintf(`"%s/services/%s/domain/entity"`, projectName, dto.services)).Id("\n").
			Id(fmt.Sprintf(`"%s/services/%s/domain/repository"`, projectName, dto.services)).Id("\n").
			Id(fmt.Sprintf(`"%s/services/%s/internal/repository/models"`, projectName, dto.services)).Id("\n"),
	))

	file.Type().Id(interactorName).Struct(
		jen.Id(modelsName).Id("*"+modelVar),
		jen.Id(entityName).Id("*"+entityVar),
	)

	file.Commentf("New%sMapper create new %s models mapper", upperName, title)
	file.Func().Id("New"+upperName+"Mapper").Params(jen.Id(entityName).Id("*"+entityVar), jen.Id(modelsName).Id("*"+modelVar)).Id("repository.MapperCommon").Block(
		jen.Return(jen.Id("&"+interactorName).Block(
			jen.Id(entityName).Op(":").Id(entityName).Id(","),
			jen.Id(modelsName).Op(":").Id(modelsName).Id(","),
		)),
	)

	var (
		generatedDomToMod = make([]jen.Code, 0)
		generatedModToDom = make([]jen.Code, 0)
	)

	for _, field := range dto.arrFields {
		var (
			lowerCaseField = lowerize(field)
			upperCaseField = capitalize(field)
		)

		if lowerCaseField == "ctx" {
			continue
		}

		generatedDomToMod = append(
			generatedDomToMod,
			jen.Id(upperCaseField).Op(":").Id("mapper").Dot(entityName).Dot("Get"+upperCaseField+"()").Id(","),
		)

		generatedModToDom = append(
			generatedModToDom,
			jen.Id(upperCaseField).Op(":").Id("mapper").Dot(modelsName).Dot(upperCaseField).Id(","),
		)
	}

	generatedDomToMod = append(
		generatedDomToMod,
		jen.Id("CreatedAt").Op(":").Id("mapper").Dot(entityName).Dot("GetCreatedAt()").Id(","),
		jen.Id("UpdatedAt").Op(":").Id("time.Now(),"),
	)

	file.Commentf("MapDomainToModels map domain to models %s", title)
	file.Func().Params(jen.Id("mapper").Id("*"+interactorName)).Id("MapDomainToModels()").Id("repository.ModelsCommon").Block(
		jen.Id("repoModels").Id(":=").Id(modelVar).Block(
			generatedDomToMod...,
		),
		jen.Return(jen.Id("repoModels")),
	)

	file.Commentf("MapModelsToDomain map models to domain %s", title)
	file.Func().Params(jen.Id("mapper").Id("*"+interactorName)).Id("MapModelsToDomain").Parens(jen.Id("entityStruct").Interface()).Block(
		jen.Id("domain, err").Id(":=").Id("entity").Dot("New"+upperName).Parens(
			jen.Id("entity").Dot("DTO"+upperName).Block(
				generatedModToDom...,
			),
		),
		jen.If(jen.Id("err").Op("!=").Nil()).Block(
			jen.Id("logger").Dot("Logger").Dot("Warnf").Params(jen.Lit(defaultErr), jen.Id("err")),
		),
		jen.Line(),

		jen.Id("domain, err").Id("=").Id("domain").Dot("SetCreatedAt").Parens(jen.Id("mapper").Dot(modelsName).Dot("CreatedAt")),
		jen.If(jen.Id("err").Op("!=").Nil()).Block(
			jen.Id("logger").Dot("Logger").Dot("Warnf").Params(jen.Lit(defaultErr), jen.Id("err")),
		),
		jen.Line(),

		jen.Id("domain, err").Id("=").Id("domain").Dot("SetUpdatedAt").Parens(jen.Id("mapper").Dot(modelsName).Dot("UpdatedAt")),
		jen.If(jen.Id("err").Op("!=").Nil()).Block(
			jen.Id("logger").Dot("Logger").Dot("Warnf").Params(jen.Lit(defaultErr), jen.Id("err")),
		),
		jen.Line(),

		jen.Id("domain, err").Id("=").Id("domain").Dot("SetDeletedAt").Parens(jen.Id("mapper").Dot(modelsName).Dot("DeletedAt")),
		jen.If(jen.Id("err").Op("!=").Nil()).Block(
			jen.Id("logger").Dot("Logger").Dot("Warnf").Params(jen.Lit(defaultErr), jen.Id("err")),
		),

		jen.Line(),
		jen.Id("entityDomain").Op(":=").Id("entityStruct.").Parens(jen.Id("*"+entityVar)),
		jen.Id("*entityDomain").Op("=").Id("*domain"),
	)

	return file.Save(dto.path + "/" + dir + ".go")
}
