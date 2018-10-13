package server

import (
	"../../models"
)

type PersistenceServiceConfiguration struct {
	Port             int    `json:"port"`
	ConnectionString string `json:"connectionString"`
}

type PersistenceServiceEndpoint struct {
	Configuration PersistenceServiceConfiguration
}

func NewPersistenceServiceEndpoint(configuration PersistenceServiceConfiguration) *PersistenceServiceEndpoint {
	result := new(PersistenceServiceEndpoint)
	result.Configuration = configuration
	return result
}

func (p *PersistenceServiceEndpoint) SetKeyValueCache() (*models.KeyValueResult, error) {
	return nil, nil
}

func (p *PersistenceServiceEndpoint) GetKeyValueCache() (*models.KeyValueResult, error) {
	return nil, nil
}

func (p *PersistenceServiceEndpoint) GetInfrastructureMetadata() (*models.InfrastructureMetadata, error) {
	return nil, nil
}

func (p *PersistenceServiceEndpoint) SetLogRecord() error {
	return nil
}
