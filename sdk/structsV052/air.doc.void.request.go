package structsV052

import (
	"encoding/xml"
	"strings"

	"gitlab.teamc.io/tm-consulting/tmc24/avia/layer3/s7-agent-go/pkg/models"
)

// AirDocVoidRQ is a body of S7 voidTicket request
type AirDocVoidRQ struct {
	XMLName  xml.Name `xml:"http://www.iata.org/IATA/EDIST AirDocVoidRQ"`
	Version  string   `xml:"Version,attr"`
	Document *Document
	Party    *Party
	Query    *QueryAirVoid
}

// QueryAirVoid is a Query section of S7 voidTicket request
type QueryAirVoid struct {
	TicketDocQuantity int
	TicketDocument    *TicketDocument
}
