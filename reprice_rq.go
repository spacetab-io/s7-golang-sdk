package sdk

import (
	"encoding/xml"
)

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
}

func MakeItinReshopRQ() (request *Envelope) {
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
							ActionType: &ActionType{},
						},
					},
				},
			},
		},
	}

	return
}
