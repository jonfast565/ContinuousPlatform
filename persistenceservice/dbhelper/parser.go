package dbhelper

import (
	"../../constants"
	"../../logging"
	"../../networking"
	"../../stringutil"
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"net/url"
	"reflect"
	"strconv"
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
	if err != nil {
		panic(err)
	}

	return Database{DatabaseConnection: db}, err
}

func paramsToInterface(args []sql.NamedArg) []interface{} {
	items := make([]interface{}, len(args))
	for i, v := range args {
		items[i] = v
	}
	return items
}

func (d Database) RunStatement(statement SqlStatement) ([]map[string]interface{}, error) {
	result := statement.GetStatementString()
	logging.LogInfoMultiline("Executing SQL Statement:", result)

	statement.LogParameters()
	params := statement.GetNamedParameters()
	interfaceType := paramsToInterface(params)

	stmt, err := d.DatabaseConnection.Prepare(result)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(interfaceType...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	counter := 1
	resultMaps := make([]map[string]interface{}, 0)
	for rows.Next() {
		columns := make([]interface{}, len(cols))
		columnPointers := make([]interface{}, len(cols))

		for i := range columns {
			columnPointers[i] = &columns[i]
		}

		if err := rows.Scan(columnPointers...); err != nil {
			panic(err)
		}

		m := make(map[string]interface{})
		for i, colName := range cols {
			val := columnPointers[i].(*interface{})
			m[colName] = *val
		}

		logging.LogMapPretty("Result Row #"+strconv.Itoa(counter), m)
		resultMaps = append(resultMaps, m)
		counter++
	}

	return resultMaps, nil
}

type SqlParameter struct {
	Name  string
	Type  reflect.Type
	Value interface{}
}

func (s SqlParameter) GetNamedParameter() sql.NamedArg {
	return sql.Named(s.Name, s.Value)
}

func (s SqlParameter) GetNamedOutParameter(dest interface{}) sql.NamedArg {
	// dest should be a pointer here
	return sql.Named(s.Name, sql.Out{Dest: dest})
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
	Used           bool
	Type           StatementType
	TableName      string
	ColumnNames    []string
	ValueParams    [][]UnnamedSqlParameter
	Params         []SqlParameter
	WhereStatement string
	SetStatements  []string
}

func NewSqlStatement() SqlStatement {
	return SqlStatement{}
}

func (s SqlStatement) Insert(intoTable string) SqlStatement {
	if s.Used {
		panic("Statement cannot be re-used")
	}
	s.TableName = intoTable
	s.Type = InsertStatement
	s.Used = true
	return s
}

func (s SqlStatement) Update(updateTable string) SqlStatement {
	if s.Used {
		panic("Statement cannot be re-used")
	}
	s.TableName = updateTable
	s.Type = UpdateStatement
	s.Used = true
	return s
}

func (s SqlStatement) Select(fromTable string) SqlStatement {
	if s.Used {
		panic("Statement cannot be re-used")
	}
	s.TableName = fromTable
	s.Type = SelectStatement
	s.Used = true
	return s
}

func (s SqlStatement) SimpleWhere(expr string) SqlStatement {
	if s.Type != SelectStatement && s.Type != UpdateStatement {
		panic("Cannot use WHERE on this statement")
	}
	s.WhereStatement = expr
	return s
}

func (s SqlStatement) SimpleSet(columnExpr string, valueExpr string) SqlStatement {
	if s.Type != UpdateStatement {
		panic("Cannot use SET on this statement")
	}
	s.SetStatements = append(s.SetStatements, stringutil.ConcatMultiple(columnExpr, "=", valueExpr))
	return s
}

func (s SqlStatement) Columns(colNames ...string) SqlStatement {
	s.ColumnNames = colNames
	return s
}

func (s SqlStatement) Values(valuesRows ...interface{}) SqlStatement {
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
	return s
}

func (s *SqlStatement) getInsertStatement() string {
	prefix := stringutil.ConcatMultiple(
		"INSERT INTO",
		s.TableName,
		"(",
		stringutil.ConcatDelimitMultiple(", ", "[", "]", s.ColumnNames),
		")",
		"VALUES")
	mapperFn := func(parameter UnnamedSqlParameter) string {
		return parameter.Value.(string)
	}
	var outerResult []string
	for _, row := range s.ValueParams {
		var innerResult []string
		for _, col := range row {
			innerResult = append(innerResult, mapperFn(col))
		}
		mappedColumn := stringutil.ConcatMultiple("(",
			stringutil.ConcatMultipleWithSeparator(", ", innerResult...),
			")")
		outerResult = append(outerResult, mappedColumn)
	}
	result := stringutil.ConcatMultipleWithSeparator(", ", outerResult...)
	result2 := stringutil.ConcatMultiple(prefix, result)
	return result2
}

func (s *SqlStatement) getUpdateStatement() string {
	result := stringutil.ConcatMultiple(
		"UPDATE",
		s.TableName,
		"SET")
	result2 := stringutil.ConcatMultipleWithSeparator(", ", s.SetStatements...)
	result3 := stringutil.ConcatMultiple("WHERE",
		s.WhereStatement)
	finalResult := stringutil.ConcatMultiple(result, result2, result3)
	return finalResult
}

func (s *SqlStatement) getSelectStatement() string {
	result := stringutil.ConcatMultiple(
		"SELECT * FROM",
		s.TableName,
		"WHERE",
		s.WhereStatement)
	return result
}

func (s *SqlStatement) GetStatementString() string {
	var result string
	if s.Type == InsertStatement {
		result = s.getInsertStatement()
	} else if s.Type == UpdateStatement {
		result = s.getUpdateStatement()
	} else if s.Type == SelectStatement {
		result = s.getSelectStatement()
	} else {
		panic("Invalid or not implemented statement type")
	}
	return result
}

func (s SqlStatement) LogParameters() {
	result := make([]string, 0)
	result = append(result, "SQL Params: ")
	for _, param := range s.Params {
		result = append(result, param.Name+" -> "+fmt.Sprintf("%v", param.Value))
	}
	logging.LogInfoMultiline(result...)
}

func (s SqlStatement) AddParameterWithValue(parameter string, value interface{}) SqlStatement {
	s.Params = append(s.Params, SqlParameter{
		Name:  parameter,
		Value: value,
		Type:  reflect.TypeOf(value),
	})
	return s
}

func (s SqlStatement) GetNamedParameters() []sql.NamedArg {
	args := make([]sql.NamedArg, 0)
	for _, param := range s.Params {
		args = append(args, param.GetNamedParameter())
	}
	return args
}
