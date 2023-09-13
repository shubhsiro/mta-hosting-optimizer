package service

type Service interface {
	GetHostnamesWithActiveIPs(threshold int) ([]string, error)
}
