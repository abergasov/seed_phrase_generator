package txtparser

import (
	"io/ioutil"
	"os"
	"seed_phrase_generator/internal/logger"
	"seed_phrase_generator/internal/utils"
	"seed_phrase_generator/internal/utils/txtparser/epub"
	"strings"

	"go.uber.org/zap"
)

type TextParser struct {
	preparedBooks []epub.BookEpub
	validTexts    []string
}

func InitParser() *TextParser {
	return &TextParser{
		validTexts:    make([]string, 0, 10),
		preparedBooks: make([]epub.BookEpub, 0, 10),
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
	for i := range b.preparedBooks {
		result = append(result, ParserResponse{
			FileID: i,
			File:   b.validTexts[i],
			Title:  b.preparedBooks[i].Opf.Metadata.Title[0],
			Author: b.preparedBooks[i].Opf.Metadata.Creator[0].Data,
		})
	}
	return result
}

func (b *TextParser) canBookBeUsed(dir, file string) (*epub.BookEpub, error) {
	return epub.NewEpubBook(dir + string(os.PathSeparator) + file)
}

func (b *TextParser) GetTextChapters(resp *ParserResponse) []Chapter {
	result := make([]Chapter, 0)
	for i := range b.preparedBooks[resp.FileID].Ncx.NavMap.NavPoint {
		// book
		for j := range b.preparedBooks[resp.FileID].Ncx.NavMap.NavPoint[i].NavPoint {
			// book chapter
			tmp := Chapter{
				ChapterTitle: b.preparedBooks[resp.FileID].Ncx.NavMap.NavPoint[i].NavPoint[j].NavLabel.Text,
			}
			paragraphList := make([]string, 0, 5)
			for _, k := range b.preparedBooks[resp.FileID].Ncx.NavMap.NavPoint[i].NavPoint[j].ChapterContent.P {
				if len(paragraphList) >= 3 {
					break
				}
				txt := utils.CleanString(k.Text)
				if len(txt) > 0 {
					paragraphList = append(paragraphList, txt)
				}
			}
			tmp.FullText = utils.CutText(strings.Join(paragraphList, "\n"), 200)
			tmp.Offset = i + j + 1
			result = append(result, tmp)
		}
	}
	return result
}

func (b *TextParser) GetOffsetData(resp *ParserResponse, offset int) []string {
	res := make([]string, 0, 4000)
	for i := range b.preparedBooks[resp.FileID].Ncx.NavMap.NavPoint {
		// book
		for j := range b.preparedBooks[resp.FileID].Ncx.NavMap.NavPoint[i].NavPoint {
			// book chapter
			if j < offset {
				continue
			}
			for _, k := range b.preparedBooks[resp.FileID].Ncx.NavMap.NavPoint[i].NavPoint[j].ChapterContent.P {
				txt := utils.CleanString(k.Text)
				if len(txt) > 0 {
					res = append(res, txt)
				}
			}
		}
	}
	return res
}
