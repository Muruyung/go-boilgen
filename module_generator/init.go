package modulegenerator

import (
	"bytes"
	"os/exec"

	"github.com/Muruyung/go-utilities/logger"
	"github.com/spf13/cobra"
)

var initBoilgen = &cobra.Command{
	Use:   "init",
	Short: "initialize project",
	Long:  `This subcommand used for initialize required package for this project`,
	Run: func(cmd *cobra.Command, args []string) {
		var (
			out     bytes.Buffer
			err     error
			command *exec.Cmd
		)

		command = exec.Command("go", "get", "github.com/Muruyung/go-utilities@latest")
		command.Stdout = &out
		err = command.Run()
		if err != nil {
			logger.Logger.Errorf(defaultErr, err)
		}

		command = exec.Command("go", "get", "github.com/golang/mock/gomock")
		command.Stdout = &out
		err = command.Run()
		if err != nil {
			logger.Logger.Errorf(defaultErr, err)
		}
	},
}

func init() {
	rootCmd.AddCommand(initBoilgen)
}
