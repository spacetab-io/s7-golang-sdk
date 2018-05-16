package s7_api_sdk

import (
	"encoding/xml"
)

type OrderCreateRQ struct {
	XMLName  xml.Name `xml:"http://www.iata.org/IATA/EDIST OrderCreateRQ"`
	Version  string   `xml:"Version,attr"`
	Document *Document
	Party    *Party
	Query    *Query
}

type Query struct {
	Passengers        *Passengers
	OrderItems        *OrderItems
	TicketDocQuantity int              `xml:",omitempty"`
	TicketDocInfo     []*TicketDocInfo `xml:",omitempty"`
	DataLists         *DataLists
	Reshop            *Reshop `xml:",omitempty"`
}

type ShoppingResponse struct {
	Owner      string
	ResponseID string
	Offers     *Offers
}

type Offers struct {
	Offer []*Offer
}

type Offer struct {
	OfferID    *OfferID
	OfferItems *OfferItems
}

type OfferItems struct {
	OfferItem        []*OfferItem
	ApplicableFlight *ApplicableFlight
}

type OfferItem struct {
	OfferItemID      *OfferID          `xml:",omitempty"`
	Passengers       *Passengers       `xml:",omitempty"`
	OfferItemType    *OfferItemType    `xml:",omitempty"`
	ApplicableFlight *ApplicableFlight `xml:",omitempty"`
}

type InstructionsList struct {
	Instruction []*Instruction
}

type Instruction struct {
	ListKey                   string `xml:",attr,omitempty"`
	SpecialBookingInstruction *SpecialBookingInstruction
}

type SpecialBookingInstruction struct {
	Code       string `xml:",omitempty"`
	Definition string `xml:",omitempty"`
}

type OfferItemType struct {
	DetailedFlightItem *DetailedFlightItem
}

type DetailedFlightItem struct {
	OriginDestination []*OriginDestinationFlight
}

func MakeSessionBookRQ() (request *Envelope) {
	request = &Envelope{
		Header: new(Header),
		Body: &Body{
			OrderCreateRQ: &OrderCreateRQ{
				Document: &Document{
					Name: "1.0",
				},
				Party: &Party{
					Sender: new(Sender),
				},
				Query: new(Query),
			},
		},
	}

	return
}
