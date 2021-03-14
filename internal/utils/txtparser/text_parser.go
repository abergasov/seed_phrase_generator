package txtparser

import (
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"seed_phrase_generator/internal/logger"
	"seed_phrase_generator/internal/utils"

	"go.uber.org/zap"
)

type TextParser struct {
	preparedBooks []FB2
	validTexts    []string
}

func InitParser() *TextParser {
	return &TextParser{
		validTexts:    make([]string, 0, 10),
		preparedBooks: make([]FB2, 0, 10),
	}
}

func (b *TextParser) ParseTexts(dir string) []string {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		logger.Fatal("error parse files", err)
	}
	for i := range files {
		if book, errB := b.canBookBeUsed(dir, files[i].Name()); errB == nil {
			b.validTexts = append(b.validTexts, files[i].Name())
			b.preparedBooks = append(b.preparedBooks, *book)
			logger.Info("can use", zap.String("book", files[i].Name()))
		}
	}
	return b.validTexts
}

func (b *TextParser) GetPossibleTexts() []ParserResponse {
	result := make([]ParserResponse, 0, len(b.validTexts))
	for i, v := range b.preparedBooks {
		result = append(result, ParserResponse{
			FileID: i,
			File:   b.validTexts[i],
			Title:  v.Description.TitleInfo.BookTitle,
			Author: fmt.Sprintf("%s %s %s",
				v.Description.TitleInfo.Author.FirstName,
				v.Description.TitleInfo.Author.MiddleName,
				v.Description.TitleInfo.Author.LastName,
			),
		})
	}
	return result
}

func (b *TextParser) canBookBeUsed(dir, file string) (*FB2, error) {
	xmlFile, err := os.Open(dir + string(os.PathSeparator) + file)
	if err != nil {
		logger.Error("error open file", err, zap.String("file", file))
		return nil, err
	}
	defer xmlFile.Close()

	var book FB2
	d := xml.NewDecoder(xmlFile)
	d.CharsetReader = b.charsetReader
	err = d.Decode(&book)

	if err != nil {
		// logger.Error("can't parse file", err, zap.String("file", file))
		return nil, err
	}
	return &book, nil
}

func (b *TextParser) charsetReader(c string, i io.Reader) (r io.Reader, e error) {
	switch c {
	case "windows-1251":
		r = utils.DecodeWin1251(i)
	}
	return
}

func (b *TextParser) GetTextChapters(resp *ParserResponse) []Chapter {
	result := make([]Chapter, 0, len(b.preparedBooks[resp.FileID].Body))
	for i, c := range b.preparedBooks[resp.FileID].Body {
		text := "awdawdawdawdawdawdawdawdawdawdawdawdawdawdawdawdawdawdawdawd" + c.Text
		result = append(result, Chapter{
			Text:     utils.CutText(text, 20),
			FullText: text,
			Offset:   i,
		})
	}
	return result
}
