package structsV052

const ActionTypeContextDiscount = "discount"

type ItinReshopRQ struct {
	// XMLName    xml.Name `xml:"http://www.iata.org/IATA/EDIST ItinReshopRQ"`
	Xmlns      string      `xml:"xmlns,attr"`
	Version    string      `xml:"Version,attr"`
	Qualifiers *Qualifiers `xml:",omitempty"`
	Document   *Document
	Party      *Party
	Query      *QueryItinReshop
	OrderItems *OrderItems `xml:",omitempty"`
}

type QueryItinReshop struct {
	Passengers        *Passengers
	OrderItems        *OrderItems
	TicketDocQuantity int              `xml:",omitempty"`
	TicketDocInfo     []*TicketDocInfo `xml:",omitempty"`
	DataLists         *DataLists
	Reshop            *Reshop `xml:",omitempty"`
}

type Reshop struct {
	Actions *Actions
}

type Actions struct {
	ActionType        *ActionType
	BookingReferences *BookingReferences `xml:",omitempty"`
	OrderItems        *OrderItems
	Passengers        *PassengersAirDocIssueRQ
	Qualifiers        *Qualifiers `xml:",omitempty"`
}

type ActionType struct {
	Context string `xml:",attr,omitempty"`
	Value   string `xml:",chardata"`
}
