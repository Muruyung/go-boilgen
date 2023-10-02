package modulegenerator

import "strings"

var (
	commandCondition = []string{
		"Create",
		"Store",
		"Update",
		"Delete",
		"Change",
		"Put",
		"Post",
		"Insert",
	}
)

func cqrsTypeCheck(name string) string {
	var cqrs = "query"
	name = capitalize(name)

	for _, condition := range commandCondition {
		if strings.Contains(name, condition) {
			cqrs = "command"
			break
		}
	}

	return cqrs
}
