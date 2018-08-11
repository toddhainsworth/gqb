package gqb

import "testing"

var query = NewSelectQuery("person", []string{"name"})
var allQuery = NewSelectQuery("person", []string{})

func TestNewSelectQuery(t *testing.T) {
	if column := query.Columns[0]; column != "name" {
		t.Errorf("Expected first column to equal \"name\" but got \"%s\"", column)
	}
}

func TestToString(t *testing.T) {
	expected := "SELECT name FROM person;"

	if str := query.ToString(); str != expected {
		t.Errorf("Expected query to equal \"%s\" but got \"%s\"", expected, str)
	}

	expected = "SELECT * FROM person;"

	if str := allQuery.ToString(); str != expected {
		t.Errorf("Expected query to equal \"%s\" but got \"%s\"", expected, str)
	}
}
