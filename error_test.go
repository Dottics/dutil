package dutil

import (
	"testing"
)

func TestNewErr(t *testing.T) {
	xs := []string{"error one", "error two"}
	e := NewErr(400, "error-identifier", xs)

	if len(e.Errors["error-identifier"]) != 2 {
		t.Errorf("expected '%v' got '%v'", 2, len(e.Errors["error-identifier"]))
	}

	if e.Status != 400 {
		t.Errorf("expected %d got %d", 400, e.Status)
	}

	for _, values := range e.Errors {
		for i, v := range xs {
			if values[i] != v {
				t.Errorf("expected '%v' got '%v'", v, values[i])
			}
		}
	}
}

func TestErr_Error(t *testing.T) {
	xs := []string{"error one", "error two"}
	e := NewErr(400, "error-identifier", xs)
	errorString := "map[error-identifier:[error one error two]]"
	if e.Error() != errorString {
		t.Errorf("expected '%v' got '%v'", errorString, e.Error())
	}

}

func TestErr_recover(t *testing.T) {
	e := &Err{
		Status: 405,
		Errors: map[string][]string{
			"identifier": {"error one", "error two"},
		},
	}

	e2 := e.recover()

	if e != e2 {
		t.Errorf("expected %v got %v", e, e2)
	}
}

func TestInst(t *testing.T) {
	f := func(b bool) Error {
		if b {
			return &Err{
				Status: 405,
				Errors: map[string][]string{
					"identifier": {"error one", "error two"},
				},
			}
		}
		return nil
	}

	e1 := f(true)  // *Err
	e2 := f(false) // nil

	e3 := Inst(e1)
	if e3 != e1 {
		t.Errorf("expected '%v' got '%v'", e1, e3)
	}
	if e3.Error() != e1.Error() {
		t.Errorf("expected '%v' got '%v'", e1.Error(), e3.Error())
	}
	if e3.Status != 405 {
		t.Errorf("expected %d got %d", 500, e3.Status)
	}
	if e3.Errors == nil {
		t.Errorf("expected %v got %v", nil, e3.Errors)
	}

	e4 := Inst(e2)
	if e4 == e2 {
		t.Errorf("expected new address '%v' got '%v'", e1, e4)
	}
	if e4.Status != 500 {
		t.Errorf("expected %d got %d", 500, e4.Status)
	}
	if e4.Errors != nil {
		t.Errorf("expected %v got %v", nil, e4.Errors)
	}
}
