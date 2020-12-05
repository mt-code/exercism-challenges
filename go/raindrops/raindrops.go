package raindrops

import (
	"strconv"
	"strings"
)

func Convert(input int) string {
	result := strings.Builder{}

	if input%3 == 0 {
		result.WriteString("Pling")
	}
	if input%5 == 0 {
		result.WriteString("Plang")
	}
	if input%7 == 0 {
		result.WriteString("Plong")
	}

	if result.String() == "" {
		return strconv.Itoa(input)
	}

	return result.String()
}
