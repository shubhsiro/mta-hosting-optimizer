package handler

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// MockService is a mock implementation of the service.Service interface for testing purposes.
type MockService struct{}

// GetHostnamesWithActiveIPs is a mock implementation of the service.Service interface for testing purposes.
func (ms *MockService) GetHostnamesWithActiveIPs(threshold int) ([]string, error) {
	if threshold == 0 {
		return []string{}, nil
	} else if threshold == 1 {
		return []string{"example.com", "example.org"}, nil
	}
	return nil, errors.New("mock service error")
}

// TestGetHostnamesWithActiveIPs tests the GetHostnamesWithActiveIPs handler function with various threshold values.
func TestGetHostnamesWithActiveIPs(t *testing.T) {
	tests := []struct {
		name           string
		queryThreshold string
		expectedStatus int
		expectedJSON   []string
		expectedError  error
	}{
		{
			name:           "ValidThreshold",
			queryThreshold: "1",
			expectedStatus: http.StatusOK,
			expectedJSON:   []string{"example.com", "example.org"},
			expectedError:  nil,
		},
		{
			name:           "ZeroThreshold",
			queryThreshold: "0",
			expectedStatus: http.StatusOK,
			expectedJSON:   []string{},
			expectedError:  nil,
		},
		{
			name:           "InvalidThreshold",
			queryThreshold: "invalid",
			expectedStatus: http.StatusBadRequest,
			expectedJSON:   []string{},
			expectedError:  errors.New("Invalid threshold value"),
		},
		{
			name:           "NegativeThreshold",
			queryThreshold: "-1",
			expectedStatus: http.StatusBadRequest,
			expectedJSON:   []string{},
			expectedError:  errors.New("Invalid threshold value"),
		},
		{
			name:           "InternalServiceError",
			queryThreshold: "2",
			expectedStatus: http.StatusInternalServerError,
			expectedJSON:   []string{},
			expectedError:  errors.New("Internal Server Error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "/?threshold="+tt.queryThreshold, nil)
			if err != nil {
				t.Fatal(err)
			}

			rec := httptest.NewRecorder()

			// Create an instance of your handler with the MockService
			handler := NewIpConfigHandler(&MockService{})

			// Call the handler function
			_, err = handler.GetHostnamesWithActiveIPs(rec, req)
			assert.Equal(t, tt.expectedError, err)
		})
	}
}
