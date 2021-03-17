package txtparser

import "encoding/xml"

type ParserResponse struct {
	Title  string
	Author string
	File   string
	FileID int
}

type Chapter struct {
	ChapterTitle string
	FullText     string
	Offset       int
}

type AuthorType struct {
	Text     string `xml:",chardata"`
	Nickname string `xml:"nickname"`
	Email    string `xml:"email"`

	FirstName  string `xml:"first-name"`
	MiddleName string `xml:"middle-name"`
	LastName   string `xml:"last-name"`
	HomePage   string `xml:"home-page"`
}

type Description struct {
	XMLName   xml.Name `xml:"description"`
	Text      string   `xml:",chardata"`
	TitleInfo struct {
		Text       string     `xml:",chardata"`
		Genre      []string   `xml:"genre"`
		Author     AuthorType `xml:"author"`
		BookTitle  string     `xml:"book-title"`
		Annotation struct {
			Text string   `xml:",chardata"`
			P    []string `xml:"p"`
		} `xml:"annotation"`
		Date struct {
			Text  string `xml:",chardata"`
			Value string `xml:"value,attr"`
		} `xml:"date"`
		Lang     string `xml:"lang"`
		SrcLang  string `xml:"src-lang"`
		Sequence struct {
			Text   string `xml:",chardata"`
			Name   string `xml:"name,attr"`
			Number string `xml:"number,attr"`
		} `xml:"sequence"`
	} `xml:"title-info"`

	PublishInfo struct {
		Text      string `xml:",chardata"`
		BookName  string `xml:"book-name"`
		Publisher string `xml:"publisher"`
		City      string `xml:"city"`
		Year      string `xml:"year"`
	} `xml:"publish-info"`
}

type Body struct {
	XMLName xml.Name `xml:"body"`
	Text    string   `xml:",chardata"`
	Title   []struct {
		Text string   `xml:",chardata"`
		P    []string `xml:"p"`
	} `xml:"title"`
	Section []struct {
		Text  string `xml:",chardata"`
		Title struct {
			Text string `xml:",chardata"`
			P    string `xml:"p"`
		} `xml:"title"`
		Section []struct {
			Text  string `xml:",chardata"`
			Title struct {
				Text string `xml:",chardata"`
				P    string `xml:"p"`
			} `xml:"title"`
			P []string `xml:"p"`
		} `xml:"section"`
	} `xml:"section"`
}

type FB2 struct {
	Description Description `xml:"description"`
	Body        []Body      `xml:"body"`
}
