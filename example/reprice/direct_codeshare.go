package main

import (
	"encoding/xml"
	"log"

	sdk "github.com/tmconsulting/s7-golang-sdk"
)

func main() {
	request := sdk.MakeItinReshopRQ(&sdk.ActionType{Context: sdk.ActionTypeContextDiscount, Value: 10})

	request.Body.ItinReshopRQ.Party.Sender.AgentUserSender = sdk.MakeAgentUserSender(
		"OVB", "OVB902/1OVB2TR",
		"RU", "00000011Z", "1",
		"OVB902", "1001/1001A", "115")

	request.Body.ItinReshopRQ.Query.Reshop.Actions.OrderItems = &sdk.OrderItems{
		OrderItem: &sdk.OrderItem{
			FlightItem: &sdk.FlightItem{
				OriginDestination: []*sdk.OriginDestinationFlight{
					{
						Flight: []*sdk.OriginDestination{
							{
								SegmentKey:       "FL1",
								Departure:        sdk.MakePoint("DME", "2018-08-24", "00:00", ""),
								Arrival:          sdk.MakePoint("MSQ", "2018-08-24", "00:00", ""),
								MarketingCarrier: sdk.MakeCarrier("S7", "", "4431"),
							},
						},
					},
					{
						Flight: []*sdk.OriginDestination{
							{
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
						{
							Reference: "FL1",
							FareBasis: &sdk.FareBasis{
								FareBasisCode: &sdk.FareBasisCode{
									Code: "OFLRTCS",
								},
								RBD: "O",
							},
						},
						{
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
			{
				ID: "SH1",
				PTC: &sdk.PTC{
					Value: sdk.PassengerTypeCodeAdult,
				},
				Name: &sdk.Name{},
			},
		},
	}

	output, err := xml.MarshalIndent(request, "", "  ")
	log.Printf("out: \n%s\nerr: %v\n", output, err)
}
