package txtparser

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"seed_phrase_generator/internal/logger"
	"seed_phrase_generator/internal/utils"
	"seed_phrase_generator/internal/utils/txtparser/epub"
	"strings"

	"go.uber.org/zap"
)

type TextParser struct {
	preparedBooks []epub.EpubBook
	validTexts    []string
}

func InitParser() *TextParser {
	return &TextParser{
		validTexts:    make([]string, 0, 10),
		preparedBooks: make([]epub.EpubBook, 0, 10),
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
			Author: fmt.Sprintf("%s",
				b.preparedBooks[i].Opf.Metadata.Creator[0].Data,
			),
		})
	}
	return result
}

func (b *TextParser) canBookBeUsed(dir, file string) (*epub.EpubBook, error) {
	return epub.NewEpubBook(dir + string(os.PathSeparator) + file)
}

func (b *TextParser) charsetReader(c string, i io.Reader) (r io.Reader, e error) {
	if c == "windows-1251" {
		r = utils.DecodeWin1251(i)
	}
	return
}

func (b *TextParser) GetTextChapters(resp *ParserResponse) []Chapter {
	result := make([]Chapter, 0)
	for i := range b.preparedBooks[resp.FileID].Ncx.NavMap.NavPoint {
		// book
		for j, c := range b.preparedBooks[resp.FileID].Ncx.NavMap.NavPoint[i].NavPoint {
			// book chapter
			tmp := Chapter{
				ChapterTitle: c.NavLabel.Text,
			}
			paragraphList := make([]string, 0, 5)
			for _, k := range c.ChapterContent.P {
				if len(paragraphList) >= 3 {
					break
				}
				txt := strings.TrimSpace(k.Text)
				if len(txt) > 0 {
					paragraphList = append(paragraphList, txt)
				}
				for _, e := range k.Em {
					txt = strings.TrimSpace(e)
					if len(txt) > 0 {
						paragraphList = append(paragraphList, txt)
					}
				}
			}
			tmp.FullText = utils.CutText(strings.Join(paragraphList, "\n"), 200)
			tmp.Offset = i + j + 1
			result = append(result, tmp)
		}
	}
	return result
}
