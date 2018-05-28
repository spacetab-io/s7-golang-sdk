package s7_api_sdk

import (
	"encoding/xml"
)

const ActionTypeContextDiscount = "discount"

type ItinReshopRQ struct {
	XMLName  xml.Name `xml:"http://www.iata.org/IATA/EDIST ItinReshopRQ"`
	Version  string   `xml:"Version,attr"`
	Document *Document
	Party    *Party
	Query    *Query
}

type Reshop struct {
	Actions *Actions
}

type Actions struct {
	ActionType *ActionType
	OrderItems *OrderItems
	Passengers *Passengers
	Qualifiers *Qualifiers `xml:",omitempty"`
}

type ActionType struct {
	Context string `xml:",attr,omitempty"`
	Value   int    `xml:",chardata"`
}

func MakeItinReshopRQ(actionType *ActionType) (request *Envelope) {
	request = &Envelope{
		Header: new(Header),
		Body: &Body{
			ItinReshopRQ: &ItinReshopRQ{
				Version:  "1.0",
				Document: new(Document),
				Party: &Party{
					Sender: new(Sender),
				},
				Query: &Query{
					Reshop: &Reshop{
						Actions: &Actions{
							ActionType: actionType,
						},
					},
				},
			},
		},
	}

	return
}
