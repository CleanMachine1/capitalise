package capitalise

import (
	"unicode"
)

// Capitalises specified character in string
func Specified(str string, num int) string {
	if len(str) == 0 {
		return ""
	}
	if num < 0 || num > len(str) {
		return str
	}
	tmp := []rune(str)
	tmp[num] = unicode.ToUpper(tmp[num])
	return string(tmp)
}

// Capitalises the first character of a string
func First(str string) string {
	return Specified(str, 0)
}
