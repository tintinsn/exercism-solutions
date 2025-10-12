package bob

import (
	"strings"
	"unicode"
)

func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	isQuestion := strings.HasSuffix(remark, "?")
	isYell := remark == strings.ToUpper(remark) && strings.IndexFunc(remark, unicode.IsLetter) >= 0

	switch {

	case isYell && isQuestion:
		return "Calm down, I know what I'm doing!"
	case isYell:
		return "Whoa, chill out!"
	case isQuestion:
		return "Sure."
	case remark == "":
		return "Fine. Be that way!"
	default:
		return "Whatever."
	}
}
