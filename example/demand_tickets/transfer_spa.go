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
				Surname: "Doe",
				Given:   "John",
				PTC:     sdk.PASSENGER_TYPE_CODE_ADULT,
			},
			BookingReference: &sdk.BookingReference{
				ObjectKey: "RPH1",
				Type: &sdk.Type{
					Code: "14",
				},
				ID:        "SRFFF",
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
			Reference: "SEG1",
			AllowanceDescription: &sdk.AllowanceDescription{
				ApplicableParty: "Traveller",
				Descriptions: &sdk.Descriptions{
					Description: []*sdk.Description{
						&sdk.Description{
							Text: "1PC",
						},
					},
				},
			},
		},
		&sdk.CheckedBagAllowance{
			ID:        "BG2",
			Reference: "SEG2",
			AllowanceDescription: &sdk.AllowanceDescription{
				ApplicableParty: "Traveller",
				Descriptions: &sdk.Descriptions{
					Description: []*sdk.Description{
						&sdk.Description{
							Text: "1PC",
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
							Reference: "SEG1",
							PriceBreakdown: &sdk.PriceBreakdown{
								Price: &sdk.Price{
									BaseAmount: &sdk.BaseAmount{
										Code:  "RUB",
										Value: 11840,
									},
									FareFiledIn: &sdk.FareFiledIn{
										BaseAmount: &sdk.BaseAmount{
											Code:  "RUB",
											Value: 11840,
										},
									},
									Taxes: &sdk.Taxes{
										Total: &sdk.Total{
											Code:  "RUB",
											Value: 3968,
										},
										Breakdown: &sdk.Breakdown{
											Tax: []*sdk.Tax{
												&sdk.Tax{
													Amount: &sdk.Total{
														Code:  "RUB",
														Value: 3968,
													},
													TaxCode: "YR",
												},
											},
										},
									},
								},
							},
							FareBasis: &sdk.FareBasis{
								FareBasisCode: &sdk.FareBasisCode{
									Code: "TFLMUC",
								},
							},
							FareRules: &sdk.FareRules{
								Ticketing: &sdk.Ticketing{
									Endorsements: &sdk.Endorsements{
										Endorsement: []string{
											"s7 only/ref/chng/restr",
											"INCL VAT 3860.55RUB",
										},
									},
								},
							},
						},
						&sdk.FareComponent{
							ID:        "FC2",
							Reference: "SEG2",
							FareBasis: &sdk.FareBasis{
								FareBasisCode: &sdk.FareBasisCode{
									Code: "TFLMUC",
								},
							},
						},
					},
					Remarks: &sdk.Remarks{
						Remark: []string{
							"mow s7 x/muc s7 ham195.76nuc195.76end roe0.944989",
						},
					},
				},
			},
			FareBasisCode: &sdk.FareBasisCode{
				Code: "TFLMUC",
			},
		},
	}

	request.Body.AirDocIssueRQ.Query.DataLists.FlightSegmentList.FlightSegment = []*sdk.FlightSegment{
		&sdk.FlightSegment{
			SegmentKey:       "SEG1",
			Departure:        sdk.MakePoint("DME", "2018-08-02", "08:55", ""),
			Arrival:          sdk.MakePoint("MUC", "2018-08-02", "11:05", ""),
			MarketingCarrier: sdk.MakeCarrier("S7", "", "795"),
			ClassOfService: &sdk.ClassOfService{
				Code: &sdk.Code{
					Value: "T",
				},
				MarketingName: "TFLMUC",
			},
		},
		&sdk.FlightSegment{
			SegmentKey:       "SEG2",
			Departure:        sdk.MakePoint("MUC", "2018-08-02", "15:40", ""),
			Arrival:          sdk.MakePoint("HAM", "2018-08-02", "17:00", ""),
			MarketingCarrier: sdk.MakeCarrier("S7", "", "4725"),
			ClassOfService: &sdk.ClassOfService{
				Code: &sdk.Code{
					Value: "T",
				},
				MarketingName: "TFLMUC",
			},
		},
	}

	request.Body.AirDocIssueRQ.Query.DataLists.TermsList = &sdk.TermsList{
		ID: "TL1",
		Term: []*sdk.Term{
			&sdk.Term{
				ID:        "T1",
				Reference: "SEG1",
				AvailablePeriod: &sdk.AvailablePeriod{
					Earliest: new(sdk.Earliest),
					Latest:   new(sdk.Latest),
				},
			},
			&sdk.Term{
				ID:        "T2",
				Reference: "SEG2",
				AvailablePeriod: &sdk.AvailablePeriod{
					Earliest: new(sdk.Earliest),
					Latest:   new(sdk.Latest),
				},
			},
		},
	}

	output, err := xml.MarshalIndent(request, "", "  ")
	log.Printf("out: \n%s\nerr: %v\n", output, err)
}
