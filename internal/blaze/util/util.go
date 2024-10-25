package util

import (
	"strconv"
)

func StrToByteSlice(s string) []byte {
	buffer := make([]byte, len(s))
	nonzero := false
	for _, char := range s {
		if byte(char) == byte(0) && !nonzero {
			continue
		}

		buffer = append(buffer, byte(char))
		nonzero = true
	}

	return buffer
}

func Atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func StartsWith(s string, r string) bool {
	if len(r) > len(s) {
		return false
	}

	for i := range len(r) {
		if s[i] != r[i] {
			return false
		}
	}

	return true
}
