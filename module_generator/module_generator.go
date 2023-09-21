package modulegenerator

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strconv"

	"github.com/iancoleman/strcase"
	"github.com/spf13/cobra"

	"github.com/Muruyung/go-utilities/logger"
)

func modGen(cmd *cobra.Command, args []string) {
	var (
		svcName            = cmd.Flag("service").Value.String()
		name               = strcase.ToSnake(cmd.Flag("name").Value.String())
		isUseEntity        = cmd.Flag("fields").Value.String() != ""
		fields, _          = parseFields(cmd.Flag("fields").Value.String(), true)
		methods            = parseMethods(cmd.Flag("methods").Value.String())
		methodName         = cmd.Flag("custom-method").Value.String()
		isUseReturn        = cmd.Flag("return").Value.String() != ""
		params, arrParams  = parseFields(cmd.Flag("params").Value.String(), false)
		returns, arrReturn = parseFields(cmd.Flag("return").Value.String(), false)
		separator          = string(os.PathSeparator)
		isRepoOnly, _      = strconv.ParseBool(cmd.Flag("repo-only").Value.String())
		isServiceOnly, _   = strconv.ParseBool(cmd.Flag("service-only").Value.String())
		isUseCaseOnly, _   = strconv.ParseBool(cmd.Flag("usecase-only").Value.String())
		isAll              = !isRepoOnly && !isServiceOnly && !isUseCaseOnly
	)

	if !isUseEntity {
		fields = nil
	}

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
			methods:    methods,
			methodName: methodName,
			params:     params,
			arrParams:  arrParams,
			returns:    returns,
			arrReturn:  arrReturn,
		}
	)

	if err := validate(dto); err != nil {
		logger.Logger.Error(err)
		return
	}

	//===============internal generator===============
	dto.path = internalPath
	err = internalUcGenerator(dto, isAll, isUseCaseOnly)
	if err != nil {
		logger.Logger.Error(fmt.Sprintf(defaultErr, err))
		return
	}
	//================================================

	//================domain generator================
	dto.path = domainPath
	if isUseEntity {
		err = entityGenerator(domainPath, separator, name, fields)
		if err != nil {
			logger.Logger.Error(err)
			return
		}

		if cmd.Flag("entity-only").Value.String() == "true" {
			return
		}
	}

	err = domainRepoGenerator(dto, isAll, isRepoOnly)
	if err != nil {
		logger.Logger.Error(fmt.Sprintf(defaultErr, err))
		return
	}

	err = domainSvcGenerator(dto, isAll, isServiceOnly)
	if err != nil {
		logger.Logger.Error(fmt.Sprintf(defaultErr, err))
		return
	}

	err = domainUcGenerator(dto, isAll, isUseCaseOnly)
	if err != nil {
		logger.Logger.Error(fmt.Sprintf(defaultErr, err))
		return
	}
	//================================================

	var (
		out     bytes.Buffer
		command = exec.Command("go", "fmt", "./...")
	)
	command.Stdout = &out
	err = command.Run()
	if err != nil {
		logger.Logger.Errorf(defaultErr, err)
	}
}
