package structsV052

import (
	"encoding/xml"
	"strings"
)

// AirDocVoidRQ is a body of S7 voidTicket request
type AirDocVoidRQ struct {
	XMLName  xml.Name `xml:"http://www.iata.org/IATA/EDIST AirDocVoidRQ"`
	Version  string   `xml:"Version,attr"`
	Document *Document
	Party    *Party
	Query    *QueryAirVoid
}

// QueryAirVoid is a Query section of S7 voidTicket request
type QueryAirVoid struct {
	TicketDocQuantity int
	TicketDocument    *TicketDocument
}

func MakeAirDocVoidRQ(AgentUserSender AgentUserSender, ticketNbr, ticketType string) (request *Envelope) {
	request = &Envelope{
		Header: new(Header),
		Body: &Body{
			AirDocVoidRQ: &AirDocVoidRQ{
				Version:  Version,
				Document: new(Document),
				Party: &Party{
					Sender: &Sender{
						&AgentUserSender,
					},
				},
				Query: &QueryAirVoid{
					TicketDocQuantity: 1,
					TicketDocument: &TicketDocument{
						TicketDocNbr: strings.TrimPrefix(ticketNbr, "421"),
						Type:         &Type{Code: ticketType},
					},
				},
			},
		},
	}

	return request
}
