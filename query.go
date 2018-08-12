package gqb

type Query interface {
	String() string
	Construct() string
	Table() string
	Columns() []string
	Conditions() map[string]map[string]string
}
