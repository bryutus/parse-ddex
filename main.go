package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

type NewReleaseMessage struct {
	LanguageAndScriptCode  string        `xml:"LanguageAndScriptCode,attr"`
	MessageSchemaVersionID string        `xml:"MessageSchemaVersionId,attr"`
	MessageHeader          MessageHeader `xml:"MessageHeader"`
	ResourceList           ResourceList  `xml:"ResourceList"`
	ReleaseList            ReleaseList   `xml:"ReleaseList"`
	DealList               DealList      `xml:"DealList"`
}

type MessageHeader struct {
	MessageID string `xml:"MessageId"`
}

type ResourceList struct {
	Resources []SoundRecording `xml:"SoundRecording"`
}

type SoundRecording struct {
	SoundRecordingType string           `xml:"SoundRecordingType"`
	SoundRecordingID   SoundRecordingID `xml:"SoundRecordingId"`
}

type SoundRecordingID struct {
	ISRC string `xml:"ISRC"`
}

type ReleaseList struct {
	Releases []Release `xml:"Release"`
}

type Release struct {
	IsMainRelease    bool      `xml:"IsMainRelease"`
	ReleaseID        ReleaseID `xml:"ReleaseId"`
	ReleaseReference string    `xml:"ReleaseReference"`
}

type ReleaseID struct {
	ICPN string `xml:"ICPN"`
}

type DealList struct {
	Deals []ReleaseDeal `xml:"ReleaseDeal"`
}

type ReleaseDeal struct {
}

func main() {
	bytes, err := ioutil.ReadFile("resources/721620118165_combined.xml")
	if err != nil {
		panic(err)
	}

	message := new(NewReleaseMessage)

	if err := xml.Unmarshal(bytes, message); err != nil {
		panic(err)
	}

	fmt.Println(message.LanguageAndScriptCode)
	fmt.Println(message.MessageSchemaVersionID)
	fmt.Println(message.MessageHeader.MessageID)
	fmt.Printf("%+v\n", message.ResourceList.Resources[0])
	fmt.Printf("%+v\n", message.ReleaseList.Releases[0])
	fmt.Printf("%+v\n", message.DealList.Deals[0])
}
