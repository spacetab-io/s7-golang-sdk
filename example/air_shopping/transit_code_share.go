package main

import (
	"encoding/xml"
	"log"
	"time"

	sdk "github.com/tmconsulting/s7-golang-sdk"
)

func main() {
	request := sdk.MakeSessionAirShoppingRQ()

	request.Body.AirShoppingRQ.Party.Sender.AgentUserSender = sdk.MakeAgentUserSender(
		"OVB", "OVB902/1OVB2TR",
		"RU", "00000011Z", "1",
		"OVB902", "1001/1001A", "115")

	request.Body.AirShoppingRQ.Travelers = sdk.MakeTravelers(map[string]int32{
		sdk.PassengerTypeCodeAdult: 1,
	})

	request.Body.AirShoppingRQ.CoreQuery.OriginDestinations = &sdk.OriginDestinations{
		OriginDestination: []*sdk.OriginDestination{
			sdk.MakeOriginDestination("SLY", time.Now().Add(10*24*time.Hour), "OVB", time.Time{}, 0, 0),
		},
	}

	request.Body.AirShoppingRQ.Metadata = sdk.MakeResultType(sdk.ResultTypeSmartchoice)

	output, err := xml.MarshalIndent(request, "", "  ")
	log.Printf("out: \n%s\nerr: %v\n", output, err)
}
