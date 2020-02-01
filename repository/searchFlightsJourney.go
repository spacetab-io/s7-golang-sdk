package repository

import (
	"encoding/xml"
	"github.com/tmconsulting/s7-golang-sdk/structsV052"
)

// Description: https://s7airlines.atlassian.net/wiki/spaces/GAAPI/pages/436764970/searchFlightsJourney+operation
//
func (r *Repository) SearchFlightJourney(logAttributes map[string]string, request *structsV052.Envelope) (*structsV052.Envelope, error) {

	requestBytes, err := xml.MarshalIndent(request, "  ", "    ")
	if err != nil {
		return nil, err
	}

	response, err := r.transport.Request("SearchFlightsJourney", requestBytes, logAttributes)
	if err != nil {
		return nil, err
	}

	envelope := new(structsV052.Envelope)

	if err := xml.Unmarshal(response, envelope); err != nil {
		return nil, err
	}

	return envelope, nil
}
