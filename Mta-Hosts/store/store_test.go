package store

import (
	"testing"

	"Mta-Hosts/models"

	"github.com/stretchr/testify/assert"
)

// TestGetAllIpConfigsErrorCases tests error cases for the GetAllIpConfigs method of the InMemoryStore.
func TestGetAllIpConfigsErrorCases(t *testing.T) {
	// Create an empty store
	store := NewInMemoryStore()

	// Test case: No IP configurations in the store
	ipConfigs, err := store.GetAllIpConfigs()
	assert.Nil(t, ipConfigs)                                   // Ensure that ipConfigs is nil
	assert.NotNil(t, err)                                      // Ensure that err is not nil
	assert.Equal(t, "no IP configurations found", err.Error()) // Check the error message
}

// TestGetAllIpConfigsSuccess tests the success case for the GetAllIpConfigs method of the InMemoryStore.
func TestGetAllIpConfigsSuccess(t *testing.T) {
	// Create a store with sample data
	store := &InMemoryStore{
		IpConfigs: []models.IpConfig{
			{IP: "127.0.0.1", Hostname: "mta-prod-1", Active: true},
			{IP: "127.0.0.2", Hostname: "mta-prod-2", Active: false},
		},
	}

	// Test case: IP configurations are present
	ipConfigs, err := store.GetAllIpConfigs()
	assert.NotNil(t, ipConfigs) // Ensure that ipConfigs is not nil
	assert.Nil(t, err)          // Ensure that err is nil

	// Check the number of IP configurations
	assert.Equal(t, 2, len(ipConfigs))
}
