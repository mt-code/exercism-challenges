package acronym

import (
	"regexp"
	"strings"
)

func Abbreviate(s string) (result string) {
	regex := regexp.MustCompile("(^|[^'])\\b[_]?([A-Za-z])")
	matches := regex.FindAllStringSubmatch(s, -1)

	for _, char := range matches {
		result += strings.ToUpper(char[2])
	}

	return result
}
