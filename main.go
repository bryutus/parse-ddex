package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

type NewReleaseMessage struct {
	XMLName                xml.Name      `xml:"NewReleaseMessage"`
	LanguageAndScriptCode  string        `xml:"LanguageAndScriptCode,attr"`
	MessageSchemaVersionID string        `xml:"MessageSchemaVersionId,attr"`
	MessageHeader          MessageHeader `xml:"MessageHeader"`
	ResourceList           ResourceList  `xml:"ResourceList"`
	ReleaseList            ReleaseList   `xml:"ReleaseList"`
	DealList               DealList      `xml:"DealList"`
	// Description            string        `xml:",innerxml"`
}

type MessageHeader struct {
	MessageID              string             `xml:"MessageId"`
	MessageSender          MessageSender      `xml:"MessageSender"`
	SentOnBehalfOf         MessageSender      `xml:"SentOnBehalfOf"`
	MessageRecipients      []MessageRecipient `xml:"MessageRecipient"`
	MessageCreatedDateTime string             `xml:"MessageCreatedDateTime"`
	MessageControlType     string             `xml:"MessageControlType"`
}

type MessageSender struct {
	PartyID   string `xml:"PartyId"`
	PartyName string `xml:"PartyName>FullName"`
}

type MessageRecipient struct {
	PartyID   string `xml:"PartyId"`
	PartyName string `xml:"PartyName>FullName"`
}

type ResourceList struct {
	SoundRecordings []SoundRecording `xml:"SoundRecording"`
	Images          []Image          `xml:"Image"`
}

type SoundRecording struct {
	SoundRecordingType                string                           `xml:"SoundRecordingType"`
	ISWC                              string                           `xml:"IndirectSoundRecordingId>ISWC"`
	ResourceReference                 string                           `xml:"ResourceReference"`
	Duration                          string                           `xml:"Duration"`
	SoundRecordingDetailsByTerritorry SoundRecordingDetailsByTerritory `xml:"SoundRecordingDetailsByTerritory"`
}

type SoundRecordingDetailsByTerritory struct {
	TechnicalSoundRecordingDetails TechnicalSoundRecordingDetail `xml:"TechnicalSoundRecordingDetails"`
}

type TechnicalSoundRecordingDetail struct {
	TechnicalResourceDetailsReference string `xml:"TechnicalResourceDetailsReference"`
	FileName                          string `xml:"File>FileName"`
	FilePath                          string `xml:"File>FilePath"`
}

type Image struct {
	ImageType         string `xml:"ImageType"`
	ImageID           string `xml:"ImageId>ProprietaryId"`
	ResourceReference string `xml:"ResourceReference"`
}

type ReleaseList struct {
	Releases []Release `xml:"Release"`
}

type Release struct {
	IsMainRelease    bool   `xml:"IsMainRelease"`
	ICPN             string `xml:"ReleaseId>ICPN"`
	ReleaseReference string `xml:"ReleaseReference"`
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

	m := new(NewReleaseMessage)
	if err := xml.Unmarshal(bytes, m); err != nil {
		panic(err)
	}

	fmt.Println("##################")
	fmt.Println("# Message")
	fmt.Println("##################")
	fmt.Printf("XMLName: %v\n", m.XMLName)
	fmt.Printf("LanguageAndScriptCode: %v\n", m.LanguageAndScriptCode)
	fmt.Printf("MessageSchemaVersionID: %v\n", m.MessageSchemaVersionID)

	fmt.Printf("\n")

	fmt.Println("##################")
	fmt.Println("# MessageHeader")
	fmt.Println("##################")
	fmt.Printf("MessageId: %v\n", m.MessageHeader.MessageID)
	fmt.Printf("MessageCreatedDateTime: %v\n", m.MessageHeader.MessageCreatedDateTime)
	fmt.Printf("MessageControlType: %v\n", m.MessageHeader.MessageControlType)
	fmt.Printf("MessageSender PartyID: %v, PartyName: %v\n", m.MessageHeader.MessageSender.PartyID, m.MessageHeader.MessageSender.PartyName)
	fmt.Printf("SentOnBehalfOf PartyID: %v, PartyName: %v\n", m.MessageHeader.SentOnBehalfOf.PartyID, m.MessageHeader.SentOnBehalfOf.PartyName)
	fmt.Println("MessageRecipients:")
	for _, r := range m.MessageHeader.MessageRecipients {
		fmt.Printf("\tPartyID: %v, PartyName: %v\n", r.PartyID, r.PartyName)
	}

	fmt.Printf("\n")

	fmt.Println("##################")
	fmt.Println("# ResourceList")
	fmt.Println("##################")
	fmt.Println("SoundRecordings:")
	for _, s := range m.ResourceList.SoundRecordings {
		fmt.Printf("\tResourceReference: %v\n", s.ResourceReference)
		fmt.Printf("\tSoundRecordingType: %v\n", s.SoundRecordingType)
		fmt.Printf("\tISWC: %v\n", s.ISWC)
		fmt.Printf("\tDuration: %v\n", s.Duration)
		t := s.SoundRecordingDetailsByTerritorry.TechnicalSoundRecordingDetails
		fmt.Printf("\tTechnicalResourceDetailsReference: %v\n", t.TechnicalResourceDetailsReference)
		fmt.Printf("\tFileName: %v\n", t.FileName)
		fmt.Printf("\tFilePath: %v\n", t.FilePath)
		fmt.Printf("\n")
	}
	fmt.Println("Images:")
	for _, i := range m.ResourceList.Images {
		fmt.Printf("\tResourceReference: %v\n", i.ResourceReference)
		fmt.Printf("\tImageID: %v\n", i.ImageID)
		fmt.Printf("\tImageType: %v\n", i.ImageType)
		fmt.Printf("\n")
	}

	fmt.Printf("\n")

	fmt.Println("##################")
	fmt.Println("# ReleaseList")
	fmt.Println("##################")
	fmt.Printf("%+v\n", m.ReleaseList.Releases[0])

	fmt.Printf("\n")

	fmt.Println("##################")
	fmt.Println("# DealList")
	fmt.Println("##################")
	fmt.Printf("%+v\n", m.DealList.Deals[0])
}
