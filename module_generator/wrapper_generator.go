package modulegenerator

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/Muruyung/go-utilities/logger"
)

func generateWrapper(interfaceShort, interfaceName, path string) error {
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
