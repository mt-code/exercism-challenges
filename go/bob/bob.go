package bob

import (
	"strings"
)

// Constants
const alpha = "abcdefghijklmnopqrstuvwxz"

// Hey should have a comment documenting it.
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)

	if remark == "" {
		return "Fine. Be that way!"
	}

	isAlpha := strings.ContainsAny(strings.ToLower(remark), alpha)
	isQuestion := remark[len(remark) - 1 :] == "?"
	isUpper := isAlpha && (strings.ToUpper(remark) == remark)

	if isUpper && isQuestion {
		return "Calm down, I know what I'm doing!"
	}
	if isQuestion {
		return "Sure."
	}
	if isUpper {
		return "Whoa, chill out!"
	}

	return "Whatever."
}