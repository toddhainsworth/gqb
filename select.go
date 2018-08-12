package gqb

import (
	"strings"
	"fmt"
)

// SelectQuery defines a query that selects data from the database
type SelectQuery struct {
	Table      string
	Columns    []string
	Conditions map[string]map[string]string
}

// NewSelectQuery creates a SelectQuery with the given columns
func NewSelectQuery(table string, columns []string) SelectQuery {
	return SelectQuery { Table: table, Columns: columns }
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

func (s SelectQuery) getColumns() (columns string) {
	columns = "*"

	if len(s.Columns) > 0 {
		columns = strings.Join(s.Columns, ",")
	}

	return
}
