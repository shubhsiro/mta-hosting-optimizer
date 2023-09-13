package main

import (
	"Mta-Hosts/handler"
	"Mta-Hosts/models"
	"Mta-Hosts/service"
	"Mta-Hosts/store"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

//func TestIntegrationMain(t *testing.T) {
//	// Create a test server with the router for this test case
//	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		// Initialize your store, service, and handler as in your main function
//		store := store.NewInMemoryStore()
//		store.IpConfigs = []models.IpConfig{
//			{IP: "127.0.0.1", Hostname: "mta-prod-1", Active: true},
//			{IP: "127.0.0.2", Hostname: "mta-prod-1", Active: false},
//			{IP: "127.0.0.3", Hostname: "mta-prod-2", Active: true},
//			{IP: "127.0.0.4", Hostname: "mta-prod-2", Active: true},
//			{IP: "127.0.0.5", Hostname: "mta-prod-2", Active: false},
//			{IP: "127.0.0.6", Hostname: "mta-prod-3", Active: false},
//		}
//
//		ipConfigService := service.NewIpConfigService(store)
//		ipConfigHandler := handler.NewIpConfigHandler(ipConfigService)
//
//		// Handle the request
//		data, err := ipConfigHandler.GetHostnamesWithActiveIPs(w, r)
//		if err != nil {
//			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
//			return
//		}
//
//		// Encode the response data to JSON
//		w.Header().Set("Content-Type", "application/json")
//		if err := json.NewEncoder(w).Encode(data); err != nil {
//			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
//			return
//		}
//	}))
//
//	defer testServer.Close()
//
//	t.Run("ValidInput", func(t *testing.T) {
//		// Send a GET request to the test server
//		resp, err := http.Get(testServer.URL)
//		if err != nil {
//			t.Fatal(err)
//		}
//		defer resp.Body.Close()
//
//		// Check the response status code
//		if resp.StatusCode != http.StatusOK {
//			t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
//		}
//
//		// Parse the JSON response body
//		var response struct {
//			Hostnames []string `json:"hostnames"`
//		}
//		if _ = json.NewDecoder(resp.Body).Decode(&response); err != nil {
//			t.Fatalf("Failed to decode JSON response: %v", err)
//		}
//
//		// Perform assertions on the response data
//		expectedHostnames := []string{"mta-prod-1", "mta-prod-2", "mta-prod-2", "mta-prod-3"}
//		if !reflect.DeepEqual(response.Hostnames, expectedHostnames) {
//			t.Errorf("Expected hostnames %v, got %v", expectedHostnames, response.Hostnames)
//		}
//	})
//}

func TestIntegration(t *testing.T) {
	// Create a new router
	r := mux.NewRouter()

	// Initialize a temporary in-memory store with sample data
	store := store.NewInMemoryStore()
	store.IpConfigs = []models.IpConfig{
		{IP: "127.0.0.1", Hostname: "mta-prod-1", Active: true},
		{IP: "127.0.0.2", Hostname: "mta-prod-1", Active: false},
		{IP: "127.0.0.3", Hostname: "mta-prod-2", Active: true},
		{IP: "127.0.0.4", Hostname: "mta-prod-2", Active: true},
		{IP: "127.0.0.5", Hostname: "mta-prod-2", Active: false},
		{IP: "127.0.0.6", Hostname: "mta-prod-3", Active: false},
	}

	// Create a new service and handler
	ipConfigService := service.NewIpConfigService(store)
	ipConfigHandler := handler.NewIpConfigHandler(ipConfigService)

	// Define routes
	r.HandleFunc("/hostnames", func(w http.ResponseWriter, r *http.Request) {
		_, err := ipConfigHandler.GetHostnamesWithActiveIPs(w, r)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	})

	// Start a test server
	server := httptest.NewServer(r)
	defer server.Close()

	// Define test cases
	testCases := []struct {
		name          string
		url           string
		expectedCode  int
		expectedJSON  []string
		expectedError string
	}{
		{
			name:         "ValidThreshold",
			url:          server.URL + "/hostnames?threshold=1",
			expectedCode: http.StatusOK,
			expectedJSON: []string{"mta-prod-1", "mta-prod-3"},
		},
		{
			name:         "ZeroThreshold",
			url:          server.URL + "/hostnames?threshold=0",
			expectedCode: http.StatusOK,
			expectedJSON: []string{"mta-prod-3"},
		},
	}

	// Run the test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resp, err := http.Get(tc.url)
			assert.NoError(t, err)
			defer resp.Body.Close()

			// Check the HTTP status code
			assert.Equal(t, tc.expectedCode, resp.StatusCode)

			if tc.expectedCode == http.StatusOK {
				// Check the JSON response
				var response struct {
					Hostnames []string `json:"hostnames"`
				}
				err := json.NewDecoder(resp.Body).Decode(&response)
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedJSON, response.Hostnames)
			} else {
				// Check the error message
				var response struct {
					Error string `json:"error"`
				}
				err := json.NewDecoder(resp.Body).Decode(&response)
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedError, response.Error)
			}
		})
	}
}
