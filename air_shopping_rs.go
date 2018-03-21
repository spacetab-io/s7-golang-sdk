package sdk

import (
	"fmt"
	"time"
)

type AirShoppingRS struct {
	Version             string `xml:"Version,attr"`
	Document            *Document
	ShoppingResponseIDs *ShoppingResponseIDs
	OffersGroup         *OffersGroup
	DataLists           *DataLists
	Metadata            *Metadata
	Errors              *Errors
}

type OffersGroup struct {
	AirlineOffers *AirlineOffers
}

type AirlineOffers struct {
	TotalOfferQuantity int
	Owner              string
	AirlineOffer       []*AirlineOffer
}

type AirlineOffer struct {
	OfferID     *OfferID
	TotalPrice  *TotalPrice
	PricedOffer *PricedOffer
}

type TotalPrice struct {
	DetailCurrencyPrice *DetailCurrencyPrice
}

type DetailCurrencyPrice struct {
	Total   *Total
	Details *DetailsPrice
	Fees    *Fees
	Taxes   *Taxes
}

type Fees struct {
	Total *Total
}

type FlightSegmentReference struct {
	Reference            string `xml:"ref,attr"`
	ClassOfService       *ClassOfService
	BagDetailAssociation *BagDetailAssociation
}

type AnonymousTravelerList struct {
	AnonymousTraveler []*User
}

type CheckedBagAllowanceList struct {
	CheckedBagAllowance []*CheckedBagAllowance
}

type CheckedBagAllowance struct {
	Reference            string                `xml:"refs,attr,omitempty"`
	ID                   string                `xml:"ListKey,attr,omitempty"`
	WeightAllowance      *WeightAllowance      `xml:",omitempty"`
	AllowanceDescription *AllowanceDescription `xml:",omitempty"`
	PieceAllowance       *PieceAllowance       `xml:",omitempty"`
}

type WeightAllowance struct {
	MaximumWeight *ValueUOM
	Descriptions  *Descriptions
}

type AllowanceDescription struct {
	ApplicableParty string
	Descriptions    *Descriptions
}

type PieceAllowance struct {
	ApplicableParty   string
	TotalQuantity     int
	Descriptions      *Descriptions
	PieceMeasurements *PieceMeasurements
}

type PieceMeasurements struct {
	Quantity int `xml:",attr"`
}

type FareList struct {
	FareGroup []*FareGroup
}

type FareGroup struct {
	Refs          string `xml:"refs,attr,omitempty"`
	ID            string `xml:"ListKey,attr"`
	Fare          *Fare
	FareBasisCode *FareBasisCode
}

type Fare struct {
	FareCode   *FareCode
	FareDetail *FareDetail
}

type FareCode struct {
	Code string
}

type ParametersRS struct {
	Quantity int `xml:",attr"`
}

type PriceBreakdown struct {
	Price *Price
}

type Price struct {
	BaseAmount  *Total
	FareFiledIn *FareFiledIn
	Taxes       *Taxes
}

type FlightSegmentList struct {
	FlightSegment []*FlightSegment
}

type FlightSegment struct {
	SegmentKey       string `xml:",attr"`
	Departure        *Point
	Arrival          *Point
	MarketingCarrier *Carrier
	OperatingCarrier *Carrier
	Equipment        *Equipment
	ClassOfService   *ClassOfService
	FlightDetail     *FlightDetail
}

type FlightDetail struct {
	FlightDistance *ValueUOM
	FlightDuration *FlightDuration
	Stops          *Stops `xml:",omitempty"`
}

type Stops struct {
	StopQuantity  int
	StopLocations []*StopLocation
}

type StopLocation struct {
	AirportCode   string
	ArrivalDate   string
	ArrivalTime   string
	DepartureDate string
	DepartureTime string
	GroundTime    string
}

func (l *StopLocation) GetArrivalDate() time.Time {
	date, _ := time.Parse(longForm, fmt.Sprintf("%s %s", l.ArrivalDate, l.ArrivalTime))
	return date
}

func (l *StopLocation) GetDepartureDate() time.Time {
	date, _ := time.Parse(longForm, fmt.Sprintf("%s %s", l.DepartureDate, l.DepartureTime))
	return date
}

type ValueUOM struct {
	Value int
	UOM   string
}

type FlightDuration struct {
	Value string
}

type FlightList struct {
	Flight []*Flight
}

type OriginDestinationList struct {
	OriginDestination []*OriginDestinationRS
}

type OriginDestinationRS struct {
	ID               string `xml:"OriginDestinationKey,attr"`
	DepartureCode    string
	ArrivalCode      string
	FlightReferences string
}

type PenaltyList struct {
	Penalty []*Penalty
}

type Penalty struct {
	FareGroupReference string `xml:"refs,attr"`
	ID                 string `xml:"ObjectKey,attr"`
	Details            *Details
}

type Details struct {
	Detail []*Detail
}

type Detail struct {
	Type    string
	Amounts *Amounts
}

type Amounts struct {
	Amount []*Amount
}

type Amount struct {
	CurrencyAmountValue *CurrencyAmountValue
	AmountApplication   string
}

type CurrencyAmountValue struct {
	Code  string `xml:",attr"`
	Value string `xml:",chardata"`
}

type PriceClassList struct {
	PriceClass []*PriceClass
}

type PriceClass struct {
	ID             string `xml:"ObjectKey,attr"`
	Name           string
	FareBasisCode  *FareBasisCode
	ClassOfService *ClassOfService
}

type BagDetailAssociation struct {
	CheckedBagReferences string
}
