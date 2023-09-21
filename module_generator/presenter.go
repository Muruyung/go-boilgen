package modulegenerator

var (
	projectName   string
	defaultErr    = "Error: %v"
	loggerInfo    = "logger.DetailLoggerInfo"
	loggerErr     = "logger.DetailLoggerError"
	loggerCtx     = "\nctx,"
	loggerCmdName = "\ncommandName,"
	ctx           = "context.Context"
)

type dtoModule struct {
	path       string
	sep        string
	name       string
	services   string
	fields     map[string]string
	methods    map[string]bool
	methodName string
	params     map[string]string
	arrParams  []string
	returns    map[string]string
	arrReturn  []string
}

type isExists struct {
	isTimeExists   bool
	isUtilsExists  bool
	isEntityExists bool
	isError        bool
}
