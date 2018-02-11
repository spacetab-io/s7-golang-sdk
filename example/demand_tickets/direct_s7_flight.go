package main

import (
	"encoding/xml"
	"log"

	sdk "github.com/tmconsulting/s7-golang-sdk"
)

func main() {
	request := sdk.MakeAirDocIssueRQ()

	request.Body.AirDocIssueRQ.Party.Sender.AgentUserSender = sdk.MakeAgentUserSender(
		"OVB", "OVB902/1OVB2TR",
		"RU", "00000011Z", "1",
		"OVB902", "1001/1001A", "115")

	request.Body.AirDocIssueRQ.Query.TicketDocQuantity = 1

	request.Body.AirDocIssueRQ.Query.TicketDocInfo = []*sdk.TicketDocInfo{
		&sdk.TicketDocInfo{
			TravelerInfo: &sdk.TravelerInfo{
				Surname: "Trefilov",
				Given:   "Danil",
				PTC: &sdk.PTC{
					Value: sdk.PASSENGER_TYPE_CODE_ADULT,
				},
			},
			BookingReference: &sdk.BookingReference{
				ObjectKey: "RPH1",
				Type: &sdk.Type{
					Code: "14",
				},
				ID:        "SRFLN",
				AirlineID: "S7",
			},
			Payments: &sdk.Payments{
				Payment: []*sdk.Payment{
					&sdk.Payment{
						Type: &sdk.Type{
							Code: "MS",
						},
						Other: &sdk.Other{
							Remarks: &sdk.Remarks{
								Remark: []string{
									"*A*TEXT",
								},
							},
						},
					},
				},
			},
		},
	}

	request.Body.AirDocIssueRQ.Query.DataLists.CheckedBagAllowanceList.CheckedBagAllowance = []*sdk.CheckedBagAllowance{
		&sdk.CheckedBagAllowance{
			ID:        "BG1",
			Reference: "SG1",
			WeightAllowance: &sdk.WeightAllowance{
				MaximumWeight: &sdk.ValueUOM{
					Value: 30,
					UOM:   "Kilogram",
				},
				Descriptions: &sdk.Descriptions{
					Description: []*sdk.Description{
						&sdk.Description{
							Text: "30K",
						},
					},
				},
			},
		},
		&sdk.CheckedBagAllowance{
			ID:        "BG2",
			Reference: "SG2",
			WeightAllowance: &sdk.WeightAllowance{
				MaximumWeight: &sdk.ValueUOM{
					Value: 30,
					UOM:   "Kilogram",
				},
				Descriptions: &sdk.Descriptions{
					Description: []*sdk.Description{
						&sdk.Description{
							Text: "30K",
						},
					},
				},
			},
		},
	}

	request.Body.AirDocIssueRQ.Query.DataLists.FareList.FareGroup = []*sdk.FareGroup{
		&sdk.FareGroup{
			ID: "ETK",
			Fare: &sdk.Fare{
				FareCode: &sdk.FareCode{
					Code: "ANY",
				},
				FareDetail: &sdk.FareDetail{
					FareComponent: []*sdk.FareComponent{
						&sdk.FareComponent{
							ID:        "FC1",
							Reference: "SG1",
							PriceBreakdown: &sdk.PriceBreakdown{
								Price: &sdk.Price{
									BaseAmount: &sdk.BaseAmount{
										Code:  "RUB",
										Value: 27265,
									},
									FareFiledIn: &sdk.FareFiledIn{
										BaseAmount: &sdk.BaseAmount{
											Code:  "EUR",
											Value: 27265,
										},
									},
									Taxes: &sdk.Taxes{
										Breakdown: &sdk.Breakdown{
											Tax: []*sdk.Tax{
												&sdk.Tax{
													Amount: &sdk.Total{
														Code:  "RUB",
														Value: 1536,
													},
													TaxCode: "YR",
												},
												&sdk.Tax{
													Amount: &sdk.Total{
														Code:  "RUB",
														Value: 285,
													},
													TaxCode: "B8",
												},
												&sdk.Tax{
													Amount: &sdk.Total{
														Code:  "RUB",
														Value: 252,
													},
													TaxCode: "WH",
												},
												&sdk.Tax{
													Amount: &sdk.Total{
														Code:  "RUB",
														Value: 252,
													},
													TaxCode: "WH",
												},
												&sdk.Tax{
													Amount: &sdk.Total{
														Code:  "RUB",
														Value: 670,
													},
													TaxCode: "WI",
												},
												&sdk.Tax{
													Amount: &sdk.Total{
														Code:  "RUB",
														Value: 327,
													},
													TaxCode: "WJ",
												},
											},
										},
									},
								},
							},
							FareBasis: &sdk.FareBasis{
								FareBasisCode: &sdk.FareBasisCode{
									Code: "DFLRTCS",
								},
							},
							FareRules: &sdk.FareRules{
								Ticketing: &sdk.Ticketing{
									Endorsements: &sdk.Endorsements{
										Endorsement: []string{
											"s7 only/ref/chng/restr",
										},
									},
								},
							},
						},
						&sdk.FareComponent{
							ID:        "FC2",
							Reference: "SG2",
							FareBasis: &sdk.FareBasis{
								FareBasisCode: &sdk.FareBasisCode{
									Code: "DFLRTCS",
								},
							},
						},
					},
					Remarks: &sdk.Remarks{
						Remark: []string{
							"mow s7 msq45.50s7 mow45.50nuc91.00end roe0.944989",
						},
					},
				},
			},
			FareBasisCode: &sdk.FareBasisCode{
				Code: "DFLRTCS",
			},
		},
	}

	request.Body.AirDocIssueRQ.Query.DataLists.FlightSegmentList.FlightSegment = []*sdk.FlightSegment{
		&sdk.FlightSegment{
			SegmentKey:       "SG1",
			Departure:        sdk.MakePoint("DME", "2018-08-31", "05:40", ""),
			Arrival:          sdk.MakePoint("MSQ", "2018-08-31", "07:00", ""),
			MarketingCarrier: sdk.MakeCarrier("S7", "", "4431"),
			ClassOfService: &sdk.ClassOfService{
				Code: &sdk.Code{
					Value: "D",
				},
				MarketingName: "DFLRTCS",
			},
		},
		&sdk.FlightSegment{
			SegmentKey:       "SG2",
			Departure:        sdk.MakePoint("MSQ", "2018-08-31", "06:45", ""),
			Arrival:          sdk.MakePoint("DME", "2018-08-31", "08:05", ""),
			MarketingCarrier: sdk.MakeCarrier("S7", "", "4434"),
			ClassOfService: &sdk.ClassOfService{
				Code: &sdk.Code{
					Value: "D",
				},
				MarketingName: "DFLRTCS",
			},
		},
	}

	request.Body.AirDocIssueRQ.Query.DataLists.TermsList = &sdk.TermsList{
		ID: "TL1",
		Term: []*sdk.Term{
			&sdk.Term{
				ID:        "T1",
				Reference: "SG1",
				AvailablePeriod: &sdk.AvailablePeriod{
					Earliest: new(sdk.Earliest),
					Latest: &sdk.Latest{
						ShortDate: "2018-10-14",
					},
				},
			},
			&sdk.Term{
				ID:        "T2",
				Reference: "SG2",
				AvailablePeriod: &sdk.AvailablePeriod{
					Earliest: new(sdk.Earliest),
					Latest: &sdk.Latest{
						ShortDate: "2018-10-14",
					},
				},
			},
		},
	}

	output, err := xml.MarshalIndent(request, "", "  ")
	log.Printf("out: \n%s\nerr: %v\n", output, err)
}
