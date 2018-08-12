package gqb

import (
	"reflect"
	"testing"
)

var specificQuery Query = NewSelectQuery("person", []string{"name", "age"})
var allQuery Query = NewSelectQuery("person", []string{})

func TestNewSelectQuery(t *testing.T) {
	if column := specificQuery.Columns()[0]; column != "name" {
		t.Errorf("Expected first column to equal \"name\" but got \"%s\"", column)
	}
}

func TestConstruct(t *testing.T) {
	expected := "SELECT name,age FROM person;"

	if str := specificQuery.Construct(); str != expected {
		t.Errorf("Expected query to equal \"%s\" but got \"%s\"", expected, str)
	}

	expected = "SELECT * FROM person;"

	if str := allQuery.Construct(); str != expected {
		t.Errorf("Expected query to equal \"%s\" but got \"%s\"", expected, str)
	}
}

func TestWhere(t *testing.T) {
	expectedConditions := map[string]map[string]string{
		"name": map[string]string{
			"eq": "Todd",
		},
		"gender": map[string]string{
			"eq":  "male",
			"neq": "female",
		},
	}
	query := NewSelectQuery("person", []string{"name", "age", "gender"})
	query.Where("name", map[string]string{"eq": "Todd"})
	query.Where("gender", map[string]string{"eq": "male"})
	query.Where("gender", map[string]string{"neq": "female"})

	if conditions := query.conditions; !reflect.DeepEqual(conditions, expectedConditions) {
		t.Errorf("Expected conditions to equal: %v but got %v", expectedConditions, conditions)
	}
}
