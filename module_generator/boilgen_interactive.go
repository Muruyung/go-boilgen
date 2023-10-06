package modulegenerator

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Muruyung/go-utilities/logger"
	"github.com/spf13/cobra"
	"github.com/stoewer/go-strcase"
)

var (
	interactive = &cobra.Command{
		Use:   "run",
		Short: "Generate core modules with fields",
		Long:  `This subcommand used to creating core modules (usecase, service, repository, entity) with interactive input`,
		Run:   execexecBoilgenInteractive,
	}
)

func init() {
	initConfig()
	rootCmd.AddCommand(interactive)
}

func execexecBoilgenInteractive(cmd *cobra.Command, args []string) {
	var (
		reader           = bufio.NewReader(os.Stdin)
		availableService = true
		dirName          []string
		err              error
		text             string
		char             rune
		index            int
	)
	isWithoutUT = true
	methodName = "example"
	varReturn = "err:error"

	if _, err = os.Stat("./services"); os.IsNotExist(err) {
		availableService = false
	}

	if availableService {
		dirName, err = directoryScan("./services/")
		if err != nil {
			logger.Logger.Errorf(defaultErr, err)
			return
		}

		for key, val := range dirName {
			fmt.Printf("%d. %s\n", key+1, val)
		}
		fmt.Printf("%d. Create new services\n", len(dirName)+1)
		fmt.Print("Choose services (input number): ")
		text, _ = reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		index, err = strconv.Atoi(text)
		if err != nil {
			logger.Logger.Errorf(defaultErr, err)
			return
		}

		if index < 1 || index > len(dirName)+1 {
			logger.Logger.Error("invalid input value, out of range")
			return
		}

		if index >= 1 && index <= len(dirName) {
			svcName = dirName[index-1]
		}
		fmt.Println()
	}

	if !availableService || index == len(dirName)+1 {
		fmt.Print("Input services name: ")
		text, _ = reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		if text == "" {
			logger.Logger.Error("invalid input value, cannot be empty")
			return
		}
		svcName = strings.ToLower(strcase.SnakeCase(text))
		fmt.Println()
	}

	fmt.Print("Input modules name: ")
	text, _ = reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	if text == "" {
		logger.Logger.Error("invalid input value, cannot be empty")
		return
	}
	name = strings.ToLower(strcase.SnakeCase(text))
	fmt.Println()

	fmt.Println("Input fields struct (optional, empty this line if you won't generate struct)")
	fmt.Println(`Example: id:string,age:int,startDate:time.Time`)
	fmt.Print("-> ")
	text, _ = reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	varField = text
	if varField != "" {
		fmt.Print("Generate entity?(y/n)")
		char, _, _ = reader.ReadRune()
		if char != 10 {
			_, _ = reader.ReadString('\n')
		}
		isWithoutEntity = !yesOrNo(char)
		fmt.Println(!isWithoutEntity)

		if !isWithoutEntity {
			fmt.Print("Generate entity only?(y/n)")
			char, _, _ = reader.ReadRune()
			if char != 10 {
				_, _ = reader.ReadString('\n')
			}
			isEntityOnly = yesOrNo(char)
			fmt.Println(isEntityOnly)
			if isEntityOnly {
				modGen()
				return
			}
		} else {
			fmt.Print("Generate models and mapper only?(y/n)")
			char, _, _ = reader.ReadRune()
			if char != 10 {
				_, _ = reader.ReadString('\n')
			}
			isModelsOnly = yesOrNo(char)
			fmt.Println(isModelsOnly)
			if isModelsOnly {
				modGen()
				return
			}
		}
	}
	fmt.Println()

	fmt.Println("1. Get")
	fmt.Println("2. Get List")
	fmt.Println("3. Create")
	fmt.Println("4. Update")
	fmt.Println("5. Delete")
	fmt.Println("Choose methods (optional, empty this line if you want to custom your method)")
	fmt.Println(`Example: 1,3,4`)
	fmt.Print("-> ")
	text, _ = reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	varMethod = "custom"
	if text != "" {
		varMethod, err = parseMethodsByIndex(text)
		if err != nil {
			return
		}
		if varMethod == "" {
			logger.Logger.Error("invalid value")
			return
		}
	}
	fmt.Println()

	if varMethod == "custom" {
		fmt.Print("Input custom method name: ")
		text, _ = reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		if text == "" {
			logger.Logger.Error("invalid input value, cannot be empty")
			return
		}
		methodName = strings.ToLower(strcase.SnakeCase(text))
		fmt.Println()

		fmt.Println("Input parameter (optional, empty this line will only generate context parameter)")
		fmt.Println(`Example: id:string,age:int,startDate:time.Time`)
		fmt.Print("-> ")
		text, _ = reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		varParam = text
		fmt.Println()

		fmt.Println("Input return (optional, empty this line will only generate error return)")
		fmt.Println(`Example: id:string,age:int,startDate:time.Time`)
		fmt.Print("-> ")
		text, _ = reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		if text != "" {
			varReturn = text
		}
		fmt.Println()
	}

	fmt.Print("Generate repository only?(y/n)")
	char, _, _ = reader.ReadRune()
	if char != 10 {
		_, _ = reader.ReadString('\n')
	}
	isRepoOnly = yesOrNo(char)
	fmt.Println(isRepoOnly)
	if isRepoOnly {
		modGen()
		return
	}

	fmt.Print("Generate service only?(y/n)")
	char, _, _ = reader.ReadRune()
	if char != 10 {
		_, _ = reader.ReadString('\n')
	}
	isServiceOnly = yesOrNo(char)
	fmt.Println(isServiceOnly)
	if isServiceOnly {
		modGen()
		return
	}

	fmt.Print("Generate usecase only?(y/n)")
	char, _, _ = reader.ReadRune()
	if char != 10 {
		_, _ = reader.ReadString('\n')
	}
	isUseCaseOnly = yesOrNo(char)
	fmt.Println(isUseCaseOnly)

	fmt.Print("Using CQRS pattern?(y/n)")
	char, _, _ = reader.ReadRune()
	if char != 10 {
		_, _ = reader.ReadString('\n')
	}
	isCqrs = yesOrNo(char)
	fmt.Println(isCqrs)

	fmt.Print("Generate mocks?(y/n)")
	char, _, _ = reader.ReadRune()
	if char != 10 {
		_, _ = reader.ReadString('\n')
	}
	isWithoutUT = !yesOrNo(char)
	fmt.Println(!isWithoutUT)

	modGen()
}
