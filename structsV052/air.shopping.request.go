package structsV052

import (
	"encoding/xml"
)

// AirShoppingRQ is a body of S7 searchflights request
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

// Parameters section of S7 searchflightsJourney request
type Parameters struct {
	//ServiceFilters		*ServiceFilters		`xml:",omitempty"`
	CurrCodes *CurrCodes `xml:",omitempty"`
}

// CurrCodes is a subsection in Parametrs section holding currency code
type CurrCodes struct {
	//CurrCode			[]string
	CurrCode string
}

// Travelers section of S7 searchflightsJourney request
type Travelers struct {
	Traveler []*Traveler
}

// Traveler is an element in Travelers section of S7 searchflightsJourney request
type Traveler struct {
	AnonymousTraveler *TravelerType `xml:",omitempty"`
}

// CoreQuery section of S7 searchflightsJourney request
type CoreQuery struct {
	OriginDestinations  *OriginDestinations  `xml:",omitempty"`
	ShoppingResponseIDs *ShoppingResponseIDs `xml:",omitempty"`
}

// OriginDestinations section of S7 searchflightsJourney request
type OriginDestinations struct {
	OriginDestination []*OriginDestination
}

// ShoppingResponseIDs
type ShoppingResponseIDs struct {
	ResponseID    string
	AssociatedIDs *AssociatedIDs
}

// AssociatedIDs section
type AssociatedIDs struct {
	AssociatedID *AssociatedID
}

// AssociatedID is an element of AssociatedIDs section
type AssociatedID struct {
	OfferItemID *OfferItemID
}

// OfferItemID element
type OfferItemID struct {
	Owner string `xml:",attr"`
	Value string `xml:",chardata"`
}

// Preferences section
type Preferences struct {
	Preference []*Preference
}

// Preference element of Preferences section
type Preference struct {
	FlightPreferences *FlightPreferences
}

// FlightPreferences section in Preference element
type FlightPreferences struct {
	Aircraft       *Aircraft
	Characteristic *Characteristic
}

// Aircraft element of FlightPreferences section
type Aircraft struct {
	Cabins *Cabins
}

// Cabins section
type Cabins struct {
	Cabin *Cabin
}

// Cabin element of Cabins section
type Cabin struct {
	Code string
}

// Characteristic element of FlightPreferences section
type Characteristic struct {
	DirectPreferences string
}

// MakePreference generates Preference for S7 searchflightsJourney request
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

func MakeAirShoppingRQ(
	agent *AgentUserSender,
	query *CoreQuery,
	version,
	currency string,
) *AirShoppingRQ {
	var parameters *Parameters

	parameters = &Parameters{
		CurrCodes: &CurrCodes{
			CurrCode: currency,
		},
	}

	airShoppingRQ := &AirShoppingRQ{
		Version:  Version, // "1.0" - v0.50, "2.0" - v0.52
		Document: &Document{},
		Party: &Party{
			Sender: &Sender{
				AgentUserSender: agent,
			},
		},
		Parameters: parameters,
		CoreQuery:  query,
	}
	// if qualifiers != nil {
	// 	airShoppingRQ.Qualifiers = qualifiers
	// }

	return airShoppingRQ
}
