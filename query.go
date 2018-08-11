package gqb

type Query interface {
	ToString() string
	Conditions(string) map[string]string
}
