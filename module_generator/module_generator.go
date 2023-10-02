package modulegenerator

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strconv"

	"github.com/Muruyung/go-utilities/logger"
	"github.com/iancoleman/strcase"
	"github.com/spf13/cobra"
)

func modGen(cmd *cobra.Command, args []string) {
	var (
		svcName            = cmd.Flag("service").Value.String()
		name               = strcase.ToSnake(cmd.Flag("name").Value.String())
		fields, arrFields  = parseFields(cmd.Flag("fields").Value.String(), true)
		methods            = parseMethods(cmd.Flag("methods").Value.String())
		methodName         = cmd.Flag("custom-method").Value.String()
		isUseReturn        = cmd.Flag("return").Value.String() != ""
		params, arrParams  = parseFields(cmd.Flag("params").Value.String(), false)
		returns, arrReturn = parseFields(cmd.Flag("return").Value.String(), false)
		separator          = string(os.PathSeparator)
		isModelsOnly, _    = strconv.ParseBool(cmd.Flag("models-only").Value.String())
		isEntityOnly, _    = strconv.ParseBool(cmd.Flag("entity-only").Value.String())
		isRepoOnly, _      = strconv.ParseBool(cmd.Flag("repo-only").Value.String())
		isServiceOnly, _   = strconv.ParseBool(cmd.Flag("service-only").Value.String())
		isUseCaseOnly, _   = strconv.ParseBool(cmd.Flag("usecase-only").Value.String())
		isWithoutUT, _     = strconv.ParseBool(cmd.Flag("no-unit-test").Value.String())
		isWithoutEntity, _ = strconv.ParseBool(cmd.Flag("no-entity").Value.String())
		isCqrs, _          = strconv.ParseBool(cmd.Flag("cqrs").Value.String())
		isCqrsCommand, _   = strconv.ParseBool(cmd.Flag("is-command").Value.String())
		isCqrsQuery, _     = strconv.ParseBool(cmd.Flag("is-query").Value.String())
		isEmptyFields      = cmd.Flag("fields").Value.String() == ""
		isUseEntity        = !isEmptyFields && !isWithoutEntity
		isAll              = !isRepoOnly && !isServiceOnly && !isUseCaseOnly && !isModelsOnly && !isEntityOnly
		cqrsType           string
	)

	if _, ok := methods["custom"]; ok && isCqrs {
		if isCqrsCommand && isCqrsQuery {
			logger.Logger.Error("choose one cqrs type")
			return
		} else if isCqrsQuery {
			cqrsType = "query"
		} else if isCqrsCommand {
			cqrsType = "command"
		} else {
			cqrsType = cqrsTypeCheck(methodName)
		}
	}

	// if !isUseEntity {
	// 	fields = nil
	// }

	if !isUseReturn {
		returns = nil
	}

	currDir, err := os.Getwd()
	if err != nil {
		logger.Logger.Warn(fmt.Sprintf(defaultErr, err))
	}

	err = getProjectName(currDir, separator)
	if err != nil {
		logger.Logger.Error(fmt.Sprintf(defaultErr, err))
		return
	}

	var (
		internalPath = currDir + separator + "services" + separator + svcName + separator + "internal" + separator
		domainPath   = currDir + separator + "services" + separator + svcName + separator + "domain" + separator
		dto          = dtoModule{
			sep:        separator,
			name:       name,
			services:   svcName,
			fields:     fields,
			arrFields:  arrFields,
			methods:    methods,
			methodName: methodName,
			params:     params,
			arrParams:  arrParams,
			returns:    returns,
			arrReturn:  arrReturn,
			entityOnly: isEntityOnly,
		}
	)

	if err := validate(dto); err != nil {
		logger.Logger.Error(err)
		return
	}

	//===============internal generator===============
	dto.path = internalPath

	if isCqrs {
		err = internalCqrsUcGenerator(dto, cqrsType, isAll, isUseCaseOnly)
	} else {
		err = internalUcGenerator(dto, isAll, isUseCaseOnly)
	}
	if err != nil {
		logger.Logger.Error(fmt.Sprintf(defaultErr, err))
		return
	}

	err = internalSvcGenerator(dto, isAll, isServiceOnly)
	if err != nil {
		logger.Logger.Error(fmt.Sprintf(defaultErr, err))
		return
	}

	err = internalRepoGenerator(dto, isAll, isRepoOnly)
	if err != nil {
		logger.Logger.Error(fmt.Sprintf(defaultErr, err))
		return
	}

	if !isEmptyFields {
		err = modelsGenerator(dto, isAll, isModelsOnly || isRepoOnly)
		if err != nil {
			logger.Logger.Error(fmt.Sprintf(defaultErr, err))
			return
		}

		err = mapperGenerator(dto, domainPath, isAll, isModelsOnly || isRepoOnly)
		if err != nil {
			logger.Logger.Error(fmt.Sprintf(defaultErr, err))
			return
		}
	}
	//================================================

	//================domain generator================
	dto.path = domainPath
	if isUseEntity {
		err = entityGenerator(dto, isAll, isEntityOnly)
		if err != nil {
			logger.Logger.Error(fmt.Sprintf(defaultErr, err))
			return
		}
	}

	if isCqrs {
		err = domainCqrsUcGenerator(dto, cqrsType, isAll, isUseCaseOnly)
	} else {
		err = domainUcGenerator(dto, isAll, isUseCaseOnly)
	}
	if err != nil {
		logger.Logger.Error(fmt.Sprintf(defaultErr, err))
		return
	}

	err = domainSvcGenerator(dto, isAll, isServiceOnly)
	if err != nil {
		logger.Logger.Error(fmt.Sprintf(defaultErr, err))
		return
	}

	err = domainRepoGenerator(dto, isAll, isRepoOnly)
	if err != nil {
		logger.Logger.Error(fmt.Sprintf(defaultErr, err))
		return
	}
	//================================================

	var (
		out     bytes.Buffer
		ioErr   bytes.Buffer
		command *exec.Cmd
	)

	if !isWithoutUT {
		command = exec.Command("genut", "--config")
		command.Stdout = &out
		command.Stderr = &ioErr
		err = command.Run()
		if err != nil {
			logger.Logger.Errorf("Genut config %s", fmt.Sprintf(defaultErr, command.Stderr))
		}

		command = exec.Command("genut", "mocks")
		command.Stdout = &out
		command.Stderr = &ioErr
		err = command.Run()
		if err != nil {
			logger.Logger.Errorf("Genut mocks %s", fmt.Sprintf(defaultErr, command.Stderr))
		}
		logger.Logger.Info("mocks created")
	}

	command = exec.Command("go", "fmt", "./...")
	command.Stdout = &out
	command.Stderr = &ioErr
	err = command.Run()
	if err != nil {
		logger.Logger.Errorf("Fmt %s", fmt.Sprintf(defaultErr, command.Stderr))
	}

	command = exec.Command("go", "get", "./...")
	command.Stdout = &out
	command.Stderr = &ioErr
	err = command.Run()
	if err != nil {
		logger.Logger.Errorf("Go get %s", fmt.Sprintf(defaultErr, command.Stderr))
	}

	command = exec.Command("sh", "-c", "goimports -w **/*.go")
	command.Stdout = &out
	command.Stderr = &ioErr
	err = command.Run()
	if err != nil {
		logger.Logger.Errorf("Goimports %s", fmt.Sprintf(defaultErr, command.Stderr))
	}
}
