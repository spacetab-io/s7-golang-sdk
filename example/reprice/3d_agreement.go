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
								Departure:        sdk.MakePoint("MUC", "2017-12-27", "10:35", "1"),
								Arrival:          sdk.MakePoint("DME", "2017-12-27", "15:45", ""),
								MarketingCarrier: sdk.MakeCarrier("S7", "", "796"),
								OperatingCarrier: sdk.MakeCarrier("S7", "", "796"),
								Equipment:        sdk.MakeEquipment("ref", "Airbus A320-100/200"),
							},
						},
					},
					&sdk.OriginDestinationFlight{
						OriginDestinationKey: "OD2",
						Flight: []*sdk.OriginDestination{
							&sdk.OriginDestination{
								SegmentKey:       "FL2",
								Departure:        sdk.MakePoint("DME", "2018-01-27", "08:30", ""),
								Arrival:          sdk.MakePoint("MUC", "2018-01-27", "09:40", "1"),
								MarketingCarrier: sdk.MakeCarrier("S7", "", "795"),
								OperatingCarrier: sdk.MakeCarrier("S7", "", "795"),
								Equipment:        sdk.MakeEquipment("ref", "Airbus A320-100/200"),
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
									Code:        "OBSRT",
									Application: "XEX",
								},
								RBD: "O",
							},
						},
						&sdk.FareComponent{
							Reference: "FL2",
							FareBasis: &sdk.FareBasis{
								FareBasisCode: &sdk.FareBasisCode{
									Code:        "OBSRT",
									Application: "XAT",
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

	request.Body.ItinReshopRQ.Query.Reshop.Actions.Qualifiers = &sdk.Qualifiers{
		Qualifier: []*sdk.Qualifier{
			sdk.MakeQualifier("QUW1725", "10187"),
		},
	}

	output, err := xml.MarshalIndent(request, "", "  ")
	log.Printf("out: \n%s\nerr: %v\n", output, err)
}
