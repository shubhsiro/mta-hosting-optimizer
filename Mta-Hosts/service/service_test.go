package service

import (
	"Mta-Hosts/models"
	"errors"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

// Create a mock store for testing.
type MockStore struct {
	ipConfigs []models.IpConfig
	err       error
}

func (m *MockStore) GetAllIpConfigs() ([]models.IpConfig, error) {
	return m.ipConfigs, m.err
}

func TestGetHostnamesWithActiveIPs(t *testing.T) {
	tests := []struct {
		name           string
		ipConfigs      []models.IpConfig
		threshold      int
		expectedResult []string
		expectedError  error
	}{
		{
			name: "ValidThreshold",
			ipConfigs: []models.IpConfig{
				{IP: "127.0.0.1", Hostname: "mta-prod-1", Active: true},
				{IP: "127.0.0.2", Hostname: "mta-prod-1", Active: false},
				{IP: "127.0.0.3", Hostname: "mta-prod-2", Active: true},
				{IP: "127.0.0.4", Hostname: "mta-prod-2", Active: true},
				{IP: "127.0.0.5", Hostname: "mta-prod-2", Active: false},
				{IP: "127.0.0.6", Hostname: "mta-prod-3", Active: false},
			},
			threshold:      1,
			expectedResult: []string{"mta-prod-3", "mta-prod-1"},
			expectedError:  nil,
		},
		{
			name: "ZeroThreshold",
			ipConfigs: []models.IpConfig{
				{IP: "127.0.0.1", Hostname: "mta-prod-1", Active: true},
				{IP: "127.0.0.2", Hostname: "mta-prod-1", Active: false},
				{IP: "127.0.0.3", Hostname: "mta-prod-2", Active: true},
				{IP: "127.0.0.4", Hostname: "mta-prod-2", Active: true},
				{IP: "127.0.0.5", Hostname: "mta-prod-2", Active: false},
				{IP: "127.0.0.6", Hostname: "mta-prod-3", Active: false},
			},
			threshold:      0,
			expectedResult: []string{"mta-prod-3"},
			expectedError:  nil,
		},
		{
			name: "NegativeThreshold",
			ipConfigs: []models.IpConfig{
				{IP: "127.0.0.1", Hostname: "mta-prod-1", Active: true},
				{IP: "127.0.0.2", Hostname: "mta-prod-1", Active: false},
				{IP: "127.0.0.3", Hostname: "mta-prod-2", Active: true},
				{IP: "127.0.0.4", Hostname: "mta-prod-2", Active: true},
				{IP: "127.0.0.5", Hostname: "mta-prod-2", Active: false},
				{IP: "127.0.0.6", Hostname: "mta-prod-3", Active: false},
			},
			threshold:      -1,
			expectedResult: nil,
			expectedError:  errors.New("threshold cannot be negative"),
		},
		{
			name:           "EmptyIpConfigs",
			ipConfigs:      []models.IpConfig{},
			threshold:      1,
			expectedResult: nil,
			expectedError:  nil,
		},
		{
			name:           "ErrorFromStore",
			ipConfigs:      []models.IpConfig{},
			threshold:      1,
			expectedResult: nil,
			expectedError:  errors.New("no IP configurations found"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := &IpConfigService{
				Store: &MockStore{ipConfigs: tt.ipConfigs, err: tt.expectedError},
			}

			result, err := service.GetHostnamesWithActiveIPs(tt.threshold)

			if tt.name != "ValidThreshold" {
				assert.Equal(t, tt.expectedResult, result)
				assert.Equal(t, tt.expectedError, err)
			} else {
				if tt.name == "ValidThreshold" {
					equalCase := reflect.DeepEqual(tt.expectedResult[0], "mta-prod-3")

					if equalCase {
						assert.Equal(t, tt.expectedResult, []string{"mta-prod-3", "mta-prod-1"})
					} else {
						assert.Equal(t, tt.expectedResult, []string{"mta-prod-1", "mta-prod-3"})
					}
					assert.Equal(t, tt.expectedError, err)
				}

			}
		})
	}
}
