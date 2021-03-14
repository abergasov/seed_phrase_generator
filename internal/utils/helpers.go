package utils

import (
	"io"

	"golang.org/x/text/encoding/charmap"
)

func ReverseSlice(src []string) []string {
	result := make([]string, 0, len(src))
	for i := len(src) - 1; i >= 0; i-- {
		result = append(result, src[i])
	}
	return result
}

func PositionInArray(target string, strArr []string) int {
	for i := range strArr {
		if strArr[i] == target {
			return i
		}
	}
	return -1
}

// decode windows-1251
func DecodeWin1251(i io.Reader) (r io.Reader) {
	decoder := charmap.Windows1251.NewDecoder()
	r = decoder.Reader(i)
	return
}

func CutText(text string, size int) string {
	if len(text) < size {
		return text
	}
	return text[0:size] + "..."
}
