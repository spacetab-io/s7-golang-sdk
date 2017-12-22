package main

import (
	"gitlab.tm-consulting.ru/avia/services/s7-sdk"

	"encoding/xml"
	"log"
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
								Departure:        sdk.MakePoint("DME", "2018-08-24", "00:00", ""),
								Arrival:          sdk.MakePoint("MSQ", "2018-08-24", "00:00", ""),
								MarketingCarrier: sdk.MakeCarrier("S7", "", "4431"),
							},
						},
					},
					&sdk.OriginDestinationFlight{
						Flight: []*sdk.OriginDestination{
							&sdk.OriginDestination{
								SegmentKey:       "FL2",
								Departure:        sdk.MakePoint("MSQ", "2018-08-31", "00:00", ""),
								Arrival:          sdk.MakePoint("DME", "2018-08-31", "00:00", ""),
								MarketingCarrier: sdk.MakeCarrier("S7", "", "4434"),
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
									Code: "OFLRTCS",
								},
								RBD: "O",
							},
						},
						&sdk.FareComponent{
							Reference: "FL2",
							FareBasis: &sdk.FareBasis{
								FareBasisCode: &sdk.FareBasisCode{
									Code: "OFLRTCS",
								},
								RBD: "O",
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
