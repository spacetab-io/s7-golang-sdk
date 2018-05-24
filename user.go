package sdk

const (
	PASSENGER_TYPE_CODE_ADULT  = "ADT"
	PASSENGER_TYPE_CODE_CHILD  = "CHD"
	PASSENGER_TYPE_CODE_INFANT = "INF"
)

type User struct {
	ID              string `xml:"ObjectKey,attr,omitempty"`
	PTC             *PTC
	Age             *Age             `xml:",omitempty"`
	Name            *Name            `xml:",omitempty"`
	Contacts        *Contacts        `xml:",omitempty"`
	FQTVs           *FQTVs           `xml:",omitempty"`
	Gender          string           `xml:",omitempty"`
	PassengerIDInfo *PassengerIDInfo `xml:",omitempty"`
	Document        *Document        `xml:",omitempty"`
}

type PTC struct {
	Quantity int    `xml:",attr,omitempty"`
	Value    string `xml:",chardata"`
}

func MakePTC(qty int, passengerType string, qtyAdultChild, qtyInfant, qtyAdult *int) (ptc *PTC) {
	if qty < 1 {
		return
	}
	switch {
	case passengerType != PASSENGER_TYPE_CODE_ADULT &&
		passengerType != PASSENGER_TYPE_CODE_CHILD &&
		passengerType != PASSENGER_TYPE_CODE_INFANT:
		return
	case qty > 9:
		return
	case passengerType == PASSENGER_TYPE_CODE_ADULT:
		*qtyAdultChild += qty
		*qtyAdult += qty
	case passengerType == PASSENGER_TYPE_CODE_CHILD:
		*qtyAdultChild += qty
	case passengerType == PASSENGER_TYPE_CODE_INFANT:
		*qtyInfant += qty
	}

	// switch {
	// case *qtyAdultChild > 9:
	// 	log.Println(1)
	// 	return
	// case *qtyInfant > *qtyAdult:
	// 	log.Println(2)
	// 	return
	// }

	ptc = &PTC{
		Quantity: qty,
		Value:    passengerType,
	}

	return
}

type Age struct {
	BirthDate string
}

func MakeAge(date string) *Age {
	return &Age{
		BirthDate: date,
	}
}

type Name struct {
	Surname string
	Given   string
	Title   string `xml:",omitempty"`
	Middle  string `xml:",omitempty"`
}

func MakeName(title, given, surname, middle string) *Name {
	return &Name{
		Title:   title,
		Given:   given,
		Surname: surname,
		Middle:  middle,
	}
}

type Contacts struct {
	Contact []*Contact
}

type Contact struct {
	EmailContact *EmailContact `xml:",omitempty"`
	PhoneContact *PhoneContact `xml:",omitempty"`
}

func MakeContact(email, countryCode, areaCode, phoneNumber string) *Contact {
	var emailContact *EmailContact
	if email != "" {
		emailContact = &EmailContact{
			Address: email,
		}
	}

	var phoneContact *PhoneContact
	if countryCode != "" && areaCode != "" && phoneNumber != "" {
		phoneContact = &PhoneContact{
			Number: PhoneNumber{
				CountryCode: countryCode,
				AreaCode:    areaCode,
				Number:      phoneNumber,
			},
		}
	}

	return &Contact{
		EmailContact: emailContact,
		PhoneContact: phoneContact,
	}
}

type EmailContact struct {
	Address string
}

type PhoneContact struct {
	Number PhoneNumber
}

type PhoneNumber struct {
	CountryCode string `xml:",attr"`
	AreaCode    string `xml:",attr"`
	Number      string `xml:",chardata"`
}

type FQTVs struct {
	TravelerFQTVInformation *TravelerFQTVInformation `xml:"TravelerFQTV_Information"`
}

type TravelerFQTVInformation struct {
	AirlineID string   `xml:",omitempty"`
	Account   *Account `xml:",omitempty"`
	ProgramID string   `xml:",omitempty"`
}

func MakeTravelerFQTVInformation(airlineID, account, programID string) *TravelerFQTVInformation {
	return &TravelerFQTVInformation{
		AirlineID: airlineID,
		Account: &Account{
			Number: account,
		},
		ProgramID: programID,
	}
}

type Account struct {
	Number string
}

type PassengerIDInfo struct {
	PassengerDocument *PassengerDocument
}

type PassengerDocument struct {
	Type               string
	ID                 string
	BirthCountry       string
	DateOfIssue        string
	DateOfExpiration   string
	CountryOfResidence string
}

type Passengers struct {
	Passenger           []*User `xml:",omitempty"`
	PassengerReference  string  `xml:",omitempty"`
	PassengerReferences string  `xml:",omitempty"`
}
