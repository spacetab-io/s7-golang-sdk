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

func MakeOrderRetrieveRQ(
	id,
	airline,
	S7AgentName,
	S7POSType,
	S7RequestorType,
	S7Password,
	S7PseudoCity,
	S7AgentUserID,
	S7UserRole string,
) (request *Envelope) {

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
						AgentUserSender: MakeS7AgentUserSender(
							S7AgentName,
							S7POSType,
							S7RequestorType,
							S7Password,
							S7PseudoCity,
							S7AgentUserID,
							S7UserRole,
						),
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
