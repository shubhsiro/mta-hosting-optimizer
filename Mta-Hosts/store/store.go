package store

import (
	"errors"

	"Mta-Hosts/models"
)

type InMemoryStore struct {
	IpConfigs []models.IpConfig
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{}
}

// GetAllIpConfigs retrieves all IP configurations stored in memory.
func (s *InMemoryStore) GetAllIpConfigs() ([]models.IpConfig, error) {
	if len(s.IpConfigs) == 0 {
		return nil, errors.New("no IP configurations found")
	}

	return s.IpConfigs, nil
}
