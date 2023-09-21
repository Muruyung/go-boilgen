package modulegenerator

import (
	"strings"

	"github.com/stoewer/go-strcase"
)

func parseMethods(args string) (res map[string]bool) {
	var availableMethods = map[string]bool{
		"get":     true,
		"getList": true,
		"create":  true,
		"update":  true,
		"delete":  true,
		"custom":  true,
	}

	res = make(map[string]bool)
	methods := strings.Split(args, ",")
	for _, method := range methods {
		method = strcase.LowerCamelCase(method)
		if _, ok := availableMethods[method]; ok {
			res[method] = true
		}
	}

	return
}
