package mdrive

import "testing"

func TestQuery_Parents(t *testing.T) {
	q := NewQuery().Parents().In("id")
	if q.String() != "parents in 'id'" {
		t.Errorf("expected: parents in 'id' got: %s", q.String())
	}
	if q.state != OpState {
		t.Errorf("expected state %d got %d", OpState, q.state)
	}
	q.Reset()
	q.Parents().Eq("some id")
	if q.err == nil {
		t.Errorf("expected error as operand does not exist for term")
	}
}

func TestQuery_Not(t *testing.T) {
	q := NewQuery().Not().Parents().In("id")
	if q.String() != "not parents in 'id'" {
		t.Errorf("expected: not parents in 'id' got: %s", q.String())
	}
	q.Reset()
	q.Parents().In("id").And().Not().Name().Eq("name")
	if q.String() != "parents in 'id' and not name = 'name'" {
		t.Errorf("expected: parents in 'id' and not name = 'name' got: %s", q.String())
	}
	q.Reset()
	q.Not().Not().Parents().In("id")
	if q.err == nil {
		t.Errorf("expected error because of double not")
	}
}

func TestQuery_And(t *testing.T) {
	q := NewQuery().Parents().In("id").And().Name().Eq("name")
	if q.String() != "parents in 'id' and name = 'name'" {
		t.Errorf("wrong and query: %s", q.String())
	}
	q.Reset()
	q.And().Parents().In("id")
	if q.err == nil {
		t.Errorf("expected error as a query cannot start with and")
	}
}

func TestQuery_Or(t *testing.T) {
	q := NewQuery().Parents().In("id").Or().Name().Eq("name")
	if q.String() != "parents in 'id' or name = 'name'" {
		t.Errorf("wrong and query: %s", q.String())
	}
	q.Reset()
	q.Or().Parents().In("id")
	if q.err == nil {
		t.Errorf("expected error as a query cannot start with or")
	}
}

