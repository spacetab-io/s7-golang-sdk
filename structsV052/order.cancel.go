package structsV052

import (
	"encoding/xml"
)

// OrderCancelRQ is a body of S7 bookingCancel request
type OrderCancelRQ struct {
	XMLName  xml.Name `xml:"http://www.iata.org/IATA/EDIST OrderCancelRQ"`
	Version  string   `xml:"Version,attr"`
	Document *Document
	Party    *Party
	Query    *QueryOrderCancel
}

// QueryOrderCancel is a Query section of S7 bookingCancel request
type QueryOrderCancel struct {
	BookingReferences *BookingReferences
}

// OrderCancelRS is a body of S7 bookingCancel response
type OrderCancelRS struct {
	Version  string `xml:"Version,attr"`
	Document *Document
	Success  *Success
	Response *ResponseOrderCancel
	Errors   *Errors
}

// ResponseOrderCancel is a body of S7 bookingCancel response
type ResponseOrderCancel struct {
	OrderCancelProcessing bool
	OrderReference        string
}

func MakeOrderCancelRQ(agentSender *AgentUserSender, pnr, airline string) (request *Envelope) {
	request = &Envelope{
		Header: new(Header),
		Body: &Body{
			OrderCancelRQ: &OrderCancelRQ{
				Version:  "1.0",
				Document: new(Document),
				Party: &Party{
					Sender: &Sender{
						AgentUserSender: agentSender,
					},
				},
				Query: &QueryOrderCancel{
					BookingReferences: &BookingReferences{
						BookingReference: []*BookingReference{
							&BookingReference{
								ID:        pnr,
								AirlineID: airline,
							},
						},
					},
				},
			},
		},
	}

	return
}
