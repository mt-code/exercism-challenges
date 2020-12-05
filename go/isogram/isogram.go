package isogram

import "strings"

func IsIsogram(input string) bool {
	var chars []string

	for _, char := range input {
		lowerCaseChar := strings.ToLower(string(char))

		if lowerCaseChar == "-" || lowerCaseChar == " " {
			continue
		}
		if ContainsChar(chars, lowerCaseChar) {
			return false
		}

		chars = append(chars, lowerCaseChar)
	}

	return true
}

func ContainsChar(slice []string, char string) bool {
	for _, c := range slice {
		if c == char {
			return true
		}
	}

	return false
}
