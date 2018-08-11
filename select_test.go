package gqb

import "testing"

func TestToString(t *testing.T) {
	query := SelectQuery{}

	if query.ToString() != "" {
		t.Errorf("Expected query to equal \"\" but got \"%s\"", query)
	}
}
