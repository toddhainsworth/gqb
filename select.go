package gqb

// SelectQuery defines a query that selects data from the database
type SelectQuery struct {
	Columns    []string
	Conditions map[string]map[string]string
}

func (s SelectQuery) ToString() string {
	return "" // todo: construct from parts
}
