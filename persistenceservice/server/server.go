package server

import (
	"../../models/inframodel"
	"../../models/loggingmodel"
	"../../models/persistmodel"
	"../../networking"
	"../../timeutil"
	"../dbhelper"
	"github.com/satori/go.uuid"
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
	db, err := p.Configuration.GetSqlServerConnection()
	if err != nil {
		return nil, err
	}

	hostname, err := networking.GetMyHostName()
	if err != nil {
		return nil, err
	}

	insertKeyValueCache := dbhelper.
		NewSqlStatement().
		Insert("dbo.KeyValueCache").
		Columns("Key", "Value", "ValueType", "MachineName").
		Values("@Key", "@Value", "@ValueType", "@MachineName").
		AddParameterWithValue("Key", setRequest.Key).
		AddParameterWithValue("Value", setRequest.Value).
		AddParameterWithValue("ValueType", "Binary").
		AddParameterWithValue("MachineName", hostname)
	db.RunStatement(insertKeyValueCache)
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
	db, err := p.Configuration.GetSqlServerConnection()
	if err != nil {
		return err
	}

	hostname, err := networking.GetMyHostName()
	if err != nil {
		return err
	}

	uid, err := uuid.NewV4()
	if err != nil {
		return err
	}

	currentTime := timeutil.GetCurrentSqlTime()
	insertKeyValueCache := dbhelper.
		NewSqlStatement().
		Insert("dbo.Logs").
		Columns("LogId", "Date", "MachineName", "ApplicationName", "LogLevel", "Message").
		Values("@LogId", "@Date", "@MachineName", "@ApplicationName", "@LogLevel", "@Message").
		AddParameterWithValue("LogId", uid.String()).
		AddParameterWithValue("Date", currentTime).
		AddParameterWithValue("MachineName", hostname).
		AddParameterWithValue("ApplicationName", logRecord.ApplicationName).
		AddParameterWithValue("LogLevel", logRecord.LogLevel).
		AddParameterWithValue("Message", logRecord.Message)

	db.RunStatement(insertKeyValueCache)
	return nil
}
