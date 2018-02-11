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
			Age:  sdk.MakeAge("1991-10-09"),
			Name: sdk.MakeName("MR", "John", "Doe", ""),
			Contacts: &sdk.Contacts{
				Contact: []*sdk.Contact{
					sdk.MakeContact("john_doe@mail.com", "7", "999", "4644694"),
				},
			},
			Gender: "Male",
			PassengerIDInfo: &sdk.PassengerIDInfo{
				PassengerDocument: &sdk.PassengerDocument{
					Type:               "PP",
					ID:                 "111111",
					BirthCountry:       "RU",
					DateOfIssue:        "2010-10-09",
					DateOfExpiration:   "2020-10-09",
					CountryOfResidence: "RU",
				},
			},
		},
		&sdk.User{
			ID: "SH2",
			PTC: &sdk.PTC{
				Quantity: 1,
				Value:    sdk.PASSENGER_TYPE_CODE_ADULT,
			},
			Age:  sdk.MakeAge("1991-10-26"),
			Name: sdk.MakeName("MRS", "Kate", "Doe", ""),
			Contacts: &sdk.Contacts{
				Contact: []*sdk.Contact{
					sdk.MakeContact("kate_doe@mail.com", "7", "999", "4623394"),
				},
			},
			Gender: "Female",
			PassengerIDInfo: &sdk.PassengerIDInfo{
				PassengerDocument: &sdk.PassengerDocument{
					Type:               "PP",
					ID:                 "111111",
					BirthCountry:       "RU",
					DateOfIssue:        "2010-10-08",
					DateOfExpiration:   "2020-10-08",
					CountryOfResidence: "RU",
				},
			},
		},
		&sdk.User{
			ID: "SH3",
			PTC: &sdk.PTC{
				Quantity: 1,
				Value:    sdk.PASSENGER_TYPE_CODE_CHILD,
			},
			Age:  sdk.MakeAge("2010-10-09"),
			Name: sdk.MakeName("MS", "Sevda", "Doe", ""),
			Contacts: &sdk.Contacts{
				Contact: []*sdk.Contact{
					sdk.MakeContact("sevda_doe@mail.com", "7", "999", "1144694"),
				},
			},
			Gender: "Female",
			PassengerIDInfo: &sdk.PassengerIDInfo{
				PassengerDocument: &sdk.PassengerDocument{
					Type:               "PP",
					ID:                 "111111",
					BirthCountry:       "RU",
					DateOfIssue:        "2010-10-09",
					DateOfExpiration:   "2020-10-09",
					CountryOfResidence: "RU",
				},
			},
		},
		&sdk.User{
			ID: "SH4",
			PTC: &sdk.PTC{
				Quantity: 1,
				Value:    sdk.PASSENGER_TYPE_CODE_CHILD,
			},
			Age:  sdk.MakeAge("2010-10-09"),
			Name: sdk.MakeName("MS", "Vlada", "Doe", ""),
			Contacts: &sdk.Contacts{
				Contact: []*sdk.Contact{
					sdk.MakeContact("vlada_doe@mail.com", "7", "999", "1244694"),
				},
			},
			Gender: "Female",
			PassengerIDInfo: &sdk.PassengerIDInfo{
				PassengerDocument: &sdk.PassengerDocument{
					Type:               "PP",
					ID:                 "111111",
					BirthCountry:       "RU",
					DateOfIssue:        "2010-10-09",
					DateOfExpiration:   "2020-10-09",
					CountryOfResidence: "RU",
				},
			},
		},
		&sdk.User{
			ID: "SH5",
			PTC: &sdk.PTC{
				Quantity: 1,
				Value:    sdk.PASSENGER_TYPE_CODE_INFANT,
			},
			Age:  sdk.MakeAge("2017-01-09"),
			Name: sdk.MakeName("MR", "Viktor", "Doe", ""),
			Contacts: &sdk.Contacts{
				Contact: []*sdk.Contact{
					sdk.MakeContact("viktor_doe@mail.com", "7", "999", "0044694"),
				},
			},
			Gender: "Male",
			PassengerIDInfo: &sdk.PassengerIDInfo{
				PassengerDocument: &sdk.PassengerDocument{
					Type:               "PP",
					ID:                 "111111",
					BirthCountry:       "RU",
					DateOfIssue:        "2017-01-09",
					DateOfExpiration:   "2020-10-09",
					CountryOfResidence: "RU",
				},
			},
		},
		&sdk.User{
			ID: "SH6",
			PTC: &sdk.PTC{
				Quantity: 1,
				Value:    sdk.PASSENGER_TYPE_CODE_INFANT,
			},
			Age:  sdk.MakeAge("2017-01-09"),
			Name: sdk.MakeName("MR", "Vladimir", "Doe", ""),
			Contacts: &sdk.Contacts{
				Contact: []*sdk.Contact{
					sdk.MakeContact("vladimir_doe@mail.com", "7", "999", "1344694"),
				},
			},
			Gender: "Male",
			PassengerIDInfo: &sdk.PassengerIDInfo{
				PassengerDocument: &sdk.PassengerDocument{
					Type:               "PP",
					ID:                 "111111",
					BirthCountry:       "RU",
					DateOfIssue:        "2017-01-09",
					DateOfExpiration:   "2020-10-09",
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
										PassengerReference: "SH1 SH2 SH3 SH4 SH5 SH6",
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
									Departure:        sdk.MakePoint("DME", "2017-08-02", "08:55", ""),
									Arrival:          sdk.MakePoint("MUC", "2017-08-02", "11:05", "1"),
									MarketingCarrier: sdk.MakeCarrier("S7", "", "795"),
									OperatingCarrier: sdk.MakeCarrier("S7", "", "795"),
									Equipment:        sdk.MakeEquipment("ref", "Airbus A319"),
									ClassOfService:   sdk.MakeClassOfService(1, "T", "QFLOW"),
								},
								&sdk.OriginDestination{
									SegmentKey:       "FL2",
									Departure:        sdk.MakePoint("MUC", "2017-08-02", "15:40", "1"),
									Arrival:          sdk.MakePoint("HAM", "2017-08-02", "17:00", "1"),
									MarketingCarrier: sdk.MakeCarrier("S7", "", "4725"),
									OperatingCarrier: sdk.MakeCarrier("AB", "", ""),
									Equipment:        sdk.MakeEquipment("ref", "Airbus A320-100/200"),
									ClassOfService:   sdk.MakeClassOfService(1, "T", "QFLOW"),
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
