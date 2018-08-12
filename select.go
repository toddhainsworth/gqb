package gqb

import (
	"strings"
	"fmt"
)

// SelectQuery defines a query that selects data from the database
type SelectQuery struct {
	Table      string
	Columns    []string
	Conditions map[string]map[string]string // TODO: condition structs?
}

// NewSelectQuery creates a SelectQuery with the given columns
func NewSelectQuery(table string, columns []string) SelectQuery {
	return SelectQuery {
		Table: table,
		Columns: columns,
		Conditions: make(map[string]map[string]string),
	}
}

// String makes it to that SelectQuery satisfies the Stringer interface
// by returning the query as a usable string
func (s SelectQuery) String() string {
	return s.Construct()
}

// Construct puts the columns, conditions and table name together to make the query
func (s SelectQuery) Construct() string {
	return fmt.Sprintf("SELECT %s FROM %s;", s.getColumns(), s.Table)
}

// Where adds a condition to the where clause relating to the given column
func (s *SelectQuery) Where(column string, condition map[string]string) {
	_, ok := s.Conditions[column]

	if !ok {
		s.Conditions[column] = make(map[string]string)
	}

	for operator, value := range condition {
		//           "name"  "eq"        "Todd"
		s.Conditions[column][operator] = value
	}
}

// getColumns gets the SQL friendly column list
func (s SelectQuery) getColumns() (columns string) {
	columns = "*"

	if len(s.Columns) > 0 {
		columns = strings.Join(s.Columns, ",")
	}

	return
}
