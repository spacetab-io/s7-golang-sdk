package repository

import "github.com/tmconsulting/s7-golang-sdk/structsV052"

// Repository it is entity containing storage logic
type Repository interface {
	Void(logAttributes map[string]string, voidReq *structsV052.Envelope) (*structsV052.Envelope, error)
	Cancel(logAttributes map[string]string, cancelReq *structsV052.Envelope) (*structsV052.Envelope, error)
	Order(logAttributes map[string]string, orderRetrieveReq structsV052.Envelope) (*structsV052.Envelope, error)
	SearchFlightJourney(logAttributes map[string]string, request *structsV052.Envelope) (*structsV052.Envelope, error)
	Book(logAttributes map[string]string, request *structsV052.Envelope) (*structsV052.Envelope, error)
	ItinReshop(logAttributes map[string]string, request *structsV052.Envelope) (*structsV052.Envelope, error)
	DemandTickets(logAttributes map[string]string, request *structsV052.Envelope) (*structsV052.Envelope, error)
}
