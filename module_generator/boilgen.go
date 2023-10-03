package modulegenerator

import (
	"os"

	"github.com/Muruyung/go-utilities/logger"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var (
	rootCmd = &cobra.Command{
		Use:     "boilgen",
		Short:   "Generate core modules with fields",
		Long:    `This subcommand used to creating core modules (usecase, service, repository, entity)`,
		Run:     modGen,
		Version: "1.6.1",
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
	rootCmd.InitDefaultVersionFlag()
	rootCmd.PersistentFlags().StringP("service", "s", "", "Targeted service name")
	rootCmd.PersistentFlags().StringP("name", "n", "", "Module name")
	rootCmd.PersistentFlags().StringP("fields", "f", "", `"field_name1:data_type,field_name2:data_type"`)
	rootCmd.PersistentFlags().StringP("methods", "m", "custom", "The methods that you will create (get, get_list, create, update, or delete)")
	rootCmd.PersistentFlags().StringP("custom-method", "c", "", "Custom method name (required for 'custom' methods flag)")
	rootCmd.PersistentFlags().StringP("params", "p", "", `Custom method parameters (required for 'custom' methods flag), example:"field_name1:data_type,field_name2:data_type"`)
	rootCmd.PersistentFlags().StringP("return", "r", "err:error", `custom method return (required for 'custom' methods flag), example:"field_name1:data_type,field_name2:data_type"`)
	rootCmd.PersistentFlags().Bool("models-only", false, "Generate models only")
	rootCmd.PersistentFlags().Bool("entity-only", false, "Generate entity only")
	rootCmd.PersistentFlags().Bool("repo-only", false, "Generate repository only")
	rootCmd.PersistentFlags().Bool("service-only", false, "Generate service only")
	rootCmd.PersistentFlags().Bool("usecase-only", false, "Generate usecase only")
	rootCmd.PersistentFlags().Bool("no-unit-test", false, "Generate without unit test")
	rootCmd.PersistentFlags().Bool("no-entity", false, "Generate without entity")
	rootCmd.PersistentFlags().Bool("cqrs", false, "Generate using CQRS pattern")
	rootCmd.PersistentFlags().Bool("is-query", false, "Generate query for CQRS pattern")
	rootCmd.PersistentFlags().Bool("is-command", false, "Generate command for CQRS pattern")
}

func initConfig() {
	logger.InitLogger("local", "client")
}
