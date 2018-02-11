package main

import (
	"encoding/xml"
	"log"

	sdk "github.com/tmconsulting/s7-golang-sdk"
)

func main() {
	request := sdk.MakeSessionBookRQ()

	request.Body.OrderCreateRQ.Party.Sender.AgentUserSender = sdk.MakeAgentUserSender(
		"OVB", "OVB902/1OVB2TR",
		"RU", "00000011Z", "1",
		"OVB902", "1001/1001A", "115")

	passengers := []*sdk.User{
		&sdk.User{
			ID: "SH1",
			PTC: &sdk.PTC{
				Quantity: 1,
				Value:    sdk.PASSENGER_TYPE_CODE_ADULT,
			},
			Age:  sdk.MakeAge("1970-10-25"),
			Name: sdk.MakeName("Mrs", "Lila", "Conelly", ""),
			Contacts: &sdk.Contacts{
				Contact: []*sdk.Contact{
					sdk.MakeContact("l.conelly@mail.ru", "375", "29", "3277438"),
				},
			},
			Gender: "Female",
			PassengerIDInfo: &sdk.PassengerIDInfo{
				PassengerDocument: &sdk.PassengerDocument{
					Type:               "PP",
					ID:                 "131231231222",
					BirthCountry:       "RU",
					DateOfIssue:        "2008-06-06",
					DateOfExpiration:   "2028-06-06",
					CountryOfResidence: "RU",
				},
			},
		},
	}

	request.Body.OrderCreateRQ.Query.Passengers = &sdk.Passengers{
		Passenger: passengers,
	}

	request.Body.OrderCreateRQ.Query.OrderItems = &sdk.OrderItems{
		ShoppingResponse: &sdk.ShoppingResponse{
			Owner:      "S7",
			ResponseID: "UNKNOWN",
			Offers: &sdk.Offers{
				Offer: []*sdk.Offer{
					&sdk.Offer{
						OfferID: &sdk.OfferID{
							Owner: "S7",
							Value: "UNKNOWN",
						},
						OfferItems: &sdk.OfferItems{
							OfferItem: []*sdk.OfferItem{
								&sdk.OfferItem{
									OfferItemID: &sdk.OfferID{
										Owner: "S7",
										Value: "UNKNOWN",
									},
									Passengers: &sdk.Passengers{
										PassengerReference: "SH1",
									},
									ApplicableFlight: &sdk.ApplicableFlight{
										FlightReferences: "FL1 FL2",
									},
								},
							},
						},
					},
				},
			},
		},
		OfferItem: &sdk.OfferItem{
			OfferItemID: &sdk.OfferID{
				Owner: "S7",
				Value: "UNKNOWN",
			},
			OfferItemType: &sdk.OfferItemType{
				DetailedFlightItem: &sdk.DetailedFlightItem{
					OriginDestination: []*sdk.OriginDestinationFlight{
						&sdk.OriginDestinationFlight{
							Flight: []*sdk.OriginDestination{
								&sdk.OriginDestination{
									SegmentKey:       "FL1",
									Departure:        sdk.MakePoint("MUC", "2017-12-27", "10:35", ""),
									Arrival:          sdk.MakePoint("DME", "2017-12-27", "15:45", ""),
									MarketingCarrier: sdk.MakeCarrier("S7", "", "796"),
									OperatingCarrier: sdk.MakeCarrier("S7", "", "796"),
									ClassOfService:   sdk.MakeClassOfService(0, "O", ""),
								},
								&sdk.OriginDestination{
									SegmentKey:       "FL2",
									Departure:        sdk.MakePoint("DME", "2017-12-27", "08:30", ""),
									Arrival:          sdk.MakePoint("MUC", "2017-12-27", "09:40", ""),
									MarketingCarrier: sdk.MakeCarrier("S7", "", "795"),
									OperatingCarrier: sdk.MakeCarrier("S7", "", "795"),
									ClassOfService:   sdk.MakeClassOfService(0, "O", ""),
								},
							},
						},
					},
				},
			},
		},
	}

	request.Body.OrderCreateRQ.Query.DataLists = &sdk.DataLists{
		InstructionsList: &sdk.InstructionsList{
			Instruction: []*sdk.Instruction{
				&sdk.Instruction{
					ListKey: "CC",
					SpecialBookingInstruction: &sdk.SpecialBookingInstruction{
						Code:       "QUW1725",
						Definition: "10187",
					},
				},
			},
		},
	}

	output, err := xml.MarshalIndent(request, "", "  ")
	log.Printf("out: \n%s\nerr: %v\n", output, err)
}
