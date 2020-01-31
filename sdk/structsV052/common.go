package structsV052

import (
	"encoding/xml"
)

type Envelope struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Header  *Header  `xml:",omitempty"`
	Body    *Body
}

type Header struct{}

type Body struct {
	AirShoppingRQ   *AirShoppingRQ   `xml:",omitempty"`
	AirShoppingRS   *AirShoppingRS   `xml:",omitempty"`
	OrderCreateRQ   *OrderCreateRQ   `xml:",omitempty"`
	OrderViewRS     *OrderViewRS     `xml:",omitempty"`
	OrderRetrieveRQ *OrderRetrieveRQ `xml:",omitempty"`
	ItinReshopRQ    *ItinReshopRQ    `xml:",omitempty"`
	ItinReshopRS    *ItinReshopRS    `xml:",omitempty"`
	AirDocIssueRQ   *AirDocIssueRQ   `xml:",omitempty"`
	AirDocDisplayRS *AirDocDisplayRS `xml:",omitempty"`
	OrderCancelRQ   *OrderCancelRQ   `xml:",omitempty"`
	OrderCancelRS   *OrderCancelRS   `xml:",omitempty"`
	AirDocVoidRQ    *AirDocVoidRQ    `xml:",omitempty"`
	AirDocVoidRS    *AirDocVoidRS    `xml:",omitempty"`
}

type ReshopResponse struct {
	Passengers           *PassengersRefs
	Order                *Order
	ReShopOffers         *ReShopOffers
	DataList             *DataList
	TicketDocInfos       *TicketDocInfos
	ItinReshopProcessing *ItinReshopProcessing
}

type OrderViewResponse struct {
	Passengers           *PassengerOrderRS `xml:"Passengers"`
	Order                *Order
	ReShopOffers         *ReShopOffers
	DataList             *DataList
	TicketDocInfos       *TicketDocInfos
	ItinReshopProcessing *ItinReshopProcessing
}

func (o *OrderViewResponse) GetPassengersIDs() []string {
	var passengerIDs = []string{}
	for _, passenger := range o.Passengers.Passenger {
		passengerIDs = append(passengerIDs, passenger.ObjectKey)
	}

	return passengerIDs
}

type PassengerOrderRS struct {
	Passenger []PassengerOrderViewRS
}

// getPaxByID func  return passenger object accordingly to its s7 ID
func getPaxByID(paxID string, passengers *PassengersRefs) *Passenger {

	if passengers == nil {
		return nil
	}

	for _, pax := range passengers.Passenger {

		if pax.ObjectKey == paxID {
			return &pax
		} // pax.ObjectKey: SH1 etc
	}

	return nil
}

func (r *ReshopResponse) GetPriceByPax(paxType string) *OfferPrice {

	for _, ReShopOffer := range r.ReShopOffers.ReShopOffer {
		for _, OfferPrice := range ReShopOffer.ReShopPricedOffer.OfferPrice {
			for _, Association := range OfferPrice.RequestedDate.Associations {

				if Association.AssociatedTraveler == nil {

					return nil
				}

				if getPaxByID(Association.AssociatedTraveler.TravelerReferences, r.Passengers) != nil && getPaxByID(Association.AssociatedTraveler.TravelerReferences, r.Passengers).PTC.Value == paxType {
					return OfferPrice
				}
			}
		}
	}

	return nil
}

type Passenger struct {
	ObjectKey string        `xml:"ObjectKey,attr"`
	PTC       PassengerPTC  `xml:"PTC"`
	Name      PassengerName `xml:"Name"`
}

type PassengersRefs struct {
	Passenger []Passenger
}

type PassengerPTC struct {
	Quantity string `xml:"Quantity,attr"`
	Value    string `xml:",chardata"` // ADT | CHD | INF
}

type PassengerName struct {
	Surname Surname
}

type Surname struct {
	Value string `xml:",chardata"`
}

type ItinReshopProcessing struct{}

type DataList struct {
	FlightSegmentList       *FlightSegmentList
	OriginDestinationList   *OriginDestinationList
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
	Earliest *Period
	Latest   *Period
}

type Period struct {
	ShortDate string `xml:",attr,omitempty"`
}

type Order struct {
	OrderID           *OfferID
	BookingReferences *BookingReferences
	TimeLimits        *TimeLimits
	OrderItems        *OrderItemsRS
}

type OfferID struct {
	Owner string `xml:",attr"`     // S7 | Any
	Value string `xml:",chardata"` //
}

type Success struct{}

type Document struct {
	Name string `xml:",omitempty"`
	//Type				string				`xml:",omitempty"`
	//ID					string				`xml:",omitempty"`
	//BirthCountry		string				`xml:",omitempty"`
	//DateOfIssue			string				`xml:",omitempty"`
	//DateOfExpiration	string				`xml:",omitempty"`
	//CountryOfResidence	string				`xml:",omitempty"`
}

type Party struct {
	Sender *Sender
}

type Sender struct {
	AgentUserSender *AgentUserSender
}

type Flight struct {
	SegmentKey       string          `xml:",omitempty"`
	Status           *Status         `xml:",omitempty"`
	Departure        *Point          `xml:",omitempty"`
	Arrival          *Point          `xml:",omitempty"`
	MarketingCarrier *Carrier        `xml:",omitempty"`
	OperatingCarrier *Carrier        `xml:",omitempty"`
	CabinType        *CabinType      `xml:",omitempty"`
	ClassOfService   *ClassOfService `xml:",omitempty"`
	// ID					string				`xml:"FlightKey,attr,omitempty"`
	// SegmentReferences	string				`xml:",omitempty"`
}

func (f *Flight) FillSegmentKey(segmentKey string) {
	f.SegmentKey = segmentKey
}

func (f *Flight) FillStatusCode(statusCode string) {
	f.Status.StatusCode.Code = statusCode
}

func (f *Flight) FillDeparture(time, date, airportCode string) {

	departure := &Point{
		AirportCode: airportCode,
		Date:        date,
		Time:        time,
	}

	f.Departure = departure
}

func (f *Flight) FillArrival(time, date, airportCode string) {

	arrival := &Point{
		AirportCode: airportCode,
		Date:        date,
		Time:        time,
	}

	f.Arrival = arrival

}

func (f *Flight) FillMarketingCarrier(airlineID, flightNumber string) {

	carrier := &Carrier{
		AirlineID:    airlineID,
		FlightNumber: flightNumber,
	}

	f.MarketingCarrier = carrier

}

func (f *Flight) FillOperatingCarrier(airlineID, flightNumber string) {
	carrier := &Carrier{
		AirlineID:    airlineID,
		FlightNumber: flightNumber,
	}

	f.OperatingCarrier = carrier
}

func (f *Flight) FillCabinType(cabinType string) {

	cabin := &CabinType{
		Code: &Code{
			Value: &cabinType,
		},
	}

	f.CabinType = cabin
}

func (f *Flight) FillClassOfService(classOfServiceCode string) {

	classOfService := &ClassOfService{
		Code: &Code{
			Value: &classOfServiceCode,
		},
		MarketingName: nil,
	}

	f.ClassOfService = classOfService
}

type Carrier struct {
	AirlineID    string `xml:",omitempty"`
	FlightNumber string `xml:",omitempty"`
}

type CalendarDates struct {
	XMLName    xml.Name `xml:"CalendarDates"`
	DaysBefore int      `xml:"DaysBefore,attr,omitempty"`
	DaysAfter  int      `xml:"DaysAfter,attr,omitempty"`
}

type CabinType struct {
	Code *Code
}

type Code struct {
	SeatsLeft *int    `xml:",attr,omitempty"`
	Value     *string `xml:",chardata"`
}

type Total struct {
	Code  string  `xml:",attr"`
	Value float64 `xml:",chardata"`
}

type DetailsPrice struct {
	Detail []*DetailPrice // array ?
}

type DetailPrice struct {
	SubTotal    *Total
	Application string
}

type Taxes struct {
	Total     *Total
	Breakdown *Breakdown `xml:",omitempty"`
}

type Breakdown struct {
	Tax []*Tax
}

type Tax struct {
	Amount  *Total
	TaxCode string
}

type Fees struct {
	Total *Total
}

type Surcharges struct {
	Surcharge *Surcharge
}

type Surcharge struct {
	Total *Total
}

type Descriptions struct {
	Description []*Description
}

type Description struct {
	Text string
}

type Metadata struct {
	Other *Other
}

type Other struct {
	OtherMetadata *OtherMetadata `xml:",omitempty"`
	Remarks       *Remarks       `xml:",omitempty"`
}

type OtherMetadata struct {
	CurrencyMetadatas *CurrencyMetadatas
	RuleMetadatas     *RuleMetadatas
}

type CurrencyMetadatas struct {
	CurrencyMetadata []*CurrencyMetadata
}

type CurrencyMetadata struct {
	MetadataKey string `xml:"MetadataKey,attr"`
	Application string
	Decimals    int
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

type OrderItemsRS struct {
	OrderItem []*OrderItem
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

// MakeQualifier constructs Qualifier element in S7 search request (3D agreement shopping scenario)
func MakeQualifier(companyIndex, account string) *Qualifier {

	specialFareQualifiers := &SpecialFareQualifiers{
		AirlineID:    "S7",
		CompanyIndex: companyIndex,
		Account:      account,
	}
	qualifier := &Qualifier{SpecialFareQualifiers: specialFareQualifiers}

	return qualifier
}

type PaxFareItem struct {
	TotalAmount *Total `json:"total_amount" bson:"totalAmount"`
	FareAmount  *Total `json:"fare_amount" bson:"fareAmount"`
	TaxAmount   *Total `json:"tax_amount" bson:"taxAmount"`
	FeeAmount   *Fees  `json:"fee_amount" bson:"feeAmount"`
	Taxes       []*Tax `json:"taxes,omitempty" bson:"taxes"`
}

type BookingPrice struct {
	TotalAmount *Total     `json:"total_amount" bson:"totalAmount"`
	FareAmount  *Total     `json:"fare_amount" bson:"fareAmount"`
	TaxAmount   *Tax       `json:"tax_amount" bson:"taxAmount"`
	FeeAmount   *Fees      `json:"fee_amount" bson:"feeAmount"`
	PaxFare     *PaxesFare `json:"pax_fare" bson:"paxFare"`
}

type PricePassenger struct {
	BaseAmount    Total    `json:"-"`
	TotalAmount   Total    `json:"total_amount"`
	FareAmount    Total    `json:"fare_amount"`
	TaxesAmount   Total    `json:"tax_amount"`
	FeeAmount     Fees     `json:"fee_amount"`
	Taxes         []Tax    `json:"taxes,omitempty"`
	Baggage       *Baggage `json:"-"`
	FareBasisCode string   `json:"-"`
	Brand         string   `json:"-"`
	UPT           string   `json:"-"`
}

type Baggage struct {
	Value int    `json:"value"`
	Unit  string `json:"unit"`
}

type Amount struct {
	Currency string `json:"currency" bson:"currency"`
	Amount   int64  `json:"amount" bson:"amount"`
}

type PaxesFare struct {
	ADT *PaxFareItem `json:"ADT" bson:"ADT"`
	CHD *PaxFareItem `json:"CHD,omitempty" bson:"CHD"`
	INF *PaxFareItem `json:"INF,omitempty" bson:"INF"`
}
