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
		Long:    `This command used to creating core modules (usecase, service, repository, entity)`,
		Version: "1.8.2",
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
	// rootCmd.PersistentFlags().StringVarP(&svcName, "service", "s", "", "Targeted service name")
	// rootCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "Module name")
	// rootCmd.PersistentFlags().StringVarP(&varField, "fields", "f", "", `"field_name1:data_type,field_name2:data_type"`)
	// rootCmd.PersistentFlags().StringVarP(&varMethod, "methods", "m", "custom", "The methods that you will create (get, get_list, create, update, or delete)")
	// rootCmd.PersistentFlags().StringVarP(&methodName, "custom-method", "c", "", "Custom method name (required for 'custom' methods flag)")
	// rootCmd.PersistentFlags().StringVarP(&varParam, "params", "p", "", `Custom method parameters (required for 'custom' methods flag), example:"field_name1:data_type,field_name2:data_type"`)
	// rootCmd.PersistentFlags().StringVarP(&varReturn, "return", "r", "err:error", `custom method return (required for 'custom' methods flag), example:"field_name1:data_type,field_name2:data_type"`)
	// rootCmd.PersistentFlags().BoolVar(&isModelsOnly, "models-only", false, "Generate models only")
	// rootCmd.PersistentFlags().BoolVar(&isEntityOnly, "entity-only", false, "Generate entity only")
	// rootCmd.PersistentFlags().BoolVar(&isRepoOnly, "repo-only", false, "Generate repository only")
	// rootCmd.PersistentFlags().BoolVar(&isServiceOnly, "service-only", false, "Generate service only")
	// rootCmd.PersistentFlags().BoolVar(&isUseCaseOnly, "usecase-only", false, "Generate usecase only")
	// rootCmd.PersistentFlags().BoolVar(&isWithoutUT, "no-unit-test", false, "Generate without unit test")
	// rootCmd.PersistentFlags().BoolVar(&isWithoutEntity, "no-entity", false, "Generate without entity")
	// rootCmd.PersistentFlags().BoolVar(&isCqrs, "cqrs", false, "Generate using CQRS pattern")
	// rootCmd.PersistentFlags().BoolVar(&isCqrsQuery, "is-query", false, "Generate query for CQRS pattern")
	// rootCmd.PersistentFlags().BoolVar(&isCqrsCommand, "is-command", false, "Generate command for CQRS pattern")
}

func initConfig() {
	logger.InitLogger("local", "client")
}

// func execBoilgen(cmd *cobra.Command, args []string) {
// 	modGen()
// }
