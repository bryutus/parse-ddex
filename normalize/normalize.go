package normalize

import (
	"encoding/xml"
	"fmt"
)

type NewReleaseMessage struct {
	ReleaseList ReleaseList `xml:"ReleaseList"`
}

type ReleaseList struct {
	Releases []Release `xml:"Release"`
}

type Release struct {
	IsMainRelease bool   `xml:"IsMainRelease,attr"`
	ICPN          string `xml:"ReleaseId>ICPN"`
	ISRC          string `xml:"ReleaseId>ISRC"`
}

func Exec(b []byte) {
	a := Album{}
	a.Show(b)
}

type Album struct{}

func (a Album) Show(b []byte) {
	n := new(NewReleaseMessage)
	if err := xml.Unmarshal(b, n); err != nil {
		panic(err)
	}

	for _, r := range n.ReleaseList.Releases {
		fmt.Println(r)
	}
}
