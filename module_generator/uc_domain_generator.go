package modulegenerator

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Muruyung/go-utilities/logger"
	"github.com/dave/jennifer/jen"
)

func domainUcGenerator(dto dtoModule, isGenerate bool) error {
	if !isGenerate {
		return nil
	}

	dto.path += "usecase" + dto.sep
	var err error

	if _, err = os.Stat(dto.path); os.IsNotExist(err) {
		err = os.MkdirAll(dto.path, 0777)
		if err != nil {
			return err
		}
	}

	if _, err = os.Stat(dto.path + dto.name + ".go"); os.IsNotExist(err) {
		err = generateDomainUc(dto)
		if err != nil {
			logger.Logger.Error(fmt.Sprintf(defaultErr, err))
			return err
		} else {
			logger.Logger.Info("domain usecase created")
		}

		var (
			upperName      = capitalize(dto.name)
			interfaceShort = fmt.Sprintf("%sUC", upperName)
			interfaceName  = fmt.Sprintf("%sUseCase", upperName)
		)
		err = appendWrapper(interfaceShort, interfaceName, dto.path)
		if err != nil {
			logger.Logger.Error(fmt.Sprintf(defaultErr, err))
			return err
		}
		logger.Logger.Info("domain usecase wrapper created")
	} else {
		return appendDomainUc(dto.path+dto.name+".go", dto)
	}

	return nil
}

func generateDomainUc(dto dtoModule) error {
	var (
		dtoName                      string
		file                         = jen.NewFilePathName(dto.path, "usecase")
		upperName                    = capitalize(dto.name)
		dir                          = dto.name
		interfaceName                = fmt.Sprintf("%sUseCase", upperName)
		ctx                          = "context.Context"
		entityName                   = fmt.Sprintf("*entity.%s", upperName)
		generatedMethods             = make([]jen.Code, 0)
		generatedDtoFields, isExists = parseJenCodeFields(dto.fields)
		generatedCustomParams        = parseCustomJenCodeFields(dto.params, dto.arrParams, isExists, false)
		generatedCustomReturn        = parseCustomJenCodeFields(dto.returns, dto.arrReturn, isExists, true)
	)
	dto.name = sentences(dto.name)
	dto.methodName = capitalize(dto.methodName)

	importList := jen.Id("\n").Id(`"context"`).Id("\n")

	if isExists.isTimeExists {
		importList = importList.Id(`"time"`).Id("\n")
	}

	_, ok1 := dto.methods["get"]
	_, ok2 := dto.methods["getList"]
	if ok1 || ok2 || isExists.isEntityExists {
		importList = importList.Id(fmt.Sprintf(`"%s/services/%s/domain/entity"`, projectName, dto.services)).Id("\n")
	}

	if ok2 || isExists.isUtilsExists {
		importList = importList.Id(fmt.Sprintf(`"%s/pkg/utils"`, projectName)).Id("\n")
	}

	file.Add(jen.Id("import").Parens(
		importList,
	))

	if len(generatedDtoFields) > 0 {
		dtoName = fmt.Sprintf("DTO%s", upperName)
		file.Commentf("%s dto for %s usecase", dtoName, dto.name)
		file.Type().Id(dtoName).Struct(
			generatedDtoFields...,
		)
	}

	if _, ok := dto.methods["get"]; ok {
		generatedMethods = append(
			generatedMethods,
			jen.Id("Get"+upperName+"ByID").Params(jen.Id("ctx").Id(ctx), jen.Id("id").Id(dto.fields["id"])).
				Parens(jen.List(jen.Id(entityName), jen.Error())),
		)
	}

	if _, ok := dto.methods["getList"]; ok {
		generatedMethods = append(
			generatedMethods,
			jen.Id("GetList"+upperName).Params(jen.Id("ctx").Id(ctx), jen.Id("request").Id("*utils.RequestOption")).
				Parens(jen.List(jen.Id("[]"+entityName), jen.Id("*utils.MetaResponse"), jen.Error())),
		)
	}

	if _, ok := dto.methods["create"]; ok {
		generatedMethods = append(
			generatedMethods,
			jen.Id("Create"+upperName).Params(jen.Id("ctx").Id(ctx), jen.Id("dto").Id(dtoName)).
				Parens(jen.Error()),
		)
	}

	if _, ok := dto.methods["update"]; ok {
		generatedMethods = append(
			generatedMethods,
			jen.Id("Update"+upperName).Params(jen.Id("ctx").Id(ctx), jen.Id("id").Id(dto.fields["id"]), jen.Id("dto").Id(dtoName)).
				Parens(jen.Error()),
		)
	}

	if _, ok := dto.methods["delete"]; ok {
		generatedMethods = append(
			generatedMethods,
			jen.Id("Delete"+upperName).Params(jen.Id("ctx").Id(ctx), jen.Id("id").Id(dto.fields["id"])).
				Parens(jen.Error()),
		)
	}

	if _, ok := dto.methods["custom"]; ok {
		params := []jen.Code{jen.Id("ctx").Id(ctx)}
		params = append(params, generatedCustomParams...)

		generatedMethods = append(
			generatedMethods,
			jen.Id(dto.methodName).Params(params...).
				Parens(jen.List(generatedCustomReturn...)),
		)
	}

	file.Commentf("%s %s usecase template", interfaceName, dto.name)
	file.Type().Id(interfaceName).Interface(
		generatedMethods...,
	)

	return file.Save(dto.path + "/" + dir + ".go")
}

func appendDomainUc(path string, dto dtoModule) error {
	f, err := os.OpenFile(path, os.O_RDWR, 0666)
	if err != nil {
		logger.Logger.Error(fmt.Sprintf(defaultErr, err))
		return err
	}
	defer f.Close()

	var (
		scanner      = bufio.NewScanner(f)
		isFound      = false
		strBefore    string
		insertText   string
		upperName    = capitalize(dto.name)
		dtoName      = fmt.Sprintf("DTO%s", upperName)
		entityName   = fmt.Sprintf("*entity.%s", upperName)
		defaultError = "error"
	)

	if _, ok := dto.methods["get"]; ok {
		insertText += "\nGet" + upperName + fmt.Sprintf("ByID (ctx %s, id %s)(%s, %s)", ctx, dto.fields["id"], entityName, defaultError)
	}

	if _, ok := dto.methods["getList"]; ok {
		insertText += "\nGetList" + upperName + fmt.Sprintf("(ctx %s, request *utils.RequestOption)([]%s, *utils.MetaResponse, %s)", ctx, entityName, defaultError)
	}

	if _, ok := dto.methods["create"]; ok {
		insertText += "\nCreate" + upperName + fmt.Sprintf("(ctx %s, dto %s) %s", ctx, dtoName, defaultError)
	}

	if _, ok := dto.methods["update"]; ok {
		insertText += "\nUpdate" + upperName + fmt.Sprintf("(ctx %s, id %s, dto %s) %s", ctx, dto.fields["id"], dtoName, defaultError)
	}

	if _, ok := dto.methods["delete"]; ok {
		insertText += "\nDelete" + upperName + fmt.Sprintf("(ctx %s, id %s) %s", ctx, dto.fields["id"], defaultError)
	}

	if _, ok := dto.methods["custom"]; ok {
		var (
			par = fmt.Sprintf("(ctx %s", ctx)
			ret string
		)

		for _, field := range dto.arrParams {
			par += fmt.Sprintf(", %s %s", field, dto.params[field])
		}
		par += ")"

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

		insertText += "\n" + capitalize(dto.methodName) + par + ret
	}
	insertText += "\n"

	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "UseCase interface {") {
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

	logger.Logger.Info("domain usecase created")
	return nil
}
