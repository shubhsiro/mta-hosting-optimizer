package service

import (
	"Mta-Hosts/store"
)

type IpConfigService struct {
	Store store.Store
}

func NewIpConfigService(s store.Store) *IpConfigService {
	return &IpConfigService{
		Store: s,
	}
}

// GetHostnamesWithActiveIPs retrieves hostnames with active IP addresses based on the specified threshold.
func (s *IpConfigService) GetHostnamesWithActiveIPs(threshold int) ([]string, error) {
	ipConfigs, err := s.Store.GetAllIpConfigs()
	if err != nil {
		return nil, err
	}

	hostnameActiveCount := make(map[string]int)

	for _, ipConfig := range ipConfigs {
		if ipConfig.Active {
			hostnameActiveCount[ipConfig.Hostname]++
		} else if _, ok := hostnameActiveCount[ipConfig.Hostname]; !ok {
			// Initialize count for hostnames with no active IPs
			hostnameActiveCount[ipConfig.Hostname] = 0
		}
	}

	var result []string

	for hostname, activeCount := range hostnameActiveCount {
		if activeCount <= threshold {
			result = append(result, hostname)
		}
	}

	return result, nil
}
