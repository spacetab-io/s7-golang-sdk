package main

import (
	"gitlab.tm-consulting.ru/avia/services/s7-sdk"

	"encoding/xml"
	"log"
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
			Age:  sdk.MakeAge("1991-10-10"),
			Name: sdk.MakeName("MR", "Danil", "Trefilov", "Alexandrovich"),
			Contacts: &sdk.Contacts{
				Contact: []*sdk.Contact{
					sdk.MakeContact("test.test@s7.ru", "7", "999", "9999999"),
				},
			},
			FQTVs: &sdk.FQTVs{
				TravelerFQTVInformation: sdk.MakeTravelerFQTVInformation("S7", "612966192", "S7"),
			},
			Gender: "Male",
			PassengerIDInfo: &sdk.PassengerIDInfo{
				PassengerDocument: &sdk.PassengerDocument{
					Type:               "PP",
					ID:                 "111111",
					BirthCountry:       "RU",
					DateOfIssue:        "2010-10-10",
					DateOfExpiration:   "2021-10-10",
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
									Departure:        sdk.MakePoint("DME", "2017-08-24", "05:40", ""),
									Arrival:          sdk.MakePoint("MSQ", "2017-08-24", "07:00", ""),
									MarketingCarrier: sdk.MakeCarrier("S7", "", "4431"),
									OperatingCarrier: sdk.MakeCarrier("B2", "", ""),
									Equipment:        sdk.MakeEquipment("ref", "Boeing 737-500 Passenger"),
									ClassOfService:   sdk.MakeClassOfService(1, "D", "DFLRTCS"),
								},
							},
						},
						&sdk.OriginDestinationFlight{
							Flight: []*sdk.OriginDestination{
								&sdk.OriginDestination{
									SegmentKey:       "FL2",
									Departure:        sdk.MakePoint("MSQ", "2017-08-31", "06:45", ""),
									Arrival:          sdk.MakePoint("DME", "2017-08-31", "08:05", ""),
									MarketingCarrier: sdk.MakeCarrier("S7", "", "4434"),
									OperatingCarrier: sdk.MakeCarrier("B2", "", ""),
									Equipment:        sdk.MakeEquipment("ref", "Boeing 737-300 Passenger"),
									ClassOfService:   sdk.MakeClassOfService(1, "D", "DFLRTCS"),
								},
							},
						},
					},
				},
			},
		},
	}

	request.Body.OrderCreateRQ.Query.DataLists = &sdk.DataLists{
		ServiceList: &sdk.ServiceList{
			Service: []*sdk.Service{
				&sdk.Service{
					ID:        "SRV1",
					ServiceID: sdk.MakeOfferID("S7", "SRV1"),
					Name:      "Empty",
					Descriptions: &sdk.Descriptions{
						Description: []*sdk.Description{
							&sdk.Description{
								Text: "Empty",
							},
						},
					},
				},
			},
		},
	}

	output, err := xml.MarshalIndent(request, "", "  ")
	log.Printf("out: \n%s\nerr: %v\n", output, err)
}
