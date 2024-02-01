package services

import (
	"coin-ant/api"
	"coin-ant/config"
)

type ServiceChain struct {
}

func NewServiceChain(cfg *config.Config) api.ManagerApi {
	return &ServiceChain{}
}

func (m *ServiceChain) Run() error {
	return nil
}

func (m *ServiceChain) Close() {
}
