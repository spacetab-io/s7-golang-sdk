package repositoryImplementation

import (
	"encoding/xml"

	"github.com/tmconsulting/s7-golang-sdk/structsV052"
)

func (r *Repository) Void(logAttributes map[string]string, voidReq *structsV052.Envelope) (*structsV052.Envelope, error) {
	s7VoidTicketRequestBytes, err := xml.MarshalIndent(voidReq, "  ", "    ")
	if err != nil {
		// loggerWorker.Error(err)
		// respond.With(w, r, http.StatusInternalServerError, e.Error(err))
		return nil, err
	}

	// Send void ticket request to S7 service
	s7VoidTicketResponse, err := r.transport.Request("VoidTicket", s7VoidTicketRequestBytes, logAttributes)
	if err != nil {
		// loggerWorker.Error(err)
		// respond.With(w, r, http.StatusInternalServerError, e.Error(err))
		return nil, err
	}

	voidTicketResponse := structsV052.Envelope{}
	if err := xml.Unmarshal(s7VoidTicketResponse, &voidTicketResponse); err != nil {
		// loggerWorker.Error(err)
		// respond.With(w, r, http.StatusInternalServerError, e.Error(err))
		return nil, err
	}

	return &voidTicketResponse, nil
}
