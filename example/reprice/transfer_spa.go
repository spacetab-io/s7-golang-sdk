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
						OriginDestinationKey: "OD1",
						Flight: []*sdk.OriginDestination{
							{
								SegmentKey:       "FL1",
								Departure:        sdk.MakePoint("DME", "2018-09-16", "00:00", ""),
								Arrival:          sdk.MakePoint("MUC", "2018-09-16", "00:00", ""),
								MarketingCarrier: sdk.MakeCarrier("S7", "", "797"),
							},
							{
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
						{
							Reference: "FL1",
							FareBasis: &sdk.FareBasis{
								FareBasisCode: &sdk.FareBasisCode{
									Code: "QLHMUCV",
								},
								RBD: "Q",
							},
						},
						{
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
			{
				ID: "SH1",
				PTC: &sdk.PTC{
					Value: sdk.PassengerTypeCodeAdult,
				},
				Name: &sdk.Name{},
			},
			{
				ID: "SH2",
				PTC: &sdk.PTC{
					Value: sdk.PassengerTypeCodeChild,
				},
				Name: &sdk.Name{},
			},
			{
				ID: "SH3",
				PTC: &sdk.PTC{
					Value: sdk.PassengerTypeCodeInfant,
				},
				Name: &sdk.Name{},
			},
		},
	}

	output, err := xml.MarshalIndent(request, "", "  ")
	log.Printf("out: \n%s\nerr: %v\n", output, err)
}
