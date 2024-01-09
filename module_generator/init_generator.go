package modulegenerator

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/Muruyung/go-utilities/logger"
	"github.com/dave/jennifer/jen"
)

func generateInit(name, services, path string) error {
	var (
		pkgName    string
		returnData jen.Code
		params     = make([]jen.Code, 0)
		wrapper    string
		importList = jen.Id("\n").Id(fmt.Sprintf(`"%s/services/%s/domain/repository"`, projectName, services)).Id("\n")
	)

	if strings.Contains(path, "/repository/") {
		pkgName = "repository_mysql"
		params = append(params, jen.Id("db").Id("repository.SqlTx"))

		if name == "SqlTx" {
			returnData = jen.Id("SqlTx:").Id("db,")
		} else {
			importList = importList.Id(fmt.Sprintf("%s_repo", name)).Lit(fmt.Sprintf("%s/services/%s/internal/repository/mysql/%s", projectName, services, name))
			returnData = jen.Id(fmt.Sprintf("%sRepo:", capitalize(name))).Id(fmt.Sprintf("%s_repo", name)).Dot(fmt.Sprintf("New%s(db),", capitalize(name)))
		}

		wrapper = "repository.Wrapper"
	}

	if strings.Contains(path, "/service/") {
		pkgName = "service"
		importList = importList.Id(fmt.Sprintf(`"%s/services/%s/domain/service"`, projectName, services)).Id("\n")

		params = append(params,
			jen.Id("repo").Id("*repository.Wrapper"),
			jen.Id("svcTx").Id("service.SvcTx"),
		)

		if name == "SvcTx" {
			returnData = jen.Id("SvcTx:").Id("svcTx,")
		} else {
			importList = importList.Id(fmt.Sprintf("%s_repo", name)).Lit(fmt.Sprintf("%s/services/%s/internal/service/%s", projectName, services, name))
			returnData = jen.Id(fmt.Sprintf("%sSvc:", capitalize(name))).Id(fmt.Sprintf("%s_svc", name)).Dot(fmt.Sprintf("New%s(repo),", capitalize(name)))
		}

		wrapper = "service.Wrapper"
	}

	var (
		file = jen.NewFilePathName(path, pkgName)
	)

	file.Add(jen.Id("import").Params(importList))
	file.Func().Id("Init").Params(params...).Id("*" + wrapper).Block(
		jen.Return(jen.Id("&" + wrapper).Block(returnData)),
	)

	return file.Save(path)
}

func appendInit(name, services, path string) error {
	var (
		short      string
		full       string
		param      string
		importPath string
	)

	if strings.Contains(path, "/repository/") {
		short = "repo"
		full = "repository"
		param = "db"
		importPath = fmt.Sprintf("\n	%s_%s \"%s/services/%s/internal/repository/mysql/%s\"\n", name, short, projectName, services, name)
	}

	if strings.Contains(path, "/service/") {
		short = "svc"
		full = "service"
		param = "repo"
		importPath = fmt.Sprintf("\n	%s_%s \"%s/services/%s/internal/service/%s\"\n", name, short, projectName, services, name)
	}

	if _, err := os.Stat(path + "init.go"); os.IsNotExist(err) {
		return generateInit(name, services, path+"init.go")
	}

	err := appendImportInit(name, short, services, importPath, path)
	if err != nil {
		logger.Logger.Error(fmt.Sprintf(defaultErr, err))
		return err
	}

	f, err := os.OpenFile(path+"init.go", os.O_RDWR, 0666)
	if err != nil {
		logger.Logger.Error(fmt.Sprintf(defaultErr, err))
		return err
	}
	defer f.Close()

	var (
		scanner   = bufio.NewScanner(f)
		isFound   = false
		strBefore string
	)

	insertText := fmt.Sprintf("		%s%s: %s_%s.New%s%s(%s),\n", capitalize(name), capitalize(short), name, short, capitalize(name), capitalize(full), param)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), fmt.Sprintf("return &%s.Wrapper{", full)) {
			isFound = true
			strBefore += scanner.Text()
			continue
		}
		if isFound {
			insertText += scanner.Text() + "\n"
		} else {
			strBefore += scanner.Text() + "\n"
		}
	}

	strBeforeBytes := bytes.Count([]byte(strBefore), []byte{})

	if err = scanner.Err(); err != nil {
		logger.Logger.Error(fmt.Sprintf(defaultErr, err))
		return err
	}

	_, err = f.WriteAt([]byte(insertText), int64(strBeforeBytes+2))
	if err != nil {
		logger.Logger.Error(fmt.Sprintf(defaultErr, err))
		return err
	}

	return nil
}

func appendImportInit(name, short, services, insertText, path string) error {
	f, err := os.OpenFile(path+"init.go", os.O_RDWR, 0666)
	if err != nil {
		logger.Logger.Error(fmt.Sprintf(defaultErr, err))
		return err
	}
	defer f.Close()

	var (
		scanner   = bufio.NewScanner(f)
		isFound   = false
		strBefore string
	)

	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "import (") {
			isFound = true
			strBefore += scanner.Text()
			continue
		}
		if isFound {
			insertText += scanner.Text() + "\n"
		} else {
			strBefore += scanner.Text()
		}
	}

	strBeforeBytes := bytes.Count([]byte(strBefore), []byte{})

	if err = scanner.Err(); err != nil {
		logger.Logger.Error(fmt.Sprintf(defaultErr, err))
		return err
	}

	_, err = f.WriteAt([]byte(insertText), int64(strBeforeBytes+2))
	if err != nil {
		logger.Logger.Error(fmt.Sprintf(defaultErr, err))
		return err
	}

	return nil
}
