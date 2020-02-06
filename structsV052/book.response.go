package structsV052

type OrderViewRS struct {
	Version  string `xml:"Version,attr"`
	Document *Document
	Response *OrderViewResponse
	Errors   *Errors
}

type BookingReferences struct {
	BookingReference []*BookingReference `xml:",omitempty"`
}

type BookingReference struct {
	ObjectKey string `xml:",attr,omitempty"`
	Type      *Type  `xml:",omitempty"`
	ID        string
	AirlineID string
}
