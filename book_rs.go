package s7_api_sdk

type OrderViewRS struct {
	Version  string `xml:"Version,attr"`
	Document *Document
	Response *Response
	Errors   *Errors
}

type BookingReferences struct {
	BookingReference []*BookingReference
}

type BookingReference struct {
	ObjectKey string `xml:",attr,omitempty"`
	Type      *Type  `xml:",omitempty"`
	ID        string
	AirlineID string
}
