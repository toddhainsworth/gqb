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

// ToString returns the query as a usable string
func (s SelectQuery) ToString() string {
	return fmt.Sprintf("SELECT %s FROM %s;", s.getColumns(), s.Table)
}

func (s SelectQuery) getColumns() (columns string) {
	columns = "*"

	if len(s.Columns) > 0 {
		columns = strings.Join(s.Columns, ",")
	}

	return
}
