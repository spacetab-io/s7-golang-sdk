package repositoryImplementation

import (
	"encoding/xml"

	"github.com/tmconsulting/s7-golang-sdk/structsV052"
)

func (r *Repository) Cancel(logAttributes map[string]string, cancelReq *structsV052.Envelope) (*structsV052.Envelope, error) {
	s7CancelBookingRequestBytes, err := xml.MarshalIndent(cancelReq, "  ", "    ")
	if err != nil {

		return nil, err
	}

	// Send cancel booking request to S7 service
	s7CancelBookingResponse, err := r.transport.Request("CancelBooking", s7CancelBookingRequestBytes, logAttributes)
	if err != nil {
		// loggerWorker.Error(err)
		// respond.With(w, r, http.StatusInternalServerError, e.Error(err))
		return nil, err
	}

	cancelBookingResponse := structsV052.Envelope{}
	if err := xml.Unmarshal(s7CancelBookingResponse, &cancelBookingResponse); err != nil {
		// loggerWorker.Error(err)
		// respond.With(w, r, http.StatusInternalServerError, e.Error(err))
		return nil, err
	}

	return &cancelBookingResponse, nil
}
