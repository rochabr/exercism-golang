package isogram

import (
	"strings"
	"unicode"
)

//IsIsogram needs
func IsIsogram(word string) bool {
	word = strings.ToLower(word)
	for _, c := range word {
		if !unicode.IsLetter(c) {
			continue
		}
		if strings.Count(word, string(c)) > 1 {
			return false
		}
	}
	return true
}
