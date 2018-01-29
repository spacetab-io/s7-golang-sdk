package sdk

import (
	"encoding/xml"
	"time"
)

type Envelope struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Header  *Header
	Body    *Body
}

type Header struct{}

type Body struct {
	AirShoppingRQ   *AirShoppingRQ   `xml:",omitempty"`
	AirShoppingRS   *AirShoppingRS   `xml:",omitempty"`
	OrderCreateRQ   *OrderCreateRQ   `xml:",omitempty"`
	OrderViewRS     *OrderViewRS     `xml:",omitempty"`
	ItinReshopRQ    *ItinReshopRQ    `xml:",omitempty"`
	ItinReshopRS    *ItinReshopRS    `xml:",omitempty"`
	AirDocIssueRQ   *AirDocIssueRQ   `xml:",omitempty"`
	AirDocDisplayRS *AirDocDisplayRS `xml:",omitempty"`
}

type Document struct {
	Name               string `xml:",omitempty"`
	Type               string `xml:",omitempty"`
	ID                 string `xml:",omitempty"`
	BirthCountry       string `xml:",omitempty"`
	DateOfIssue        string `xml:",omitempty"`
	DateOfExpiration   string `xml:",omitempty"`
	CountryOfResidence string `xml:",omitempty"`
}

type Party struct {
	Sender *Sender
}

type Sender struct {
	AgentUserSender *AgentUserSender
}

type Metadata struct {
	Other *Other
}

type Other struct {
	OtherMetadata *OtherMetadata `xml:",omitempty"`
	Remarks       *Remarks       `xml:",omitempty"`
}

type OtherMetadata struct {
	RuleMetadatas *RuleMetadatas
}

type RuleMetadatas struct {
	RuleMetadata []*RuleMetadata
}

type RuleMetadata struct {
	MetadataKey string `xml:",attr"`
	RuleID      string
	Status      string
	Remarks     *Remarks
}

type Code struct {
	SeatsLeft int32  `xml:",attr,omitempty"`
	Value     string `xml:",chardata"`
}

func MakeCode(seat int32, value string) *Code {
	return &Code{
		SeatsLeft: seat,
		Value:     value,
	}
}

type Errors struct {
	Error []*Error
}

type Error struct {
	Type      string `xml:",attr"`
	Tag       string `xml:",attr"`
	Status    string `xml:",attr"`
	Owner     string `xml:",attr"`
	ShortText string `xml:",attr"`
	Code      string `xml:",attr"`
	Value     string `xml:",chardata"`
}

type Remarks struct {
	Remark []string
}

type DataLists struct {
	AnonymousTravelerList   *AnonymousTravelerList   `xml:",omitempty"`
	CheckedBagAllowanceList *CheckedBagAllowanceList `xml:",omitempty"`
	FareList                *FareList                `xml:",omitempty"`
	FlightSegmentList       *FlightSegmentList       `xml:",omitempty"`
	FlightList              *FlightList              `xml:",omitempty"`
	OriginDestinationList   *OriginDestinationList   `xml:",omitempty"`
	PenaltyList             *PenaltyList             `xml:",omitempty"`
	PriceClassList          *PriceClassList          `xml:",omitempty"`
	InstructionsList        *InstructionsList        `xml:",omitempty"`
	ServiceList             *ServiceList             `xml:",omitempty"`
	TermsList               *TermsList               `xml:",omitempty"`
}

type OfferID struct {
	Owner string `xml:",attr"`
	Value string `xml:",chardata"`
}

func MakeOfferID(owner, value string) *OfferID {
	return &OfferID{
		Owner: owner,
		Value: value,
	}
}

type ApplicableFlight struct {
	OriginDestinationReferences string                    `xml:",omitempty"`
	FlightReferences            string                    `xml:",omitempty"`
	FlightSegmentReference      []*FlightSegmentReference `xml:",omitempty"`
}

type CalendarDates struct {
	XMLName    xml.Name `xml:"CalendarDates"`
	DaysBefore int      `xml:"DaysBefore,attr,omitempty"`
	DaysAfter  int      `xml:"DaysAfter,attr,omitempty"`
}

type Carrier struct {
	AirlineID    string `xml:",omitempty"`
	Name         string `xml:",omitempty"`
	FlightNumber string `xml:",omitempty"`
}

func MakeCarrier(airlineID, name, flightNumber string) *Carrier {
	return &Carrier{
		AirlineID:    airlineID,
		Name:         name,
		FlightNumber: flightNumber,
	}
}

type ClassOfService struct {
	Code          *Code
	MarketingName string `xml:",omitempty"`
}

func MakeClassOfService(seat int32, value, name string) *ClassOfService {
	return &ClassOfService{
		Code:          MakeCode(seat, value),
		MarketingName: name,
	}
}

type ServiceList struct {
	Service []*Service
}

type Service struct {
	ID           string        `xml:"ObjectKey,attr"`
	ServiceID    *OfferID      `xml:",omitempty"`
	Name         string        `xml:",omitempty"`
	TimeLimits   *TimeLimits   `xml:",omitempty"`
	Descriptions *Descriptions `xml:",omitempty"`
}

type TimeLimits struct {
	TicketingTimeLimits *TicketingTimeLimits `xml:",omitempty"`
	PaymentTimeLimit    *PaymentTimeLimit    `xml:",omitempty"`
}

type TicketingTimeLimits struct {
	Timestamp string `xml:",attr"`
}

type PaymentTimeLimit struct {
	DateTime string `xml:",attr"`
}

func (l *PaymentTimeLimit) GetDate() time.Time {
	date, _ := time.Parse(longFormT, l.DateTime)
	return date
}

type Equipment struct {
	AircraftCode     string
	AirlineEquipCode string
}

func MakeEquipment(aircraftCode, airlineEquipCode string) *Equipment {
	return &Equipment{
		AircraftCode:     aircraftCode,
		AirlineEquipCode: airlineEquipCode,
	}
}

type Descriptions struct {
	Description []*Description
}

type Description struct {
	Text string
}

type Associations struct {
	AssociatedTraveler     *AssociatedTraveler     `xml:",omitempty"`
	ApplicableFlight       *ApplicableFlight       `xml:",omitempty"`
	OfferDetailAssociation *OfferDetailAssociation `xml:",omitempty"`
	OtherAssociation       *OtherAssociations      `xml:",omitempty"`
	Passengers             *Passengers             `xml:",omitempty"`
}

type AssociatedTraveler struct {
	TravelerReferences string
}

type OtherAssociations struct {
	OtherAssociation []*OtherAssociation
}

type OtherAssociation struct {
	Type           string
	ReferenceValue string
}

type OfferDetailAssociation struct {
	OfferPenaltyReferences string
	OfferTermReferences    string
}

type Qualifiers struct {
	Qualifier []*Qualifier
}

type Qualifier struct {
	SpecialFareQualifiers *SpecialFareQualifiers
}

type SpecialFareQualifiers struct {
	AirlineID    string
	CompanyIndex string
	Account      string
}

func MakeQualifier(companyIndex, account string) *Qualifier {
	specialFareQualifiers := new(SpecialFareQualifiers)
	specialFareQualifiers.AirlineID = "S7"
	specialFareQualifiers.CompanyIndex = companyIndex
	specialFareQualifiers.Account = account

	qualifier := new(Qualifier)
	qualifier.SpecialFareQualifiers = specialFareQualifiers

	return qualifier
}

type OrderItems struct {
	DataLists        *DataLists
	ShoppingResponse *ShoppingResponse
	OrderItem        *OrderItem `xml:",omitempty"`
	OfferItem        *OfferItem `xml:",omitempty"`
}

type OrderItem struct {
	OrderItemID  *OfferID
	FlightItem   *FlightItem
	Associations *Associations
}

type FlightItem struct {
	OriginDestination []*OriginDestinationFlight
	FareDetail        *FareDetail
}

type Flight struct {
	ID                string   `xml:"FlightKey,attr,omitempty"`
	SegmentReferences string   `xml:",omitempty"`
	Departure         *Point   `xml:",omitempty"`
	Arrival           *Point   `xml:",omitempty"`
	MarketingCarrier  *Carrier `xml:",omitempty"`
}

type FareDetail struct {
	FareIndicatorCode   *FareIndicatorCode `xml:",omitempty"`
	PriceClassReference string             `xml:",omitempty"`
	FareComponent       []*FareComponent   `xml:",omitempty"`
	Remarks             *Remarks           `xml:",omitempty"`
}

type FareIndicatorCode struct {
	Code *Code
}

type FareComponent struct {
	Reference        string          `xml:"refs,attr,omitempty"`
	ID               string          `xml:"ObjectKey,attr,omitempty"`
	SegmentReference string          `xml:",omitempty"`
	Parameters       *ParametersRS   `xml:",omitempty"`
	PriceBreakdown   *PriceBreakdown `xml:",omitempty"`
	FareBasis        *FareBasis      `xml:",omitempty"`
	TicketDesig      *TicketDesig    `xml:",omitempty"`
	FareRules        *FareRules      `xml:",omitempty"`
}

type FareBasis struct {
	FareBasisCode *FareBasisCode `xml:",omitempty"`
	RBD           string         `xml:",omitempty"`
}

type FareBasisCode struct {
	Code        string
	Application string `xml:",omitempty"`
}

type TicketDesig struct {
	Application string `xml:",attr,omitempty"`
	Value       string `xml:",chardata"`
}

type FareRules struct {
	Ticketing *Ticketing
	Remarks   *Remarks
}

type Ticketing struct {
	Endorsements *Endorsements
}

type Endorsements struct {
	Endorsement []string
}

type Response struct {
	Passengers     *Passengers
	Order          *Order
	ReShopOffers   *ReShopOffers
	DataList       *DataList
	TicketDocInfos *TicketDocInfos
}

type DataList struct {
	CheckedBagAllowanceList *CheckedBagAllowanceList
	TermsList               *TermsList
}

type TermsList struct {
	ID   string `xml:"ListKey,attr,omitempty"`
	Term []*Term
}

type Term struct {
	Reference       string `xml:"refs,attr,omitempty"`
	ID              string `xml:"ObjectKey,attr,omitempty"`
	AvailablePeriod *AvailablePeriod
}

type AvailablePeriod struct {
	Earliest *Earliest
	Latest   *Latest
}

type Latest struct {
	ShortDate string `xml:",attr,omitempty"`
}

type Earliest struct{}

type Order struct {
	OrderID           *OfferID
	BookingReferences *BookingReferences
	TimeLimits        *TimeLimits
	OrderItems        *OrderItemsRS
}

type PricedOffer struct {
	OfferPrice []*OfferPrice
}

type OfferPrice struct {
	OfferItemID   string `xml:",attr"`
	RequestedDate *RequestedDate
	FareDetail    *FareDetail
}

type RequestedDate struct {
	PriceDetail  *PriceDetail
	Associations []*Associations
}

type PriceDetail struct {
	TotalAmount *TotalAmount
	BaseAmount  *BaseAmount
	FareFiledIn *FareFiledIn
	Discount    *Discount
	Taxes       *Taxes
}

type Discount struct {
	DiscountAmount  *Total
	DiscountPercent float64
}

type FareFiledIn struct {
	BaseAmount   *BaseAmount
	ExchangeRate float64 `xml:",omitempty"`
}

type BaseAmount struct {
	Code  string  `xml:",attr"`
	Value float64 `xml:",chardata"`
}

type TotalAmount struct {
	DetailCurrencyPrice *DetailCurrencyPrice
}
type Total struct {
	Code  string  `xml:",attr"`
	Value float64 `xml:",chardata"`
}

type Taxes struct {
	Total     *Total
	Breakdown *Breakdown
}

type Breakdown struct {
	Tax []*Tax
}

type Tax struct {
	Amount  *Total
	TaxCode string
}

type OrderItemsRS struct {
	OrderItem []*OrderItem
}
