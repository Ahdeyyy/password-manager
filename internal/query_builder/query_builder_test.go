package query_builder

import "testing"

func TestSelectStatement(t *testing.T) {

	test1 := "SELECT * FROM"
	test2 := "SELECT id, name FROM"

	query1, _ := NewSqlBuilder().Select().Build()
	query2, _ := NewSqlBuilder().Select("id", "name").Build()

	if query1 != test1 {
		t.Fatalf("expected %s, got %s", test1, query1)
	}

	if query2 != test2 {
		t.Fatalf("expected %s, got %s", test2, query2)
	}

}
