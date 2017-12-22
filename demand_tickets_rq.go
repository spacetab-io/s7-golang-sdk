package sdk

import (
	"encoding/xml"
)

type AirDocIssueRQ struct {
	XMLName  xml.Name `xml:"http://www.iata.org/IATA/EDIST AirDocIssueRQ"`
	Version  string   `xml:"Version,attr"`
	Document *Document
	Party    *Party
	Query    *Query
}

func MakeAirDocIssueRQ() (request *Envelope) {
	request = &Envelope{
		Header: new(Header),
		Body: &Body{
			AirDocIssueRQ: &AirDocIssueRQ{
				Version:  "1.0",
				Document: new(Document),
				Party: &Party{
					Sender: new(Sender),
				},
				Query: &Query{
					DataLists: &DataLists{
						CheckedBagAllowanceList: new(CheckedBagAllowanceList),
						FareList:                new(FareList),
						FlightSegmentList:       new(FlightSegmentList),
					},
				},
			},
		},
	}

	return
}

type TicketDocInfo struct {
	TravelerInfo          *TravelerInfo
	Traveler              *TravelerInfo       `xml:",omitempty"`
	IssuingAirlineInfo    *IssuingAirlineInfo `xml:",omitempty"`
	BookingReference      *BookingReference
	BookingReferences     *BookingReferences `xml:",omitempty"`
	Payments              *Payments
	OriginDestination     *OriginDestinationDemand `xml:",omitempty"`
	FareInfo              *FareInfo                `xml:",omitempty"`
	TicketDocument        *TicketDocument          `xml:",omitempty"`
	Commission            *Commission              `xml:",omitempty"`
	PassengerSpecificData string                   `xml:",omitempty"`
}

type TravelerInfo struct {
	Surname string
	Given   string
	PTC     *PTC
}

type Payments struct {
	Payment []*Payment
}

type Payment struct {
	Type   *Type
	Other  *Other
	Amount *Total `xml:",omitempty"`
}

type Type struct {
	Code string
}
