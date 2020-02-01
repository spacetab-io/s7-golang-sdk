package structsV052

type TravelerType struct {
	PTC             *PTC
	Age             *Age             `xml:",omitempty"`
	Name            *Name            `xml:",omitempty"`
	Contacts        *Contacts        `xml:",omitempty"`
	Gender          string           `xml:",omitempty"`
	PassengerIDInfo *PassengerIDInfo `xml:",omitempty"`
	ID              string           `xml:"ObjectKey,attr,omitempty"` // SH1
	//EntryPointID    string           `xml:"-"`                        // 242
	//FQTVs				*FQTVs				`xml:",omitempty"`
	//Document			*Document			`xml:",omitempty"`
}

type PassengerOrderViewRS struct {
	ObjectKey       string `xml:",attr"`
	PTC             *PTC
	Age             *Age             `xml:",omitempty"`
	Name            *Name            `xml:",omitempty"`
	Gender          string           `xml:",omitempty"`
	PassengerIDInfo *PassengerIDInfo `xml:",omitempty"`
}

type PTC struct {
	Quantity int    `xml:",attr,omitempty"` // 1
	Value    string `xml:",chardata"`       // ADT

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
	Given   string `xml:",omitempty"`
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
	Application string `xml:",omitempty"`
	Number      PhoneNumber
}

type PhoneNumber struct {
	CountryCode string `xml:",attr"`
	AreaCode    string `xml:",attr"`
	Number      string `xml:",chardata"`
}

type Account struct {
	Number string
}

type PassengerIDInfo struct {
	PassengerDocument *PassengerDocument
}

type PassengerDocument struct {
	Type               string `xml:",omitempty"`
	ID                 string
	BirthCountry       string `xml:",omitempty"`
	DateOfIssue        string `xml:",omitempty"`
	DateOfExpiration   string `xml:",omitempty"`
	CountryOfResidence string `xml:",omitempty"`
}

type Passengers struct {
	Passenger           []*TravelerType `xml:"Passenger,omitempty"`
	PassengerReference  string          `xml:",omitempty"`
	PassengerReferences string          `xml:",omitempty"`
}

type PassengersAirDocIssueRQ struct {
	Passenger []*Passenger `xml:"Passenger,omitempty"`
}

// MakePassengerIDInfo is a constructor of PassengerIDInfo
func MakePassengerIDInfo(
	docType,
	docID,
	birthCountry,
	dateOfIssue,
	dateOfExpiration,
	countryOfResidence string,
) *PassengerIDInfo {
	var passengerDocument *PassengerDocument
	if docType != "" && docID != "" {
		passengerDocument = &PassengerDocument{
			Type:               docType,
			ID:                 docID,
			BirthCountry:       birthCountry,
			DateOfIssue:        dateOfIssue,
			DateOfExpiration:   dateOfExpiration,
			CountryOfResidence: countryOfResidence,
		}
	}

	return &PassengerIDInfo{PassengerDocument: passengerDocument}
}
