package structsV052

// AirShoppingRS is a body of S7 searchflights response
type AirShoppingRS struct {
	Version               string `xml:"Version,attr"`
	Document              *Document
	Success               *Success
	AirShoppingProcessing *AirShoppingProcessing
	OffersGroup           *OffersGroup
	DataLists             *DataLists
	Metadata              *Metadata
	Errors                *Errors `xml:",omitempty"`
}

// getPaxByID func  return passenger object accordingly to its s7 ID
func airShoppingGetPaxByID(paxID string, passengers AnonymousTravelerList) (passenger *Passenger) {
	for _, pax := range passengers.AnonymousTraveler {
		if pax.ObjectKey == paxID {
			return &pax

		}
	}

	return nil
}

type AnonymousTravelerList struct {
	AnonymousTraveler []Passenger `xml:"AnonymousTraveler"`
}

// AirShoppingProcessing is an empty section of S7 searchflights response
type AirShoppingProcessing struct{}

// OffersGroup section contains airline offers in S7 searchflights response
type OffersGroup struct {
	AirlineOffers []*AirlineOffers
}

// AirlineOffers section holds an array of AirlineOffer
type AirlineOffers struct {
	Owner        string // S7
	AirlineOffer []*AirlineOffer
}

// GetPaxPriceFareFromSearch func counts price FareAmount for pax of passed type
// FareAmount relating to Pax type, paxType value contains only in offerPrice objects
func (a *AirlineOffer) GetPaxPriceFareFromSearch(passengers AnonymousTravelerList, paxType string) float64 {

	// range over OfferPrices (2 for RT, 1 for OW)
	// every offerPrice contains fareComponents that represent passengers price (one fareComponent for one passenger)
	var fareAmount float64
	for _, OfferPrice := range a.PricedOffer.OfferPrice {
		for _, fareComponent := range OfferPrice.FareDetail.FareComponent {

			if airShoppingGetPaxByID(fareComponent.Refs, passengers).PTC.Value == paxType {
				fareAmount += fareComponent.PriceBreakdown.Price.BaseAmount.Value
			}
		}
	}

	return fareAmount
}

func (a *AirlineOffer) GetTaxFromPrice(passengers AnonymousTravelerList, paxType string) *float64 {

	// range over OfferPrices (2 for RT, 1 for OW)
	// every offerPrice contains fareComponents that represent passengers price (one fareComponent for one passenger)
	var taxAmount *float64
	for _, OfferPrice := range a.PricedOffer.OfferPrice {
		for _, fareComponent := range OfferPrice.FareDetail.FareComponent {

			if airShoppingGetPaxByID(fareComponent.Refs, passengers) == nil {

				return nil
			}

			if airShoppingGetPaxByID(fareComponent.Refs, passengers).PTC.Value == paxType {

				//allocate a new zero-valued taxAmount
				taxAmount = new(float64)

				*taxAmount += fareComponent.PriceBreakdown.Price.Taxes.Total.Value
			}
		}
	}

	return taxAmount
}

// TotalAmountFromPrice func counts price amount for pax of passed type
// price amount relating to Pax type, contains only in offerPrice objects
func (a *AirlineOffer) TotalAmountFromPrice(passengers AnonymousTravelerList, paxType string) *float64 {
	// range over OfferPrices (2 for RT, 1 for OW)
	// every offerPrice contains fareComponents that represent passengers price (one fareComponent for one passenger)
	var totalAmount *float64
	for _, OfferPrice := range a.PricedOffer.OfferPrice {
		for _, fareComponent := range OfferPrice.FareDetail.FareComponent {
			if airShoppingGetPaxByID(fareComponent.Refs, passengers) == nil {

				return nil
			}
			if airShoppingGetPaxByID(fareComponent.Refs, passengers).PTC.Value == paxType {

				//allocate a new zero-valued taxAmount
				totalAmount = new(float64)

				// Tax total adds to BaseAmount total
				*totalAmount += fareComponent.PriceBreakdown.Price.Taxes.Total.Value + fareComponent.PriceBreakdown.Price.BaseAmount.Value
			}
		}

	}

	return totalAmount
}

// TotalTaxFromAirlineOffer func return total price of AirlineOffer
func (a *AirlineOffer) TotalPriceFromAirlineOffer() *Total {

	return a.TotalPrice.DetailCurrencyPrice.Total
}

// TotalTaxFromAirlineOffer func return taxes amount of AirlineOffer
func (a *AirlineOffer) TotalTaxFromAirlineOffer() *Taxes {

	return a.TotalPrice.DetailCurrencyPrice.Taxes
}

// TotalFareFromAirlineOffer func return fare amount of AirlineOffer
func (a *AirlineOffer) TotalFareFromAirlineOffer() *DetailPrice {
	for _, detail := range a.TotalPrice.DetailCurrencyPrice.Details.Detail {
		if detail.Application == "Fare" {

			return detail
		}
	}

	return nil
}

func (a *AirlineOffer) CurrencyFromPrice() *string {
	for _, OfferPrice := range a.PricedOffer.OfferPrice {

		return &OfferPrice.RequestedDate.PriceDetail.BaseAmount.Code
	}

	return nil
}

// AirlineOffer is an elenemt of AirlineOffer array
type AirlineOffer struct {
	OfferID     *OfferID     // S7, OF1
	TotalPrice  *TotalAmount // Application = Fare
	Disclosure  *Disclosure  // Text = BASIC ECONOMY
	PricedOffer *PricedOffer
}

// TotalPrice section in AirlineOffer
type TotalPrice struct {
	DetailCurrencyPrice *DetailCurrencyPrice
}

// TotalAmount section in AirlineOffer
type TotalAmount struct {
	DetailCurrencyPrice *DetailCurrencyPrice
}

// DetailCurrencyPrice element of AirlineOffer
type DetailCurrencyPrice struct {
	Total   *Total
	Details *DetailsPrice `xml:",omitempty"`
	Fees    *Fees         `xml:",omitempty"`
	Taxes   *Taxes        `xml:",omitempty"`
}

// Disclosure section in AirlineOffer
type Disclosure struct {
	Description *Description // basicEconomy
}

// PricedOffer section in AirlineOffer holds OfferPrice array
type PricedOffer struct {
	OfferPrice []*OfferPrice
}

// OfferPrice section contains price information for passenger
type OfferPrice struct {
	OfferItemID   string `xml:",attr"` // OF1SL1
	RequestedDate *RequestedDate
	FareDetail    *FareDetail
}

type PaxSegmentsAssociations struct {
	SegmentPax map[string]*SegmentPax
}

func (o *OfferPrice) GetAllSegments() {

}

type SegmentPax struct {
	ADT *FlightSegmentReference
	CHD *FlightSegmentReference
	INF *FlightSegmentReference
}

// getSegmentsForPax func attaches segments information to pax type
func getSegmentsForPax(association *Associations, segments *PaxSegmentsAssociations) {

	// ApplicableFlight is a segment
	for index, flightSegmentReference := range association.ApplicableFlight.FlightSegmentReference {

		// As ApplicableFlight contains 2 objects of FlightSegmentReference, with a same key so combine them into singe one
		// One of FlightSegmentReference object contains info about cabin, other about bag and marketing code
		if index != 0 && association.ApplicableFlight.FlightSegmentReference[index-1] != nil && flightSegmentReference.Reference == association.ApplicableFlight.FlightSegmentReference[index-1].Reference {

			segments.SegmentPax[flightSegmentReference.Reference] = &SegmentPax{
				ADT: &FlightSegmentReference{},
				CHD: &FlightSegmentReference{},
				INF: &FlightSegmentReference{},
			}

			segments.SegmentPax[flightSegmentReference.Reference].ADT = &FlightSegmentReference{
				Reference:            flightSegmentReference.Reference,
				Cabin:                association.ApplicableFlight.FlightSegmentReference[index-1].Cabin,
				ClassOfService:       flightSegmentReference.ClassOfService,
				BagDetailAssociation: flightSegmentReference.BagDetailAssociation,
			}

			segments.SegmentPax[flightSegmentReference.Reference].CHD = &FlightSegmentReference{
				Reference:            flightSegmentReference.Reference,
				Cabin:                association.ApplicableFlight.FlightSegmentReference[index-1].Cabin,
				ClassOfService:       flightSegmentReference.ClassOfService,
				BagDetailAssociation: flightSegmentReference.BagDetailAssociation,
			}

			segments.SegmentPax[flightSegmentReference.Reference].INF = &FlightSegmentReference{
				Reference:            flightSegmentReference.Reference,
				Cabin:                association.ApplicableFlight.FlightSegmentReference[index-1].Cabin,
				ClassOfService:       flightSegmentReference.ClassOfService,
				BagDetailAssociation: flightSegmentReference.BagDetailAssociation,
			}

		}
	}
}

// GetApplicableFlights func return map of Paxes with attached segments information
func (o *OfferPrice) GetApplicableFlights(passengers AnonymousTravelerList) *PaxSegmentsAssociations {
	var segments = PaxSegmentsAssociations{SegmentPax: map[string]*SegmentPax{}}

	for _, association := range o.RequestedDate.Associations {

		// there can be associations without ApplicableFlight object || AssociatedTraveler
		if association.AssociatedTraveler == nil || association.ApplicableFlight == nil {

			continue
		}

		getSegmentsForPax(association, &segments)
	}

	return &segments
}

// GetOfferPriceSegmentID func return if of flight segment
// FlightSegmentReferences contains same ids of segment
func (o *OfferPrice) GetOfferPriceSegmentID() *string {
	for _, association := range o.RequestedDate.Associations {

		// there can be associations without ApplicableFlight object
		if association.ApplicableFlight == nil {

			continue
		}

		return &association.ApplicableFlight.FlightSegmentReference[0].Reference
	}

	return nil
}

// GetOfferPriceSegmentID func return if of flight segment
// FlightSegmentReferences contains same ids of segment
func (o *OfferPrice) GetOfferPriceSegmentIDs() []string {

	var segmentsIDs []string
	for _, association := range o.RequestedDate.Associations {

		// there can be associations without ApplicableFlight object
		if association.ApplicableFlight == nil {

			continue
		}

		for _, ApplicableFlight := range association.ApplicableFlight.FlightSegmentReference {
			segmentsIDs = append(segmentsIDs, ApplicableFlight.Reference)
		}

		// Delete duplicate segmentsIDs
		for index, segmentID := range segmentsIDs {
			for internalIndex, internalSegmentID := range segmentsIDs {

				if internalIndex != index && segmentID == internalSegmentID {
					segmentsIDs[index] = segmentsIDs[len(segmentsIDs)-1]
					segmentsIDs[len(segmentsIDs)-1] = ""
					segmentsIDs = segmentsIDs[:len(segmentsIDs)-1]
				}
			}
		}

		return segmentsIDs
	}

	return nil
}

// GetAssociationPaxInformation func returns associations object for passed pax type
// associations objects is contained in offer price objects
func (o *OfferPrice) GetAssociationPaxInformation(paxType string, passengers AnonymousTravelerList) *Associations {
	for _, association := range o.RequestedDate.Associations {
		if paxType == airShoppingGetPaxByID(association.AssociatedTraveler.TravelerReferences, passengers).PTC.Value {
			return association
		}
	}

	return nil
}

// GetAssociationPaxInformation func returns cabin type for passed pax type
func (o *OfferPrice) GetCabinTypePax(paxType string, passengers AnonymousTravelerList) *CabinFlightSegment {
	for _, association := range o.RequestedDate.Associations {
		if association.AssociatedTraveler == nil {

			//One of flightSegmentReference objects does not contain information about AssociatedTraveler
			continue
		}

		if paxType == airShoppingGetPaxByID(association.AssociatedTraveler.TravelerReferences, passengers).PTC.Value {
			for _, flightSegmentReference := range association.ApplicableFlight.FlightSegmentReference {

				// One of flightSegmentReference objects does not contain information about Cabin
				if flightSegmentReference.Cabin == nil {
					continue
				}

				return flightSegmentReference.Cabin
			}
		}
	}

	return nil
}

// GetAssociationPaxInformation func returns class of service for passed pax type
func (o *OfferPrice) GetClassTypePax(paxType string, passengers AnonymousTravelerList) *ClassOfService {
	for _, association := range o.RequestedDate.Associations {
		if association.AssociatedTraveler == nil || association.ApplicableFlight == nil {

			//One of flightSegmentReference objects does not contain information about AssociatedTraveler || ApplicableFlight
			continue
		}

		if paxType == airShoppingGetPaxByID(association.AssociatedTraveler.TravelerReferences, passengers).PTC.Value {
			for _, flightSegmentReference := range association.ApplicableFlight.FlightSegmentReference {

				// One of flightSegmentReference objects does not contain information about class of service
				if flightSegmentReference.ClassOfService == nil {

					continue
				}

				return flightSegmentReference.ClassOfService
			}
		}
	}

	return nil
}

func (o *OfferPrice) GetBaggageInformation(checkedBagAllowanceList CheckedBagAllowanceList) {

}

func (o *OfferPrice) GetSegmentAdditionalInfo(list FlightSegmentList) *FlightSegment {
	segID := o.GetOfferPriceSegmentID()

	for _, segment := range list.FlightSegment {

		if segment.SegmentKey == *segID {
			return segment
		}
	}

	return nil
}

func (o *OfferPrice) GetSegmentsAdditionalInfo(list FlightSegmentList) []FlightSegment {
	//segID := o.GetOfferPriceSegmentID()

	flightSegments := make([]FlightSegment, 0)

	segmentsIDs := o.GetOfferPriceSegmentIDs()

	for _, segment := range list.FlightSegment {
		for _, flightSegmentID := range segmentsIDs {

			if segment.SegmentKey == flightSegmentID {
				flightSegments = append(flightSegments, *segment)
			}
		}
	}

	return flightSegments
}

// RequestedDate section in OfferPrice element
type RequestedDate struct {
	PriceDetail  *PriceDetail
	Associations []*Associations
}

// PriceDetail section in RequestedDate holds price details
type PriceDetail struct {
	TotalAmount *TotalAmount
	BaseAmount  *Total
	FareFiledIn *FareFiledIn
	Taxes       *Taxes
	Fees        *Fees
}

// FareFiledIn element of Price section
type FareFiledIn struct {
	BaseAmount   *BaseAmount
	ExchangeRate float64 `xml:",omitempty"`
}

// BaseAmount price element
type BaseAmount struct {
	Code  string  `xml:",attr"`
	Value float64 `xml:",chardata"`
}

type Associations struct {
	AssociatedTraveler *AssociatedTraveler `xml:",omitempty"`
	ApplicableFlight   *ApplicableFlight   `xml:",omitempty"`
	OtherAssociation   *OtherAssociations  `xml:",omitempty"`
	Passengers         *Passengers         `xml:",omitempty"`
}

type AssociatedTraveler struct {
	TravelerReferences string // SH1
}

type ApplicableFlight struct {
	OriginDestinationReferences string                    `xml:",omitempty"` // OD1
	FlightReferences            string                    `xml:",omitempty"`
	FlightSegmentReference      []*FlightSegmentReference `xml:",omitempty"`
}

type FlightSegmentReference struct {
	Reference            string              `xml:"ref,attr"` // SEG1
	Cabin                *CabinFlightSegment `xml:",omitempty"`
	ClassOfService       *ClassOfService
	BagDetailAssociation *BagDetailAssociation
}

type CabinFlightSegment struct {
	CabinDesignator *string // Y
	MarketingName   *string // BASIC ECONOMY
}

type BagDetailAssociation struct {
	CheckedBagReferences string // BG1
	CarryOnReferences    string // CO1
}

type OtherAssociations struct {
	OtherAssociation []*OtherAssociation
}

type OtherAssociation struct {
	Type           string // SEG1 OVB-KUF | SBSRT
	ReferenceValue string // Meal_S | PL1
}

type FareDetail struct {
	FareComponent []*FareComponent

	PriceClassReference string   `xml:",omitempty"`
	Remarks             *Remarks `xml:",omitempty"`
}

type FareComponent struct {
	ID             string                   `xml:"ObjectKey,attr,omitempty"` // FC399
	Refs           string                   `xml:"refs,attr,omitempty"`      // SH1
	Parameters     *ParametersFareComponent `xml:",omitempty"`
	PriceBreakdown *PriceBreakdown          `xml:",omitempty"`
	FareBasis      *FareBasis               `xml:",omitempty"`
	TicketDesig    *TicketDesig             `xml:",omitempty"`
}

type ParametersFareComponent struct {
	Quantity int `xml:",attr"` // 1
}

type PriceBreakdown struct {
	Price *Price
}

type Price struct {
	BaseAmount  *Total
	FareFiledIn *FareFiledIn `xml:",omitempty"`
	Taxes       *Taxes
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
	Value       int    `xml:",chardata"`
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

type DataLists struct {
	AnonymousTravelerList   *AnonymousTravelerList   `xml:",omitempty"`
	CarryOnAllowanceList    *CarryOnAllowanceList    `xml:",omitempty"`
	CheckedBagAllowanceList *CheckedBagAllowanceList `xml:",omitempty"`
	FareList                *FareList                `xml:",omitempty"`
	FlightSegmentList       *FlightSegmentList       `xml:",omitempty"`
	OriginDestinationList   *OriginDestinationList   `xml:",omitempty"`
	PenaltyList             *PenaltyList             `xml:",omitempty"`
	InstructionsList        *InstructionsList        `xml:",omitempty"`
	ServiceList             *ServiceList             `xml:",omitempty"`
}

type CarryOnAllowanceList struct {
	CarryOnAllowance []*CarryOnAllowance
}

type CarryOnAllowance struct {
	ID                   string `xml:"ListKey,attr"` // CO1
	AllowanceDescription *AllowanceDescription
}

type AllowanceDescription struct {
	ApplicableParty string        // Traveler
	ApplicableBag   string        // 1PC
	Descriptions    *Descriptions // Text = 10KG
}

type CheckedBagAllowanceList struct {
	CheckedBagAllowance []*CheckedBagAllowance
}

type CheckedBagAllowance struct {
	ID                   string                `xml:"ListKey,attr"`
	AllowanceDescription *AllowanceDescription `xml:",omitempty"`
}

type FareList struct {
	FareGroup []*FareGroup
}

type FareGroup struct {
	ID            string `xml:"ListKey,attr"`
	Refs          string `xml:"refs,attr,omitempty"`
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

type FlightSegmentList struct {
	FlightSegment []*FlightSegment
}

type FlightSegment struct {
	SegmentKey       string `xml:",attr"` // SEG1
	Departure        *Point
	Arrival          *Point
	MarketingCarrier *Carrier
	OperatingCarrier *Carrier
	Equipment        *Equipment
	FlightDetail     *FlightDetail
}

type FlightDetail struct {
	FlightDistance *ValueUOM
	FlightDuration *FlightDuration
	Stops          *Stops `xml:",omitempty"`
}

type ValueUOM struct {
	Value int    // 1303
	UOM   string // Miles
}

type FlightDuration struct {
	Value string // PT3H30M
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

type OriginDestinationList struct {
	OriginDestination []*OriginDestinationRS
}

type OriginDestinationRS struct {
	ID               string `xml:"OriginDestinationKey,attr"` // OD1
	DepartureCode    string
	ArrivalCode      string
	FlightReferences string // SEG1 SEG2
}

type PenaltyList struct {
	Penalty []*Penalty
}

type Penalty struct {
	ID      string `xml:"ObjectKey,attr"` // PL1
	Details *PenaltyDetails
}

type PenaltyDetails struct {
	Detail []*PenaltyDetail
}

type PenaltyDetail struct {
	Type    string // BEFORE-DEPARTURE | NO-SHOW
	Amounts *PenaltyAmounts
}

type PenaltyAmounts struct {
	Amount []*PenaltyAmount
}

type PenaltyAmount struct {
	CurrencyAmountValue *PenaltyAmountValue
	AmountApplication   string // MinimumPenaltyAmount | MaximumPenaltyAmount
}

type PenaltyAmountValue struct {
	Code  string  `xml:",attr"`     // RUB
	Value float32 `xml:",chardata"` // 5000
}

type ServiceList struct {
	Service []*Service
}

type Service struct {
	ID           string        `xml:"ObjectKey,attr"` // Meal_S | Meal_L
	ServiceID    *OfferID      `xml:",omitempty"`     // Any, S | Any, L
	Name         string        `xml:",omitempty"`     // Meal
	Descriptions *Descriptions `xml:",omitempty"`
}

const (
	ResultTypeSmartchoice = "smartchoice"
	ResultTypeLowfare     = "lowfare"
	ResultTypeFlightinfo  = "flightinfo"
)
