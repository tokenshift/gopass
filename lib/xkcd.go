package lib

import (
	"strings"
)

func XKCDPassword(length int) string {
	var result []string
	count := 0

	for count < length {
		word := randWord()
		result = append(result, word)
		count += len(word)
	}

	return strings.Join(result, " ")
}

func randWord() string {
	return DICTIONARY[randInt(len(DICTIONARY))]
}