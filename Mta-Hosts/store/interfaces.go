package store

import "Mta-Hosts/models"

type Store interface {
	GetAllIpConfigs() ([]models.IpConfig, error)
}
