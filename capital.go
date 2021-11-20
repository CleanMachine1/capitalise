package capitalise

import (
	"unicode"
)

// Capitalises the first character of a string
func First(str string) string {
	if len(str) == 0 {
		return ""
	}
	tmp := []rune(str)
	tmp[0] = unicode.ToUpper(tmp[0])
	return string(tmp)
}
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
