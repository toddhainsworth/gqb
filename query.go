package gqb

type Query interface {
	String() string
	Construct() string
	Conditions(string) map[string]string
}
