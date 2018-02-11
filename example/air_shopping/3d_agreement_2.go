package main

import (
	"encoding/xml"
	"log"

	sdk "github.com/tmconsulting/s7-golang-sdk"
)

func main() {
	request := sdk.MakeSessionAirShoppingRQ()

	request.Body.AirShoppingRQ.Party.Sender.AgentUserSender = sdk.MakeAgentUserSender(
		"OVB", "OVB902/1OVB2TR",
		"RU", "00000011Z", "1",
		"OVB902", "1001/1001A", "115")

	request.Body.AirShoppingRQ.CoreQuery.ShoppingResponseIDs = sdk.MakeShoppingResponseIDs("S7",
		"yeZ0QbBbSQGvOmScCIdNE27r8/0CR4HXWVB8SHOaG3wNYmS41/fjOIe0mFn17QEiqZyQAR3C00N/10/1")

	request.Body.AirShoppingRQ.Qualifiers = &sdk.Qualifiers{
		Qualifier: []*sdk.Qualifier{
			sdk.MakeQualifier("QUW1725", "10187"),
		},
	}

	request.Body.AirShoppingRQ.Metadata = sdk.MakeResultType(sdk.RESULT_TYPE_SMARTCHOICE)

	output, err := xml.MarshalIndent(request, "", "  ")
	log.Printf("out: \n%s\nerr: %v\n", output, err)
}
