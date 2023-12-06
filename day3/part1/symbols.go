package main

import "unicode"

func IsSymbol(r rune) bool {
	return r != '.' && !unicode.IsDigit(r)
}
