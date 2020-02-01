package repository

import (
	"github.com/tmconsulting/s7-golang-sdk/connectionClients"
)

// client is S7 client
type Repository struct {
	transport connectionClients.Client
}

// NewStorage Create new storage client for communicating with s7gds
func NewStorage(transport connectionClients.Client) (*Repository, error) {

	repo := &Repository{transport: transport}

	return repo, nil
}
