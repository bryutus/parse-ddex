package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

type NewReleaseMessage struct {
	// XMLName                xml.Name      `xml:"NewReleaseMessage"`
	// LanguageAndScriptCode  string        `xml:"LanguageAndScriptCode,attr"`
	// MessageSchemaVersionID string        `xml:"MessageSchemaVersionId,attr"`
	// Description            string        `xml:",innerxml"`
	MessageHeader MessageHeader `xml:"MessageHeader"`
	ResourceList  ResourceList  `xml:"ResourceList"`
	ReleaseList   ReleaseList   `xml:"ReleaseList"`
	DealList      DealList      `xml:"DealList"`
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
	Resources []SoundRecording `xml:"SoundRecording"`
}

type SoundRecording struct {
	SoundRecordingType string `xml:"SoundRecordingType"`
	ISRC               string `xml:"SoundRecordingId>ISRC"`
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

	message := new(NewReleaseMessage)

	if err := xml.Unmarshal(bytes, message); err != nil {
		panic(err)
	}

	fmt.Println("##################")
	fmt.Println("# MessageHeader")
	fmt.Println("##################")
	fmt.Printf("MessageId: %v\n", message.MessageHeader.MessageID)
	fmt.Printf("MessageCreatedDateTime: %v\n", message.MessageHeader.MessageCreatedDateTime)
	fmt.Printf("MessageControlType: %v\n", message.MessageHeader.MessageControlType)
	fmt.Printf("MessageSender PartyID: %v, PartyName: %v\n", message.MessageHeader.MessageSender.PartyID, message.MessageHeader.MessageSender.PartyName)
	fmt.Printf("SentOnBehalfOf PartyID: %v, PartyName: %v\n", message.MessageHeader.SentOnBehalfOf.PartyID, message.MessageHeader.SentOnBehalfOf.PartyName)
	fmt.Println("MessageRecipients:")
	for _, r := range message.MessageHeader.MessageRecipients {
		fmt.Printf("\tPartyID: %v, PartyName: %v\n", r.PartyID, r.PartyName)
	}

	fmt.Printf("\n")

	fmt.Println("##################")
	fmt.Println("# ResourceList")
	fmt.Println("##################")
	fmt.Printf("%+v\n", message.ResourceList.Resources[0])

	fmt.Printf("\n")

	fmt.Println("##################")
	fmt.Println("# ReleaseList")
	fmt.Println("##################")
	fmt.Printf("%+v\n", message.ReleaseList.Releases[0])

	fmt.Printf("\n")

	fmt.Println("##################")
	fmt.Println("# DealList")
	fmt.Println("##################")
	fmt.Printf("%+v\n", message.DealList.Deals[0])
}
