package structsV052

// AirDocVoidRS is a body of S7 voidTicket response
type AirDocVoidRS struct {
	Version  string `xml:"Version,attr"`
	Document *Document
	Success  bool
	Response *ResponseAirDocVoid
	Errors   *Errors
}

// ResponseAirDocVoid is a body of S7 voidTicket response
type ResponseAirDocVoid struct {
	TicketDocument *TicketDocument
}
