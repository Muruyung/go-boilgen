package modulegenerator

import (
	"errors"
	"strconv"
	"strings"

	"github.com/stoewer/go-strcase"
)

func parseMethodsByIndex(args string) (res string, err error) {
	var (
		availableMethods = []string{
			"get",
			"getList",
			"create",
			"update",
			"delete",
		}
		index int
	)

	methods := strings.Split(args, ",")
	for key, method := range methods {
		index, err = strconv.Atoi(method)
		if err != nil {
			break
		}

		if index < 1 || index > 5 {
			err = errors.New("invalid input value, out of range")
			break
		}

		if key > 0 {
			res += ","
		}
		res += availableMethods[index-1]
	}

	return
}

func parseMethods(args string) (res map[string]bool) {
	var availableMethods = map[string]bool{
		"get":     true,
		"getList": true,
		"create":  true,
		"update":  true,
		"delete":  true,
		"custom":  true,
	}

	var methodsChanger = map[string]string{
		"getlist":   "getList",
		"list":      "getList",
		"save":      "create",
		"store":     "create",
		"put":       "update",
		"getDetail": "get",
		"del":       "delete",
	}

	res = make(map[string]bool)
	methods := strings.Split(args, ",")
	for _, method := range methods {
		method = strcase.LowerCamelCase(method)
		if _, ok := methodsChanger[method]; ok {
			res[methodsChanger[method]] = true
		} else if _, ok := availableMethods[method]; ok {
			res[method] = true
		}
	}

	return
}
