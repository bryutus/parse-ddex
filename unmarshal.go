package main

import (
	"encoding/xml"
	"fmt"
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
	TechnicalSoundRecordingDetail TechnicalDetail `xml:"TechnicalSoundRecordingDetails"`
}

type Image struct {
	ImageType               string                  `xml:"ImageType"`
	ImageID                 string                  `xml:"ImageId>ProprietaryId"`
	ResourceReference       string                  `xml:"ResourceReference"`
	ImageDetailsByTerritory ImageDetailsByTerritory `xml:"ImageDetailsByTerritory"`
}

type ImageDetailsByTerritory struct {
	TechnicalImageDetail TechnicalDetail `xml:"TechnicalImageDetails"`
}

type TechnicalDetail struct {
	TechnicalResourceDetailsReference string `xml:"TechnicalResourceDetailsReference"`
	FileName                          string `xml:"File>FileName"`
	FilePath                          string `xml:"File>FilePath"`
}

type ReleaseList struct {
	Releases []Release `xml:"Release"`
}

type Release struct {
	IsMainRelease               bool                        `xml:"IsMainRelease,attr"`
	ICPN                        string                      `xml:"ReleaseId>ICPN"`
	ISRC                        string                      `xml:"ReleaseId>ISRC"`
	ReleaseReference            string                      `xml:"ReleaseReference"`
	ReleaseType                 string                      `xml:"ReleaseType"`
	ReleaseResourceReferences   []ReleaseResourceReference  `xml:"ReleaseResourceReferenceList>ReleaseResourceReference"`
	ReleaseDetailsByTerritories []ReleaseDetailsByTerritory `xml:"ReleaseDetailsByTerritory"`
	PLine                       PLine                       `xml:"PLine"`
}

type ReleaseResourceReference struct {
	ReleaseResourceReference string `xml:",innerxml"`
	ReleaseResourceType      string `xml:"ReleaseResourceType,attr"`
}

type ReleaseDetailsByTerritory struct {
	TerritoryCode       string          `xml:"TerritoryCode"`
	DisplayArtistName   string          `xml:"DisplayArtistName"`
	LabelName           string          `xml:"LabelName"`
	Titles              []Title         `xml:"Title"`
	DisplayArtists      []DisplayArtist `xml:"DisplayArtist"`
	ParentalWarningType string          `xml:"ParentalWarningType"`
	Genre               string          `xml:"Genre>GenreText"`
	ResourceGroup       ResourceGroup   `xml:"ResourceGroup"`
}

type Title struct {
	LanguageAndScriptCode string `xml:"LanguageAndScriptCode,attr"`
	Type                  string `xml:"TitleType,attr"`
	Text                  string `xml:"TitleText"`
}

type DisplayArtist struct {
	SequenceNumber uint16    `xml:"SequenceNumber,attr"`
	PartyName      string    `xml:"PartyName>FullName"`
	PartyIds       []PartyId `xml:"PartyId"`
	Roles          []string  `xml:"ArtistRole"`
}

type PartyId struct {
	PartyId string `xml:",innerxml"`
	IsISNI  bool   `xml:"IsISNI,attr"`
	IsDPID  bool   `xml:"IsDPID,attr"`
}

type ResourceGroup struct {
	TrackResources []TrackResource `xml:"ResourceGroup"`
	AlbumContent   AlbumContent    `xml:"ResourceGroupContentItem"`
}

type TrackResource struct {
	DiscNo        uint16         `xml:"SequenceNumber"`
	TrackContents []TrackContent `xml:"ResourceGroupContentItem"`
}

type TrackContent struct {
	TrackNo                  uint16                   `xml:"SequenceNumber"`
	ResourceType             string                   `xml:"ResourceType"`
	ReleaseResourceReference ReleaseResourceReference `xml:"ReleaseResourceReference"`
}

type AlbumContent struct {
	SequenceNumber           uint16                   `xml:"SequenceNumber"`
	ResourceType             string                   `xml:"ResourceType"`
	ReleaseResourceReference ReleaseResourceReference `xml:"ReleaseResourceReference"`
}

type PLine struct {
	Year string `xml:"Year"`
	Text string `xml:"PLineText"`
}

type DealList struct {
	ReleaseDeals []ReleaseDeal `xml:"ReleaseDeal"`
}

type ReleaseDeal struct {
	DealReleaseReferences []string `xml:"DealReleaseReference"`
	Deals                 []Deal   `xml:"Deal"`
}

type Deal struct {
	DealTerms     DealTerms `xml:"DealTerms"`
	DealReference string    `xml:"DealReference"`
}

type DealTerms struct {
	CommercialModelTypes []string       `xml:"CommercialModelType"`
	UseTypes             []string       `xml:"Usage>UseType"`
	TerritoryCode        string         `xml:"TerritoryCode"`
	ValidityPeriod       ValidityPeriod `xml:"ValidityPeriod"`
	TakeDown             bool           `xml:"TakeDown"`
}

type ValidityPeriod struct {
	StartDate     string `xml:"StartDate"`
	EndDate       string `xml:"EndDate"`
	StartDateTime string `xml:"StartDateTime"`
	EndDateTime   string `xml:"EndDateTime"`
}

func Unmarshal(b []byte) {
	m := new(NewReleaseMessage)
	if err := xml.Unmarshal(b, m); err != nil {
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
		t := s.SoundRecordingDetailsByTerritorry.TechnicalSoundRecordingDetail
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
		fmt.Printf("\tTechnicalResourceDetailsReference: %v\n", i.ImageDetailsByTerritory.TechnicalImageDetail.TechnicalResourceDetailsReference)
		fmt.Printf("\tFileName: %v\n", i.ImageDetailsByTerritory.TechnicalImageDetail.FileName)
		fmt.Printf("\tFilePath: %v\n", i.ImageDetailsByTerritory.TechnicalImageDetail.FilePath)

		fmt.Printf("\n")
	}

	fmt.Printf("\n")

	fmt.Println("##################")
	fmt.Println("# ReleaseList")
	fmt.Println("##################")
	fmt.Println("Releases:")
	for _, r := range m.ReleaseList.Releases {
		id := r.ISRC
		if r.IsMainRelease == true {
			id = r.ICPN
		}
		fmt.Printf("\tReleaseReference: %v\n", r.ReleaseReference)
		fmt.Printf("\tReleaseType: %v\n", r.ReleaseType)
		fmt.Printf("\tICPN/ISRC: %v\n", id)
		fmt.Printf("\tPLine - Year: %v\n", r.PLine.Year)
		fmt.Printf("\tPLine - Text: %v\n", r.PLine.Text)
		fmt.Println("\tReleaseResourceReferences:")
		for _, ref := range r.ReleaseResourceReferences {
			fmt.Printf("\t\tReleaseResourceReference: %v\n", ref.ReleaseResourceReference)
			fmt.Printf("\t\tReleaseResourceType: %v\n", ref.ReleaseResourceType)
			fmt.Printf("\n")
		}

		fmt.Println("\tReleaseDetailsByTerritories:")
		for _, rel := range r.ReleaseDetailsByTerritories {
			fmt.Printf("\t\tTerritoryCode: %v\n", rel.TerritoryCode)
			fmt.Printf("\t\tDisplayArtistName: %v\n", rel.DisplayArtistName)
			fmt.Println("\t\tTitles:")
			for _, t := range rel.Titles {
				fmt.Printf("\t\t\tLanguageAndScriptCode: %v\n", t.LanguageAndScriptCode)
				fmt.Printf("\t\t\tType: %v\n", t.Type)
				fmt.Printf("\t\t\tText: %v\n", t.Text)
				fmt.Printf("\n")
			}
			fmt.Println("\t\tDisplayArtists:")
			for _, a := range rel.DisplayArtists {
				fmt.Printf("\t\t\tSequenceNumber: %v\n", a.SequenceNumber)
				fmt.Printf("\t\t\tPartyName: %v\n", a.PartyName)

				fmt.Println("\t\t\tPartyIds:")
				for _, p := range a.PartyIds {
					fmt.Printf("\t\t\t\tPartyId: %v\n", p.PartyId)
					fmt.Printf("\t\t\t\tIsDPID: %v\n", p.IsDPID)
					fmt.Printf("\t\t\t\tIsISNI: %v\n", p.IsISNI)
					fmt.Printf("\n")
				}

				fmt.Println("\t\t\tRoles:")
				for _, role := range a.Roles {
					fmt.Printf("\t\t\t\tRole: %v\n", role)
				}

				fmt.Printf("\n")
			}

			fmt.Printf("\t\tLabelName: %v\n", rel.LabelName)
			fmt.Printf("\t\tParentalWarningType: %v\n", rel.ParentalWarningType)
			fmt.Printf("\t\tGenre: %v\n", rel.Genre)

			fmt.Printf("\n")

			if r.IsMainRelease == true {
				fmt.Println("\t\tTrackResources:")
				for _, tr := range rel.ResourceGroup.TrackResources {
					fmt.Printf("\t\t\tDiscNo: %v\n", tr.DiscNo)
					fmt.Println("\t\t\tTrackContents:")
					for _, tc := range tr.TrackContents {
						fmt.Printf("\t\t\t\tTrackNo: %v\n", tc.TrackNo)
						fmt.Printf("\t\t\t\tResourceType: %v\n", tc.ResourceType)
						fmt.Printf("\t\t\t\tReleaseResourceReference: %v\n", tc.ReleaseResourceReference.ReleaseResourceReference)
						fmt.Printf("\t\t\t\tReleaseResourceType: %v\n", tc.ReleaseResourceReference.ReleaseResourceType)
						fmt.Printf("\n")
					}
				}

				fmt.Println("\t\tAlbumContent:")
				fmt.Printf("\t\t\tSequenceNumber: %v\n", rel.ResourceGroup.AlbumContent.SequenceNumber)
				fmt.Printf("\t\t\tResourceType: %v\n", rel.ResourceGroup.AlbumContent.ResourceType)
				fmt.Printf("\t\t\tReleaseResourceReference: %v\n", rel.ResourceGroup.AlbumContent.ReleaseResourceReference.ReleaseResourceReference)
				fmt.Printf("\t\t\tReleaseResourceType: %v\n", rel.ResourceGroup.AlbumContent.ReleaseResourceReference.ReleaseResourceType)

				fmt.Printf("\n")
			}
		}
		fmt.Printf("\n")
	}

	fmt.Printf("\n")

	fmt.Println("##################")
	fmt.Println("# DealList")
	fmt.Println("##################")
	fmt.Println("ReleaseDeals:")
	for _, r := range m.DealList.ReleaseDeals {
		fmt.Println("\tDealReleaseReferences:")
		fmt.Printf("\t\tDealReleaseReference: %v\n", r.DealReleaseReferences)
		fmt.Println("\tDeals>DealTerms:")
		for _, d := range r.Deals {
			fmt.Printf("\t\tCommercialModelType: %v\n", d.DealTerms.CommercialModelTypes)
			fmt.Printf("\t\tUseType: %v\n", d.DealTerms.UseTypes)
			fmt.Printf("\t\tTerritoryCode: %v\n", d.DealTerms.TerritoryCode)
			fmt.Printf("\t\tStartDate: %v\n", d.DealTerms.ValidityPeriod.StartDate)
			fmt.Printf("\t\tEndDate: %v\n", d.DealTerms.ValidityPeriod.EndDate)
			fmt.Printf("\t\tDealReference: %v\n", d.DealReference)
			fmt.Printf("\n")
		}
	}

}
