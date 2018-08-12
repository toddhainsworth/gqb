package gqb

import "testing"

var specificQuery = NewSelectQuery("person", []string{"name"})
var allQuery = NewSelectQuery("person", []string{})

func TestNewSelectQuery(t *testing.T) {
	if column := specificQuery.Columns[0]; column != "name" {
		t.Errorf("Expected first column to equal \"name\" but got \"%s\"", column)
	}
}

func TestConstruct(t *testing.T) {
	expected := "SELECT name FROM person;"

	if str := specificQuery.Construct(); str != expected {
		t.Errorf("Expected query to equal \"%s\" but got \"%s\"", expected, str)
	}

	expected = "SELECT * FROM person;"

	if str := allQuery.Construct(); str != expected {
		t.Errorf("Expected query to equal \"%s\" but got \"%s\"", expected, str)
	}
}
