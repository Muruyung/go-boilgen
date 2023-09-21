package modulegenerator

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/Muruyung/go-utilities/logger"
)

// rootCmd represents the base command when called without any subcommands
var (
	rootCmd = &cobra.Command{
		Use:   "boilgen",
		Short: "Generate core modules with fields",
		Long:  `This subcommand used to creating core modules (usecase, service, repository, entity)`,
		Run:   modGen,
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	initConfig()
	rootCmd.PersistentFlags().StringP("service", "s", "", "targeted service name")
	rootCmd.PersistentFlags().StringP("name", "n", "", "module name")
	rootCmd.PersistentFlags().StringP("fields", "f", "", `"field_name1:data_type,field_name2:data_type"`)
	rootCmd.PersistentFlags().StringP("methods", "m", "custom", "the methods that you will create (get, get_list, create, update, or delete)")
	rootCmd.PersistentFlags().StringP("custom-method", "c", "", "custom method name (required for 'custom' methods flag)")
	rootCmd.PersistentFlags().StringP("params", "p", "", `custom method parameters (required for 'custom' methods flag), example:"field_name1:data_type,field_name2:data_type"`)
	rootCmd.PersistentFlags().StringP("return", "r", "err:error", `custom method return (required for 'custom' methods flag), example:"field_name1:data_type,field_name2:data_type"`)
	rootCmd.PersistentFlags().Bool("entity-only", false, "generate entity only")
	rootCmd.PersistentFlags().Bool("repo-only", false, "generate repository only")
	rootCmd.PersistentFlags().Bool("service-only", false, "generate service only")
	rootCmd.PersistentFlags().Bool("usecase-only", false, "generate usecase only")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {
	logger.InitLogger("local", "client")
}
