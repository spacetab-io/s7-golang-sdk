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

	request.Body.AirShoppingRQ.Parameters = &sdk.Parameters{
		CurrCodes: &sdk.CurrCodes{
			CurrCode: []string{"RUB"},
		},
	}

	request.Body.AirShoppingRQ.Travelers = sdk.MakeTravelers(map[string]int32{
		sdk.PassengerTypeCodeAdult:  2,
		sdk.PassengerTypeCodeChild:  2,
		sdk.PassengerTypeCodeInfant: 2,
	})

	request.Body.AirShoppingRQ.CoreQuery.OriginDestinations = &sdk.OriginDestinations{
		OriginDestination: []*sdk.OriginDestination{
			sdk.MakeOriginDestination("MOW", time.Now().Add(10*24*time.Hour), "SVX", time.Now().Add(10*24*time.Hour), 0, 0),
		},
	}

	request.Body.AirShoppingRQ.Preferences = &sdk.Preferences{
		Preference: []*sdk.Preference{
			sdk.MakePreference("Y", "Exclude"),
		},
	}

	request.Body.AirShoppingRQ.Metadata = sdk.MakeResultType(sdk.ResultTypeSmartchoice)

	output, err := xml.MarshalIndent(request, "", "  ")
	log.Printf("out: \n%s\nerr: %v\n", output, err)
}
