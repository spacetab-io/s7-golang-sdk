package repositoryImplementation

import (
	"encoding/xml"

	"github.com/tmconsulting/s7-golang-sdk/structsV052"
)

func (r *Repository) Book(logAttributes map[string]string, request *structsV052.Envelope) (*structsV052.Envelope, error) {

	requestBytes, err := xml.MarshalIndent(request, "  ", "    ")
	if err != nil {
		return nil, err
	}

	s7BookingResponse, err := r.transport.Request("Book", requestBytes, logAttributes)
	if err != nil {
		return nil, err
	}

	envelope := new(structsV052.Envelope)
	if err := xml.Unmarshal(s7BookingResponse, envelope); err != nil {
		return nil, err
	}

	return envelope, nil
}
