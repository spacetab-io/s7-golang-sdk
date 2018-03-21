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
				PTC:     sdk.PASSENGER_TYPE_CODE_ADULT,
			},
			BookingReference: &sdk.BookingReference{
				ObjectKey: "RPH1",
				Type: &sdk.Type{
					Code: "14",
				},
				ID:        "SREP5",
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
			ID:        "BG2",
			Reference: "SG1",
			PieceAllowance: &sdk.PieceAllowance{
				ApplicableParty: "Traveller",
				TotalQuantity:   1,
				Descriptions: &sdk.Descriptions{
					Description: []*sdk.Description{
						&sdk.Description{
							Text: "0PC",
						},
					},
				},
				PieceMeasurements: &sdk.PieceMeasurements{},
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
									BaseAmount: &sdk.Total{
										Code:  "RUB",
										Value: 14500,
									},
									FareFiledIn: &sdk.FareFiledIn{
										BaseAmount: &sdk.Title{
											Code:  "RUB",
											Value: 14500,
										},
									},
									Taxes: &sdk.Taxes{
										Breakdown: &sdk.Breakdown{
											Tax: []*sdk.Tax{
												&sdk.Tax{
													Amount: &sdk.Total{
														Code:  "RUB",
														Value: 1500,
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
									Code: "MFLOW",
								},
							},
							FareRules: &sdk.FareRules{
								Ticketing: &sdk.Ticketing{
									Endorsements: &sdk.Endorsements{
										Endorsement: []string{
											"s7 only/ref/chng/restr",
											"INCL VAT 1454.55RUB",
										},
									},
								},
							},
						},
					},
					Remarks: &sdk.Remarks{
						Remark: []string{
							"sly s7 ovb14500rub14500end",
						},
					},
				},
			},
			FareBasisCode: &sdk.FareBasisCode{
				Code: "MFLOW",
			},
		},
	}

	request.Body.AirDocIssueRQ.Query.DataLists.FlightSegmentList.FlightSegment = []*sdk.FlightSegment{
		&sdk.FlightSegment{
			SegmentKey:       "SG1",
			Departure:        sdk.MakePoint("SLY", "2018-08-31", "14:30", ""),
			Arrival:          sdk.MakePoint("OVB", "2018-08-31", "21:30", ""),
			MarketingCarrier: sdk.MakeCarrier("S7", "", "3028"),
			ClassOfService: &sdk.ClassOfService{
				Code: &sdk.Code{
					Value: "M",
				},
				MarketingName: "MFLOW",
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
					Latest:   new(sdk.Latest),
				},
			},
		},
	}

	output, err := xml.MarshalIndent(request, "", "  ")
	log.Printf("out: \n%s\nerr: %v\n", output, err)
}
