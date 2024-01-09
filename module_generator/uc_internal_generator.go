package modulegenerator

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/Muruyung/go-utilities/logger"
	"github.com/dave/jennifer/jen"
	"github.com/iancoleman/strcase"
)

func internalUcGenerator(dto dtoModule, isGenerate bool) error {
	if !isGenerate {
		return nil
	}

	dto.path += "usecase" + dto.sep + dto.name
	var (
		err error
	)

	if _, err = os.Stat(dto.path); os.IsNotExist(err) {
		err = os.MkdirAll(dto.path, 0777)
		if err != nil {
			return err
		} else {
			logger.Logger.Info("internal directory usecase created")
		}
	}

	if _, err = os.Stat(dto.path + "/init.go"); os.IsNotExist(err) {
		err = generateInitUc(dto.name, dto.path, dto.services, dto.fields)
		if err != nil {
			return err
		} else {
			logger.Logger.Info("internal init usecase created")
		}
	}

	if _, ok := dto.methods["get"]; ok {
		err = generateGetUc(dto.name, dto.path, dto.services, dto.fields["id"])
		if err != nil {
			return err
		} else {
			logger.Logger.Info("internal get usecase created")
		}
	}

	if _, ok := dto.methods["getList"]; ok {
		err = generateGetListUc(dto.name, dto.path, dto.services)
		if err != nil {
			return err
		} else {
			logger.Logger.Info("internal get list usecase created")
		}
	}

	if _, ok := dto.methods["create"]; ok {
		err = generateCreateUc(dto.name, dto.path, dto.services, dto.fields)
		if err != nil {
			return err
		} else {
			logger.Logger.Info("internal create usecase created")
		}
	}

	if _, ok := dto.methods["update"]; ok {
		err = generateUpdateUc(dto.name, dto.path, dto.services, dto.fields)
		if err != nil {
			return err
		} else {
			logger.Logger.Info("internal update usecase created")
		}
	}

	if _, ok := dto.methods["delete"]; ok {
		err = generateDeleteUc(dto.name, dto.path, dto.services, dto.fields)
		if err != nil {
			return err
		} else {
			logger.Logger.Info("internal update usecase created")
		}
	}

	if _, ok := dto.methods["custom"]; ok {
		err = generateCustomUc(dto)
		if err != nil {
			return err
		} else {
			logger.Logger.Info("internal custom usecase created")
		}
	}

	return nil
}

func generateInitUc(name, path, services string, fields map[string]string) error {
	var (
		file        = jen.NewFilePathName(path, strings.ToLower(name)+"_usecase")
		upperName   = capitalize(name)
		title       = sentences(name)
		dir         = "init"
		initReturn  = "usecase"
		structField = jen.Id("svc").Id("*service.Wrapper")
		returnUC    = jen.Id("svc").Op(":").Id("svc,")
		cqrsImport  string
	)

	if strings.Contains(path, "/query/") {
		initReturn = "query"
		cqrsImport = "/query"
	} else if strings.Contains(path, "/command/") {
		initReturn = "command"
		cqrsImport = "/command"
		structField = jen.Id("tx").Id("service.SvcTx")
		returnUC = jen.Id("tx").Op(":").Id("svc.SvcTx,")
	}

	name = lowerize(name)
	interactorName := fmt.Sprintf("%sInteractor", name)

	file.Add(jen.Id("import").Parens(
		jen.Id(fmt.Sprintf(`"%s/services/%s/domain/service"`, projectName, services)).Id("\n").
			Id(fmt.Sprintf(`"%s/services/%s/domain/usecase%s"`, projectName, services, cqrsImport)).Id("\n"),
	))

	var (
		ucName = fmt.Sprintf("%sUseCase", upperName)
	)
	file.Type().Id(name + "Interactor").Struct(
		structField,
	)

	initName := fmt.Sprintf("New%sUseCase", upperName)
	file.Commentf("%s initialize new %s use case", initName, title)
	file.Func().Id(initName).Params(jen.Id("svc").Id("*service.Wrapper")).Id(initReturn).Dot(ucName).
		Block(
			jen.Return(
				jen.Id("&" + interactorName).Block(returnUC),
			),
		)

	return file.Save(path + "/" + dir + ".go")
}

func generateGetUc(name, path, services, idFieldType string) error {
	var (
		file       = jen.NewFilePathName(path, strings.ToLower(name)+"_usecase")
		upperName  = capitalize(name)
		title      = sentences(name)
		dir        = name
		entityName = fmt.Sprintf("*entity.%s", upperName)
		methodName = fmt.Sprintf("Get%sByID", upperName)
		svcName    = fmt.Sprintf("%sSvc", upperName)
		err        error
	)
	name = lowerize(name)

	var (
		interactorName = fmt.Sprintf("%sInteractor", name)
		embedStruct    = fmt.Sprintf("*%s", interactorName)
	)

	file.Add(jen.Id("import").Parens(
		jen.Id("\n").
			Id(`"context"`).Id("\n").
			Line().
			Id(fmt.Sprintf(`"%s/services/%s/domain/entity"`, projectName, services)).Id("\n"),
	))

	file.Commentf("%s get %s by id", methodName, title)
	file.Func().Params(jen.Id("uc").Id(embedStruct)).Id(methodName).
		Params(jen.Id("ctx").Id(ctx), jen.Id("id").Id(idFieldType)).
		Parens(jen.List(jen.Id(entityName), jen.Error())).
		Block(
			jen.Return(jen.Id("uc.svc").Dot(svcName).Dot(methodName).Params(jen.Id("ctx"), jen.Id("id"))),
		)

	dir = path + "/get_" + dir + ".go"
	if _, err = os.Stat(dir); os.IsNotExist(err) {
		return file.Save(dir)
	}

	logger.Logger.Warn("internal get usecase already created")
	err = errors.New("duplicate internal get usecase name")
	return err
}

func generateGetListUc(name, path, services string) error {
	var (
		file       = jen.NewFilePathName(path, strings.ToLower(name)+"_usecase")
		upperName  = capitalize(name)
		title      = sentences(name)
		dir        = name
		entityName = fmt.Sprintf("[]*entity.%s", upperName)
		methodName = fmt.Sprintf("GetList%s", upperName)
		svcName    = fmt.Sprintf("%sSvc", upperName)
		err        error
	)
	name = lowerize(name)

	var (
		interactorName = fmt.Sprintf("%sInteractor", name)
		embedStruct    = fmt.Sprintf("*%s", interactorName)
	)

	file.Add(jen.Id("import").Parens(
		jen.Id("\n").
			Id(`"context"`).Id("\n").
			Line().
			Id(fmt.Sprintf(`"%s/pkg/utils"`, projectName)).Id("\n").
			Id(fmt.Sprintf(`"%s/services/%s/domain/entity"`, projectName, services)).Id("\n"),
	))

	file.Commentf("%s get list %s", methodName, title)
	file.Func().Params(jen.Id("uc").Id(embedStruct)).Id(methodName).
		Params(jen.Id("ctx").Id(ctx), jen.Id("request").Id("*utils.RequestOption")).
		Parens(jen.List(jen.Id(entityName), jen.Id("*utils.MetaResponse"), jen.Error())).
		Block(
			jen.Return(jen.Id("uc.svc").Dot(svcName).Dot(methodName).Params(jen.Id("ctx"), jen.Id("request"))),
		)

	dir = path + "/get_list_" + dir + ".go"
	if _, err = os.Stat(dir); os.IsNotExist(err) {
		return file.Save(dir)
	}

	logger.Logger.Warn("internal get list usecase already created")
	err = errors.New("duplicate internal get list usecase name")
	return err
}

func generateCreateUc(name, path, services string, fields map[string]string) error {
	var (
		file            = jen.NewFilePathName(path, strings.ToLower(name)+"_usecase")
		upperName       = capitalize(name)
		title           = sentences(name)
		dir             = name
		methodName      = fmt.Sprintf("Create%s", upperName)
		svcName         = fmt.Sprintf("%sSvc", upperName)
		dto             = fmt.Sprintf("DTO%s", upperName)
		generatedParser = make([]jen.Code, 0)
		err             error
		dtoPath         = "usecase"
		cqrsImport      string
		returnTx        = "uc.svc.SvcTx.BeginTx"
	)
	name = lowerize(name)

	if strings.Contains(path, "/command/") {
		dtoPath = "command"
		cqrsImport = "/command"
		returnTx = "uc.tx.BeginTx"
	}

	var (
		interactorName = fmt.Sprintf("%sInteractor", name)
		embedStruct    = fmt.Sprintf("*%s", interactorName)
	)

	for field := range fields {
		if field != "id" {
			var upperCaseField = capitalize(field)
			generatedParser = append(
				generatedParser,
				jen.Id(upperCaseField).Op(":").Id("dto").Dot(upperCaseField).Id(","),
			)
		}
	}

	file.Add(jen.Id("import").Parens(
		jen.Id("\n").
			Id(`"context"`).Id("\n").
			Line().
			Id(fmt.Sprintf(`"%s/services/%s/domain/service"`, projectName, services)).Id("\n").
			Id(fmt.Sprintf(`"%s/services/%s/domain/usecase%s"`, projectName, services, cqrsImport)).Id("\n"),
	))

	file.Commentf("%s create %s", methodName, title)
	file.Func().Params(jen.Id("uc").Id(embedStruct)).Id(methodName).
		Params(jen.Id("ctx").Id(ctx), jen.Id("dto").Id(dtoPath).Dot(dto)).Error().
		Block(
			jen.Return(jen.Id(returnTx).Params(jen.Id("ctx"), jen.Func().Params(
				jen.Id("ctx").Id("context.Context"), jen.Id("svc").Id("*service.Wrapper"),
			).Error().Block(
				jen.Return().Id("svc").Dot(svcName).Dot(methodName).Parens(jen.Id("ctx, service."+dto).Block(
					generatedParser...,
				)),
			))),
		)

	dir = path + "/create_" + dir + ".go"
	if _, err = os.Stat(dir); os.IsNotExist(err) {
		return file.Save(dir)
	}

	logger.Logger.Warn("internal create usecase already created")
	err = errors.New("duplicate internal create usecase name")
	return err
}

func generateUpdateUc(name, path, services string, fields map[string]string) error {
	var (
		file            = jen.NewFilePathName(path, strings.ToLower(name)+"_usecase")
		upperName       = capitalize(name)
		title           = sentences(name)
		dir             = name
		methodName      = fmt.Sprintf("Update%s", upperName)
		svcName         = fmt.Sprintf("%sSvc", upperName)
		dto             = fmt.Sprintf("DTO%s", upperName)
		generatedParser = make([]jen.Code, 0)
		dtoPath         = "usecase"
		err             error
		cqrsImport      string
		returnTx        = "uc.svc.SvcTx.BeginTx"
	)
	name = lowerize(name)

	if strings.Contains(path, "/command/") {
		dtoPath = "command"
		cqrsImport = "/command"
		returnTx = "uc.tx.BeginTx"
	}

	var (
		interactorName = fmt.Sprintf("%sInteractor", name)
		embedStruct    = fmt.Sprintf("*%s", interactorName)
	)

	for field := range fields {
		if field != "id" {
			var upperCaseField = capitalize(field)
			generatedParser = append(
				generatedParser,
				jen.Id(upperCaseField).Op(":").Id("dto").Dot(upperCaseField).Id(","),
			)
		}
	}

	file.Add(jen.Id("import").Parens(
		jen.Id("\n").
			Id(`"context"`).Id("\n").
			Line().
			Id(fmt.Sprintf(`"%s/services/%s/domain/service"`, projectName, services)).Id("\n").
			Id(fmt.Sprintf(`"%s/services/%s/domain/usecase%s"`, projectName, services, cqrsImport)).Id("\n"),
	))

	file.Commentf("%s update %s", methodName, title)
	file.Func().Params(jen.Id("uc").Id(embedStruct)).Id(methodName).
		Params(jen.Id("ctx").Id(ctx), jen.Id("id").Id(fields["id"]), jen.Id("dto").Id(dtoPath).Dot(dto)).Error().
		Block(
			jen.Return(jen.Id(returnTx).Params(jen.Id("ctx"), jen.Func().Params(
				jen.Id("ctx").Id("context.Context"), jen.Id("svc").Id("*service.Wrapper"),
			).Error().Block(
				jen.Return().Id("svc").Dot(svcName).Dot(methodName).Parens(jen.Id("ctx, id, service."+dto).Block(
					generatedParser...,
				)),
			))),
		)

	dir = path + "/update_" + dir + ".go"
	if _, err = os.Stat(dir); os.IsNotExist(err) {
		return file.Save(dir)
	}

	logger.Logger.Warn("internal update usecase already created")
	err = errors.New("duplicate internal update usecase name")
	return err
}

func generateDeleteUc(name, path, services string, fields map[string]string) error {
	var (
		file       = jen.NewFilePathName(path, strings.ToLower(name)+"_usecase")
		upperName  = capitalize(name)
		title      = sentences(name)
		dir        = name
		methodName = fmt.Sprintf("Delete%s", upperName)
		svcName    = fmt.Sprintf("%sSvc", upperName)
		err        error
	)
	name = lowerize(name)

	var (
		interactorName = fmt.Sprintf("%sInteractor", name)
		embedStruct    = fmt.Sprintf("*%s", interactorName)
		returnTx       = "uc.svc.SvcTx.BeginTx"
	)

	if strings.Contains(path, "/command/") {
		returnTx = "uc.tx.BeginTx"
	}

	file.Add(jen.Id("import").Parens(
		jen.Id("\n").Id(`"context"`).Id("\n").
			Line().
			Id(fmt.Sprintf(`"%s/services/%s/domain/service"`, projectName, services)).Id("\n"),
	))

	file.Commentf("%s update %s", methodName, title)
	file.Func().Params(jen.Id("uc").Id(embedStruct)).Id(methodName).
		Params(jen.Id("ctx").Id(ctx), jen.Id("id").Id(fields["id"])).Error().
		Block(
			jen.Return(jen.Id(returnTx).Params(jen.Id("ctx"), jen.Func().Params(
				jen.Id("ctx").Id("context.Context"), jen.Id("svc").Id("*service.Wrapper"),
			).Error().Block(
				jen.Return().Id("svc").Dot(svcName).Dot(methodName).Id("(ctx, id)"),
			))),
		)

	dir = path + "/delete_" + dir + ".go"
	if _, err = os.Stat(dir); os.IsNotExist(err) {
		return file.Save(dir)
	}

	logger.Logger.Warn("internal delete usecase already created")
	err = errors.New("duplicate internal delete usecase name")
	return err
}

func generateCustomUc(dto dtoModule) error {
	var (
		file                  = jen.NewFilePathName(dto.path, strings.ToLower(dto.name)+"_usecase")
		upperName             = capitalize(dto.name)
		svcName               = fmt.Sprintf("%sSvc", upperName)
		title                 = sentences(dto.methodName)
		dir                   = strcase.ToSnake(dto.methodName)
		methodName            = capitalize(dto.methodName)
		isExists              = new(isExists)
		generatedCustomParams = parseCustomJenCodeFields(dto.params, dto.arrParams, isExists, false)
		generatedCustomReturn = parseCustomJenCodeFields(dto.returns, dto.arrReturn, isExists, true)
		paramsVar             = strings.Join(dto.arrParams[:], ",")
		err                   error
		params                = []jen.Code{jen.Id("ctx").Id(ctx)}
		returnUC              = jen.Id("uc.svc").Dot(svcName).Dot(methodName).Id("(ctx," + paramsVar + ")")
		importList            = jen.Id("\n").Id(`"context"`).Id("\n")
	)

	if strings.Contains(dto.path, "/command/") {
		generatedCustomReturn = []jen.Code{
			jen.Error(),
		}
		importList = importList.Id(fmt.Sprintf(`"%s/services/%s/domain/service"`, projectName, dto.services)).Id("\n")
		returnUC = jen.Id("uc.tx.BeginTx").Params(jen.Id("ctx"), jen.Func().Params(
			jen.Id("ctx").Id("context.Context"), jen.Id("svc").Id("*service.Wrapper"),
		).Error().Block(
			jen.Return().Id("svc").Dot(svcName).Dot(methodName).Id("(ctx,"+paramsVar+")"),
		))
	}

	params = append(params, generatedCustomParams...)
	dto.name = lowerize(dto.name)

	var (
		interactorName = fmt.Sprintf("%sInteractor", dto.name)
		embedStruct    = fmt.Sprintf("*%s", interactorName)
	)

	if isExists.isTimeExists {
		importList = importList.Id(`"time"`)
	}

	if isExists.isEntityExists {
		importList = importList.Id(fmt.Sprintf(`"%s/services/%s/domain/entity"`, projectName, dto.services)).Id("\n")
	}

	if isExists.isUtilsExists {
		importList = importList.Id(fmt.Sprintf(`"%s/pkg/utils"`, projectName)).Id("\n")
	}

	file.Add(jen.Id("import").Parens(
		importList,
	))

	file.Commentf("%s %s use case", methodName, title)
	file.Func().Params(jen.Id("uc").Id(embedStruct)).Id(methodName).
		Params(params...).
		Parens(jen.List(generatedCustomReturn...)).
		Block(
			jen.Return(returnUC),
		)

	dir = fmt.Sprintf("%s/%s.go", dto.path, dir)
	if _, err = os.Stat(dir); os.IsNotExist(err) {
		return file.Save(dir)
	}

	logger.Logger.Warn("internal custom usecase already created")
	err = errors.New("duplicate internal custom usecase name")
	return err
}
