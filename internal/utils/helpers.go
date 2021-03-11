package utils

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
