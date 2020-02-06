package repositoryImplementation

import (
	"encoding/xml"

	"github.com/tmconsulting/s7-golang-sdk/structsV052"
)

func (r *Repository) ItinReshop(logAttributes map[string]string, request *structsV052.Envelope) (*structsV052.Envelope, error) {

	repriceWithPNRRequestBytes, err := xml.MarshalIndent(request, "  ", "    ")
	if err != nil {
		return nil, err
	}

	s7repriceWithPNRResponse, err := r.transport.Request("Reprice", repriceWithPNRRequestBytes, logAttributes)
	if err != nil {
		return nil, err
	}

	repriceWithPNRResponce := &structsV052.Envelope{}

	err = xml.Unmarshal(s7repriceWithPNRResponse, repriceWithPNRResponce)
	if err != nil {
		return nil, err
	}

	return repriceWithPNRResponce, nil
}
