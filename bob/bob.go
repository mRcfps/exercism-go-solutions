// Package bob implements a lackadaisical teenager called Bob.
// He will respond to you according to your remark.
package bob

import "strings"

// Hey returns Bob's answer.
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	var answer string

	if remark == "" {
		answer = "Fine. Be that way!"
	} else if isQuestion(remark) && isYelling(remark) {
		answer = "Calm down, I know what I'm doing!"
	} else if isQuestion(remark) {
		answer = "Sure."
	} else if isYelling(remark) {
		answer = "Whoa, chill out!"
	} else {
		answer = "Whatever."
	}
	return answer
}

// isQuestion tests whether a remark is a question.
func isQuestion(remark string) bool {
	return strings.HasSuffix(remark, "?")
}

// isYelling tests whether a remark is yelling.
func isYelling(remark string) bool {
	return remark == strings.ToUpper(remark) && remark != strings.ToLower(remark)
}
