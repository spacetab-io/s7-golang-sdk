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
