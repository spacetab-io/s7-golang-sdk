package s7_api_sdk

type AirDocDisplayRS struct {
	Version  string `xml:"Version,attr"`
	Document *Document
	Response *Response
	Errors   *Errors
}

type TicketDocInfos struct {
	TicketDocInfo []*TicketDocInfo `xml:",omitempty"`
}

type IssuingAirlineInfo struct {
	AirlineName string
}

type OriginDestinationDemand struct {
	Origin      string
	Destination string
}

type FareInfo struct {
	BaseFare                 *BaseFare
	Total                    *BaseFare
	Taxes                    *Taxes
	UnstructuredFareCalcInfo *UnstructuredFareCalcInfo
}

type BaseFare struct {
	Amount *Total
}

type UnstructuredFareCalcInfo struct {
	Info string
}

type TicketDocument struct {
	TicketDocNbr     string
	Type             *Type
	NumberofBooklets int           `xml:",omitempty"`
	DateOfIssue      string        `xml:",omitempty"`
	CouponInfo       []*CouponInfo `xml:",omitempty"`
}

type CouponInfo struct {
	CouponNumber    int
	FareBasisCode   *StatusCode
	CouponValid     *CouponValid
	Status          *StatusCode
	SoldAirlineInfo *SoldAirlineInfo
}

type CouponValid struct {
	EffectiveDatePeriod *EffectiveDatePeriod
}

type EffectiveDatePeriod struct {
	Expiration string
}

type SoldAirlineInfo struct {
	Departure                *Point
	Arrival                  *Point
	OperatingCarrier         *OperatingCarrier
	MarketingCarrier         *Carrier
	TicketedBaggageAllowance *TicketedBaggageAllowance
}

type OperatingCarrier struct {
	ResBookDesigCode string
}

type TicketedBaggageAllowance struct {
	AllowableBag *AllowableBag
}

type AllowableBag struct {
	Type   string `xml:",attr"`
	Number int    `xml:",attr"`
}

type Commission struct {
	Amount  float64
	Remarks *Remarks
}
