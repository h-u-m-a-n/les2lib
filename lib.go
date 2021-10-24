package les2lib

import (
"unicode"
)

func ChangeLetterCase(char rune) rune {
	if unicode.IsLower(char) {
		return unicode.ToUpper(char)
	} else {
		return unicode.ToLower(char)
	}
}

