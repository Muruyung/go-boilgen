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

func internalSvcGenerator(dto dtoModule, isAll, isOnly bool) error {
	if !isAll && !isOnly {
		return nil
	}

	dto.path += "service" + dto.sep + dto.name
	var (
		err error
	)

	if _, err = os.Stat(dto.path); os.IsNotExist(err) {
		err = os.MkdirAll(dto.path, 0777)
		if err != nil {
			return err
		} else {
			logger.Logger.Info("internal directory service created")
		}
	}

	if _, err = os.Stat(dto.path + "/init.go"); os.IsNotExist(err) {
		err = generateInitSvc(dto.name, dto.path, dto.services, dto.fields)
		if err != nil {
			return err
		} else {
			logger.Logger.Info("internal init service created")
		}
	}

	if _, ok := dto.methods["get"]; ok {
		err = generateGetSvc(dto.name, dto.path, dto.services, dto.fields["id"])
		if err != nil {
			return err
		} else {
			logger.Logger.Info("internal get service created")
		}
	}

	if _, ok := dto.methods["getList"]; ok {
		err = generateGetListSvc(dto.name, dto.path, dto.services)
		if err != nil {
			return err
		} else {
			logger.Logger.Info("internal get list service created")
		}
	}

	if _, ok := dto.methods["create"]; ok {
		err = generateCreatetSvc(dto.name, dto.path, dto.services, dto.fields)
		if err != nil {
			return err
		} else {
			logger.Logger.Info("internal create service created")
		}
	}

	if _, ok := dto.methods["update"]; ok {
		err = generateUpdateSvc(dto.name, dto.path, dto.services, dto.fields)
		if err != nil {
			return err
		} else {
			logger.Logger.Info("internal update service created")
		}
	}

	if _, ok := dto.methods["delete"]; ok {
		err = generateDeleteSvc(dto.name, dto.path, dto.services, dto.fields)
		if err != nil {
			return err
		} else {
			logger.Logger.Info("internal update service created")
		}
	}

	if _, ok := dto.methods["custom"]; ok {
		err = generateCustomSvc(dto)
		if err != nil {
			return err
		} else {
			logger.Logger.Info("internal custom service created")
		}
	}

	return nil
}

func generateInitSvc(name, path, services string, fields map[string]string) error {
	var (
		file      = jen.NewFilePathName(path, strings.ToLower(name)+"_service")
		upperName = capitalize(name)
		title     = sentences(name)
		dir       = "init"
	)
	name = lowerize(name)
	interactorName := fmt.Sprintf("%sInteractor", name)

	file.Add(jen.Id("import").Parens(
		jen.Id(fmt.Sprintf(`"%s/services/%s/domain/repository"`, projectName, services)).Id("\n").
			Id(fmt.Sprintf(`"%s/services/%s/domain/service"`, projectName, services)).Id("\n"),
	))

	var (
		svcName = fmt.Sprintf("%sService", upperName)
	)
	file.Type().Id(name + "Interactor").Struct(
		jen.Id("repo").Id("*repository.Wrapper"),
	)

	initName := fmt.Sprintf("New%sService", upperName)
	file.Commentf("%s initialize new %s service", initName, title)
	file.Func().Id(initName).Params(jen.Id("repo").Id("*repository.Wrapper")).Id("service").Dot(svcName).
		Block(
			jen.Return(
				jen.Id("&" + interactorName).
					Block(
						jen.Id("repo").Op(":").Id("repo,"),
					),
			),
		)

	return file.Save(path + "/" + dir + ".go")
}

func generateGetSvc(name, path, services, idFieldType string) error {
	var (
		file       = jen.NewFilePathName(path, strings.ToLower(name)+"_service")
		upperName  = capitalize(name)
		title      = sentences(name)
		dir        = name
		entityName = fmt.Sprintf("*entity.%s", upperName)
		methodName = fmt.Sprintf("Get%sByID", upperName)
		repoName   = fmt.Sprintf("%sRepo", upperName)
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
			Id(`goutils"github.com/Muruyung/go-utilities"`).Id("\n").
			Id(`"github.com/Muruyung/go-utilities/logger"`).Id("\n").
			Id(fmt.Sprintf(`"%s/services/%s/domain/entity"`, projectName, services)).Id("\n"),
	))

	file.Commentf("%s get %s by id", methodName, title)
	file.Func().Params(jen.Id("svc").Id(embedStruct)).Id(methodName).
		Params(jen.Id("ctx").Id(ctx), jen.Id("id").Id(idFieldType)).
		Parens(jen.List(jen.Id(entityName), jen.Error())).
		Block(
			jen.Const().Id("commandName").Op("=").Lit("SVC-"+strcase.ToScreamingKebab(methodName)),
			jen.Id(loggerInfo).Parens(
				jen.Id(loggerCtx).
					Id(loggerCmdName).
					Id("\n"+fmt.Sprintf(`"Get %s process...",`, title)).
					Id("\n").Nil().Id(",\n"),
			),
			jen.Line(),
			jen.Var().Id("query").Op("=").Id("goutils.NewQueryBuilder()"),
			jen.Id("query").Dot("AddWhere").Parens(jen.List(jen.Lit("id"), jen.Lit("="), jen.Lit(idFieldType))),
			jen.Id("res, err").Id(":=").Id("svc").Dot("repo").Dot(repoName).Dot("Get").Id("(ctx, query)"),
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

	logger.Logger.Warn("internal get service already created")
	err = errors.New("duplicate internal get service name")
	return err
}

func generateGetListSvc(name, path, services string) error {
	var (
		file       = jen.NewFilePathName(path, strings.ToLower(name)+"_service")
		upperName  = capitalize(name)
		title      = sentences(name)
		dir        = name
		entityName = fmt.Sprintf("[]*entity.%s", upperName)
		methodName = fmt.Sprintf("GetList%s", upperName)
		repoName   = fmt.Sprintf("%sRepo", upperName)
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
	file.Func().Params(jen.Id("svc").Id(embedStruct)).Id(methodName).
		Params(jen.Id("ctx").Id(ctx), jen.Id("request").Id("*goutils.RequestOption")).
		Parens(jen.List(jen.Id(entityName), jen.Id("*goutils.MetaResponse"), jen.Error())).
		Block(
			jen.Const().Id("commandName").Op("=").Lit("SVC-"+strcase.ToScreamingKebab(methodName)),
			jen.Id(loggerInfo).Parens(
				jen.Id(loggerCtx).
					Id(loggerCmdName).
					Id("\n"+fmt.Sprintf(`"Get list %s process...",`, title)).
					Id("\n").Nil().Id(",\n"),
			),
			jen.Line(),
			jen.Var().Parens(
				jen.Id("\n").
					Id("query").Op("=").Id("goutils.NewQueryBuilder()").Id("\n").
					Id("queryPagination").Op("=").Id("goutils.NewQueryBuilder()").Id("\n").
					Id("metaRes").Id("*goutils.MetaResponse").Id("\n").
					Id("page").Id("int").Id("\n").
					Id("limit").Id("int").Id("\n"),
			),
			jen.Line(),
			jen.If(jen.Id("request").Op("!=").Nil()).Block(
				jen.Id("query, page, limit").Id("=").Id("request").Dot("SetPaginationWithSort").Parens(jen.Id("query")),
			),
			jen.Line(),
			jen.Id("res, err").Id(":=").Id("svc").Dot("repo").Dot(repoName).Dot("GetList").Id("(ctx, query)"),
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
			jen.If(jen.Id("request").Op("!=").Nil().Op("&&").Id("request").Dot("GetPagination()").Op("!=").Nil()).Block(
				jen.Id("totalCount, err").Id(":=").Id("svc").Dot("repo").Dot(repoName).Dot("GetCount").Id("(ctx, queryPagination)"),
				jen.If(jen.Id("err").Op("!=").Nil()).Block(
					jen.Id(loggerErr).Parens(
						jen.Id(loggerCtx).
							Id(loggerCmdName).
							Id("\n").Id(`"Error get total count list",`).Id("\n").
							Id(logErr),
					),
					jen.Return(jen.Nil(), jen.Nil(), jen.Id("err")),
				),
				jen.Line(),
				jen.Var().Id("meta").Id("=").Id("goutils.MapMetaResponse").Parens(jen.List(
					jen.Id("totalCount"),
					jen.Id("len(res)"),
					jen.Id("page"),
					jen.Id("limit"),
				)),
				jen.Id("metaRes").Id("=").Id("&meta"),
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

	logger.Logger.Warn("internal get list service already created")
	err = errors.New("duplicate internal get list service name")
	return err
}

func generateCreatetSvc(name, path, services string, fields map[string]string) error {
	var (
		file            = jen.NewFilePathName(path, strings.ToLower(name)+"_service")
		upperName       = capitalize(name)
		title           = sentences(name)
		dir             = name
		methodName      = fmt.Sprintf("Create%s", upperName)
		repoName        = fmt.Sprintf("%sRepo", upperName)
		dto             = fmt.Sprintf("DTO%s", upperName)
		generatedParser = make([]jen.Code, 0)
		err             error
	)
	name = lowerize(name)

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
			Id(fmt.Sprintf(`"%s/services/%s/domain/entity"`, projectName, services)).Id("\n").
			Id(fmt.Sprintf(`"%s/services/%s/domain/service"`, projectName, services)).Id("\n"),
	))

	file.Commentf("%s create %s", methodName, title)
	file.Func().Params(jen.Id("svc").Id(embedStruct)).Id(methodName).
		Params(jen.Id("ctx").Id(ctx), jen.Id("dto").Id("service").Dot(dto)).Error().
		Block(
			jen.Const().Id("commandName").Op("=").Lit("SVC-"+strcase.ToScreamingKebab(methodName)),
			jen.Id(loggerInfo).Parens(
				jen.Id(loggerCtx).
					Id(loggerCmdName).
					Id("\n"+fmt.Sprintf(`"Create %s process...",`, title)).
					Id("\n").Nil().Id(",\n"),
			),
			jen.Line(),
			jen.Id("entityDTO, err").Id(":=").Id("entity").Dot("New"+upperName).Parens(jen.Id("entity."+dto).Block(
				generatedParser...,
			)),
			jen.If(jen.Id("err").Op("!=").Nil()).Block(
				jen.Id(loggerErr).Parens(
					jen.Id(loggerCtx).
						Id(loggerCmdName).
						Id("\n").Id(`"Error generate entity",`).
						Id(logErr),
				),
				jen.Return(jen.Id("err")),
			),
			jen.Line(),
			jen.Id("err").Id("=").Id("svc").Dot("repo").Dot(repoName).Dot("Save").Parens(jen.Id("ctx, entityDTO")),
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

	logger.Logger.Warn("internal create service already created")
	err = errors.New("duplicate internal create service name")
	return err
}
func generateUpdateSvc(name, path, services string, fields map[string]string) error {
	var (
		file            = jen.NewFilePathName(path, strings.ToLower(name)+"_service")
		upperName       = capitalize(name)
		title           = sentences(name)
		dir             = name
		methodName      = fmt.Sprintf("Update%s", upperName)
		repoName        = fmt.Sprintf("%sRepo", upperName)
		dto             = fmt.Sprintf("DTO%s", upperName)
		generatedParser = make([]jen.Code, 0)
		err             error
	)
	name = lowerize(name)

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
			Id(fmt.Sprintf(`"%s/services/%s/domain/entity"`, projectName, services)).Id("\n").
			Id(fmt.Sprintf(`"%s/services/%s/domain/service"`, projectName, services)).Id("\n"),
	))

	file.Commentf("%s update %s", methodName, title)
	file.Func().Params(jen.Id("svc").Id(embedStruct)).Id(methodName).
		Params(jen.Id("ctx").Id(ctx), jen.Id("id").Id(fields["id"]), jen.Id("dto").Id("service").Dot(dto)).Error().
		Block(
			jen.Const().Id("commandName").Op("=").Lit("SVC-"+strcase.ToScreamingKebab(methodName)),
			jen.Id(loggerInfo).Parens(
				jen.Id(loggerCtx).
					Id(loggerCmdName).
					Id("\n"+fmt.Sprintf(`"Create %s process...",`, title)).
					Id("\n").Nil().Id(",\n"),
			),
			jen.Line(),
			jen.Id("entityDTO, err").Id(":=").Id("entity").Dot("New"+upperName).Parens(jen.Id("entity."+dto).Block(
				generatedParser...,
			)),
			jen.If(jen.Id("err").Op("!=").Nil()).Block(
				jen.Id(loggerErr).Parens(
					jen.Id(loggerCtx).
						Id(loggerCmdName).
						Id("\n").Id(`"Error generate entity",`).
						Id(logErr),
				),
				jen.Return(jen.Id("err")),
			),
			jen.Line(),
			jen.Id("err").Id("=").Id("svc").Dot("repo").Dot(repoName).Dot("Update").Parens(jen.Id("ctx, id, entityDTO")),
			jen.If(jen.Id("err").Op("!=").Nil()).Block(
				jen.Id(loggerErr).Parens(
					jen.Id(loggerCtx).
						Id(loggerCmdName).
						Id("\n").Id(`"Error update",`).
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

	logger.Logger.Warn("internal update service already created")
	err = errors.New("duplicate internal update service name")
	return err
}

func generateDeleteSvc(name, path, services string, fields map[string]string) error {
	var (
		file       = jen.NewFilePathName(path, strings.ToLower(name)+"_service")
		upperName  = capitalize(name)
		title      = sentences(name)
		dir        = name
		methodName = fmt.Sprintf("Delete%s", upperName)
		repoName   = fmt.Sprintf("%sRepo", upperName)
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
	file.Func().Params(jen.Id("svc").Id(embedStruct)).Id(methodName).
		Params(jen.Id("ctx").Id(ctx), jen.Id("id").Id(fields["id"])).Error().
		Block(
			jen.Const().Id("commandName").Op("=").Lit("SVC-"+strcase.ToScreamingKebab(methodName)),
			jen.Id(loggerInfo).Parens(
				jen.Id(loggerCtx).
					Id(loggerCmdName).
					Id("\n"+fmt.Sprintf(`"Delete %s process...",`, title)).
					Id("\n").Nil().Id(",\n"),
			),
			jen.Line(),
			jen.Id("err").Id(":=").Id("svc").Dot("repo").Dot(repoName).Dot("Delete").Id("(ctx, id)"),
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

	logger.Logger.Warn("internal delete service already created")
	err = errors.New("duplicate internal delete service name")
	return err
}

func generateCustomSvc(dto dtoModule) error {
	var (
		file                  = jen.NewFilePathName(dto.path, strings.ToLower(dto.name)+"_service")
		title                 = sentences(dto.methodName)
		dir                   = strcase.ToSnake(dto.methodName)
		methodName            = capitalize(dto.methodName)
		isExists              = new(isExists)
		generatedCustomParams = parseCustomJenCodeFields(dto.params, dto.arrParams, isExists, false)
		generatedCustomReturn = parseCustomJenCodeFields(dto.returns, dto.arrReturn, isExists, true)
		returnVar             = strings.Join(dto.arrReturn[:], ",")
		err                   error
		params                = []jen.Code{jen.Id("ctx").Id(ctx)}
	)
	params = append(params, generatedCustomParams...)
	dto.name = lowerize(dto.name)

	var (
		interactorName = fmt.Sprintf("%sInteractor", dto.name)
		embedStruct    = fmt.Sprintf("*%s", interactorName)
		fieldList      = "\n"
	)

	for _, field := range dto.arrReturn {
		fieldList += fmt.Sprintf("%s %s\n", field, dto.returns[field])
	}

	importList := jen.Id("\n").Id(`"context"`).Id("\n")

	if isExists.isTimeExists {
		importList = importList.Id(`"time"`)
	}

	if isExists.isEntityExists {
		importList = importList.Id(fmt.Sprintf(`"%s/services/%s/domain/entity"`, projectName, dto.services)).Id("\n")
	}

	file.Add(jen.Id("import").Parens(
		importList.Id(`"github.com/Muruyung/go-utilities/logger"`).Id("\n"),
	))

	blockCode := []jen.Code{
		jen.Const().Id("commandName").Op("=").Lit("SVC-" + strcase.ToScreamingKebab(methodName)),
		jen.Id(loggerInfo).Parens(
			jen.Id(loggerCtx).
				Id(loggerCmdName).
				Id("\n" + fmt.Sprintf(`"%s process...",`, title)).
				Id("\n").Nil().Id(",\n"),
		),
		jen.Line(),
		jen.Var().Parens(jen.Id(fieldList)),
		jen.Line(),
		jen.Comment("TODO: Implement code here"),
	}
	// 	jen.Var().Id("query").Op("=").Id("goutils.NewQueryBuilder()"),
	// 	jen.Id(returnVar).Id(":=").Id("svc").Dot("repo").Dot(repoName).Dot(methodName).Id("(ctx, query)"),
	// }

	// if isExists.isError {
	// 	blockCode = append(blockCode,
	// 		jen.If(jen.Id("err").Op("!=").Nil()).Block(
	// 			jen.Id(loggerErr).Parens(
	// 				jen.Id(loggerCtx).
	// 					Id(loggerCmdName).
	// 					Id("\n").Id(fmt.Sprintf(`"Error %s",`, title)).
	// 					Id(logErr),
	// 			),
	// 			jen.Return().Id(returnVar),
	// 		),
	// 	)
	// }

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

	file.Commentf("%s %s service", methodName, title)
	file.Func().Params(jen.Id("svc").Id(embedStruct)).Id(methodName).
		Params(params...).
		Parens(jen.List(generatedCustomReturn...)).
		Block(blockCode...)

	dir = fmt.Sprintf("%s/%s.go", dto.path, dir)
	if _, err = os.Stat(dir); os.IsNotExist(err) {
		return file.Save(dir)
	}

	logger.Logger.Warn("internal custom service already created")
	err = errors.New("duplicate internal custom service name")
	return err
}
