package main

import (
	"Mta-Hosts/handler"
	"Mta-Hosts/models"
	"Mta-Hosts/service"
	"Mta-Hosts/store"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	// Create a new router
	r := mux.NewRouter()

	// Initialize the store with sample data
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

	// Start the server
	http.Handle("/", r)
	_ = http.ListenAndServe(":8090", nil)
}
