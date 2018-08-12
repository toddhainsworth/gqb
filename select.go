package gqb

import (
	"fmt"
	"strings"
)

// SelectQuery defines a query that selects data from the database
type SelectQuery struct {
	table      string
	columns    []string
	conditions map[string]map[string]string // TODO: condition structs?
}

// NewSelectQuery creates a SelectQuery with the given columns
func NewSelectQuery(table string, columns []string) SelectQuery {
	return SelectQuery{
		table:      table,
		columns:    columns,
		conditions: make(map[string]map[string]string),
	}
}

// String makes it to that SelectQuery satisfies the Stringer interface
// by returning the query as a usable string
func (s SelectQuery) String() string {
	return s.Construct()
}

// Construct puts the columns, conditions and table name together to make the query
func (s SelectQuery) Construct() string {
	return fmt.Sprintf("SELECT %s FROM %s;", s.columnString(), s.table)
}

// Where adds a condition to the where clause relating to the given column
func (s *SelectQuery) Where(column string, condition map[string]string) {
	if _, ok := s.conditions[column]; !ok {
		s.conditions[column] = make(map[string]string)
	}

	for operator, value := range condition {
		//           "name"  "eq"        "Todd"
		s.conditions[column][operator] = value
	}
}

// Table returns the query table
func (s SelectQuery) Table() string {
	return s.table
}

// Table returns the query columns
func (s SelectQuery) Columns() []string {
	return s.columns
}

// Table returns the query conditions
func (s SelectQuery) Conditions() map[string]map[string]string {
	return s.conditions
}

// getColumns gets the SQL friendly column list
func (s SelectQuery) columnString() (columns string) {
	columns = "*"

	if cols := s.Columns(); len(cols) > 0 {
		columns = strings.Join(cols, ",")
	}

	return
}
