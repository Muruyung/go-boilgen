package modulegenerator

import (
	"fmt"
	"os"

	"github.com/Muruyung/go-utilities/logger"
	"github.com/dave/jennifer/jen"
)

func sqlTxGenerator(dto dtoModule) error {
	path := dto.path + "repository" + dto.sep + "mysql" + dto.sep
	dto.path = path + "mysql_tx"
	var (
		err error
	)

	if _, err = os.Stat(dto.path); os.IsNotExist(err) {
		err = os.MkdirAll(dto.path, 0777)
		if err != nil {
			return err
		} else {
			logger.Logger.Info("internal directory mysql_tx created")
		}
	} else {
		logger.Logger.Warn("internal directory mysql_tx already created")
		return nil
	}

	err = appendInit("SqlTx", dto.services, path)
	if err != nil {
		return err
	} else {
		logger.Logger.Info("general init repository created")
	}

	err = generateInitMysqlTx(dto.path, dto.services)
	if err != nil {
		return err
	} else {
		logger.Logger.Info("internal init mysqlTx created")
	}

	err = generateRepoBeginTx(dto.path, dto.services)
	if err != nil {
		return err
	} else {
		logger.Logger.Info("internal beginTx repository created")
	}

	err = generateGetDB(dto.path, dto.services)
	if err != nil {
		return err
	} else {
		logger.Logger.Info("internal get DB created")
	}

	err = generateTxSession(dto.path, dto.services)
	if err != nil {
		return err
	} else {
		logger.Logger.Info("internal tx session repository created")
	}

	err = generateRepoWrapper(dto.path, dto.services)
	if err != nil {
		return err
	} else {
		logger.Logger.Info("internal wrapper repository created")
	}

	return nil
}

func generateInitMysqlTx(path, services string) error {
	var (
		file = jen.NewFilePathName(path, "mysqltx")
		dir  = "init"
	)

	file.Add(jen.Id("import").Parens(
		jen.Id(`"database/sql"`).Id("\n").
			Line().
			Id(fmt.Sprintf(`"%s/pkg/database"`, projectName)).Id("\n").
			Id(fmt.Sprintf(`"%s/services/%s/domain/repository"`, projectName, services)).Id("\n").
			Id(fmt.Sprintf(`repository_mysql "%s/services/%s/internal/repository/mysql"`, projectName, services)).Id("\n"),
	))

	var (
		repoName       = "SqlTx"
		interactorName = "mysqlTxRepository"
	)
	file.Type().Id(interactorName).Struct(
		jen.Id("db").Id("*sql.DB"),
		jen.Id("tx").Id("*database.TX"),
		jen.Id("wrapper").Id("*repository.Wrapper"),
	)

	initName := "NewMysqlTx"
	file.Commentf("%s initialize new sqltx repository", initName)
	file.Func().Id(initName).Params(jen.Id("db").Id("*sql.DB")).Id("repository").Dot(repoName).Block(
		jen.Id("tx").Id(":=").Id("&"+interactorName).Block(
			jen.Id("db: db,"),
		),
		jen.Id("tx.wrapper").Op("=").Id("repository_mysql.Init(tx)"),
		jen.Return(
			jen.Id("tx"),
		),
	)

	return file.Save(path + "/" + dir + ".go")
}

func generateRepoBeginTx(path, services string) error {
	var (
		file           = jen.NewFilePathName(path, "mysqltx")
		dir            = "begin_tx"
		methodName     = "BeginTx"
		interactorName = "mysqlTxRepository"
		embedStruct    = fmt.Sprintf("*%s", interactorName)
	)

	file.Add(jen.Id("import").Parens(
		jen.Id("\n").
			Id(`"context"`).Id("\n").
			Id(`"database/sql"`).Id("\n").
			Line().
			Id(fmt.Sprintf(`"%s/pkg/database"`, projectName)).Id("\n").
			Id(fmt.Sprintf(`"%s/pkg/logger"`, projectName)).Id("\n").
			Id(fmt.Sprintf(`"%s/services/%s/domain/repository"`, projectName, services)).Id("\n").
			Id(fmt.Sprintf(`repository_mysql "%s/services/%s/internal/repository/mysql"`, projectName, services)).Id("\n"),
	))

	file.Commentf("%s begin sql transaction repository", methodName)
	file.Func().Params(jen.Id("db").Id(embedStruct)).Id(methodName).
		Params(jen.Id("ctx").Id(ctx), jen.Id("operation").Func().Params(jen.Id(ctx), jen.Id("*repository.Wrapper")).Error()).Error().
		Block(
			jen.Const().Id("commandName").Op("=").Lit("REPO-BEGIN-TRANSACTION"),
			jen.Id(loggerInfo).Parens(
				jen.Id(loggerCtx).
					Id(loggerCmdName).
					Id("\n"+`"Begin transaction process...",`).
					Id("\n").Nil().Id(",\n"),
			),
			jen.Line(),
			jen.Id("rollback").Id(":=").Func().Params(jen.Id("tx").Id("*sql.Tx"), jen.Id("err").Error()).Error().Block(
				jen.Id(loggerWarn).Parens(
					jen.Id(loggerCtx).
						Id(loggerCmdName).
						Id("\n"+`"Transaction rollback process...",`).
						Id(logErr),
				),
				jen.Line(),
				jen.Id("err").Op("=").Id("tx.Rollback()"),
				jen.If(jen.Id("err == nil || err == sql.ErrTxDone || err == sql.ErrConnDone")).Block(
					jen.Id(loggerInfo).Parens(
						jen.Id(loggerCtx).
							Id(loggerCmdName).
							Id("\n"+`"Transaction rollback success",`).
							Id("\n").Nil().Id(",\n"),
					),
					jen.Return().Nil(),
				),
				jen.Line(),
				jen.Id(loggerErr).Parens(
					jen.Id(loggerCtx).
						Id(loggerCmdName).
						Id("\n"+`"Transaction rollback failed",`).
						Id(logErr),
				),
				jen.Return(jen.Id("err")),
			),
			jen.Line(),
			jen.Var().Parens(
				jen.Id("\n").
					Id("//ctxTx").Op("=").Id("context.WithValue").Params(jen.Id("ctx"), jen.Id("logger.IsUseES"), jen.True()).Id("\n").
					Id("ctxTx").Op("=").Id("ctx").Id("\n").
					Id("dbTx").Op("=").Id("db").Id("\n").
					Id("tx").Id("*sql.Tx").Id("\n").
					Id("err").Error().Id("\n").
					Id("isSession").Op("=").False().Id("\n"),
			),
			jen.Line(),
			jen.If(jen.Id("db.tx != nil && db.tx.UseTx")).Block(
				jen.Id("tx").Op("=").Id("db.tx.Tx"),
				jen.Id("isSession").Op("=").True(),
			).Else().Block(
				jen.Id("tx, err").Op("=").Id("db.db.Begin").Params(),
				jen.If(jen.Id("err").Op("!=").Nil()).Block(
					jen.Id(loggerErr).Parens(
						jen.Id(loggerCtx).
							Id(loggerCmdName).
							Id("\n"+`"Begin transaction failed",`).
							Id(logErr),
					),
					jen.Return(jen.Id("err")),
				),
				jen.Line(),
				jen.Id("dbTx").Op("=").Id("&"+interactorName).Block(
					jen.Id("db").Op(":").Id("db.db,"),
					jen.Id("tx").Op(":").Id("&database.TX").Block(
						jen.Id("Tx").Op(":").Id("tx,"),
						jen.Id("UseTx").Op(":").True().Id(","),
					).Id(","),
				),
				jen.Line(),
				jen.Id("dbTx.wrapper").Op("=").Id("repository_mysql.Init").Params(jen.Id("dbTx")),
			),
			jen.Line(),
			jen.Id("err").Op("=").Id("operation").Params(jen.Id("ctxTx"), jen.Id("dbTx.Wrapper").Params()),
			jen.If(jen.Id("err").Op("!=").Nil()).Block(
				jen.Id("errRollback").Id(":=").Id("rollback").Params(jen.Id("tx"), jen.Id("err")),
				jen.If(jen.Id("errRollback").Op("!=").Nil()).Block(
					jen.Return(jen.Id("errRollback")),
				),
				jen.Return(jen.Id("err")),
			),
			jen.Line(),
			jen.Id(loggerInfo).Parens(
				jen.Id(loggerCtx).
					Id(loggerCmdName).
					Id("\n"+`"Transaction commit process...",`).
					Id("\n").Nil().Id(",\n"),
			),
			jen.Line(),
			jen.If(jen.Id("isSession")).Block(
				jen.Return().Nil(),
			),
			jen.Line(),
			jen.Id("err").Op("=").Id("tx.Commit").Params(),
			jen.If(jen.Id("err").Op("!=").Nil()).Block(
				jen.Id("err").Op("=").Id("rollback").Params(jen.Id("tx"), jen.Id("err")),
				jen.Return(jen.Id("err")),
			),
			jen.Line(),
			jen.Id("dbTx.tx.UseTx").Op("=").False(),
			jen.Line(),
			jen.Id(loggerInfo).Parens(
				jen.Id(loggerCtx).
					Id(loggerCmdName).
					Id("\n"+`"Transaction commit success",`).
					Id("\n").Nil().Id(",\n"),
			),
			jen.Return().Nil(),
		)

	return file.Save(path + "/" + dir + ".go")
}

func generateGetDB(path, services string) error {
	var (
		file           = jen.NewFilePathName(path, "mysqltx")
		dir            = "get_db"
		methodName     = "DB"
		interactorName = "mysqlTxRepository"
		embedStruct    = fmt.Sprintf("*%s", interactorName)
	)

	file.Add(jen.Id("import").Id(`"database/sql"`))

	file.Commentf("%s get sql db engine", methodName)
	file.Func().Params(jen.Id("db").Id(embedStruct)).Id(methodName).Params().Id("*sql.DB").Block(
		jen.Return(jen.Id("db.db")),
	)

	return file.Save(path + "/" + dir + ".go")
}

func generateTxSession(path, services string) error {
	var (
		file           = jen.NewFilePathName(path, "mysqltx")
		dir            = "session"
		methodName     = "Session"
		interactorName = "mysqlTxRepository"
		embedStruct    = fmt.Sprintf("*%s", interactorName)
	)

	file.Add(jen.Id("import").Id(fmt.Sprintf(`"%s/pkg/database"`, projectName)))

	file.Commentf("%s get tx session", methodName)
	file.Func().Params(jen.Id("db").Id(embedStruct)).Id(methodName).Params().Id("*database.TX").Block(
		jen.If(jen.Id("db.tx.UseTx")).Block(
			jen.Return(jen.Id("db.tx")),
		),
		jen.Line(),
		jen.Return(jen.Id("&database.TX").Block(
			jen.Id("UseTx").Op(":").False().Id(","),
		)),
	)

	return file.Save(path + "/" + dir + ".go")
}

func generateRepoWrapper(path, services string) error {
	var (
		file           = jen.NewFilePathName(path, "mysqltx")
		dir            = "wrapper"
		methodName     = "Wrapper"
		interactorName = "mysqlTxRepository"
		embedStruct    = fmt.Sprintf("*%s", interactorName)
	)

	file.Add(jen.Id("import").Id(fmt.Sprintf(`"%s/services/%s/domain/repository"`, projectName, services)))

	file.Commentf("%s get repository wrapper", methodName)
	file.Func().Params(jen.Id("db").Id(embedStruct)).Id(methodName).Params().Id("*repository.Wrapper").Block(
		jen.Return(jen.Id("db.wrapper")),
	)

	return file.Save(path + "/" + dir + ".go")
}
