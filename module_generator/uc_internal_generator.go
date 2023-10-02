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

func internalUcGenerator(dto dtoModule, isAll, isOnly bool) error {
	if !isAll && !isOnly {
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
		file       = jen.NewFilePathName(path, strings.ToLower(name)+"_usecase")
		upperName  = capitalize(name)
		title      = sentences(name)
		dir        = "init"
		initReturn = "usecase"
		cqrsImport string
	)

	if strings.Contains(path, "/query/") {
		initReturn = "query"
		cqrsImport = "/query"
	} else if strings.Contains(path, "/command/") {
		initReturn = "command"
		cqrsImport = "/command"
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
		jen.Id("*service.Wrapper"),
	)

	initName := fmt.Sprintf("New%sUseCase", upperName)
	file.Commentf("%s initialize new %s use case", initName, title)
	file.Func().Id(initName).Params(jen.Id("svc").Id("*service.Wrapper")).Id(initReturn).Dot(ucName).
		Block(
			jen.Return(
				jen.Id("&" + interactorName).
					Block(
						jen.Id("Wrapper").Op(":").Id("svc,"),
					),
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
		jen.Id(`"context"`).Id("\n").
			Id(`"fmt"`).Id("\n").
			Id(`"github.com/Muruyung/go-utilities/logger"`).Id("\n").
			Id(fmt.Sprintf(`"%s/services/%s/domain/entity"`, projectName, services)).Id("\n"),
	))

	file.Commentf("%s get %s by id", methodName, title)
	file.Func().Params(jen.Id("uc").Id(embedStruct)).Id(methodName).
		Params(jen.Id("ctx").Id(ctx), jen.Id("id").Id(idFieldType)).
		Parens(jen.List(jen.Id(entityName), jen.Error())).
		Block(
			jen.Const().Id("commandName").Op("=").Lit("UC-"+strcase.ToScreamingKebab(methodName)),
			jen.Id(loggerInfo).Parens(
				jen.Id(loggerCtx).
					Id(loggerCmdName).
					Id("\n"+fmt.Sprintf(`"Get %s process...",`, title)).
					Id("\n").Nil().Id(",\n"),
			),
			jen.Line(),
			jen.Id("res, err").Id(":=").Id("uc").Dot(svcName).Dot(methodName).Id("(ctx, id)"),
			jen.If(jen.Id("err").Op("!=").Nil()).Block(
				jen.Id(loggerErr).Parens(
					jen.Id(loggerCtx).
						Id(loggerCmdName).
						Id("\n").Id(`fmt.Sprintf("Error get by id=%v", id),`).
						Id(logErr),
				),
				jen.Return(jen.Nil(), jen.Id("err")),
			),
			jen.Line(),
			jen.Id(loggerInfo).Parens(
				jen.Id(loggerCtx).
					Id(loggerCmdName).
					Id("\n"+fmt.Sprintf(`"Get %s success",`, title)).
					Id("\n").Nil().Id(",\n"),
			),
			jen.Return(jen.Id("res"), jen.Nil()),
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
		jen.Id(`"context"`).Id("\n").
			Id(`"github.com/Muruyung/go-utilities/logger"`).Id("\n").
			Id(`goutils"github.com/Muruyung/go-utilities"`).Id("\n").
			Id(fmt.Sprintf(`"%s/services/%s/domain/entity"`, projectName, services)).Id("\n"),
	))

	file.Commentf("%s get list %s", methodName, title)
	file.Func().Params(jen.Id("uc").Id(embedStruct)).Id(methodName).
		Params(jen.Id("ctx").Id(ctx), jen.Id("request").Id("*goutils.RequestOption")).
		Parens(jen.List(jen.Id(entityName), jen.Id("*goutils.MetaResponse"), jen.Error())).
		Block(
			jen.Const().Id("commandName").Op("=").Lit("UC-"+strcase.ToScreamingKebab(methodName)),
			jen.Id(loggerInfo).Parens(
				jen.Id(loggerCtx).
					Id(loggerCmdName).
					Id("\n"+fmt.Sprintf(`"Get list %s process...",`, title)).
					Id("\n").Nil().Id(",\n"),
			),
			jen.Line(),
			jen.Id("res, metaRes, err").Id(":=").Id("uc").Dot(svcName).Dot(methodName).Id("(ctx, request)"),
			jen.If(jen.Id("err").Op("!=").Nil()).Block(
				jen.Id(loggerErr).Parens(
					jen.Id(loggerCtx).
						Id(loggerCmdName).
						Id("\n").Id(`"Error get list",`).
						Id(logErr),
				),
				jen.Return(jen.Nil(), jen.Nil(), jen.Id("err")),
			),
			jen.Line(),
			jen.Id(loggerInfo).Parens(
				jen.Id(loggerCtx).
					Id(loggerCmdName).
					Id("\n"+fmt.Sprintf(`"Get list %s success",`, title)).
					Id("\n").Nil().Id(",\n"),
			),
			jen.Return(jen.Id("res"), jen.Id("metaRes"), jen.Nil()),
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
	)
	name = lowerize(name)

	if strings.Contains(path, "/query/") {
		dtoPath = "query"
		cqrsImport = "/query"
	} else if strings.Contains(path, "/command/") {
		dtoPath = "command"
		cqrsImport = "/command"
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
		jen.Id(`"context"`).Id("\n").
			Id(`"github.com/Muruyung/go-utilities/logger"`).Id("\n").
			Id(fmt.Sprintf(`"%s/services/%s/domain/service"`, projectName, services)).Id("\n").
			Id(fmt.Sprintf(`"%s/services/%s/domain/usecase%s"`, projectName, services, cqrsImport)).Id("\n"),
	))

	file.Commentf("%s create %s", methodName, title)
	file.Func().Params(jen.Id("uc").Id(embedStruct)).Id(methodName).
		Params(jen.Id("ctx").Id(ctx), jen.Id("dto").Id(dtoPath).Dot(dto)).Error().
		Block(
			jen.Const().Id("commandName").Op("=").Lit("UC-"+strcase.ToScreamingKebab(methodName)),
			jen.Id(loggerInfo).Parens(
				jen.Id(loggerCtx).
					Id(loggerCmdName).
					Id("\n"+fmt.Sprintf(`"Create %s process...",`, title)).
					Id("\n").Nil().Id(",\n"),
			),
			jen.Line(),
			jen.Id("err").Id(":=").Id("uc").Dot(svcName).Dot(methodName).Parens(jen.Id("ctx, service."+dto).Block(
				generatedParser...,
			)),
			jen.If(jen.Id("err").Op("!=").Nil()).Block(
				jen.Id(loggerErr).Parens(
					jen.Id(loggerCtx).
						Id(loggerCmdName).
						Id("\n").Id(`"Error create",`).
						Id(logErr),
				),
				jen.Return(jen.Id("err")),
			),
			jen.Line(),
			jen.Id(loggerInfo).Parens(
				jen.Id(loggerCtx).
					Id(loggerCmdName).
					Id("\n"+fmt.Sprintf(`"Create %s success",`, title)).
					Id("\n").Nil().Id(",\n"),
			),
			jen.Return(jen.Nil()),
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
	)
	name = lowerize(name)

	if strings.Contains(path, "/query/") {
		dtoPath = "query"
		cqrsImport = "/query"
	} else if strings.Contains(path, "/command/") {
		dtoPath = "command"
		cqrsImport = "/command"
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
		jen.Id(`"context"`).Id("\n").
			Id(`"fmt"`).Id("\n").
			Id(`"github.com/Muruyung/go-utilities/logger"`).Id("\n").
			Id(fmt.Sprintf(`"%s/services/%s/domain/service"`, projectName, services)).Id("\n").
			Id(fmt.Sprintf(`"%s/services/%s/domain/usecase%s"`, projectName, services, cqrsImport)).Id("\n"),
	))

	file.Commentf("%s update %s", methodName, title)
	file.Func().Params(jen.Id("uc").Id(embedStruct)).Id(methodName).
		Params(jen.Id("ctx").Id(ctx), jen.Id("id").Id(fields["id"]), jen.Id("dto").Id(dtoPath).Dot(dto)).Error().
		Block(
			jen.Const().Id("commandName").Op("=").Lit("UC-"+strcase.ToScreamingKebab(methodName)),
			jen.Id(loggerInfo).Parens(
				jen.Id(loggerCtx).
					Id(loggerCmdName).
					Id("\n"+fmt.Sprintf(`"Update %s process...",`, title)).
					Id("\n").Nil().Id(",\n"),
			),
			jen.Line(),
			jen.Id("err").Id(":=").Id("uc").Dot(svcName).Dot(methodName).Id("(ctx, id, service."+dto).Block(
				generatedParser...,
			).Id(")"),
			jen.If(jen.Id("err").Op("!=").Nil()).Block(
				jen.Id(loggerErr).Parens(
					jen.Id(loggerCtx).
						Id(loggerCmdName).
						Id("\n").Id(`fmt.Sprintf("Error update by id=%v", id),`).
						Id(logErr),
				),
				jen.Return(jen.Id("err")),
			),
			jen.Line(),
			jen.Id(loggerInfo).Parens(
				jen.Id(loggerCtx).
					Id(loggerCmdName).
					Id("\n"+fmt.Sprintf(`"Update %s success",`, title)).
					Id("\n").Nil().Id(",\n"),
			),
			jen.Return(jen.Nil()),
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
	)

	file.Add(jen.Id("import").Parens(
		jen.Id(`"context"`).Id("\n").
			Id(`"fmt"`).Id("\n").
			Id(`"github.com/Muruyung/go-utilities/logger"`).Id("\n"),
	))

	file.Commentf("%s update %s", methodName, title)
	file.Func().Params(jen.Id("uc").Id(embedStruct)).Id(methodName).
		Params(jen.Id("ctx").Id(ctx), jen.Id("id").Id(fields["id"])).Error().
		Block(
			jen.Const().Id("commandName").Op("=").Lit("UC-"+strcase.ToScreamingKebab(methodName)),
			jen.Id(loggerInfo).Parens(
				jen.Id(loggerCtx).
					Id(loggerCmdName).
					Id("\n"+fmt.Sprintf(`"Delete %s process...",`, title)).
					Id("\n").Nil().Id(",\n"),
			),
			jen.Line(),
			jen.Id("err").Id(":=").Id("uc").Dot(svcName).Dot(methodName).Id("(ctx, id)"),
			jen.If(jen.Id("err").Op("!=").Nil()).Block(
				jen.Id(loggerErr).Parens(
					jen.Id(loggerCtx).
						Id(loggerCmdName).
						Id("\n").Id(`fmt.Sprintf("Error delete by id=%v", id),`).
						Id(logErr),
				),
				jen.Return(jen.Id("err")),
			),
			jen.Line(),
			jen.Id(loggerInfo).Parens(
				jen.Id(loggerCtx).
					Id(loggerCmdName).
					Id("\n"+fmt.Sprintf(`"Delete %s success",`, title)).
					Id("\n").Nil().Id(",\n"),
			),
			jen.Return(jen.Nil()),
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
		returnVar             = strings.Join(dto.arrReturn[:], ",")
		err                   error
		params                = []jen.Code{jen.Id("ctx").Id(ctx)}
	)
	params = append(params, generatedCustomParams...)
	dto.name = lowerize(dto.name)

	var (
		interactorName = fmt.Sprintf("%sInteractor", dto.name)
		embedStruct    = fmt.Sprintf("*%s", interactorName)
	)

	importList := jen.Id("\n").Id(`"context"`).Id("\n")

	if isExists.isTimeExists {
		importList = importList.Id(`"time"`)
	}

	if isExists.isEntityExists {
		importList = importList.Id(fmt.Sprintf(`"%s/services/%s/domain/entity"`, projectName, dto.services)).Id("\n")
	}

	if isExists.isUtilsExists {
		importList = importList.Id(`goutils"github.com/Muruyung/go-utilities"`).Id("\n")
	}

	file.Add(jen.Id("import").Parens(
		importList.Id(`"github.com/Muruyung/go-utilities/logger"`).Id("\n"),
	))

	blockCode := []jen.Code{
		jen.Const().Id("commandName").Op("=").Lit("UC-" + strcase.ToScreamingKebab(methodName)),
		jen.Id(loggerInfo).Parens(
			jen.Id(loggerCtx).
				Id(loggerCmdName).
				Id("\n" + fmt.Sprintf(`"%s process...",`, title)).
				Id("\n").Nil().Id(",\n"),
		),
		jen.Line(),
		jen.Id(returnVar).Id(":=").Id("uc").Dot(svcName).Dot(methodName).Id("(ctx," + paramsVar + ")"),
	}

	if isExists.isError {
		blockCode = append(blockCode,
			jen.If(jen.Id("err").Op("!=").Nil()).Block(
				jen.Id(loggerErr).Parens(
					jen.Id(loggerCtx).
						Id(loggerCmdName).
						Id("\n").Id(fmt.Sprintf(`"Error %s",`, title)).
						Id(logErr),
				),
				jen.Return().Id(returnVar),
			),
		)
	}

	blockCode = append(blockCode,
		jen.Line(),
		jen.Id(loggerInfo).Parens(
			jen.Id(loggerCtx).
				Id(loggerCmdName).
				Id("\n"+fmt.Sprintf(`"%s success",`, title)).
				Id("\n").Nil().Id(",\n"),
		),
		jen.Return().Id(returnVar),
	)

	file.Commentf("%s %s use case", methodName, title)
	file.Func().Params(jen.Id("uc").Id(embedStruct)).Id(methodName).
		Params(params...).
		Parens(jen.List(generatedCustomReturn...)).
		Block(blockCode...)

	dir = fmt.Sprintf("%s/%s.go", dto.path, dir)
	if _, err = os.Stat(dir); os.IsNotExist(err) {
		return file.Save(dir)
	}

	logger.Logger.Warn("internal custom usecase already created")
	err = errors.New("duplicate internal custom usecase name")
	return err
}
