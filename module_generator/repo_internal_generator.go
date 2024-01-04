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

func internalRepoGenerator(dto dtoModule, isAll, isOnly bool) error {
	if !isAll && !isOnly {
		return nil
	}

	dto.path += "repository" + dto.sep + "mysql" + dto.sep + dto.name
	var (
		err error
	)

	if _, err = os.Stat(dto.path); os.IsNotExist(err) {
		err = os.MkdirAll(dto.path, 0777)
		if err != nil {
			return err
		} else {
			logger.Logger.Info("internal directory repository created")
		}
	}

	if _, err = os.Stat(dto.path + "/init.go"); os.IsNotExist(err) {
		err = generateInitRepo(dto.name, dto.path, dto.services, dto.fields)
		if err != nil {
			return err
		} else {
			logger.Logger.Info("internal init repository created")
		}
	}

	if _, ok := dto.methods["get"]; ok {
		err = generateGetRepo(dto.name, dto.path, dto.services)
		if err != nil {
			return err
		} else {
			logger.Logger.Info("internal get repository created")
		}
	}

	if _, ok := dto.methods["getList"]; ok {
		err = generateGetListRepo(dto.name, dto.path, dto.services)
		if err != nil {
			return err
		} else {
			logger.Logger.Info("internal get list repository created")
		}

		err = generateGetCountRepo(dto.name, dto.path, dto.services)
		if err != nil {
			return err
		} else {
			logger.Logger.Info("internal get count repository created")
		}
	}

	if _, ok := dto.methods["create"]; ok {
		err = generateCreateRepo(dto.name, dto.path, dto.services)
		if err != nil {
			return err
		} else {
			logger.Logger.Info("internal create repository created")
		}
	}

	if _, ok := dto.methods["update"]; ok {
		err = generateUpdateRepo(dto.name, dto.path, dto.services, dto.fields["id"])
		if err != nil {
			return err
		} else {
			logger.Logger.Info("internal update repository created")
		}
	}

	if _, ok := dto.methods["delete"]; ok {
		err = generateDeleteRepo(dto.name, dto.path, dto.services, dto.fields["id"])
		if err != nil {
			return err
		} else {
			logger.Logger.Info("internal update repository created")
		}
	}

	// if _, ok := dto.methods["custom"]; ok {
	// 	err = generateCustomRepo(dto)
	// 	if err != nil {
	// 		return err
	// 	} else {
	// 		logger.Logger.Info("internal custom repository created")
	// 	}
	// }

	return nil
}

func generateInitRepo(name, path, services string, fields map[string]string) error {
	var (
		file      = jen.NewFilePathName(path, strings.ToLower(name)+"_repo")
		upperName = capitalize(name)
		title     = sentences(name)
		dir       = "init"
	)

	file.Add(jen.Id("import").Parens(
		jen.Id(fmt.Sprintf(`"%s/services/%s/domain/repository"`, projectName, services)).Id("\n"),
	))

	var (
		repoName       = fmt.Sprintf("%sRepository", upperName)
		interactorName = fmt.Sprintf("mysql%s", repoName)
	)
	file.Type().Id(interactorName).Struct(
		jen.Id("sql").Id("repository.SqlTx"),
	)

	initName := fmt.Sprintf("New%sRepository", upperName)
	file.Commentf("%s initialize new %s repository", initName, title)
	file.Func().Id(initName).Params(jen.Id("db").Id("repository.SqlTx")).Id("repository").Dot(repoName).Block(
		jen.Return(
			jen.Id("&" + interactorName).
				Block(
					jen.Id("sql").Op(":").Id("db,"),
				),
		),
	)

	return file.Save(path + "/" + dir + ".go")
}

func generateGetRepo(name, path, services string) error {
	var (
		file       = jen.NewFilePathName(path, strings.ToLower(name)+"_repo")
		upperName  = capitalize(name)
		title      = sentences(name)
		dir        = name
		entityName = fmt.Sprintf("entity.%s", upperName)
		methodName = "Get"
		err        error
	)
	name = lowerize(name)

	var (
		modelsName     = fmt.Sprintf("models.%sModels", upperName)
		modelVar       = fmt.Sprintf("%sModel", name)
		repoName       = fmt.Sprintf("%sRepository", upperName)
		interactorName = fmt.Sprintf("mysql%s", repoName)
		embedStruct    = fmt.Sprintf("*%s", interactorName)
	)

	file.Add(jen.Id("import").Parens(
		jen.Id("\n").Id(`"context"`).Id("\n").
			Id(`"fmt"`).Id("\n").
			Line().
			Id(fmt.Sprintf(`"%s/pkg/utils"`, projectName)).Id("\n").
			Id(fmt.Sprintf(`"%s/pkg/logger"`, projectName)).Id("\n").
			Id(`"github.com/rocketlaunchr/dbq/v2"`).Id("\n").
			Line().
			Id(fmt.Sprintf(`"%s/services/%s/domain/entity"`, projectName, services)).Id("\n").
			Id(fmt.Sprintf(`"%s/services/%s/internal/repository/mapper"`, projectName, services)).Id("\n").
			Id(fmt.Sprintf(`"%s/services/%s/internal/repository/models"`, projectName, services)).Id("\n"),
	))

	file.Commentf("%s get single data %s", methodName, title)
	file.Func().Params(jen.Id("db").Id(embedStruct)).Id(methodName).
		Params(jen.Id("ctx").Id(ctx), jen.Id("query").Id(utils)).
		Parens(jen.List(jen.Id("*"+entityName), jen.Error())).
		Block(
			jen.Const().Id("commandName").Op("=").Lit("REPO-"+strcase.ToScreamingKebab(methodName)+"-"+strcase.ToScreamingKebab(name)),
			jen.Id(loggerInfo).Parens(
				jen.Id(loggerCtx).
					Id(loggerCmdName).
					Id("\n"+fmt.Sprintf(`"Get %s process...",`, title)).
					Id("\n").Nil().Id(",\n"),
			),
			jen.Line(),
			jen.Var().Parens(
				jen.Id("\nerr").Error().Id("\n").
					Id("tableName").Id("=").Id(modelsName+"{}").Dot(getTableName).Id("\n").
					Id(name).Id("=").Id("&"+entityName+"{}").Id("\n").
					Id(modelVar).Id("*"+modelsName).Id("\n"),
			),
			jen.Line(),
			jen.Id("query").Dot("AddPagination").Parens(jen.Id("utils.NewPagination(1, 1)")),
			jen.Id("query").Dot(`AddWhere("deleted_at", "=", nil)`),
			jen.Id("stmt, val, _").Id(":=").Id("query").Dot(`GetQuery(tableName, "")`),
			jen.Id("opts").Id(":=").Id(dbqOpts).Block(
				jen.Id("SingleResult:").True().Id(","),
				jen.Id("ConcreteStruct:").Id(modelsName+"{},"),
				jen.Id("DecoderConfig:").Id("dbq.StdTimeConversionConfig(),"),
			),
			jen.Line(),
			jen.Id("result, err").Id(":=").Id("dbq.Q").Parens(jen.Id("ctx, db.sql.DB(), stmt, opts, val...")),
			jen.If(jen.Id("err").Op("!=").Nil()).Block(
				jen.Id(loggerErr).Parens(
					jen.Id(loggerCtx).
						Id(loggerCmdName).
						Id(loggerErrExecQuery).
						Id(logErr),
				),
				jen.Return(jen.Nil(), jen.Id("err")),
			),
			jen.Line(),
			jen.If(jen.Id("result").Op("!=").Nil()).Block(
				jen.Id(modelVar).Id("=").Id("result.").Parens(jen.Id("*"+modelsName)),
				jen.Id(name+"Mapper").Id(":=").Id("mapper").Dot("New"+upperName+"Mapper").Params(jen.Nil(), jen.Id(modelVar)),
				jen.Id(name+"Mapper").Dot("MapModelsToDomain").Parens(jen.Id(name)),
			).Else().Block(
				jen.Id("err").Id("=").Id("fmt.Errorf").Parens(jen.Lit(title+" data not found")),
				jen.Id(loggerErr).Parens(
					jen.Id(loggerCtx).
						Id(loggerCmdName).
						Id("\n\"Data not found\",").
						Id(logErr),
				),
				jen.Return(jen.Nil(), jen.Nil()),
			),
			jen.Line(),
			jen.Id(loggerInfo).Parens(
				jen.Id(loggerCtx).
					Id(loggerCmdName).
					Id("\n"+fmt.Sprintf(`"Get %s success",`, title)).
					Id("\n").Id(modelVar+".GetModelsMap(),\n"),
			),
			jen.Return(jen.Id(name), jen.Nil()),
		)

	dir = path + "/get_" + dir + ".go"
	if _, err = os.Stat(dir); os.IsNotExist(err) {
		return file.Save(dir)
	}

	logger.Logger.Warn("internal get repository already created")
	err = errors.New("duplicate internal get repository name")
	return err
}

func generateGetListRepo(name, path, services string) error {
	var (
		file       = jen.NewFilePathName(path, strings.ToLower(name)+"_repo")
		upperName  = capitalize(name)
		title      = sentences(name)
		dir        = name
		entityName = fmt.Sprintf("entity.%s", upperName)
		methodName = "GetList"
		err        error
	)
	name = lowerize(name)

	var (
		modelsName     = fmt.Sprintf("models.%sModels", upperName)
		modelsVar      = fmt.Sprintf("%sModels", name)
		repoName       = fmt.Sprintf("%sRepository", upperName)
		interactorName = fmt.Sprintf("mysql%s", repoName)
		embedStruct    = fmt.Sprintf("*%s", interactorName)
	)

	file.Add(jen.Id("import").Parens(
		jen.Id("\n").Id(`"context"`).Id("\n").
			Id(`"fmt"`).Id("\n").
			Line().
			Id(fmt.Sprintf(`"%s/pkg/utils"`, projectName)).Id("\n").
			Id(fmt.Sprintf(`"%s/pkg/logger"`, projectName)).Id("\n").
			Id(`"github.com/rocketlaunchr/dbq/v2"`).Id("\n").
			Line().
			Id(fmt.Sprintf(`"%s/services/%s/domain/entity"`, projectName, services)).Id("\n").
			Id(fmt.Sprintf(`"%s/services/%s/internal/repository/mapper"`, projectName, services)).Id("\n").
			Id(fmt.Sprintf(`"%s/services/%s/internal/repository/models"`, projectName, services)).Id("\n"),
	))

	file.Commentf("%s get list data %s", methodName, title)
	file.Func().Params(jen.Id("db").Id(embedStruct)).Id(methodName).
		Params(jen.Id("ctx").Id(ctx), jen.Id("query").Id(utils)).
		Parens(jen.List(jen.Id("[]*"+entityName), jen.Error())).
		Block(
			jen.Const().Id("commandName").Op("=").Lit("REPO-"+strcase.ToScreamingKebab(methodName)+"-"+strcase.ToScreamingKebab(name)),
			jen.Id(loggerInfo).Parens(
				jen.Id(loggerCtx).
					Id(loggerCmdName).
					Id("\n"+fmt.Sprintf(`"Get list %s process...",`, title)).
					Id("\n").Nil().Id(",\n"),
			),
			jen.Line(),
			jen.Var().Parens(
				jen.Id("\nerr").Error().Id("\n").
					Id("tableName").Id("=").Id(modelsName+"{}").Dot(getTableName).Id("\n").
					Id("data").Id("=").Id("make([]interface{}, 0)\n"),
			),
			jen.Line(),
			jen.Id("query").Dot(`AddWhere("deleted_at", "=", nil)`),
			jen.Id("stmt, val, _").Id(":=").Id("query").Dot(`GetQuery(tableName, "")`),
			jen.Id("opts").Id(":=").Id(dbqOpts).Block(
				jen.Id("SingleResult:").False().Id(","),
				jen.Id("ConcreteStruct:").Id(modelsName+"{},"),
				jen.Id("DecoderConfig:").Id("dbq.StdTimeConversionConfig(),"),
			),
			jen.Line(),
			jen.Id("result, err").Id(":=").Id("dbq.Q").Parens(jen.Id("ctx, db.sql.DB(), stmt, opts, val...")),
			jen.If(jen.Id("err").Op("!=").Nil()).Block(
				jen.Id(loggerErr).Parens(
					jen.Id(loggerCtx).
						Id(loggerCmdName).
						Id(loggerErrExecQuery).
						Id(logErr),
				),
				jen.Return(jen.Nil(), jen.Id("err")),
			),
			jen.Line(),
			jen.Id(modelsVar).Id(":=").Id(fmt.Sprintf("result.([]*%s)", modelsName)),
			jen.If(jen.Len(jen.Id(modelsVar)).Id("== 0")).Block(
				jen.Id("err").Id("=").Id("fmt.Errorf").Parens(jen.Lit(title+" list data not found")),
				jen.Id(loggerErr).Parens(
					jen.Id(loggerCtx).
						Id(loggerCmdName).
						Id("\n\"Data not found\",").
						Id(logErr),
				),
				jen.Return(jen.Nil(), jen.Nil()),
			),
			jen.Line(),
			jen.Id(name).Id(":=").Make(jen.Id("[]*"+entityName), jen.Len(jen.Id(modelsVar))),
			jen.For(jen.Id("key, val").Op(":=").Range().Id(modelsVar)).Block(
				jen.Id("data").Id("=").Append(jen.Id("data"), jen.Id("val.GetModelsMap()")),
				jen.Id(name+"[key]").Id("=").New(jen.Id(entityName)),
				jen.Id(name+"Mapper").Id(":=").Id("mapper").Dot("New"+upperName+"Mapper").Params(jen.Nil(), jen.Id("val")),
				jen.Id(name+"Mapper").Dot("MapModelsToDomain").Parens(jen.Id(name+"[key]")),
			),
			jen.Line(),
			jen.Id(loggerInfo).Parens(
				jen.Id(loggerCtx).
					Id(loggerCmdName).
					Id("\n"+fmt.Sprintf(`"Get list %s success",`, title)).
					Id("\n").Id("data,\n"),
			),
			jen.Return(jen.Id(name), jen.Nil()),
		)

	dir = path + "/get_list_" + dir + ".go"
	if _, err = os.Stat(dir); os.IsNotExist(err) {
		return file.Save(dir)
	}

	logger.Logger.Warn("internal get list repository already created")
	err = errors.New("duplicate internal get list repository name")
	return err
}

func generateGetCountRepo(name, path, services string) error {
	var (
		file       = jen.NewFilePathName(path, strings.ToLower(name)+"_repo")
		upperName  = capitalize(name)
		title      = sentences(name)
		dir        = name
		methodName = "GetCount"
		mapInt     = "map[string]int"
		err        error
	)
	name = lowerize(name)

	var (
		modelsName     = fmt.Sprintf("models.%sModels", upperName)
		repoName       = fmt.Sprintf("%sRepository", upperName)
		interactorName = fmt.Sprintf("mysql%s", repoName)
		embedStruct    = fmt.Sprintf("*%s", interactorName)
	)

	file.Add(jen.Id("import").Parens(
		jen.Id("\n").Id(`"context"`).Id("\n").
			Id(`"fmt"`).Id("\n").
			Line().
			Id(fmt.Sprintf(`"%s/pkg/utils"`, projectName)).Id("\n").
			Id(fmt.Sprintf(`"%s/pkg/logger"`, projectName)).Id("\n").
			Id(`"github.com/rocketlaunchr/dbq/v2"`).Id("\n").
			Line().
			Id(fmt.Sprintf(`"%s/services/%s/internal/repository/models"`, projectName, services)).Id("\n"),
	))

	file.Commentf("%s get count data %s", methodName, title)
	file.Func().Params(jen.Id("db").Id(embedStruct)).Id(methodName).
		Params(jen.Id("ctx").Id(ctx), jen.Id("query").Id(utils)).
		Parens(jen.List(jen.Int(), jen.Error())).
		Block(
			jen.Const().Id("commandName").Op("=").Lit("REPO-"+strcase.ToScreamingKebab(methodName)+"-"+strcase.ToScreamingKebab(name)),
			jen.Id(loggerInfo).Parens(
				jen.Id(loggerCtx).
					Id(loggerCmdName).
					Id("\n"+fmt.Sprintf(`"Get count %s process...",`, title)).
					Id("\n").Lit(0).Id(",\n"),
			),
			jen.Line(),
			jen.Var().Parens(
				jen.Id("\nerr").Error().Id("\n").
					Id("tableName").Id("=").Id(modelsName+"{}").Dot(getTableName).Id("\n").
					Id("count").Id("*"+mapInt).Id("\n"),
			),
			jen.Line(),
			jen.Id("query").Dot(`AddCount("id", "count")`),
			jen.Id("query").Dot(`AddWhere("deleted_at", "=", nil)`),
			jen.Id("stmt, val, _").Id(":=").Id("query").Dot(`GetQuery(tableName, "")`),
			jen.Id("opts").Id(":=").Id(dbqOpts).Block(
				jen.Id("SingleResult:").True().Id(","),
				jen.Id("ConcreteStruct:").Id(mapInt+"{},"),
				jen.Id("DecoderConfig:").Id("dbq.StdTimeConversionConfig(),"),
			),
			jen.Line(),
			jen.Id("result, err").Id(":=").Id("dbq.Q").Parens(jen.Id("ctx, db.sql.DB(), stmt, opts, val...")),
			jen.If(jen.Id("err").Op("!=").Nil()).Block(
				jen.Id(loggerErr).Parens(
					jen.Id(loggerCtx).
						Id(loggerCmdName).
						Id(loggerErrExecQuery).
						Id(logErr),
				),
				jen.Return(jen.Lit(0), jen.Id("err")),
			),
			jen.Line(),
			jen.If(jen.Id("result").Op("!=").Nil()).Block(
				jen.Id("count").Id("=").Id("result.").Parens(jen.Id("*"+mapInt)),
			).Else().Block(
				jen.Id("err").Id("=").Id("fmt.Errorf").Parens(jen.Lit(title+" data not found")),
				jen.Id(loggerErr).Parens(
					jen.Id(loggerCtx).
						Id(loggerCmdName).
						Id("\n\"Data not found\",").
						Id(logErr),
				),
				jen.Return(jen.Lit(0), jen.Nil()),
			),
			jen.Line(),
			jen.Id(loggerInfo).Parens(
				jen.Id(loggerCtx).
					Id(loggerCmdName).
					Id("\n"+fmt.Sprintf(`"Get count %s success",`, title)).
					Id("\n").Id("count,\n"),
			),
			jen.Return(jen.Id(`(*count)["count"]`), jen.Nil()),
		)

	dir = path + "/get_count_" + dir + ".go"
	if _, err = os.Stat(dir); os.IsNotExist(err) {
		return file.Save(dir)
	}

	logger.Logger.Warn("internal get count repository already created")
	err = errors.New("duplicate internal get count repository name")
	return err
}

func generateCreateRepo(name, path, services string) error {
	var (
		file       = jen.NewFilePathName(path, strings.ToLower(name)+"_repo")
		upperName  = capitalize(name)
		title      = sentences(name)
		entityName = fmt.Sprintf("entity.%s", upperName)
		dir        = name
		methodName = "Save"
		err        error
	)
	name = lowerize(name)

	var (
		modelsName     = fmt.Sprintf("models.%sModels", upperName)
		repoName       = fmt.Sprintf("%sRepository", upperName)
		interactorName = fmt.Sprintf("mysql%s", repoName)
		embedStruct    = fmt.Sprintf("*%s", interactorName)
	)

	file.Add(jen.Id("import").Parens(
		jen.Id("\n").Id(`"context"`).Id("\n").
			Id(`"time"`).Id("\n").
			Line().
			Id(fmt.Sprintf(`"%s/pkg/logger"`, projectName)).Id("\n").
			Id(`"github.com/rocketlaunchr/dbq/v2"`).Id("\n").
			Line().
			Id(fmt.Sprintf(`"%s/services/%s/domain/entity"`, projectName, services)).Id("\n").
			Id(fmt.Sprintf(`"%s/services/%s/internal/repository/mapper"`, projectName, services)).Id("\n").
			Id(fmt.Sprintf(`"%s/services/%s/internal/repository/models"`, projectName, services)).Id("\n"),
	))

	file.Commentf("%s create data %s", methodName, title)
	file.Func().Params(jen.Id("db").Id(embedStruct)).Id(methodName).
		Params(jen.Id("ctx").Id(ctx), jen.Id("data").Id("*"+entityName)).
		Error().
		Block(
			jen.Const().Id("commandName").Op("=").Lit("REPO-"+strcase.ToScreamingKebab(methodName)+"-"+strcase.ToScreamingKebab(name)),
			jen.Id(loggerInfo).Parens(
				jen.Id(loggerCtx).
					Id(loggerCmdName).
					Id("\n"+fmt.Sprintf(`"Save %s process...",`, title)).
					Id("\n").Nil().Id(",\n"),
			),
			jen.Line(),
			jen.Id("data, _ = data.SetCreatedAt(time.Now())"),
			jen.Var().Parens(
				jen.Id("\nerr").Error().Id("\n").
					Id("tableName").Id("=").Id(modelsName+"{}").Dot(getTableName).Id("\n").
					Id(name+"Mapper").Id("=").Id("mapper").Dot("New"+upperName+"Mapper(data, nil).MapDomainToModels()\n").
					Id("arrColumn").Id("=").Id(name+"Mapper.GetColumns()\n").
					Id("arrValue").Id("=").Id(name+"Mapper.GetValStruct(arrColumn)\n").
					Id("sqlDB").Id("dbq.ExecContexter"),
			),
			jen.Line(),
			jen.If(jen.Id("db.sql.Session().UseTx")).Block(
				jen.Id("sqlDB").Op("=").Id("db.sql.Session().Tx"),
			).Else().Block(
				jen.Id("sqlDB").Op("=").Id("db.sql.DB()"),
			),
			jen.Line(),
			jen.Id("ctx, cancel := context.WithTimeout(ctx, 60*time.Second)"),
			jen.Defer().Id("cancel()"),
			jen.Line(),
			jen.Id("stmt := dbq.INSERTStmt(tableName, arrColumn, len(arrValue), dbq.MySQL)"),
			jen.Id("_, err = dbq.E(ctx, sqlDB, stmt, nil, arrValue)"),
			jen.If(jen.Id("err").Op("!=").Nil()).Block(
				jen.Id(loggerErr).Parens(
					jen.Id(loggerCtx).
						Id(loggerCmdName).
						Id(loggerErrExecQuery).
						Id(logErr),
				),
				jen.Return(jen.Id("err")),
			),
			jen.Line(),
			jen.Id(loggerInfo).Parens(
				jen.Id(loggerCtx).
					Id(loggerCmdName).
					Id("\n"+fmt.Sprintf(`"Save %s success",`, title)).
					Id("\n").Nil().Id(",\n"),
			),
			jen.Line(),
			jen.Return().Nil(),
		)

	dir = path + "/save_" + dir + ".go"
	if _, err = os.Stat(dir); os.IsNotExist(err) {
		return file.Save(dir)
	}

	logger.Logger.Warn("internal save repository already created")
	err = errors.New("duplicate internal save repository name")
	return err
}

func generateUpdateRepo(name, path, services string, fieldID string) error {
	var (
		file       = jen.NewFilePathName(path, strings.ToLower(name)+"_repo")
		upperName  = capitalize(name)
		title      = sentences(name)
		entityName = fmt.Sprintf("entity.%s", upperName)
		dir        = name
		methodName = "Update"
		err        error
	)
	name = lowerize(name)

	var (
		modelsName     = fmt.Sprintf("models.%sModels", upperName)
		modelsVar      = fmt.Sprintf("%sModels", name)
		repoName       = fmt.Sprintf("%sRepository", upperName)
		interactorName = fmt.Sprintf("mysql%s", repoName)
		embedStruct    = fmt.Sprintf("*%s", interactorName)
	)

	file.Add(jen.Id("import").Parens(
		jen.Id("\n").Id(`"context"`).Id("\n").
			Id(`"fmt"`).Id("\n").
			Id(`"time"`).Id("\n").
			Line().
			Id(fmt.Sprintf(`"%s/pkg/logger"`, projectName)).Id("\n").
			Id(`"github.com/rocketlaunchr/dbq/v2"`).Id("\n").
			Line().
			Id(fmt.Sprintf(`"%s/services/%s/domain/entity"`, projectName, services)).Id("\n").
			Id(fmt.Sprintf(`"%s/services/%s/internal/repository/mapper"`, projectName, services)).Id("\n").
			Id(fmt.Sprintf(`"%s/services/%s/internal/repository/models"`, projectName, services)).Id("\n"),
	))

	file.Commentf("%s update data %s", methodName, title)
	file.Func().Params(jen.Id("db").Id(embedStruct)).Id(methodName).
		Params(jen.Id("ctx").Id(ctx), jen.Id("id").Id(fieldID), jen.Id("data").Id("*"+entityName)).
		Error().
		Block(
			jen.Const().Id("commandName").Op("=").Lit("REPO-"+strcase.ToScreamingKebab(methodName)+"-"+strcase.ToScreamingKebab(name)),
			jen.Id(loggerInfo).Parens(
				jen.Id(loggerCtx).
					Id(loggerCmdName).
					Id("\n"+fmt.Sprintf(`"Update %s process...",`, title)).
					Id("\n").Nil().Id(",\n"),
			),
			jen.Line(),
			jen.Var().Parens(
				jen.Id("\nerr").Error().Id("\n").
					Id("tableName").Id("=").Id(modelsName+"{}").Dot(getTableName).Id("\n").
					Id(name+"Mapper").Id("=").Id("mapper").Dot("New"+upperName+"Mapper(data, nil).MapDomainToModels()\n").
					Id(modelsVar).Id("=").Id(name+"Mapper.GetModelsMap()\n").
					Id("arrColumn").Id("=").Id(name+"Mapper.GetColumns()\n").
					Id("values").Id("=").Make(jen.Id("[]interface{}"), jen.Lit(0)).Id("\n").
					Id("lastIndex").Id("=").Id("len(arrColumn) - 1\n").
					Id("sqlDB").Id("dbq.ExecContexter"),
			),
			jen.Line(),
			jen.If(jen.Id("db.sql.Session().UseTx")).Block(
				jen.Id("sqlDB").Op("=").Id("db.sql.Session().Tx"),
			).Else().Block(
				jen.Id("sqlDB").Op("=").Id("db.sql.DB()"),
			),
			jen.Line(),
			jen.Id("ctx, cancel := context.WithTimeout(ctx, 60*time.Second)"),
			jen.Defer().Id("cancel()"),
			jen.Line(),
			jen.Id("stmt := fmt.Sprintf(`UPDATE %s SET`, tableName)"),
			jen.For(jen.Id("key,val").Id(":=").Range().Id("arrColumn")).Block(
				jen.If(jen.Id(modelsVar+"[val]").Op("!=").Nil()).Block(
					jen.Id("stmt = fmt.Sprintf(`%s %s = ?`, stmt, val)"),
					jen.Id("values =").Append(jen.Id("values"), jen.Id(modelsVar+"[val]")),
				),
				jen.Line(),
				jen.If(jen.Id("key < lastIndex").Op("&&").Id(modelsVar+"[val]").Op("!=").Nil()).Block(
					jen.Id("stmt += `, `"),
				).Else().If(jen.Id("key == lastIndex")).Block(
					jen.Id("stmt = fmt.Sprintf(`%s WHERE id = ?`, stmt)"),
				),
			),
			jen.Id("values =").Append(jen.Id("values"), jen.Id("id")),
			jen.Line(),
			jen.Id("_, err = dbq.E(ctx, sqlDB, stmt, nil, values...)"),
			jen.If(jen.Id("err").Op("!=").Nil()).Block(
				jen.Id(loggerErr).Parens(
					jen.Id(loggerCtx).
						Id(loggerCmdName).
						Id(loggerErrExecQuery).
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
			jen.Line(),
			jen.Return().Nil(),
		)

	dir = path + "/update_" + dir + ".go"
	if _, err = os.Stat(dir); os.IsNotExist(err) {
		return file.Save(dir)
	}

	logger.Logger.Warn("internal update repository already created")
	err = errors.New("duplicate internal update repository name")
	return err
}

func generateDeleteRepo(name, path, services string, fieldID string) error {
	var (
		file       = jen.NewFilePathName(path, strings.ToLower(name)+"_repo")
		upperName  = capitalize(name)
		title      = sentences(name)
		dir        = name
		methodName = "Delete"
		err        error
	)
	name = lowerize(name)

	var (
		modelsName     = fmt.Sprintf("models.%sModels", upperName)
		repoName       = fmt.Sprintf("%sRepository", upperName)
		interactorName = fmt.Sprintf("mysql%s", repoName)
		embedStruct    = fmt.Sprintf("*%s", interactorName)
	)

	file.Add(jen.Id("import").Parens(
		jen.Id("\n").Id(`"context"`).Id("\n").
			Id(`"fmt"`).Id("\n").
			Id(`"time"`).Id("\n").
			Line().
			Id(`"github.com/Muruyung/go-utilities/converter"`).Id("\n").
			Id(fmt.Sprintf(`"%s/pkg/logger"`, projectName)).Id("\n").
			Id(`"github.com/rocketlaunchr/dbq/v2"`).Id("\n").
			Line().
			Id(fmt.Sprintf(`"%s/services/%s/internal/repository/models"`, projectName, services)).Id("\n"),
	))

	file.Commentf("%s delete data %s", methodName, title)
	file.Func().Params(jen.Id("db").Id(embedStruct)).Id(methodName).
		Params(jen.Id("ctx").Id(ctx), jen.Id("id").Id(fieldID)).
		Error().
		Block(
			jen.Const().Id("commandName").Op("=").Lit("REPO-"+strcase.ToScreamingKebab(methodName)+"-"+strcase.ToScreamingKebab(name)),
			jen.Id(loggerInfo).Parens(
				jen.Id(loggerCtx).
					Id(loggerCmdName).
					Id("\n"+fmt.Sprintf(`"Delete %s process...",`, title)).
					Id("\n").Nil().Id(",\n"),
			),
			jen.Line(),
			jen.Var().Parens(
				jen.Id("\nerr").Error().Id("\n").
					Id("tableName").Id("=").Id(modelsName+"{}").Dot(getTableName).Id("\n").
					Id("sqlDB").Id("dbq.ExecContexter"),
			),
			jen.Line(),
			jen.If(jen.Id("db.sql.Session().UseTx")).Block(
				jen.Id("sqlDB").Op("=").Id("db.sql.Session().Tx"),
			).Else().Block(
				jen.Id("sqlDB").Op("=").Id("db.sql.DB()"),
			),
			jen.Line(),
			jen.Id("ctx, cancel := context.WithTimeout(ctx, 60*time.Second)"),
			jen.Defer().Id("cancel()"),
			jen.Line(),
			jen.Id("stmt := fmt.Sprintf(`UPDATE %s SET deleted_at = ? WHERE id = ?`, tableName)"),
			jen.Id("_, err = dbq.E(ctx, sqlDB, stmt, nil, converter.ConvertDateToString(time.Now()), id)"),
			jen.If(jen.Id("err").Op("!=").Nil()).Block(
				jen.Id(loggerErr).Parens(
					jen.Id(loggerCtx).
						Id(loggerCmdName).
						Id(loggerErrExecQuery).
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
			jen.Line(),
			jen.Return().Nil(),
		)

	dir = path + "/delete_" + dir + ".go"
	if _, err = os.Stat(dir); os.IsNotExist(err) {
		return file.Save(dir)
	}

	logger.Logger.Warn("internal delete repository already created")
	err = errors.New("duplicate internal delete repository name")
	return err
}
