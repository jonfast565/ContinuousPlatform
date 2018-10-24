package server

import (
	"../../models/loggingmodel"
	"../../models/persistmodel"
	"../../networking"
	"../../timeutil"
	"../dbhelper"
	"errors"
	"github.com/satori/go.uuid"
	"os/user"
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
	setRequest *persistmodel.KeyValueRequest) error {
	db, err := p.Configuration.GetSqlServerConnection()
	if err != nil {
		return err
	}

	hostname, err := networking.GetMyHostName()
	if err != nil {
		return err
	}

	currentUser, err := user.Current()
	if err != nil {
		return err
	}

	getExistingKeyValueCache := dbhelper.NewSqlStatement().
		Select("dbo.KeyValueCache").
		SimpleWhere("[Key] = @Key AND MachineName = @MachineName").
		AddParameterWithValue("Key", setRequest.Key).
		AddParameterWithValue("MachineName", hostname)
	rows, err := db.RunStatement(getExistingKeyValueCache)
	if err != nil {
		return err
	}

	if len(rows) > 0 {
		insertKeyValueCache := dbhelper.
			NewSqlStatement().
			Update("dbo.KeyValueCache").
			SimpleSet("Value", "@Value").
			SimpleSet("ValueType", "@ValueType").
			SimpleSet("MachineName", "@MachineName").
			SimpleSet("LastModifiedBy", "@LastModifiedBy").
			SimpleSet("LastModifiedDateTime", "@LastModifiedDatetime").
			SimpleWhere("[Key] = @Key AND MachineName = @MachineName").
			AddParameterWithValue("Key", setRequest.Key).
			AddParameterWithValue("Value", setRequest.Value).
			AddParameterWithValue("ValueType", "Binary").
			AddParameterWithValue("MachineName", hostname).
			AddParameterWithValue("LastModifiedBy", currentUser.Username).
			AddParameterWithValue("LastModifiedDateTime", timeutil.GetCurrentSqlTime())
		_, err = db.RunStatement(insertKeyValueCache)
		if err != nil {
			return err
		}
	} else {
		insertKeyValueCache := dbhelper.
			NewSqlStatement().
			Insert("dbo.KeyValueCache").
			Columns("Key", "Value", "ValueType", "MachineName").
			Values("@Key", "@Value", "@ValueType", "@MachineName").
			AddParameterWithValue("Key", setRequest.Key).
			AddParameterWithValue("Value", setRequest.Value).
			AddParameterWithValue("ValueType", "Binary").
			AddParameterWithValue("MachineName", hostname)
		_, err = db.RunStatement(insertKeyValueCache)
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *PersistenceServiceEndpoint) GetKeyValueCache(
	getRequest *persistmodel.KeyRequest) (*persistmodel.KeyValueResult, error) {
	db, err := p.Configuration.GetSqlServerConnection()
	if err != nil {
		return nil, err
	}

	hostname, err := networking.GetMyHostName()
	if err != nil {
		return nil, err
	}
	getExistingKeyValueCache := dbhelper.NewSqlStatement().
		Select("dbo.KeyValueCache").
		SimpleWhere("[Key] = @Key AND MachineName = @MachineName").
		AddParameterWithValue("Key", getRequest.Key).
		AddParameterWithValue("MachineName", hostname)

	rows, err := db.RunStatement(getExistingKeyValueCache)
	if err != nil {
		return nil, err
	}

	if len(rows) == 0 {
		return nil, errors.New("value not found")
	}

	return &persistmodel.KeyValueResult{Value: rows[0]["Value"].([]byte)}, nil
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
