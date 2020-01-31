package structsV052

type OriginDestination struct {
	ID               string          `xml:"OriginDestinationKey,attr,omitempty"` // Example: origDest1
	Refs             string          `xml:"refs,attr,omitempty"`
	SegmentKey       string          `xml:",omitempty"`
	Status           *Status         `xml:",omitempty"`
	Departure        *Point          `xml:",omitempty"`
	Arrival          *Point          `xml:",omitempty"`
	MarketingCarrier *Carrier        `xml:",omitempty"`
	OperatingCarrier *Carrier        `xml:",omitempty"`
	CalendarDates    *CalendarDates  `xml:",omitempty"`
	Equipment        *Equipment      `xml:",omitempty"`
	CabinType        *StatusCode     `xml:",omitempty"`
	ClassOfService   *ClassOfService `xml:",omitempty"`
	Flight           *Flight         `xml:",omitempty"`
}

type Status struct {
	StatusCode *StatusCode
}

type StatusCode struct {
	Code string
}

type Point struct {
	AirportCode string
	Date        string    `xml:",omitempty"`
	Time        string    `xml:",omitempty"`
	Terminal    *Terminal `xml:",omitempty"`
}

func MakePoint(namePoint, datePoint, timePoint, terminal string) *Point {
	p := &Point{
		AirportCode: namePoint,
		Date:        datePoint,
		Time:        timePoint,
	}
	if terminal != "" {
		p.Terminal = &Terminal{
			Name: terminal,
		}
	}
	return p
}

type Terminal struct {
	Name string
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

type ClassOfService struct {
	Code          *Code
	MarketingName *string `xml:",omitempty"` // NBSRT
}

type OriginDestinationFlight struct {
	OriginDestinationKey string `xml:",omitempty"`
	Flight               []*OriginDestination
}
