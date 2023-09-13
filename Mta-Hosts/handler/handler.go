package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"Mta-Hosts/service"
)

type IpConfigHandler struct {
	Service service.Service
}

func NewIpConfigHandler(s service.Service) *IpConfigHandler {
	return &IpConfigHandler{
		Service: s,
	}
}

func (h *IpConfigHandler) GetHostnamesWithActiveIPs(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	thresholdStr := r.URL.Query().Get("threshold")
	threshold, err := strconv.Atoi(thresholdStr)
	if err != nil {
		http.Error(w, "Invalid threshold value", http.StatusBadRequest)
		return nil, errors.New("Invalid threshold value")
	}

	if threshold < 0 {
		http.Error(w, "Invalid threshold value", http.StatusBadRequest)
		return nil, errors.New("Invalid threshold value")
	}

	hostnames, err := h.Service.GetHostnamesWithActiveIPs(threshold)
	if err != nil {
		return nil, errors.New("Internal Server Error")
	}

	response := struct {
		Hostnames []string `json:"hostnames"`
	}{
		Hostnames: hostnames,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return nil, errors.New("Internal Server Error")
	}

	return response, nil
}
