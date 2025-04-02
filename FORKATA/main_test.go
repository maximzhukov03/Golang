package main

import("testing")
func TestCreateTableSQL(t *testing.T) {
	v := &SQLiteGenerator{}
	success := false
	defer func() {
		r := recover()
		if r != nil {
			success = r == "implement me"
		}
		if !success {
			t.Errorf("expected panic with 'implement me', got: %v", r)
		}
	}()
	v.CreateTableSQL(nil)
}

func CreateInsertSQL(t *testing.T) {
	v := &SQLiteGenerator{}
	success := false
	defer func() {
		r := recover()
		if r != nil {
			success = r == "implement me"
		}
		if !success {
			t.Errorf("expected panic with 'implement me', got: %v", r)
		}
	}()
	v.CreateTableSQL(nil)
}

func GenerateFakeUser(t *testing.T) {
	v := &SQLiteGenerator{}
	success := false
	defer func() {
		r := recover()
		if r != nil {
			success = r == "implement me"
		}
		if !success {
			t.Errorf("expected panic with 'implement me', got: %v", r)
		}
	}()
	v.CreateTableSQL(nil)
}

func TableName(t *testing.T) {
	v := &SQLiteGenerator{}
	success := false
	defer func() {
		r := recover()
		if r != nil {
			success = r == "implement me"
		}
		if !success {
			t.Errorf("expected panic with 'implement me', got: %v", r)
		}
	}()
	v.CreateTableSQL(nil)
}