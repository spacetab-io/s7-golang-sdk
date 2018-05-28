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
		{
			TravelerInfo: &sdk.TravelerInfo{
				Surname: "CONELLY",
				Given:   "LILA",
				PTC:     sdk.PassengerTypeCodeAdult,
			},
			BookingReference: &sdk.BookingReference{
				ObjectKey: "RPH1",
				Type: &sdk.Type{
					Code: "14",
				},
				ID:        "SYZWH",
				AirlineID: "S7",
			},
			Payments: &sdk.Payments{
				Payment: []*sdk.Payment{
					{
						Type: &sdk.Type{
							Code: "MS",
						},
						Other: &sdk.Other{
							Remarks: &sdk.Remarks{
								Remark: []string{
									"*QUW1725",
								},
							},
						},
					},
				},
			},
		},
	}

	request.Body.AirDocIssueRQ.Query.DataLists.CheckedBagAllowanceList.CheckedBagAllowance = []*sdk.CheckedBagAllowance{
		{
			ID:        "BG1",
			Reference: "SEG1",
			AllowanceDescription: &sdk.AllowanceDescription{
				ApplicableParty: "Traveller",
				Descriptions: &sdk.Descriptions{
					Description: []*sdk.Description{
						{
							Text: "0PC",
						},
					},
				},
			},
		},
		{
			ID:        "BG2",
			Reference: "SEG2",
			AllowanceDescription: &sdk.AllowanceDescription{
				ApplicableParty: "Traveller",
				Descriptions: &sdk.Descriptions{
					Description: []*sdk.Description{
						{
							Text: "0PC",
						},
					},
				},
			},
		},
	}

	request.Body.AirDocIssueRQ.Query.DataLists.FareList.FareGroup = []*sdk.FareGroup{
		{
			ID: "ETK",
			Fare: &sdk.Fare{
				FareCode: &sdk.FareCode{
					Code: "ANY",
				},
				FareDetail: &sdk.FareDetail{
					FareComponent: []*sdk.FareComponent{
						{
							ID:        "FC1",
							Reference: "SEG1",
							PriceBreakdown: &sdk.PriceBreakdown{
								Price: &sdk.Price{
									BaseAmount: &sdk.Total{
										Code:  "RUB",
										Value: 4830,
									},
									FareFiledIn: &sdk.FareFiledIn{
										BaseAmount: &sdk.BaseAmount{
											Code:  "EUR",
											Value: 76.00,
										},
										ExchangeRate: 63.5,
									},
									Taxes: &sdk.Taxes{
										Total: &sdk.Total{
											Code:  "RUB",
											Value: 5120,
										},
										Breakdown: &sdk.Breakdown{
											Tax: []*sdk.Tax{
												{
													Amount: &sdk.Total{
														Code:  "RUB",
														Value: 2794,
													},
													TaxCode: "YR",
												},
												{
													Amount: &sdk.Total{
														Code:  "RUB",
														Value: 456,
													},
													TaxCode: "DE",
												},
												{
													Amount: &sdk.Total{
														Code:  "RUB",
														Value: 475,
													},
													TaxCode: "OY",
												},
												{
													Amount: &sdk.Total{
														Code:  "RUB",
														Value: 1395,
													},
													TaxCode: "RA",
												},
											},
										},
									},
								},
							},
							FareBasis: &sdk.FareBasis{
								FareBasisCode: &sdk.FareBasisCode{
									Code: "OBSRT",
								},
							},
							FareRules: &sdk.FareRules{
								Ticketing: &sdk.Ticketing{
									Endorsements: &sdk.Endorsements{
										Endorsement: []string{
											"s7 only/non-ref/chng/restr",
										},
									},
								},
							},
						},
						{
							ID:        "FC2",
							Reference: "SEG2",
							FareBasis: &sdk.FareBasis{
								FareBasisCode: &sdk.FareBasisCode{
									Code: "OBSRT",
								},
							},
						},
					},
					Remarks: &sdk.Remarks{
						Remark: []string{
							"muc s7 mow40.00s7 muc40.00nuc80.00end roe0.944989",
						},
					},
				},
			},
			FareBasisCode: &sdk.FareBasisCode{
				Code: "Empty",
			},
		},
	}

	request.Body.AirDocIssueRQ.Query.DataLists.FlightSegmentList.FlightSegment = []*sdk.FlightSegment{
		{
			SegmentKey:       "SEG1",
			Departure:        sdk.MakePoint("MUC", "2017-11-17", "10:35", ""),
			Arrival:          sdk.MakePoint("DME", "2017-11-17", "15:45", ""),
			MarketingCarrier: sdk.MakeCarrier("S7", "", "796"),
			OperatingCarrier: sdk.MakeCarrier("S7", "", "796"),
			ClassOfService: &sdk.ClassOfService{
				Code: &sdk.Code{
					SeatsLeft: 1,
					Value:     "O",
				},
				MarketingName: "OBSRT",
			},
		},
		{
			SegmentKey:       "SEG2",
			Departure:        sdk.MakePoint("DME", "2017-12-17", "08:30", ""),
			Arrival:          sdk.MakePoint("MUC", "2017-12-17", "09:40", ""),
			MarketingCarrier: sdk.MakeCarrier("S7", "", "795"),
			OperatingCarrier: sdk.MakeCarrier("S7", "", "795"),
			ClassOfService: &sdk.ClassOfService{
				Code: &sdk.Code{
					SeatsLeft: 1,
					Value:     "O",
				},
				MarketingName: "OBSRT",
			},
		},
	}

	request.Body.AirDocIssueRQ.Query.DataLists.InstructionsList = &sdk.InstructionsList{
		Instruction: []*sdk.Instruction{
			{
				ListKey: "CC",
				SpecialBookingInstruction: &sdk.SpecialBookingInstruction{
					Code:       "QUW1725",
					Definition: "10187",
				},
			},
		},
	}

	request.Body.AirDocIssueRQ.Query.DataLists.TermsList = &sdk.TermsList{
		ID: "TL1",
		Term: []*sdk.Term{
			{
				ID:        "T1",
				Reference: "SEG1",
				AvailablePeriod: &sdk.AvailablePeriod{
					Latest: &sdk.Latest{
						ShortDate: "2018-09-08",
					},
				},
			},
			{
				ID:        "T2",
				Reference: "SEG2",
				AvailablePeriod: &sdk.AvailablePeriod{
					Latest: &sdk.Latest{
						ShortDate: "2018-09-08",
					},
				},
			},
		},
	}

	output, err := xml.MarshalIndent(request, "", "  ")
	log.Printf("out: \n%s\nerr: %v\n", output, err)
}
