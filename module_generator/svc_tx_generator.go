package modulegenerator

import (
	"fmt"
	"os"

	"github.com/Muruyung/go-utilities/logger"
	"github.com/dave/jennifer/jen"
)

func svcTxGenerator(dto dtoModule) error {
	dto.path += "service" + dto.sep + "svc_tx"
	var (
		err error
	)

	if _, err = os.Stat(dto.path); os.IsNotExist(err) {
		err = os.MkdirAll(dto.path, 0777)
		if err != nil {
			return err
		} else {
			logger.Logger.Info("internal directory svc_tx created")
		}
	} else {
		logger.Logger.Warn("internal directory svc_tx already created")
		return nil
	}

	err = generateInitSvcTx(dto.path, dto.services)
	if err != nil {
		return err
	} else {
		logger.Logger.Info("internal init svcTx created")
	}

	err = generateSvcBeginTx(dto.path, dto.services)
	if err != nil {
		return err
	} else {
		logger.Logger.Info("internal init svcTx created")
	}

	return nil
}

func generateInitSvcTx(path, services string) error {
	var (
		file = jen.NewFilePathName(path, "svctx")
		dir  = "init"
	)

	file.Add(jen.Id("import").Parens(
		jen.Id("\n").
			Id(fmt.Sprintf(`"%s/services/%s/domain/repository"`, projectName, services)).Id("\n").
			Id(fmt.Sprintf(`"%s/services/%s/domain/service"`, projectName, services)).Id("\n"),
	))

	var (
		svcName        = "SvcTx"
		interactorName = "svcTxInteractor"
		initName       = "NewSvcTx"
	)
	file.Type().Id(interactorName).Struct(
		jen.Id("repo").Id("*repository.Wrapper"),
	)

	file.Commentf("%s initialize new service transaction", initName)
	file.Func().Id(initName).Params(jen.Id("repo").Id("*repository.Wrapper")).Id("service").Dot(svcName).Block(
		jen.Return(jen.Id("&" + interactorName).Block(
			jen.Id("repo").Op(":").Id("repo,"),
		)),
	)

	return file.Save(path + "/" + dir + ".go")
}

func generateSvcBeginTx(path, services string) error {
	var (
		file           = jen.NewFilePathName(path, "svctx")
		dir            = "begin_tx"
		methodName     = "BeginTx"
		interactorName = "svcTxInteractor"
		embedStruct    = fmt.Sprintf("*%s", interactorName)
	)

	file.Add(jen.Id("import").Parens(
		jen.Id("\n").
			Lit("context").Id("\n").
			Line().
			Id(fmt.Sprintf(`"%s/services/%s/domain/repository"`, projectName, services)).Id("\n").
			Id(fmt.Sprintf(`"%s/services/%s/domain/service"`, projectName, services)).Id("\n").
			Id(fmt.Sprintf(`svc "%s/services/%s/internal/service"`, projectName, services)).Id("\n"),
	))

	file.Commentf("%s begin service transaction", methodName)
	file.Func().Params(jen.Id("svcTx").Id(embedStruct)).Id(methodName).Params(
		jen.Id("ctx").Id("context.Context"), jen.Id("operation").Func().Params(
			jen.Id("ctx").Id("context.Context"), jen.Id("svc").Id("*service.Wrapper"),
		).Error(),
	).Error().Block(
		jen.Return(jen.Id("svcTx.repo.BeginTx").Params(
			jen.Id("ctx"), jen.Func().Params(
				jen.Id("ctx").Id("context.Context"), jen.Id("repo").Id("*repository.Wrapper"),
			).Error().Block(
				jen.Return(jen.Id("operation").Params(jen.Id("ctx"), jen.Id("svc.Init").Params(jen.Id("repo"), jen.Id("svcTx")))),
			),
		)),
	)

	return file.Save(path + "/" + dir + ".go")
}
