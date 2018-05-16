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
		{
			ID: "SH1",
			PTC: &sdk.PTC{
				Quantity: 1,
				Value:    sdk.PassengerTypeCodeAdult,
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
					ID:                 "111111111",
					BirthCountry:       "RU",
					DateOfIssue:        "2011-10-10",
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
					{
						OfferID: &sdk.OfferID{
							Owner: "S7",
							Value: "UNKNOWN",
						},
						OfferItems: &sdk.OfferItems{
							OfferItem: []*sdk.OfferItem{
								{
									OfferItemID: &sdk.OfferID{
										Owner: "S7",
										Value: "UNKNOWN",
									},
									Passengers: &sdk.Passengers{
										PassengerReference: "SH1",
									},
									ApplicableFlight: &sdk.ApplicableFlight{
										FlightReferences: "FL1",
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
						{
							Flight: []*sdk.OriginDestination{
								{
									SegmentKey:       "FL1",
									Departure:        sdk.MakePoint("SLY", "2018-08-31", "14:30", ""),
									Arrival:          sdk.MakePoint("OVB", "2018-08-31", "21:30", ""),
									MarketingCarrier: sdk.MakeCarrier("S7", "", "3028"),
									OperatingCarrier: sdk.MakeCarrier("S7", "", "3028"),
									Equipment:        sdk.MakeEquipment("ref", "Embraer 170 Regional Jet"),
									ClassOfService:   sdk.MakeClassOfService(1, "M", "MBSOW"),
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
				{
					ID:        "SRV1",
					ServiceID: sdk.MakeOfferID("S7", "SRV1"),
					Name:      "Empty",
					Descriptions: &sdk.Descriptions{
						Description: []*sdk.Description{
							{
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
