package repository

import (
	"encoding/xml"
	"github.com/tmconsulting/s7-golang-sdk/structsV052"
)

func (r *Repository) DemandTickets(logAttributes map[string]string, request *structsV052.Envelope) (*structsV052.Envelope, error) {

	demandTicketsRequestBytes, err := xml.MarshalIndent(request, "  ", "    ")
	if err != nil {
		return nil, err
	}

	s7demandTicketsResponse, err := r.transport.Request("DemandTickets", demandTicketsRequestBytes, logAttributes)
	if err != nil {
		return nil, err
	}

	demandTicketsResponse := structsV052.Envelope{}

	if err := xml.Unmarshal(s7demandTicketsResponse, &demandTicketsResponse); err != nil {
		return nil, err
	}

	return &demandTicketsResponse, nil
}
