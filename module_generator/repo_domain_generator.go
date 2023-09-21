package modulegenerator

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/dave/jennifer/jen"

	"github.com/Muruyung/go-utilities/logger"
)

func domainRepoGenerator(dto dtoModule, isAll, isOnly bool) error {
	if !isAll && !isOnly {
		return nil
	}

	dto.path += "repository" + dto.sep
	var err error

	if _, err = os.Stat(dto.path + dto.name + ".go"); os.IsNotExist(err) {
		err = generateDomainRepo(dto)
		if err != nil {
			logger.Logger.Error(fmt.Sprintf(defaultErr, err))
			return err
		} else {
			logger.Logger.Info("domain repository created")
		}

		var (
			upperName      = capitalize(dto.name)
			interfaceShort = fmt.Sprintf("%sRepo", upperName)
			interfaceName  = fmt.Sprintf("%sRepository", upperName)
		)
		err = generateWrapper(interfaceShort, interfaceName, dto.path)
		if err != nil {
			logger.Logger.Error(fmt.Sprintf(defaultErr, err))
			return err
		}
		logger.Logger.Info("domain repository wrapper created")
	} else {
		return appendDomainRepo(dto.path+dto.name+".go", dto)
	}

	return nil
}

func generateDomainRepo(dto dtoModule) error {
	var (
		file                  = jen.NewFilePathName(dto.path, "repository")
		upperName             = capitalize(dto.name)
		dir                   = dto.name
		interfaceName         = fmt.Sprintf("%sRepository", upperName)
		ctx                   = "context.Context"
		entityName            = fmt.Sprintf("*entity.%s", upperName)
		isExists              = new(isExists)
		generatedCustomReturn = parseCustomJenCodeFields(dto.returns, dto.arrReturn, isExists, true)
	)
	dto.name = sentences(dto.name)
	dto.methodName = capitalize(dto.methodName)

	importList := jen.Id("\n").
		Id(`"context"`).Id("\n").
		Id(`"github.com/Muruyung/go-utilities"`).Id("\n")

	if isExists.isTimeExists {
		importList = importList.Id(`"time"`)
	}

	_, ok1 := dto.methods["get"]
	_, ok2 := dto.methods["getList"]
	if ok1 || ok2 || isExists.isEntityExists {
		importList = importList.Id(fmt.Sprintf(`"%s/services/%s/domain/entity"`, projectName, dto.services)).Id("\n")
	}

	file.Add(jen.Id("import").Parens(
		importList,
	))

	var (
		generatedMethods = make([]jen.Code, 0)
	)

	if _, ok := dto.methods["get"]; ok {
		generatedMethods = append(
			generatedMethods,
			jen.Id("Get"+upperName).Params(jen.Id("ctx").Id(ctx), jen.Id("query").Id("utils.QueryBuilderInteractor")).
				Parens(jen.List(jen.Id(entityName), jen.Error())),
		)
	}

	if _, ok := dto.methods["getList"]; ok {
		generatedMethods = append(
			generatedMethods,
			jen.Id("GetList"+upperName).Params(jen.Id("ctx").Id(ctx), jen.Id("query").Id("utils.QueryBuilderInteractor")).
				Parens(jen.List(jen.Id("[]"+entityName), jen.Error())),
		)
	}

	if _, ok := dto.methods["create"]; ok {
		generatedMethods = append(
			generatedMethods,
			jen.Id("Save").Params(jen.Id("ctx").Id(ctx), jen.Id("data").Id(entityName)).
				Parens(jen.Error()),
		)
	}

	if _, ok := dto.methods["update"]; ok {
		generatedMethods = append(
			generatedMethods,
			jen.Id("Update").Params(jen.Id("ctx").Id(ctx), jen.Id("data").Id(entityName)).
				Parens(jen.Error()),
		)
	}

	if _, ok := dto.methods["delete"]; ok {
		generatedMethods = append(
			generatedMethods,
			jen.Id("Delete").Params(jen.Id("ctx").Id(ctx), jen.Id("data").Id(entityName)).
				Parens(jen.Error()),
		)
	}

	if _, ok := dto.methods["custom"]; ok {
		generatedMethods = append(
			generatedMethods,
			jen.Id(dto.methodName).Params(jen.Id("ctx").Id(ctx), jen.Id("query").Id("utils.QueryBuilderInteractor")).
				Parens(jen.List(generatedCustomReturn...)),
		)
	}

	file.Commentf("%s %s repository wrapper", interfaceName, dto.name)
	file.Type().Id(interfaceName).Interface(
		generatedMethods...,
	)

	return file.Save(dto.path + "/" + dir + ".go")
}

func appendDomainRepo(path string, dto dtoModule) error {
	f, err := os.OpenFile(path, os.O_RDWR, 0666)
	if err != nil {
		logger.Logger.Error(fmt.Sprintf(defaultErr, err))
		return err
	}
	defer f.Close()

	var (
		scanner         = bufio.NewScanner(f)
		isFound         = false
		strBefore       string
		insertText      string
		upperName       = capitalize(dto.name)
		entityName      = fmt.Sprintf("*entity.%s", upperName)
		defaultParamGet = fmt.Sprintf("(ctx %s, query utils.QueryBuilderInteractor)", ctx)
		defaultParam    = fmt.Sprintf("(ctx %s, data %s)", ctx, entityName)
		defaultError    = "error"
	)

	if _, ok := dto.methods["get"]; ok {
		insertText += "\nGet" + upperName + defaultParamGet + fmt.Sprintf("(%s, %s)", entityName, defaultError)
	}

	if _, ok := dto.methods["getList"]; ok {
		insertText += "\nGetList" + upperName + defaultParamGet + fmt.Sprintf("([]%s, %s)", entityName, defaultError)
	}

	if _, ok := dto.methods["create"]; ok {
		insertText += "\nCreate" + upperName + defaultParam + defaultError
	}

	if _, ok := dto.methods["update"]; ok {
		insertText += "\nUpdate" + upperName + defaultParam + defaultError
	}

	if _, ok := dto.methods["delete"]; ok {
		insertText += "\nDelete" + upperName + defaultParam + defaultError
	}

	if _, ok := dto.methods["custom"]; ok {
		var ret string
		if len(dto.arrReturn) > 1 {
			ret = "("
			for index, field := range dto.arrReturn {
				if index > 0 {
					ret += ", "
				}
				ret += dto.returns[field]
			}
			ret += ")"
		} else {
			ret = dto.returns[dto.arrReturn[0]]
		}

		insertText += "\n" + capitalize(dto.methodName) + defaultParamGet + ret
	}
	insertText += "\n"

	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "Repository interface {") {
			isFound = true
			strBefore += scanner.Text()
			continue
		}
		if isFound {
			insertText += scanner.Text() + "\n"
		} else {
			strBefore += scanner.Text() + " "
		}
	}

	strBeforeBytes := len([]rune(strBefore))

	if err = scanner.Err(); err != nil {
		logger.Logger.Error(fmt.Sprintf(defaultErr, err))
		return err
	}

	_, err = f.WriteAt([]byte(insertText), int64(strBeforeBytes))
	if err != nil {
		logger.Logger.Error(fmt.Sprintf(defaultErr, err))
		return err
	}

	logger.Logger.Info("domain repository created")
	return nil
}