package ltrswitcher

type ILetterSwitcher interface {
	RotateWords([]string) []string
	ReplaceLetters([]string) []string
	EncodeString([]string) string
}
