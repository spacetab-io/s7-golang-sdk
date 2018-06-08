package s7_api_sdk

import "encoding/xml"

type AirDocVoidRQ struct {
	XMLName  xml.Name `xml:"http://www.iata.org/IATA/EDIST AirDocVoidRQ"`
	Version  string   `xml:"Version,attr"`
	Document *Document
	Party    *Party
	Query    *QueryAirVoid
}

type QueryAirVoid struct {
	TicketDocQuantity int
	TicketDocument    *TicketDocument
}

type AirDocVoidRS struct {
	Version  string `xml:"Version,attr"`
	Document *Document
	Success  bool
	Response *ResponseAirDocVoid
	Errors   *Errors
}

type ResponseAirDocVoid struct {
	TicketDocument *TicketDocument
}

func MakeAirDocVoidRQ(agentSender *AgentUserSender, ticketNbr, ticketType string) (request *Envelope) {
	request = &Envelope{
		Header: new(Header),
		Body: &Body{
			AirDocVoidRQ: &AirDocVoidRQ{
				Version:  "1.0",
				Document: new(Document),
				Party: &Party{
					Sender: &Sender{
						AgentUserSender: agentSender,
					},
				},
				Query: &QueryAirVoid{
					TicketDocQuantity: 1,
					TicketDocument: &TicketDocument{
						TicketDocNbr: ticketNbr,
						Type:         &Type{Code: ticketType},
					},
				},
			},
		},
	}

	return
}
