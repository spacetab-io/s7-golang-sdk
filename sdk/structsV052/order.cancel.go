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
