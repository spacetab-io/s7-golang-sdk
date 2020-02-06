package main

import (
	"encoding/xml"
	"io/ioutil"
	"log"

	sdk "github.com/tmconsulting/s7-golang-sdk"
)

func main() {
	// name := "test/response_not_smartchoice.xml"
	// name := "test/response_not_lowfare.xml"
	// name := "test/response_basic_shopping.xml"
	// name := "test/response_3d_agreement.xml"
	// name := "test/response_transit_codeshare.xml"
	// name := "test/response_transfer_spa.xml"
	// name := "test/response_3d_agreement_2.xml"
	// name := "test/response_direct_flight.xml"
	name := "test/response_error.xml"

	data, err := ioutil.ReadFile(name)
	if err != nil {
		log.Fatalln(err)
	}

	env := &sdk.Envelope{
		Body: &sdk.Body{
			AirShoppingRS: new(sdk.AirShoppingRS),
		},
	}

	err = xml.Unmarshal(data, env)
	if err != nil {
		log.Fatalln(err)
	}
}
