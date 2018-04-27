package sdk

import (
	"encoding/xml"
)

type AirShoppingRQ struct {
	XMLName     xml.Name `xml:"http://www.iata.org/IATA/EDIST AirShoppingRQ"`
	Version     string   `xml:"Version,attr"`
	Document    *Document
	Party       *Party
	Parameters  *Parameters `xml:",omitempty"`
	Travelers   *Travelers
	CoreQuery   *CoreQuery
	Qualifiers  *Qualifiers  `xml:",omitempty"`
	Preferences *Preferences `xml:",omitempty"`
	Metadata    *Metadata    `xml:",omitempty"`
}

type Travelers struct {
	Traveler []*Traveler
}

type Traveler struct {
	AnonymousTraveler *User `xml:",omitempty"`
}

func MakeTravelers(trs map[string]int32) *Travelers {
	travelers := make([]*Traveler, 0)
	var (
		qtyAdultChild int
		qtyInfant     int
		qtyAdult      int
	)

	for code, qty := range trs {
		ptc := MakePTC(int(qty), code, &qtyAdultChild, &qtyInfant, &qtyAdult)
		if ptc != nil {
			travelers = append(travelers, &Traveler{&User{
				PTC: ptc,
			}})
		}
	}
	return &Travelers{travelers}
}

type CoreQuery struct {
	OriginDestinations  *OriginDestinations  `xml:",omitempty"`
	ShoppingResponseIDs *ShoppingResponseIDs `xml:",omitempty"`
}

type OriginDestinations struct {
	OriginDestination []*OriginDestination
}

type Parameters struct {
	ServiceFilters *ServiceFilters `xml:",omitempty"`
	CurrCodes      *CurrCodes      `xml:",omitempty"`
}

type CurrCodes struct {
	CurrCode []string
}

type ServiceFilters struct {
	ServiceFilter []*ServiceFilter `xml:",omitempty"`
}

type ServiceFilter struct {
	GroupCode    string
	SubGroupCode string `xml:",omitempty"`
}

const (
	FILTER_GROUP_CIRCUITY_LIMIT = "circuityLimit"
	FILTER_GROUP_DURATION_LIMIT = "durationLimit"
	FILTER_GROUP_SEARCH_METRIC  = "searchMetric"
	FILTER_GROUP_SORT           = "sort"
)

const (
	FILTER_SUBGROUP_QUICKEST        = "QUICKEST"
	FILTER_SUBGROUP_MOST_CONVENIENT = "MOST-CONVENIENT"
	FILTER_SUBGROUP_CHEAPEST        = "CHEAPEST"
	FILTER_SUBGROUP_LEISURE         = "LEISURE"
	FILTER_SUBGROUP_BUSINESS        = "BUSINESS"
	FILTER_SUBGROUP_LUXURY          = "LUXURY"
	FILTER_SUBGROUP_MIN_PRICE       = "MIN_PRICE"
)

func MakeServiceFilters(params map[string]string) *ServiceFilters {
	serviceFilters := new(ServiceFilters)
	for group, subGroup := range params {
		filter := MakeServiceFilter(group, subGroup)
		if filter != nil {
			serviceFilters.ServiceFilter = append(serviceFilters.ServiceFilter, filter)
		}
	}
	return serviceFilters
}

func MakeServiceFilter(group, subGroup string) *ServiceFilter {
	switch {
	case group != FILTER_GROUP_CIRCUITY_LIMIT &&
		group != FILTER_GROUP_DURATION_LIMIT &&
		group != FILTER_GROUP_SEARCH_METRIC &&
		group != FILTER_GROUP_SORT:
		return nil
	case group == FILTER_GROUP_SEARCH_METRIC &&
		subGroup != FILTER_SUBGROUP_QUICKEST &&
		subGroup != FILTER_SUBGROUP_MOST_CONVENIENT &&
		subGroup != FILTER_SUBGROUP_CHEAPEST &&
		subGroup != FILTER_SUBGROUP_LEISURE &&
		subGroup != FILTER_SUBGROUP_BUSINESS &&
		subGroup != FILTER_SUBGROUP_LUXURY:
		return nil
	case group == FILTER_GROUP_SORT &&
		subGroup != FILTER_SUBGROUP_MIN_PRICE:
	}
	serviceFilter := new(ServiceFilter)
	serviceFilter.GroupCode = group
	serviceFilter.SubGroupCode = subGroup
	return serviceFilter
}

type Preferences struct {
	Preference []*Preference
}

type Preference struct {
	FlightPreferences *FlightPreferences
}

type FlightPreferences struct {
	Aircraft       *Aircraft
	Characteristic *Characteristic
}

type Aircraft struct {
	Cabins *Cabins
}

type Cabins struct {
	Cabin *Cabin
}

type Cabin struct {
	Code string
}

type Characteristic struct {
	DirectPreferences string
}

func MakePreference(businessCabinOnly, directPreferences string) *Preference {
	preference := &Preference{
		&FlightPreferences{},
	}

	if businessCabinOnly != "" {
		preference.FlightPreferences.Aircraft = &Aircraft{
			&Cabins{
				&Cabin{
					businessCabinOnly,
				},
			},
		}
	}
	if directPreferences != "" {
		preference.FlightPreferences.Characteristic = &Characteristic{directPreferences}
	}

	return preference
}

const (
	RESULT_TYPE_SMARTCHOICE = "smartchoice"
	RESULT_TYPE_LOWFARE     = "lowfare"
	RESULT_TYPE_FLIGHTINFO  = "flightinfo"
)

func MakeResultType(status string) *Metadata {
	if status != RESULT_TYPE_SMARTCHOICE &&
		status != RESULT_TYPE_LOWFARE &&
		status != RESULT_TYPE_FLIGHTINFO {
		return nil
	}

	ruleMetadata := []*RuleMetadata{
		&RuleMetadata{
			MetadataKey: "RM1",
			RuleID:      "result_type",
			Status:      status,
		},
	}

	metadata := &Metadata{
		&Other{
			OtherMetadata: &OtherMetadata{
				RuleMetadatas: &RuleMetadatas{
					RuleMetadata: ruleMetadata,
				},
			},
		},
	}

	return metadata
}

type ShoppingResponseIDs struct {
	ResponseID    string
	AssociatedIDs *AssociatedIDs
}

type AssociatedIDs struct {
	AssociatedID *AssociatedID
}

type AssociatedID struct {
	OfferItemID *OfferItemID
}

type OfferItemID struct {
	Owner string `xml:",attr"`
	Value string `xml:",chardata"`
}

func MakeShoppingResponseIDs(owner, itemID string) *ShoppingResponseIDs {
	offerItemID := new(OfferItemID)
	offerItemID.Owner = owner
	offerItemID.Value = itemID

	return &ShoppingResponseIDs{
		ResponseID: "R1",
		AssociatedIDs: &AssociatedIDs{
			&AssociatedID{offerItemID},
		},
	}
}

func MakeSessionAirShoppingRQ() (request *Envelope) {
	request = &Envelope{
		Header: new(Header),
		Body: &Body{
			AirShoppingRQ: &AirShoppingRQ{
				Version:  "1.0",
				Document: new(Document),
				Party: &Party{
					Sender: new(Sender),
				},
				Travelers: &Travelers{
					Traveler: []*Traveler{new(Traveler)},
				},
				CoreQuery: new(CoreQuery),
			},
		},
	}

	return
}
