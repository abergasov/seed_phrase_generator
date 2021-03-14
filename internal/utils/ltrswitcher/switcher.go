package ltrswitcher

import (
	"encoding/base64"
	"seed_phrase_generator/internal/utils"
	"strings"
)

type LetterSwitcher struct {
	ruAlphabet string
	enAlphabet string

	alphabetsList      [][]string
	enAlphabetPrepared []string
}

func NewSwitcher() *LetterSwitcher {
	s := &LetterSwitcher{
		ruAlphabet: "абвгдеёжзийклмнопрстуфхцчшщъыьэюя",
		enAlphabet: "abcdefghijklmnopqrstuvwxyz",
	}
	s.enAlphabetPrepared = strings.Split(s.enAlphabet, "")
	s.alphabetsList = [][]string{
		strings.Split(s.ruAlphabet, ""),
	}
	return s
}

func (l *LetterSwitcher) ReplaceLetters(src []string) []string {
	res := make([]string, 0, len(src))
	for i := range src {
		var tmp strings.Builder
		for _, ch := range src[i] {
			tmp.WriteString(l.switchLetter(string(ch)))
		}
		res = append(res, tmp.String())
	}
	return res
}

func (l *LetterSwitcher) switchLetter(letter string) string {
	lowerLetter := strings.ToLower(letter)
	if utils.PositionInArray(lowerLetter, l.enAlphabetPrepared) != -1 {
		return letter
	}
	var pos int
	for i := range l.alphabetsList {
		pos = utils.PositionInArray(lowerLetter, l.alphabetsList[i])
		if pos == -1 {
			continue
		}
		if pos > len(l.enAlphabetPrepared) {
			return ""
		}
		key := l.enAlphabetPrepared[pos]
		if lowerLetter != letter {
			return strings.ToUpper(key)
		}
		return key
	}
	return ""
}

func (l *LetterSwitcher) EncodeString(src []string) string {
	message := strings.Join(src, " ")
	base64Text := make([]byte, base64.StdEncoding.EncodedLen(len(message)))
	base64.StdEncoding.Encode(base64Text, []byte(message))
	text := string(base64Text)
	text = strings.ReplaceAll(text, "=", "")
	perWordKeys := len(text) / len(src)
	resultWords := make([]string, 0, len(src))
	for i := 0; i < len(text); i += perWordKeys {
		var txt string
		if i+perWordKeys > len(text) {
			txt = text[i:]
		} else {
			txt = text[i : i+perWordKeys]
		}
		resultWords = append(resultWords, txt)
	}
	return strings.Join(resultWords, " ")
}

func (l *LetterSwitcher) RotateWords(srcWordList []string) []string {
	limit := 3
	// split words to 3 groups
	data := make([][]string, limit)
	for i := range srcWordList {
		data[i%limit] = append(data[i%limit], srcWordList[i])
	}
	result := make([]string, 0, len(srcWordList))
	// from the end concat into one slice, with reversing single group
	for i := len(data) - 1; i >= 0; i-- {
		result = append(result, utils.ReverseSlice(data[i])...)
	}
	return result
}
