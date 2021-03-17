package epub

import (
	"encoding/xml"
)

//Ncx OPS/toc.ncx
type NcxL struct {
	Points []NavPoint `xml:"navMap>navPoint" json:"points"`
}

//NavPoint nav point
type NavPoint struct {
	Text    string     `xml:"navLabel>text" json:"text"`
	Content Content    `xml:"content" json:"content"`
	Points  []NavPoint `xml:"navPoint" json:"points"`
}

//Content nav-point content
type Content struct {
	Src string `xml:"src,attr" json:"src"`
}

type Ncx struct {
	XMLName xml.Name `xml:"ncx"`
	Text    string   `xml:",chardata"`
	Xmlns   string   `xml:"xmlns,attr"`
	Version string   `xml:"version,attr"`
	Lang    string   `xml:"lang,attr"`
	Head    struct {
		Text string `xml:",chardata"`
		Meta []struct {
			Text    string `xml:",chardata"`
			Name    string `xml:"name,attr"`
			Content string `xml:"content,attr"`
		} `xml:"meta"`
	} `xml:"head"`
	DocTitle struct {
		Chardata string `xml:",chardata"`
		Text     string `xml:"text"`
	} `xml:"docTitle"`
	NavMap struct {
		Text     string `xml:",chardata"`
		NavPoint []struct {
			Text      string `xml:",chardata"`
			PlayOrder string `xml:"playOrder,attr"`
			ID        string `xml:"id,attr"`
			NavLabel  struct {
				Chardata string `xml:",chardata"`
				Text     string `xml:"text"`
			} `xml:"navLabel"`
			Content struct {
				Text string `xml:",chardata"`
				Src  string `xml:"src,attr"`
			} `xml:"content"`
			NavPoint []struct {
				Text      string `xml:",chardata"`
				PlayOrder string `xml:"playOrder,attr"`
				ID        string `xml:"id,attr"`
				NavLabel  struct {
					Chardata string `xml:",chardata"`
					Text     string `xml:"text"`
				} `xml:"navLabel"`
				Content struct {
					Text string `xml:",chardata"`
					Src  string `xml:"src,attr"`
				} `xml:"content"`
				ChapterContent ChapterSpan
			} `xml:"navPoint"`
		} `xml:"navPoint"`
	} `xml:"navMap"`
}

type Opf struct {
	Metadata Metadata   `xml:"metadata" json:"metadata"`
	Manifest []Manifest `xml:"manifest>item" json:"manifest"`
	Spine    Spine      `xml:"spine" json:"spine"`
}

//Metadata metadata
type Metadata struct {
	Title       []string     `xml:"title" json:"title"`
	Language    []string     `xml:"language" json:"language"`
	Identifier  []Identifier `xml:"identifier" json:"identifier"`
	Creator     []Author     `xml:"creator" json:"creator"`
	Subject     []string     `xml:"subject" json:"subject"`
	Description []string     `xml:"description" json:"description"`
	Publisher   []string     `xml:"publisher" json:"publisher"`
	Contributor []Author     `xml:"contributor" json:"contributor"`
	Date        []Date       `xml:"date" json:"date"`
	Type        []string     `xml:"type" json:"type"`
	Format      []string     `xml:"format" json:"format"`
	Source      []string     `xml:"source" json:"source"`
	Relation    []string     `xml:"relation" json:"relation"`
	Coverage    []string     `xml:"coverage" json:"coverage"`
	Rights      []string     `xml:"rights" json:"rights"`
	Meta        []Metafield  `xml:"meta" json:"meta"`
}

// Identifier identifier
type Identifier struct {
	Data   string `xml:",chardata" json:"data"`
	ID     string `xml:"id,attr" json:"id"`
	Scheme string `xml:"scheme,attr" json:"scheme"`
}

// Author author
type Author struct {
	Data   string `xml:",chardata" json:"author"`
	FileAs string `xml:"file-as,attr" json:"file_as"`
	Role   string `xml:"role,attr" json:"role"`
}

// Date date
type Date struct {
	Data  string `xml:",chardata" json:"data"`
	Event string `xml:"event,attr" json:"event"`
}

// Metafield metafield
type Metafield struct {
	Name    string `xml:"name,attr" json:"name"`
	Content string `xml:"content,attr" json:"content"`
}

//Manifest manifest
type Manifest struct {
	ID           string `xml:"id,attr" json:"id"`
	Href         string `xml:"href,attr" json:"href"`
	MediaType    string `xml:"media-type,attr" json:"type"`
	Fallback     string `xml:"media-fallback,attr" json:"fallback"`
	Properties   string `xml:"properties,attr" json:"properties"`
	MediaOverlay string `xml:"media-overlay,attr" json:"overlay"`
}

// Spine spine
type Spine struct {
	ID              string      `xml:"id,attr" json:"id"`
	Toc             string      `xml:"toc,attr" json:"toc"`
	PageProgression string      `xml:"page-progression-direction,attr" json:"progression"`
	Items           []SpineItem `xml:"itemref" json:"items"`
}

// SpineItem spine item
type SpineItem struct {
	IDref      string `xml:"idref,attr" json:"id_ref"`
	Linear     string `xml:"linear,attr" json:"linear"`
	ID         string `xml:"id,attr" json:"id"`
	Properties string `xml:"properties,attr" json:"properties"`
}

type ChapterSpan struct {
	Text string `xml:",chardata"`
	ID   string `xml:"id,attr"`
	Div  struct {
		Text  string `xml:",chardata"`
		Class string `xml:"class,attr"`
		P     []struct {
			Text  string `xml:",chardata"`
			Class string `xml:"class,attr"`
		} `xml:"p"`
	} `xml:"div"`
	P []struct {
		Text  string   `xml:",chardata"`
		Class string   `xml:"class,attr"`
		Em    []string `xml:"em"`
		A     []struct {
			Text  string `xml:",chardata"`
			Href  string `xml:"href,attr"`
			Class string `xml:"class,attr"`
		} `xml:"a"`
	} `xml:"p"`
}

type ChapterContent struct {
	XMLName xml.Name `xml:"html"`
	Text    string   `xml:",chardata"`
	Xmlns   string   `xml:"xmlns,attr"`
	Head    struct {
		Text  string `xml:",chardata"`
		Title string `xml:"title"`
		Link  struct {
			Text string `xml:",chardata"`
			Rel  string `xml:"rel,attr"`
			Href string `xml:"href,attr"`
			Type string `xml:"type,attr"`
		} `xml:"link"`
	} `xml:"head"`
	Body struct {
		Text  string `xml:",chardata"`
		Class string `xml:"class,attr"`
		Span  struct {
			Text string        `xml:",chardata"`
			Span []ChapterSpan `xml:"span"`
		} `xml:"span"`
	} `xml:"body"`
}
