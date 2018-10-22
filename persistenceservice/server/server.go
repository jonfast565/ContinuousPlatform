package server

import (
	"../../models/inframodel"
	"../../models/loggingmodel"
	"../../models/persistmodel"
	"../dbhelper"
)

type PersistenceServiceConfiguration struct {
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Hostname string `json:"hostname"`
	DbPort   int    `json:"dbPort"`
	Database string `json:"database"`
}

func (c *PersistenceServiceConfiguration) GetSqlServerConnection() (dbhelper.Database, error) {
	return dbhelper.InitDatabase(c.Username, c.Password, c.Hostname, c.DbPort, c.Database)
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
	setRequest *persistmodel.KeyValueRequest) (*persistmodel.KeyValueResult, error) {
	/*
		db, err := p.Configuration.GetSqlServerConnection()
		if err != nil {
			return nil, err
		}

		hostname, err := networking.GetMyHostName()
		if err != nil {
			return nil, err
		}
	*/

	return nil, nil
}

func (p *PersistenceServiceEndpoint) GetKeyValueCache(
	getRequest *persistmodel.KeyValueRequest) (*persistmodel.KeyValueResult, error) {
	return nil, nil
}

func (p *PersistenceServiceEndpoint) GetInfrastructureMetadata() (*inframodel.InfrastructureMetadata, error) {
	return nil, nil
}

func (p *PersistenceServiceEndpoint) SetLogRecord(logRecord *loggingmodel.LogRecord) error {
	return nil
}
