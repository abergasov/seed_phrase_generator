package epub

import (
	"archive/zip"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"path"
	"strings"
)

//Container META-INF/container.xml file
type Container struct {
	Rootfile Rootfile `xml:"rootfiles>rootfile" json:"rootfile"`
}

//Rootfile root file
type Rootfile struct {
	Path string `xml:"full-path,attr" json:"path"`
	Type string `xml:"media-type,attr" json:"type"`
}

type EpubBook struct {
	Ncx       Ncx       `json:"ncx"`
	Opf       Opf       `json:"opf"`
	Container Container `json:"-"`
	Mimetype  string    `json:"-"`

	fd *zip.ReadCloser
}

func NewEpubBook(bookFile string) (*EpubBook, error) {
	fd, err := zip.OpenReader(bookFile)
	bk := EpubBook{fd: fd}
	if err != nil {
		return nil, err
	}
	defer fd.Close()
	mt, err := bk.readBytes("mimetype")
	if err != nil {
		return nil, err
	}

	bk.Mimetype = string(mt)
	err = bk.readXML("META-INF/container.xml", &bk.Container)
	err = bk.readXML(bk.Container.Rootfile.Path, &bk.Opf)

	for _, mf := range bk.Opf.Manifest {
		if mf.ID == bk.Opf.Spine.Toc {
			err = bk.readXML(bk.filename(mf.Href), &bk.Ncx)
			break
		}
	}
	for i := range bk.Ncx.NavMap.NavPoint {
		for j, nc := range bk.Ncx.NavMap.NavPoint[i].NavPoint {
			data := strings.Split(nc.Content.Src, "#")
			var chC ChapterContent
			err = bk.readXML(bk.filename(data[0]), &chC)
			for _, c := range chC.Body.Span.Span {
				if c.ID != data[1] {
					continue
				}
				bk.Ncx.NavMap.NavPoint[i].NavPoint[j].ChapterContent = c
			}
		}
	}
	return &bk, err
}

//Open open resource file
func (p *EpubBook) Open(n string) (io.ReadCloser, error) {
	return p.open(p.filename(n))
}

//Files list resource files
func (p *EpubBook) Files() []string {
	var fns []string
	for _, f := range p.fd.File {
		fns = append(fns, f.Name)
	}
	return fns
}

//Close close file reader
func (p *EpubBook) Close() {
	p.fd.Close()
}

func (p *EpubBook) filename(n string) string {
	return path.Join(path.Dir(p.Container.Rootfile.Path), n)
}

func (p *EpubBook) readXML(n string, v interface{}) error {
	fd, err := p.open(n)
	if err != nil {
		return nil
	}
	defer fd.Close()
	dec := xml.NewDecoder(fd)
	return dec.Decode(v)
}

func (p *EpubBook) readBytes(n string) ([]byte, error) {
	fd, err := p.open(n)
	if err != nil {
		return nil, nil
	}
	defer fd.Close()

	return ioutil.ReadAll(fd)

}

func (p *EpubBook) open(n string) (io.ReadCloser, error) {
	for _, f := range p.fd.File {
		if f.Name == n {
			return f.Open()
		}
	}
	return nil, fmt.Errorf("file %s not exist", n)
}
