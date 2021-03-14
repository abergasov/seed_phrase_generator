package txtparser

type ITextParser interface {
	ParseTexts(dir string) []string
	GetPossibleTexts() []ParserResponse
	GetTextChapters(*ParserResponse) []Chapter
}
