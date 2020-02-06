package structsV052

type OrderRetrieveRQ struct {
	XMLNS    string `xml:"xmlns,attr,omitempty"` // xmlns="http://www.iata.org/IATA/EDIST"
	Version  string `xml:"Version,attr"`
	Document *Document
	Party    *Party
	Query    *RetrieveQuery
}

type RetrieveQuery struct {
	Filters *Filters
}

type Filters struct {
	BookingReferences *BookingReferences
}

func MakeOrderRetrieveRQ(id, airline string, AgentUserSender AgentUserSender) (request *Envelope) {

	request = &Envelope{
		Header: new(Header),
		Body: &Body{
			OrderRetrieveRQ: &OrderRetrieveRQ{
				XMLNS: "http://www.iata.org/IATA/EDIST",
				Document: &Document{
					Name: "1.0",
				},
				Party: &Party{
					Sender: &Sender{
						AgentUserSender: &AgentUserSender,
					},
				},
				Query: &RetrieveQuery{
					Filters: &Filters{
						BookingReferences: &BookingReferences{
							BookingReference: []*BookingReference{
								{
									ID:        id,
									AirlineID: airline,
								},
							},
						},
					},
				},
			},
		},
	}

	return request
}
