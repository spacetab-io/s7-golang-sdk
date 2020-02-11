package structsV052

import (
	"encoding/xml"
)

// OrderCreateRQ is a boby of book operation request
type OrderCreateRQ struct {
	XMLName  xml.Name `xml:"http://www.iata.org/IATA/EDIST OrderCreateRQ"`
	Version  string   `xml:"Version,attr"`
	Document *Document
	Party    *Party
	Query    *QueryOrderCreate
}

type QueryOrderCreate struct {
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
	OriginDestination []*OriginDestination
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
				Query: new(QueryOrderCreate),
			},
		},
	}

	return
}

type OrderItems struct {
	// DataLists        *DataLists
	ShoppingResponse *ShoppingResponse
	OrderItem        []*OrderItem `xml:",omitempty"`
	OfferItem        *OfferItem   `xml:",omitempty"`
}

type OrderItem struct {
	OrderItemID  *OfferID
	FlightItem   *FlightItem
	Associations *Associations
}

type FlightItem struct {
	Reference         string `xml:"refs,attr,omitempty"`
	OriginDestination []*OriginDestination
	// OriginDestination []*OriginDestinationFlight
	FareDetail *FareDetail
}

func MakeBookingRequestShoppingResponse(
	owner, version string,
	passengerReferences,
	flightReferences []string,
	agent AgentUserSender,
) *OrderCreateRQ {
	var passengerReferenceStr, flightReferencesStr string
	// Construct PassengerReference
	for _, passengerRef := range passengerReferences {
		if passengerReferenceStr == "" {
			passengerReferenceStr = passengerRef
		} else {
			passengerReferenceStr += " " + passengerRef
		}
	}
	// Construct FlightReferences string
	for _, flightRef := range flightReferences {
		if flightReferencesStr == "" {
			flightReferencesStr = flightRef
		} else {
			flightReferencesStr += " " + flightRef
		}
	}

	shoppingResponse := &ShoppingResponse{
		Owner:      owner,
		ResponseID: "UNKNOWN",
		Offers: &Offers{
			Offer: []*Offer{{
				OfferID: &OfferID{
					Owner: owner,
					Value: "UNKNOWN",
				},
				OfferItems: &OfferItems{
					OfferItem: []*OfferItem{{
						OfferItemID: &OfferID{
							Owner: "S7",
							Value: "UNKNOWN",
						},
						Passengers: &Passengers{
							PassengerReference: passengerReferenceStr,
						},
						ApplicableFlight: &ApplicableFlight{
							FlightReferences: flightReferencesStr,
						},
					},
					},
				},
			},
			},
		},
	}

	order := OrderCreateRQ{
		Version:  version, // "1.0" - v0.50, "2.0" - v0.52
		Document: &Document{},
		Party: &Party{
			Sender: &Sender{
				AgentUserSender: &agent,
			},
		},
		Query: &QueryOrderCreate{
			Passengers: nil,
			OrderItems: &OrderItems{
				ShoppingResponse: shoppingResponse,
				OrderItem:        nil,
				OfferItem:        nil,
			},
			TicketDocQuantity: 0,
			TicketDocInfo:     nil,
			DataLists:         nil,
			Reshop:            nil,
		},
	}

	return &order
}
