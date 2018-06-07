package s7_api_sdk

import "encoding/xml"

type OrderCancelRQ struct {
	XMLName  xml.Name `xml:"http://www.iata.org/IATA/EDIST OrderCancelRQ"`
	Version  string   `xml:"Version,attr"`
	Document *Document
	Party    *Party
	Query    *QueryOrderCancel
}

type QueryOrderCancel struct {
	BookingReferences *BookingReferences
}

type OrderCancelRS struct {
	Version  string `xml:"Version,attr"`
	Document *Document
	Success  bool
	Response *ResponseOrderCancel
	Errors   *Errors
}

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
