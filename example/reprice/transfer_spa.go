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
						OriginDestinationKey: "OD1",
						Flight: []*sdk.OriginDestination{
							&sdk.OriginDestination{
								SegmentKey:       "FL1",
								Departure:        sdk.MakePoint("DME", "2018-09-16", "00:00", ""),
								Arrival:          sdk.MakePoint("MUC", "2018-09-16", "00:00", ""),
								MarketingCarrier: sdk.MakeCarrier("S7", "", "797"),
							},
							&sdk.OriginDestination{
								SegmentKey:       "FL2",
								Departure:        sdk.MakePoint("MUC", "2018-09-17", "00:00", ""),
								Arrival:          sdk.MakePoint("HAM", "2018-09-17", "00:00", ""),
								MarketingCarrier: sdk.MakeCarrier("LH", "", "2070"),
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
									Code: "QLHMUCV",
								},
								RBD: "Q",
							},
						},
						&sdk.FareComponent{
							Reference: "FL2",
							FareBasis: &sdk.FareBasis{
								FareBasisCode: &sdk.FareBasisCode{
									Code: "QLHMUCV",
								},
								RBD: "V",
							},
						},
					},
				},
			},
			Associations: &sdk.Associations{
				Passengers: &sdk.Passengers{
					PassengerReferences: "SH1 SH2 SH3",
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
			&sdk.User{
				ID: "SH2",
				PTC: &sdk.PTC{
					Value: sdk.PASSENGER_TYPE_CODE_CHILD,
				},
				Name: &sdk.Name{},
			},
			&sdk.User{
				ID: "SH3",
				PTC: &sdk.PTC{
					Value: sdk.PASSENGER_TYPE_CODE_INFANT,
				},
				Name: &sdk.Name{},
			},
		},
	}

	output, err := xml.MarshalIndent(request, "", "  ")
	log.Printf("out: \n%s\nerr: %v\n", output, err)
}
