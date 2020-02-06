package repositoryImplementation

import (
	"encoding/xml"

	"github.com/tmconsulting/s7-golang-sdk/structsV052"
)

// CreateRequestXML func create new request to a server
func (r *Repository) Order(logAttributes map[string]string, orderRetrieveReq structsV052.Envelope) (*structsV052.Envelope, error) {

	s7OrderRetrieveRequestBytes, err := xml.MarshalIndent(orderRetrieveReq, "  ", "    ")
	if err != nil {
		return nil, err
	}

	// Send read operation (OrderRetrieve) request to S7 service
	s7OrderRetrieveResponse, err := r.transport.Request("Read", s7OrderRetrieveRequestBytes, logAttributes)
	if err != nil {
		return nil, err
	}

	orderRetrieveResponse := structsV052.Envelope{}
	if err := xml.Unmarshal(s7OrderRetrieveResponse, &orderRetrieveResponse); err != nil {
		return nil, err
	}

	return &orderRetrieveResponse, nil
}
