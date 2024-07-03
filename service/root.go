package service

import "eCommerce/config"

type Service struct {
	config *config.Config
}

func NewService(config *config.Config) (*Service, error) {
	s := &Service{
		config: config,
	}
	return s, nil
}
