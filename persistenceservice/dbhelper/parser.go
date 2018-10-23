package dbhelper

import (
	"../../constants"
	"../../networking"
	"../../stringutil"
	"database/sql"
	_ "github.com/denisenkom/go-mssqldb"
	"net/url"
	"reflect"
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
	_ = result
	params := statement.GetParameters()
	_ = params
}

type SqlParameter struct {
	Name  string
	Type  reflect.Type
	Value interface{}
}

type UnnamedSqlParameter struct {
	Type  reflect.Type
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
	ValueParams [][]UnnamedSqlParameter
	Params      []SqlParameter
}

func (s *SqlStatement) Insert(intoTable string) *SqlStatement {
	if s.Used {
		panic("Statement cannot be re-used")
	}
	s.TableName = intoTable
	s.Type = InsertStatement
	s.Used = true
	return s
}

func (s *SqlStatement) Update(updateTable string) *SqlStatement {
	if s.Used {
		panic("Statement cannot be re-used")
	}
	s.TableName = updateTable
	s.Type = UpdateStatement
	s.Used = true
	return s
}

func (s *SqlStatement) Select(fromTable string) *SqlStatement {
	if s.Used {
		panic("Statement cannot be re-used")
	}
	s.TableName = fromTable
	s.Type = SelectStatement
	s.Used = true
	return s
}

func (s *SqlStatement) Columns(colNames ...string) {
	s.ColumnNames = colNames
}

func (s *SqlStatement) Values(valuesRows ...interface{}) {
	if len(valuesRows) != len(s.ColumnNames) {
		panic("Invalid number of items specified in a column")
	}
	values := make([]UnnamedSqlParameter, 0)
	for _, value := range valuesRows {
		value := UnnamedSqlParameter{
			Value: value,
			Type:  reflect.TypeOf(value),
		}
		values = append(values, value)
	}
	s.ValueParams = append(s.ValueParams, values)
}

func (s *SqlStatement) getInsertPrefix() string {
	return stringutil.ConcatMultiple(
		"INSERT INTO",
		s.TableName,
		"(",
		stringutil.ConcatDelimitMultiple(", ", "[", "]", s.ColumnNames),
		")",
		"VALUES")
}

func (s *SqlStatement) getInsertRows() string {
	mapperFn := func(parameter UnnamedSqlParameter) string {
		return parameter.Value.(string)
	}
	var outerResult []string
	for _, row := range s.ValueParams {
		var innerResult []string
		for _, col := range row {
			innerResult = append(innerResult, mapperFn(col))
		}
		mappedColumn := stringutil.ConcatDelimitMultiple(", ", "(", ")", innerResult)
		outerResult = append(outerResult, mappedColumn)
	}
	result := stringutil.ConcatMultipleWithSeparator(", ", outerResult...)
	return result
}

func (s *SqlStatement) GetStatementString() string {
	var result string
	if s.Type == InsertStatement {
		result = stringutil.ConcatMultiple(s.getInsertPrefix(), s.getInsertRows())
	} else if s.Type == UpdateStatement {
		panic("Select statement not implemented")
	} else if s.Type == SelectStatement {
		panic("Select statement not implemented")
	} else {
		panic("Invalid or not implemented statement type")
	}
	return result
}

func (s *SqlStatement) AddParameterWithValue(parameter string, value interface{}) {
	s.Params = append(s.Params, SqlParameter{
		Name:  parameter,
		Value: value,
		Type:  reflect.TypeOf(value),
	})
}

func (s *SqlStatement) GetParameters() []SqlParameter {
	return s.Params
}
