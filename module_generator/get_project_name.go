package modulegenerator

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Muruyung/go-utilities/logger"
)

func getProjectName(path, sep string) error {
	f, err := os.OpenFile(path+sep+"go.mod", os.O_RDWR, 0666)
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("project not found, error:%v", err))
		return err
	}
	defer f.Close()

	var (
		scanner   = bufio.NewScanner(f)
		strBefore string
	)

	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "module") {
			strBefore = scanner.Text()
			break
		}
	}

	listStr := strings.Split(strBefore, " ")
	projectName = listStr[1]

	return nil
}
