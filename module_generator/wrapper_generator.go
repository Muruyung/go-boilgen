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

func generateWrapper(interfaceShort, interfaceName, path string) error {
	var pkgName string

	if strings.Contains(interfaceName, "Repository") {
		pkgName = "repository"
	}

	if strings.Contains(interfaceName, "Service") {
		pkgName = "service"
	}

	if strings.Contains(interfaceName, "UseCase") {
		if strings.Contains(path, "/query/") {
			pkgName = "query"
		} else if strings.Contains(path, "/command/") {
			pkgName = "command"
		} else {
			pkgName = "usecase"
		}
	}

	var (
		file = jen.NewFilePathName(path, pkgName)
	)

	file.Type().Id("Wrapper").Struct(
		jen.Id(interfaceShort).Id(interfaceName),
	)

	return file.Save(path)
}

func appendWrapper(interfaceShort, interfaceName, path string) error {
	if _, err := os.Stat(path + "wrapper.go"); os.IsNotExist(err) {
		return generateWrapper(interfaceShort, interfaceName, path+"wrapper.go")
	}

	f, err := os.OpenFile(path+"wrapper.go", os.O_RDWR, 0666)
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

	insertText := fmt.Sprintf("\n	%s %s\n", interfaceShort, interfaceName)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "type Wrapper struct {") {
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
