package sdk

import (
	"fmt"
	"time"
)

const (
	longForm  = "2006-01-02 15:04"
	longFormT = "2006-01-02T15:04:05"
)

type OriginDestination struct {
	Reference        string          `xml:"refs,attr,omitempty"`
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

func (p *Point) GetTerminal() string {
	if p.Terminal != nil {
		return p.Terminal.Name
	}
	return ""
}

func (p *Point) GetDate() time.Time {
	date, _ := time.Parse(longForm, fmt.Sprintf("%s %s", p.Date, p.Time))
	return date
}

func (p *Point) GetDateISO() string {
	return fmt.Sprintf("%sT%s", p.Date, p.Time)
}

type Terminal struct {
	Name string
}

func MakePoint(namePoint, dataPoint, timePoint, terminal string) *Point {
	p := &Point{
		AirportCode: namePoint,
		Date:        dataPoint,
		Time:        timePoint,
	}
	if terminal != "" {
		p.Terminal = &Terminal{
			Name: terminal,
		}
	}
	return p
}

func MakeOriginDestination(departureAirportCode string, dateDep time.Time, arrivalAirportCode string, dateArr time.Time, daysBefore, daysAfter int) *OriginDestination {
	originDestination := new(OriginDestination)
	originDestination.Departure = &Point{
		AirportCode: departureAirportCode,
	}
	if dateDep.Year() != 1 {
		originDestination.Departure.Date = dateDep.Format("2006-01-02")
	}
	originDestination.Arrival = &Point{
		AirportCode: arrivalAirportCode,
	}
	if dateArr.Year() != 1 {
		originDestination.Arrival.Date = dateArr.Format("2006-01-02")
	}

	if daysBefore > 0 || daysAfter > 0 {
		originDestination.CalendarDates = &CalendarDates{
			DaysBefore: daysBefore,
			DaysAfter:  daysAfter,
		}
	}

	return originDestination
}

type OriginDestinationFlight struct {
	OriginDestinationKey string `xml:",omitempty"`
	Flight               []*OriginDestination
}
