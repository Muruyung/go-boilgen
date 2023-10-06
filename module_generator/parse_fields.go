package modulegenerator

import (
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/iancoleman/strcase"
)

func parseFields(args string, isEntity bool) (res map[string]string, arrRes []string) {
	res = make(map[string]string)
	arrRes = make([]string, 0)

	if isEntity {
		arrRes = append(arrRes, "id")
		res["id"] = "string"
	}

	if args == "" {
		return
	}

	fields := strings.Split(args, ",")
	for _, field := range fields {
		var (
			content   = strings.Split(field, ":")
			name      string
			fieldType = "string"
		)

		if len(content) > 0 {
			name = strcase.ToLowerCamel(content[0])

			if name == "ctx" {
				continue
			}

			if len(content) > 1 {
				fieldType = content[1]
			}

			res[name] = fieldType

			if name != "id" || !isEntity {
				arrRes = append(arrRes, name)
			}
		}
	}

	delete(res, "ctx")

	return
}

func parseJenCodeFields(fields map[string]string) ([]jen.Code, *isExists) {
	var (
		generatedDtoFields = make([]jen.Code, 0)
		exists             = new(isExists)
	)

	for field, fieldType := range fields {
		upperCaseField := capitalize(field)

		if upperCaseField != "ID" {
			generatedDtoFields = append(
				generatedDtoFields,
				jen.Id(upperCaseField).Id(fieldType),
			)
		}

		if strings.Contains(strcase.ToSnake(fields[field]), "time") {
			fields[field] = "time.Time"
			exists.isTimeExists = true
		}

		if strings.Contains(strcase.ToSnake(fields[field]), "goutils") {
			exists.isUtilsExists = true
		}

		if strings.Contains(strcase.ToSnake(fields[field]), "entity") {
			exists.isEntityExists = true
		}
	}

	return generatedDtoFields, exists
}

func parseCustomJenCodeFields(fields map[string]string, arrFields []string, exists *isExists, isReturn bool) []jen.Code {
	var (
		generatedDtoFields = make([]jen.Code, 0)
	)

	for _, field := range arrFields {
		upperCaseField := lowerize(field)

		if isReturn {
			generatedDtoFields = append(
				generatedDtoFields,
				jen.Id(fields[field]),
			)
		} else {
			generatedDtoFields = append(
				generatedDtoFields,
				jen.Id(upperCaseField).Id(fields[field]),
			)
		}

		if strings.Contains(strcase.ToSnake(fields[field]), "time") {
			fields[field] = "time.Time"
			exists.isTimeExists = true
		}

		if strings.Contains(strcase.ToSnake(fields[field]), "goutils") {
			exists.isUtilsExists = true
		}

		if strings.Contains(strcase.ToSnake(fields[field]), "entity") {
			exists.isEntityExists = true
		}

		if strings.Contains(strcase.ToSnake(fields[field]), "err") {
			exists.isError = true
		}
	}

	return generatedDtoFields
}
