package modulegenerator

import (
	"strings"

	strcase "github.com/stoewer/go-strcase"
)

func capitalize(str string) string {
	acronym := map[string]string{
		"Id":   "ID",
		"Url":  "URL",
		"Http": "HTTP",
	}

	res := strcase.UpperCamelCase(str)
	for key, val := range acronym {
		res = strings.ReplaceAll(res, key, val)
	}

	return res
}

func lowerize(str string) string {
	acronym := map[string]string{
		"Id":   "ID",
		"Url":  "URL",
		"Http": "HTTP",
	}

	res := strcase.LowerCamelCase(str)
	for key, val := range acronym {
		res = strings.ReplaceAll(res, key, val)
	}

	return res
}

func sentences(str string) string {
	res := strcase.SnakeCase(str)
	return strings.ReplaceAll(res, "_", " ")
}
