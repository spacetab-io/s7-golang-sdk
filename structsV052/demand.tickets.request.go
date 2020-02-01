package structsV052

import (
	"encoding/xml"
	"gitlab.teamc.io/tm-consulting/tmc24/avia/layer3/s7-agent-go/pkg/models"
)

type AirDocIssueRQ struct {
	XMLName  xml.Name `xml:"http://www.iata.org/IATA/EDIST AirDocIssueRQ"`
	Version  string   `xml:"Version,attr"`
	Document *Document
	Party    *Party
	Query    *QueryAirDocIssue
}

type QueryAirDocIssue struct {
	Passengers        *Passengers
	OrderItems        *OrderItems
	TicketDocQuantity int              `xml:",omitempty"`
	TicketDocInfo     []*TicketDocInfo `xml:",omitempty"`
	DataLists         *DataLists
	Reshop            *Reshop `xml:",omitempty"`
}

type TicketDocInfo struct {
	TravelerInfo          *TravelerInfo
	Traveler              *TravelerDoc        `xml:",omitempty"`
	IssuingAirlineInfo    *IssuingAirlineInfo `xml:",omitempty"`
	BookingReference      *BookingReference
	BookingReferences     *BookingReferences `xml:",omitempty"`
	Payments              *Payments
	OriginDestination     *OriginDestinationDemand `xml:",omitempty"`
	FareInfo              *FareInfo                `xml:",omitempty"`
	TicketDocument        *TicketDocument          `xml:",omitempty"`
	Price                 *TicketPrice             `xml:",omitempty"`
	Commission            *Commission              `xml:",omitempty"`
	PassengerSpecificData string                   `xml:",omitempty"`
}

type TicketPrice struct {
	Total   *Total
	Details *DetailsTicketPrice
}

type DetailsTicketPrice struct {
	Detail []*DetailTicketPrice
}

type DetailTicketPrice struct {
	Application string
	Amount      int
}

type TravelerInfo struct {
	Surname string
	Given   string
	PTC     string `xml:",omitempty"`
}

type TravelerDoc struct {
	Surname string
	Given   string
	PTC     *PTC
}

type Payments struct {
	Payment []*Payment
}

type Payment struct {
	ID     string `xml:"ObjectKey,attr"`
	Type   *Type
	Other  *Other
	Amount *Total `xml:",omitempty"`
}

type Type struct {
	Code       string
	Definition string `xml:",omitempty"`
}

func MakeAirDocIssueRQ(AgentUserSender AgentUserSender, version string, query *QueryAirDocIssue) *AirDocIssueRQ {
	return &AirDocIssueRQ{
		Version:  version,
		Document: &Document{},
		Party: &Party{
			Sender: &Sender{
				AgentUserSender: &AgentUserSender,
			},
		},
		Query: query,
	}
}

