package main

import (
	"encoding/xml"
	"log"

	sdk "github.com/tmconsulting/s7-golang-sdk"
)

func main() {
	request := sdk.MakeItinReshopRQ()

	request.Body.ItinReshopRQ.Party.Sender.AgentUserSender = sdk.MakeAgentUserSender(
		"OVB", "OVB902/1OVB2TR",
		"RU", "00000011Z", "1",
		"OVB902", "1001/1001A", "115")

	request.Body.ItinReshopRQ.Query.Reshop.Actions.OrderItems = &sdk.OrderItems{
		OrderItem: &sdk.OrderItem{
			FlightItem: &sdk.FlightItem{
				OriginDestination: []*sdk.OriginDestinationFlight{
					&sdk.OriginDestinationFlight{
						Flight: []*sdk.OriginDestination{
							&sdk.OriginDestination{
								SegmentKey:       "FL1",
								Departure:        sdk.MakePoint("SLY", "2018-07-11", "00:00", ""),
								Arrival:          sdk.MakePoint("OVB", "2018-07-11", "00:00", ""),
								MarketingCarrier: sdk.MakeCarrier("S7", "", "3028"),
							},
						},
					},
				},
				FareDetail: &sdk.FareDetail{
					FareComponent: []*sdk.FareComponent{
						&sdk.FareComponent{
							Reference: "FL1",
							FareBasis: &sdk.FareBasis{
								FareBasisCode: &sdk.FareBasisCode{
									Code: "MBSOW",
								},
								RBD: "M",
							},
						},
					},
				},
			},
			Associations: &sdk.Associations{
				Passengers: &sdk.Passengers{
					PassengerReferences: "SH1",
				},
			},
		},
	}

	request.Body.ItinReshopRQ.Query.Reshop.Actions.Passengers = &sdk.Passengers{
		Passenger: []*sdk.User{
			&sdk.User{
				ID: "SH1",
				PTC: &sdk.PTC{
					Value: sdk.PASSENGER_TYPE_CODE_ADULT,
				},
				Name: &sdk.Name{},
			},
		},
	}

	output, err := xml.MarshalIndent(request, "", "  ")
	log.Printf("out: \n%s\nerr: %v\n", output, err)
}
