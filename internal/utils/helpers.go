package utils

import (
	"io"
	"regexp"
	"sort"
	"strings"

	"golang.org/x/text/encoding/charmap"
)

var r = regexp.MustCompile(`(</?[a-zA]+?[^>]*/?>|\[(\d*?)\])*`)

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

func CleanString(in string) string {
	groups := r.FindAllString(in, -1)
	sort.Slice(groups, func(i, j int) bool {
		return len(groups[i]) > len(groups[j])
	})
	for _, group := range groups {
		if strings.TrimSpace(group) != "" {
			in = strings.ReplaceAll(in, group, "")
		}
	}
	return strings.TrimSpace(in)
}
