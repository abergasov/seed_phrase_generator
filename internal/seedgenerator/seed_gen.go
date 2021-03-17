package seedgenerator

import (
	"errors"
	"seed_phrase_generator/internal/logger"
	"seed_phrase_generator/internal/utils/ltrswitcher"
	"seed_phrase_generator/internal/utils/txtparser"
	"strconv"

	"github.com/manifoldco/promptui"
)

type SeedGen struct {
	selectedSrcFile *txtparser.ParserResponse
	userCode        int64
	chapterStart    *txtparser.Chapter
	textDir         string
	switcher        ltrswitcher.ILetterSwitcher
	bookParser      txtparser.ITextParser
}

func NewSeedGen(textDir string, parser txtparser.ITextParser, switcher ltrswitcher.ILetterSwitcher) *SeedGen {
	seedGen := &SeedGen{
		textDir:    textDir,
		switcher:   switcher,
		bookParser: parser,
	}
	seedGen.bookParser.ParseTexts(textDir)
	return seedGen
}

func (s *SeedGen) SelectSrc() {
	books := s.bookParser.GetPossibleTexts()
	template := &promptui.SelectTemplates{
		Label:    "{{ . }}",
		Active:   "-> {{ .Title | cyan }}",
		Inactive: "  {{ .Title | cyan }}",
		Selected: "selected book: {{ .Title | green }}",
		Details: `
-------------------
{{ "Name:" | faint }}	{{ .Title }}
{{ "Author:" | faint }}	{{ .Author }}
{{ "File:" | faint }}	{{ .File }}`,
	}
	i, err := s.runSelect(books, template)
	if err == nil {
		s.selectedSrcFile = &books[i]
	}
}

func (s *SeedGen) SelectNumber() {
	validate := func(input string) error {
		_, err := strconv.ParseInt(input, 10, 64)
		return err
	}

	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }} ",
		Valid:   "{{ . | green }} ",
		Invalid: "{{ . | red }} ",
		Success: "your code: {{ . | green }}",
	}

	prompt := promptui.Prompt{
		Label:     "Enter code number",
		Templates: templates,
		Validate:  validate,
	}

	result, err := prompt.Run()

	if err != nil {
		logger.Error("error get user code", err)
		return
	}
	s.userCode, err = strconv.ParseInt(result, 10, 64)
	if err != nil {
		logger.Fatal("invalid code. Digits only allowed", err)
	}
}

func (s *SeedGen) ShowChapters() {
	if s.selectedSrcFile == nil {
		logger.Error("can't generate seed phrase", errors.New("src file is not selected"))
	}
	chapters := s.bookParser.GetTextChapters(s.selectedSrcFile)
	template := &promptui.SelectTemplates{
		Label:    "{{ . }}",
		Active:   "-> {{ .ChapterTitle | cyan }}",
		Inactive: "{{ .ChapterTitle | cyan }}",
		//Selected: "\U0001F336 {{ .Title | red | cyan }}",
		Details: `-------------------
{{ "Text:" | faint }}	{{ .FullText }}`,
	}

	i, err := s.runSelect(chapters, template)
	if err == nil {
		s.chapterStart = &chapters[i]
	}
}

func (s *SeedGen) runSelect(list interface{}, tmpl *promptui.SelectTemplates) (int, error) {
	prompt := promptui.Select{
		Label:     "Select chapter to start",
		Items:     list,
		Templates: tmpl,
		Size:      6,
	}

	i, _, err := prompt.Run()

	if err != nil {
		logger.Error("error show select", err)
		return i, err
	}
	return i, nil
}
