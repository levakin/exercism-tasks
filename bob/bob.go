// Package bob provides hey function
package bob

import (
	"strings"
	"unicode"
)

// Hey returns Bob's answer.
// Bob is a lackadaisical teenager. In conversation, his responses are very limited.
//
//Bob answers 'Sure.' if you ask him a question, such as "How are you?".
//
//He answers 'Whoa, chill out!' if you YELL AT HIM (in all capitals).
//
//He answers 'Calm down, I know what I'm doing!' if you yell a question at him.
//
//He says 'Fine. Be that way!' if you address him without actually saying anything.
//
//He answers 'Whatever.' to anything else.
//
//Bob's conversational partner is a purist when it comes to written communication and always follows normal rules regarding sentence punctuation in English.
func Hey(remark string) string {
	remark = strings.Trim(remark, " ")
	if !isSayingAnything(remark) {
		return "Fine. Be that way!"
	}
	if isQuestion(remark) {
		if isYell(remark) {
			return "Calm down, I know what I'm doing!"
		}
		return "Sure."
	}
	if isYell(remark) {
		return "Whoa, chill out!"
	}
	return "Whatever."
}

func isQuestion(remark string) bool {
	return strings.HasSuffix(remark, "?")
}

func isYell(remark string) bool {
	if !hasLetters(remark) {
		return false
	}

	for _, r := range remark {
		if unicode.IsLetter(r) && unicode.IsLower(r) {
			return false
		}
	}
	return true
}

func isSayingAnything(remark string) bool {
	for _, r := range remark {
		if !unicode.IsSpace(r) {
			return true
		}
	}
	return false
}

func hasLetters(remark string) bool {
	for _, r := range remark {
		if unicode.IsLetter(r) {
			return true
		}
	}
	return false
}
