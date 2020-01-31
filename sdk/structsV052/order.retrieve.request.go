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
