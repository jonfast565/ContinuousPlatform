package server

import (
	"../../models/inframodel"
	"../../models/loggingmodel"
	"../../models/persistmodel"
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

func (p *PersistenceServiceEndpoint) SetKeyValueCache(
	setRequest *persistmodel.KeyValueSetRequest) (*persistmodel.KeyValueResult, error) {
	return nil, nil
}

func (p *PersistenceServiceEndpoint) GetKeyValueCache(
	getRequest *persistmodel.KeyValueGetRequest) (*persistmodel.KeyValueResult, error) {
	return nil, nil
}

func (p *PersistenceServiceEndpoint) GetInfrastructureMetadata() (*inframodel.InfrastructureMetadata, error) {
	return nil, nil
}

func (p *PersistenceServiceEndpoint) SetLogRecord(logRecord *loggingmodel.LogRecord) error {
	return nil
}
