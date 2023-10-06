package modulegenerator

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"

	"github.com/Muruyung/go-utilities/logger"
)

func modGen() {
	var (
		fields, arrFields  = parseFields(varField, true)
		methods            = parseMethods(varMethod)
		params, arrParams  = parseFields(varParam, false)
		returns, arrReturn = parseFields(varReturn, false)
		isUseReturn        = varReturn != ""
		isEmptyFields      = varField == ""
		isUseEntity        = !isEmptyFields && !isWithoutEntity
		isAll              = !isRepoOnly && !isServiceOnly && !isUseCaseOnly && !isModelsOnly && !isEntityOnly
		separator          = string(os.PathSeparator)
		cqrsType           string
	)

	fmt.Println(varParam)
	fmt.Println(params)
	fmt.Println(arrParams)

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
		logger.Logger.Warnf("Go get %s", fmt.Sprintf(defaultErr, command.Stderr))
	}

	command = exec.Command("sh", "-c", "goimports -w **/*.go")
	command.Stdout = &out
	command.Stderr = &ioErr
	err = command.Run()
	if err != nil {
		logger.Logger.Warnf("Goimports %s", fmt.Sprintf(defaultErr, command.Stderr))
	}
}
