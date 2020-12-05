package luhn

import (
	"strconv"
	"strings"
)

func Valid(input string) bool {
	sum := 0
	input = strings.Replace(input, " ", "", -1)

	if len(input) <= 1 {
		return false
	}

	for i := len(input); i >= 1; i-- {
		count := len(input) - i - 1
		intDigit, err := strconv.Atoi(string(input[i-1]))

		if err != nil {
			return false
		}

		if count%2 == 0 {
			intDigit = intDigit * 2

			if intDigit > 9 {
				intDigit -= 9
			}
		}

		sum += intDigit
	}

	return sum%10 == 0
}
