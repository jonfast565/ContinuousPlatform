package dbhelper

import (
	"../../constants"
	"../../networking"
	"database/sql"
	_ "github.com/denisenkom/go-mssqldb"
	"net/url"
)

type Database struct {
	DatabaseConnection *sql.DB
}

func InitDatabase(
	username string,
	password string,
	hostname string,
	port int,
	database string) (Database, error) {
	query := url.Values{}
	query.Add(constants.DatabaseKey, database)
	query.Add(constants.AppNameKey, constants.AppName)
	u := &url.URL{
		Scheme:   constants.SqlServerDatabaseDriver,
		User:     url.UserPassword(username, password),
		Host:     networking.GetHostPortCombo(hostname, port),
		RawQuery: query.Encode(),
	}
	db, err := sql.Open(constants.SqlServerDatabaseDriver, u.String())
	return Database{DatabaseConnection: db}, err
}

func (d Database) RunStatement(statement SqlStatement) {
	result := statement.GetStatementString()
}

type SqlParameter struct {
	Name  string
	Value interface{}
}

type StatementType int

const (
	InsertStatement StatementType = 1
	UpdateStatement StatementType = 2
	DeleteStatement StatementType = 3
	SelectStatement StatementType = 4
)

type SqlStatement struct {
	Used        bool
	Type        StatementType
	TableName   string
	ColumnNames []string
	ValuesNames []interface{}
	Params      []SqlParameter
}

func (s SqlStatement) Insert(fromTable string) SqlStatement {
	if s.Used {
		panic("Statement cannot be re-used")
	}
	s.TableName = fromTable
	s.Type = InsertStatement
	s.Used = true
	return s
}

func (s SqlStatement) Columns(colNames ...string) {
	s.ColumnNames = colNames
}

func (s SqlStatement) Values(valuesRows ...interface{}) {

}

func (s SqlStatement) GetStatementString() string {

}

func (s SqlStatement) GetParameters() {

}
