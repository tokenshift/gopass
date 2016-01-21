package lib

import (
	"crypto/rand"
	"fmt"
)

var DIGITS    = []rune("1234567890")
var UPPERCASE = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
var LOWERCASE = []rune("abcdefghijklmnopqrstuvwxyz")
var SPECIAL   = []rune(`!@#$%^&*()-_+=?><[]{}\/.`)
var AMBIGUOUS = []rune("B8G6I1l0OQDS5Z2")

type PasswordOptions uint
const (
	INCLUDE_DIGITS PasswordOptions = 1 << iota
	INCLUDE_UPPERCASE
	INCLUDE_LOWERCASE
	INCLUDE_SPECIAL
	AVOID_AMBIGUOUS
	REQUIRE_ALL
)

const DEFAULT_OPTIONS = INCLUDE_DIGITS | INCLUDE_UPPERCASE | INCLUDE_LOWERCASE | INCLUDE_SPECIAL | AVOID_AMBIGUOUS | REQUIRE_ALL

func RandomPassword(length uint) string {
	return RandomPasswordOpts(length, DEFAULT_OPTIONS)
}

func RandomPasswordOpts(length uint, options PasswordOptions) string {
	password := randomPassword(length, options)

	for !meetsRequirements(password, options) {
		password = randomPassword(length, options)
	}

	return password
}

func anyOverlap(a, b []rune) bool {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			if a[i] == b[j] {
				return true
			}
		}
	}

	return false
}

func meetsRequirements(password string, options PasswordOptions) bool {
	if AVOID_AMBIGUOUS&options != 0 && anyOverlap([]rune(password), AMBIGUOUS) {
		return false
	}

	if REQUIRE_ALL&options != 0 {
		if INCLUDE_DIGITS&options != 0 && !anyOverlap([]rune(password), DIGITS) {
			return false
		}

		if INCLUDE_UPPERCASE&options != 0 && !anyOverlap([]rune(password), UPPERCASE) {
			return false
		}

		if INCLUDE_LOWERCASE&options != 0 && !anyOverlap([]rune(password), LOWERCASE) {
			return false
		}

		if INCLUDE_SPECIAL&options != 0 && !anyOverlap([]rune(password), SPECIAL) {
			return false
		}
	}

	return true
}

func randomPassword(length uint, options PasswordOptions) string {
	var valid []rune

	if (INCLUDE_DIGITS&options != 0) {
		valid = append(valid, DIGITS...)
	}

	if (INCLUDE_UPPERCASE&options != 0) {
		valid = append(valid, UPPERCASE...)
	}

	if (INCLUDE_LOWERCASE&options != 0) {
		valid = append(valid, LOWERCASE...)
	}

	if (INCLUDE_SPECIAL&options != 0) {
		valid = append(valid, SPECIAL...)
	}

	if len(valid) == 0 {
		panic(fmt.Errorf("Options must include at least one character class."))
	}

	var result []rune

	buffer := make([]byte, 1)
	for i := uint(0); i < length; i++ {
		rand.Read(buffer)
		i := buffer[0] % byte(len(valid))
		result = append(result, valid[i])
	}

	return string(result)}