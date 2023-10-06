package modulegenerator

// flag field
var (
	svcName         string
	name            string
	varField        string
	varMethod       string
	methodName      string
	varParam        string
	varReturn       string
	isModelsOnly    bool
	isEntityOnly    bool
	isRepoOnly      bool
	isServiceOnly   bool
	isUseCaseOnly   bool
	isWithoutUT     bool
	isWithoutEntity bool
	isCqrs          bool
	isCqrsQuery     bool
	isCqrsCommand   bool
)

var (
	projectName        string
	defaultErr         = "Error: %v"
	loggerInfo         = "logger.DetailLoggerInfo"
	loggerErr          = "logger.DetailLoggerError"
	logErr             = "\nerr,\n"
	loggerCtx          = "\nctx,"
	loggerCmdName      = "\ncommandName,"
	loggerErrExecQuery = "\n\"Error execute query\","
	ctx                = "context.Context"
	utils              = "goutils.QueryBuilderInteractor"
	getTableName       = "GetTableName()"
	dbqOpts            = "&dbq.Options"
)

type dtoModule struct {
	path       string
	sep        string
	name       string
	services   string
	fields     map[string]string
	arrFields  []string
	methods    map[string]bool
	methodName string
	params     map[string]string
	arrParams  []string
	returns    map[string]string
	arrReturn  []string
	entityOnly bool
}

type isExists struct {
	isTimeExists   bool
	isUtilsExists  bool
	isEntityExists bool
	isError        bool
}
