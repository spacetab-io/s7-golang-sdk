package structsV052

import "encoding/xml"

var MockOfferPrice = &OfferPrice{
	RequestedDate: &RequestedDate{
		PriceDetail: MockPriceDetail,
		Associations: []*Associations{
			{
				AssociatedTraveler: &AssociatedTraveler{TravelerReferences: "SH1"},
				ApplicableFlight:   MockApplicableFlight,
				OtherAssociation:   nil,
				Passengers:         nil,
			},
		},
	},
}

var MockADTFare = 4444.0
var MockCHDFare = 2222.0

var MockADTTotalAmount = 5555.0
var MockCHDTotalAmount = 11110.0

var MockCHDTax = 8888.0
var MockADTTax = 1111.0

var MockSeatsLeft = 7
var MockCodeValue = "Y"

//var MockMarketingName = "S7"

var MockClassOfService = ClassOfService{
	Code: &Code{
		SeatsLeft: &MockSeatsLeft,
		Value:     &MockCodeValue,
	},
	MarketingName: &MockMarketingName,
}

var MockCabinDesignator = "Y"
var MockMarketingName = "BASIC ECONOMY"

var MockCabinFlightSegment = CabinFlightSegment{
	CabinDesignator: &MockCabinDesignator,
	MarketingName:   &MockMarketingName,
}

var MockApplicableFlight = &ApplicableFlight{
	OriginDestinationReferences: "",
	FlightReferences:            "",
	FlightSegmentReference: []*FlightSegmentReference{
		{
			Reference:            "SEG1",
			Cabin:                &MockCabinFlightSegment,
			ClassOfService:       &MockClassOfService,
			BagDetailAssociation: nil,
		},
	},
}

var MockAssociation = &Associations{
	AssociatedTraveler: &AssociatedTraveler{TravelerReferences: "SH1"},
	ApplicableFlight:   MockApplicableFlight,
	OtherAssociation:   nil,
	Passengers:         nil,
}

var MockPriceDetail = &PriceDetail{
	TotalAmount: &mockTotalAmount,
	BaseAmount: &Total{
		Code:  "EU",
		Value: 200.0,
	},
	FareFiledIn: &FareFiledIn{
		BaseAmount: &BaseAmount{
			Code:  "EU",
			Value: 200.0,
		},
		ExchangeRate: 0,
	},
	Taxes: &Taxes{
		Total: &Total{
			Code:  "EU",
			Value: 225.0,
		},
		Breakdown: &Breakdown{
			Tax: []*Tax{{
				Amount: &Total{
					Code:  "EU",
					Value: 225.0,
				},
			}},
		},
	},
	Fees: nil,
}

var MockTravelerType = []*TravelerType{{
	PTC: &PTC{
		Quantity: 0,
		Value:    "ADT",
	},
	Age: &Age{BirthDate: "1991-01-01"},
	Name: &Name{
		Surname: "",
		Given:   "",
		Title:   "",
		Middle:  "",
	},
	Contacts: &Contacts{
		Contact: []*Contact{{
			EmailContact: nil,
			PhoneContact: nil,
		}},
	},
	Gender: "",
	PassengerIDInfo: &PassengerIDInfo{
		PassengerDocument: &PassengerDocument{
			Type:               "",
			ID:                 "",
			BirthCountry:       "",
			DateOfIssue:        "1991-02-01",
			DateOfExpiration:   "",
			CountryOfResidence: "",
		},
	},
	ID: "SH1",
}}

var mockTotalPrice = TotalPrice{
	DetailCurrencyPrice: &DetailCurrencyPrice{
		Total: &Total{
			Code:  "EU",
			Value: 450.0,
		},
		Details: &DetailsPrice{[]*DetailPrice{
			{
				SubTotal: &Total{
					Code:  "EU",
					Value: 225.0,
				},
				Application: "ZZ",
			},
		}},
		Fees: nil,
		Taxes: &Taxes{
			Total: &Total{
				Code:  "EU",
				Value: 225.0,
			},
			Breakdown: &Breakdown{
				Tax: []*Tax{{
					Amount: &Total{
						Code:  "EU",
						Value: 225,
					},
				}},
			},
		},
	},
}

var mockTotalAmount = TotalAmount{
	DetailCurrencyPrice: &DetailCurrencyPrice{
		Total: &Total{
			Code:  "EU",
			Value: 425.0,
		},
		Details: &DetailsPrice{[]*DetailPrice{}},
		Fees:    nil,
		Taxes: &Taxes{
			Total: &Total{
				Code:  "EU",
				Value: 225,
			},
			Breakdown: &Breakdown{
				Tax: []*Tax{{
					Amount: &Total{
						Code:  "EU",
						Value: 225,
					},
				}},
			},
		},
	},
}

var MockResponse = ReshopResponse{
	ReShopOffers: &ReShopOffers{
		ReShopOffer: []*ReShopOffer{
			{
				"",
				nil,
				&mockTotalPrice,
				&PricedOffer{
					OfferPrice: []*OfferPrice{
						MockOfferPrice,
					},
				},
			},
		},
	},

	Passengers: MockPassengersRefs,
}

var MockPassengersRefs = &PassengersRefs{
	Passenger: []Passenger{
		{
			ObjectKey: "SH1",
			PTC: PassengerPTC{
				Quantity: "1",
				Value:    "ADT",
			},
		},
		{
			ObjectKey: "SH2",
			PTC: PassengerPTC{
				Quantity: "1",
				Value:    "CHD",
			},
		},
	},
}

var MockAirlineOffer = &AirlineOffer{
	nil,
	&TotalAmount{
		DetailCurrencyPrice: &DetailCurrencyPrice{
			Total: &Total{Code: "Rub", Value: 9111.0},
			Details: &DetailsPrice{
				Detail: []*DetailPrice{
					{
						SubTotal: &Total{Value: 3333.0, Code: "RUB"},
					},
				},
			},
			Fees: &Fees{
				Total: &Total{
					Code:  "RUB",
					Value: 555.0,
				}},
			Taxes: &Taxes{
				Total: &Total{
					Code: "Rub", Value: 8888.0,
				},
			},
		},
	},
	nil,
	&PricedOffer{
		OfferPrice: []*OfferPrice{
			{
				"",
				&RequestedDate{
					PriceDetail: &PriceDetail{
						TotalAmount: &TotalAmount{
							DetailCurrencyPrice: &DetailCurrencyPrice{
								Total: &Total{
									Code:  "",
									Value: 6666.0,
								},
								Details: &DetailsPrice{
									Detail: nil,
								},
								Fees: &Fees{
									Total: nil,
								},
								Taxes: &Taxes{
									Total: &Total{
										Code: "Rub", Value: 4444.0,
									},
									Breakdown: &Breakdown{Tax: []*Tax{
										{
											Amount: &Total{
												Code:  "RUB",
												Value: 6666.0,
											},
											TaxCode: "CODE",
										},
									}},
								},
							},
						},
						BaseAmount: &Total{
							Code: "RUB",
						},
						FareFiledIn: nil,
						Taxes:       nil,
						Fees:        nil,
					},
					Associations: []*Associations{
						{
							AssociatedTraveler: &AssociatedTraveler{
								TravelerReferences: "SH1",
							},
							ApplicableFlight: &ApplicableFlight{FlightSegmentReference: []*FlightSegmentReference{
								{
									Reference: "SEG1",
									Cabin: &CabinFlightSegment{
										CabinDesignator: &MockCodeValue,
										MarketingName:   &MockMarketingName,
									},
									ClassOfService: &ClassOfService{
										Code: &Code{
											SeatsLeft: &MockSeatsLeft,
											Value:     &MockCodeValue,
										},
										MarketingName: &MockMarketingName,
									},
									BagDetailAssociation: &BagDetailAssociation{
										CheckedBagReferences: "",
										CarryOnReferences:    "",
									},
								},
								{
									Reference: "SEG1",
									Cabin: &CabinFlightSegment{
										CabinDesignator: &MockCodeValue,
										MarketingName:   &MockMarketingName,
									},
									ClassOfService: &MockClassOfService,
									BagDetailAssociation: &BagDetailAssociation{
										CheckedBagReferences: "",
										CarryOnReferences:    "",
									},
								},
							},
							},
							OtherAssociation: nil,
							Passengers: &Passengers{
								Passenger: []*TravelerType{{
									PTC: &PTC{
										Quantity: 1,
										Value:    "ADT",
									},
									Age:             nil,
									Name:            nil,
									Contacts:        nil,
									Gender:          "",
									PassengerIDInfo: nil,
									ID:              "",
								}, {
									PTC: &PTC{
										Quantity: 1,
										Value:    "CHD",
									},
									Age:             nil,
									Name:            nil,
									Contacts:        nil,
									Gender:          "",
									PassengerIDInfo: nil,
									ID:              "",
								}},
								PassengerReference:  "SH1",
								PassengerReferences: "SH1",
							},
						},
					},
				},
				&FareDetail{
					FareComponent: []*FareComponent{
						{
							"",
							"SH1", // PAX code
							nil,
							&PriceBreakdown{
								Price: &Price{
									BaseAmount: &Total{
										Code:  "RUB",
										Value: MockADTFare,
									},
									FareFiledIn: &FareFiledIn{
										BaseAmount: &BaseAmount{
											Code:  "RUB",
											Value: MockADTFare,
										},
										ExchangeRate: 6666,
									},
									Taxes: &Taxes{
										Total: &Total{
											Code:  "RUB",
											Value: MockADTTax,
										},
										Breakdown: nil,
									},
								},
							},
							nil,
							nil,
						},
						{
							"",
							"SH2", // PAX code
							nil,
							&PriceBreakdown{
								Price: &Price{
									BaseAmount: &Total{
										Code:  "RUB",
										Value: MockCHDFare,
									},
									FareFiledIn: &FareFiledIn{
										BaseAmount: &BaseAmount{
											Code:  "RUB",
											Value: MockCHDFare,
										},
										ExchangeRate: 5555,
									},
									Taxes: &Taxes{
										Total: &Total{
											Code:  "RUB",
											Value: MockCHDTax,
										},
										Breakdown: nil,
									},
								},
							},
							nil,
							nil,
						},
					},
					PriceClassReference: "",
					Remarks:             nil,
				},
			},
		},
	},
}

var MockTravelers = Travelers{
	Traveler: []*Traveler{
		{
			AnonymousTraveler: &TravelerType{
				PTC: &PTC{
					Quantity: 2,
					Value:    "ADT",
				},
				Age:             nil,
				Name:            nil,
				Contacts:        nil,
				Gender:          "",
				PassengerIDInfo: nil,
				ID:              "",
			},
		},
		{
			AnonymousTraveler: &TravelerType{
				PTC: &PTC{
					Quantity: 1,
					Value:    "CHD",
				},
				Age:             nil,
				Name:            nil,
				Contacts:        nil,
				Gender:          "",
				PassengerIDInfo: nil,
				ID:              "",
			},
		},
		{
			AnonymousTraveler: &TravelerType{
				PTC: &PTC{
					Quantity: 1,
					Value:    "INF",
				},
				Age:             nil,
				Name:            nil,
				Contacts:        nil,
				Gender:          "",
				PassengerIDInfo: nil,
				ID:              "",
			},
		},
	},
}

var MockPassengerOrderViewRS = []PassengerOrderViewRS{
	{
		ObjectKey: "SH2",
		PTC:       nil,
		Age:       nil,
		Name:      nil,
		Gender:    "",
		PassengerIDInfo: &PassengerIDInfo{
			PassengerDocument: &PassengerDocument{
				Type:               "",
				ID:                 "800",
				BirthCountry:       "",
				DateOfIssue:        "",
				DateOfExpiration:   "",
				CountryOfResidence: "",
			}},
	},
	{
		ObjectKey: "SH3",
		PTC:       nil,
		Age:       nil,
		Name:      nil,
		Gender:    "",
		PassengerIDInfo: &PassengerIDInfo{
			PassengerDocument: &PassengerDocument{
				Type:               "",
				ID:                 "900",
				BirthCountry:       "",
				DateOfIssue:        "",
				DateOfExpiration:   "",
				CountryOfResidence: "",
			}},
	},
}

var MockPassengerKeys = []string{"SH1", "SH2"}

var MockOrderViewResponse = OrderViewResponse{
	Passengers: &PassengerOrderRS{
		Passenger: MockPassengerOrderViewRS,
	},
}

var MockOrderViewRS = OrderViewRS{
	Response: &MockOrderViewResponse,
}

var MockFlighSegmentList = FlightSegmentList{[]*FlightSegment{
	{
		SegmentKey: "SEG1",
		Departure: &Point{
			AirportCode: "",
			Date:        "",
			Time:        "",
			Terminal:    nil,
		},
		Arrival: &Point{
			AirportCode: "",
			Date:        "",
			Time:        "",
			Terminal:    nil,
		},
		MarketingCarrier: nil,
		OperatingCarrier: nil,
		Equipment:        nil,
		FlightDetail: &FlightDetail{
			FlightDistance: nil,
			FlightDuration: &FlightDuration{Value: "PT1H30M"},
			Stops:          nil,
		},
	},
}}

var MocksAnonymousTravelerList = AnonymousTravelerList{
	AnonymousTraveler: []Passenger{
		{
			ObjectKey: "SH1",
			PTC: PassengerPTC{
				Quantity: "1",
				Value:    "ADT",
			},
		},
		{
			ObjectKey: "SH2",
			PTC: PassengerPTC{
				Quantity: "1",
				Value:    "CHD",
			},
		},
	},
}

var MockAirShoppingEnvelope = Envelope{
	XMLName: xml.Name{},
	Header:  nil,
	Body: &Body{
		AirShoppingRS: &AirShoppingRS{
			Version:               "",
			Document:              nil,
			Success:               nil,
			AirShoppingProcessing: nil,
			OffersGroup: &OffersGroup{
				AirlineOffers: []*AirlineOffers{
					{
						Owner: "",
						AirlineOffer: []*AirlineOffer{
							MockAirlineOffer,
						},
					},
				},
			},
			DataLists: &DataLists{
				AnonymousTravelerList: &MocksAnonymousTravelerList,
				FlightSegmentList:     &MockFlighSegmentList,
			},
			Metadata: nil,
			Errors:   nil,
		},
	},
}

var MockOrderItems = OrderItems{
	ShoppingResponse: nil,
	//OrderItem:        []*OrderItem{&MockOrderItem},
	OfferItem: &MockOfferItem,
}

var MockOfferItem = OfferItem{
	OfferItemID: &OfferID{
		Owner: "S7",
		Value: "UNKNOWN",
	},
	Passengers: nil,
	OfferItemType: &OfferItemType{
		DetailedFlightItem: &DetailedFlightItem{
			OriginDestination: []*OriginDestination{{
				ID:               "",
				Refs:             "SEG1",
				SegmentKey:       "",
				Status:           nil,
				Departure:        nil,
				Arrival:          nil,
				MarketingCarrier: nil,
				OperatingCarrier: nil,
				CalendarDates:    nil,
				Equipment:        nil,
				CabinType:        nil,
				ClassOfService:   nil,
				Flight: &Flight{
					SegmentKey: "SEG1",
					Status:     nil,
					Departure: &Point{
						AirportCode: "",
						Date:        "0001-01-01",
						Time:        "00:00",
						Terminal:    nil,
					},
					Arrival: &Point{
						AirportCode: "",
						Date:        "0001-01-01",
						Time:        "00:00",
						Terminal:    nil,
					},
					MarketingCarrier: &Carrier{
						AirlineID:    "",
						FlightNumber: "",
					},
					OperatingCarrier: &Carrier{
						AirlineID:    "",
						FlightNumber: "",
					},
					CabinType: &CabinType{
						Code: &Code{
							SeatsLeft: nil,
							Value:     &codeValue,
						},
					},
					ClassOfService: &ClassOfService{
						Code: &Code{
							SeatsLeft: nil,
							Value:     &codeValue,
						},
						MarketingName: nil,
					},
				},
			}},
		},
	},
	ApplicableFlight: nil,
}

var MockFlightSegmentList = FlightSegmentList{[]*FlightSegment{
	{
		SegmentKey: "SEG1",
		FlightDetail: &FlightDetail{
			FlightDistance: nil,
			FlightDuration: &FlightDuration{
				Value: "PT1H40M",
			},
			Stops: nil,
		},
	},
}}

var codeValue = "Y"

var MockOrderItem = OrderItem{
	FlightItem: &FlightItem{
		Reference: "",
		OriginDestination: []*OriginDestination{
			{
				ID:               "",
				Refs:             "",
				SegmentKey:       "",
				Status:           nil,
				Departure:        nil,
				Arrival:          nil,
				MarketingCarrier: nil,
				OperatingCarrier: nil,
				CalendarDates:    nil,
				Equipment:        nil,
				CabinType:        nil,
				ClassOfService:   nil,
				Flight: &Flight{
					SegmentKey: "",
					Status:     nil,
					Departure: &Point{
						AirportCode: "",
						Date:        "2006-01-02",
						Time:        "15:04",
						Terminal:    nil,
					},
					Arrival: &Point{
						AirportCode: "",
						Date:        "2006-01-08",
						Time:        "15:04",
						Terminal:    nil,
					},
					MarketingCarrier: nil,
					OperatingCarrier: nil,
					CabinType: &CabinType{
						Code: &Code{
							SeatsLeft: nil,
							Value:     &MockCodeValue,
						},
					},
					ClassOfService: nil,
				},
			},
		},
		FareDetail: nil,
	},

	OrderItemID:  nil,
	Associations: nil,
}

var MockPaxSegmentsAssociations = PaxSegmentsAssociations{SegmentPax: map[string]*SegmentPax{
	"SEG1": {
		ADT: &FlightSegmentReference{
			Reference: "SEG1",
			Cabin: &CabinFlightSegment{
				CabinDesignator: &MockCodeValue,
				MarketingName:   &MockMarketingName,
			},
			ClassOfService: &ClassOfService{
				Code: &Code{
					SeatsLeft: &MockSeatsLeft,
					Value:     &MockCodeValue,
				},
				MarketingName: &MockCodeValue,
			},
			BagDetailAssociation: &BagDetailAssociation{
				CheckedBagReferences: "",
				CarryOnReferences:    "",
			},
		},
		CHD: nil,
		INF: nil,
	},
}}
