package modulegenerator

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Muruyung/go-utilities/logger"
	"github.com/dave/jennifer/jen"
)

func domainRepoGenerator(dto dtoModule, isAll, isOnly bool) error {
	if !isAll && !isOnly {
		return nil
	}

	dto.path += "repository" + dto.sep
	var err error

	if _, err = os.Stat(dto.path); os.IsNotExist(err) {
		err = os.MkdirAll(dto.path, 0777)
		if err != nil {
			return err
		}
	}

	if _, err = os.Stat(dto.path + dto.name + ".go"); os.IsNotExist(err) {
		err = generateDomainRepo(dto, isOnly)
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
		err = appendWrapper(interfaceShort, interfaceName, dto.path)
		if err != nil {
			logger.Logger.Error(fmt.Sprintf(defaultErr, err))
			return err
		}
		logger.Logger.Info("domain repository wrapper created")
	} else {
		return appendDomainRepo(dto.path+dto.name+".go", dto, isOnly)
	}

	return nil
}

func generateCommonDomainRepo(dto dtoModule) error {
	var (
		file             = jen.NewFilePathName(dto.path, "repository")
		dir              = "common"
		mapperCommonName = "MapperCommon"
		modelsCommonName = "ModelsCommon"
	)
	dto.name = sentences(dto.name)
	dto.methodName = capitalize(dto.methodName)

	file.Commentf("%s template for common mapper models", mapperCommonName)
	file.Type().Id(mapperCommonName).Interface(
		jen.Id("MapDomainToModels()").Id(modelsCommonName),
		jen.Id("MapModelsToDomain").Parens(jen.Id("entityStruct").Interface()),
	)

	file.Commentf("%s template for common models repository", modelsCommonName)
	file.Type().Id(modelsCommonName).Interface(
		jen.Id("GetTableName()").String(),
		jen.Id("GetModels()").Interface(),
		jen.Id("GetModelsMap()").Id("map[string]interface{}"),
		jen.Id("GetColumns()").Id("[]string"),
		jen.Id("GetValStruct").Parens(jen.Id("arrColumn").Id("[]string")).Id("[]interface{}"),
	)

	return file.Save(dto.path + "/" + dir + ".go")
}

func generateDomainRepo(dto dtoModule, isOnly bool) error {
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
		Id(`goutils"github.com/Muruyung/go-utilities"`).Id("\n").
		Id(fmt.Sprintf(`"%s/services/%s/domain/entity"`, projectName, dto.services)).Id("\n")

	if isExists.isTimeExists {
		importList = importList.Id(`"time"`).Id("\n")
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
			jen.Id("Get").Params(jen.Id("ctx").Id(ctx), jen.Id("query").Id(utils)).
				Parens(jen.List(jen.Id(entityName), jen.Error())),
		)
	}

	if _, ok := dto.methods["getList"]; ok {
		generatedMethods = append(
			generatedMethods,
			jen.Id("GetList").Params(jen.Id("ctx").Id(ctx), jen.Id("query").Id(utils)).
				Parens(jen.List(jen.Id("[]"+entityName), jen.Error())),
			jen.Id("GetCount").Params(jen.Id("ctx").Id(ctx), jen.Id("query").Id(utils)).
				Parens(jen.List(jen.Int(), jen.Error())),
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
			jen.Id("Update").Params(jen.Id("ctx").Id(ctx), jen.Id("id").Id(dto.fields["id"]), jen.Id("data").Id(entityName)).
				Parens(jen.Error()),
		)
	}

	if _, ok := dto.methods["delete"]; ok {
		generatedMethods = append(
			generatedMethods,
			jen.Id("Delete").Params(jen.Id("ctx").Id(ctx), jen.Id("id").Id(dto.fields["id"])).
				Parens(jen.Error()),
		)
	}

	if _, ok := dto.methods["custom"]; ok && isOnly {
		generatedMethods = append(
			generatedMethods,
			jen.Id(dto.methodName).Params(jen.Id("ctx").Id(ctx), jen.Id("query").Id(utils)).
				Parens(jen.List(generatedCustomReturn...)),
		)
	}

	file.Commentf("%s %s repository template", interfaceName, dto.name)
	file.Type().Id(interfaceName).Interface(
		generatedMethods...,
	)

	return file.Save(dto.path + "/" + dir + ".go")
}

func appendDomainRepo(path string, dto dtoModule, isOnly bool) error {
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
		defaultParamGet = fmt.Sprintf("(ctx %s, query goutils.QueryBuilderInteractor)", ctx)
		defaultError    = "error"
	)

	if _, ok := dto.methods["get"]; ok {
		insertText += "\nGet" + defaultParamGet + fmt.Sprintf("(%s, %s)", entityName, defaultError)
	}

	if _, ok := dto.methods["getList"]; ok {
		insertText += "\nGetList" + defaultParamGet + fmt.Sprintf("([]%s, %s)", entityName, defaultError)
		insertText += "\nGetCount" + defaultParamGet + fmt.Sprintf("(int, %s)", defaultError)
	}

	if _, ok := dto.methods["create"]; ok {
		insertText += "\nSave" + fmt.Sprintf("(ctx %s, data %s) %s", ctx, entityName, defaultError)
	}

	if _, ok := dto.methods["update"]; ok {
		insertText += "\nUpdate" + fmt.Sprintf("(ctx %s, id %s, data %s) %s", ctx, dto.fields["id"], entityName, defaultError)
	}

	if _, ok := dto.methods["delete"]; ok {
		insertText += "\nDelete" + fmt.Sprintf("(ctx %s, id %s) %s", ctx, dto.fields["id"], defaultError)
	}

	if _, ok := dto.methods["custom"]; ok && isOnly {
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
