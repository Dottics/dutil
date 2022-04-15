package dutil

import (
	"fmt"
)

/* ERROR */

type Errors map[string][]string

type Err struct {
	Status int    `json:"status"`
	Errors Errors `json:"errors"`
}

func (e *Err) Error() string {
	return fmt.Sprintf("%s", e.Errors)
}

func (e *Err) recover() *Err {
	return e
}

func NewErr(status int, key string, errors []string) *Err {
	e := &Err{
		Status: status,
		Errors: make(map[string][]string),
	}
	e.Errors[key] = errors
	return e
}

type Error interface {
	Error() string
	recover() *Err
}

func Inst(e Error) *Err {
	if e != nil {
		return e.recover()
	}
	return &Err{
		Status: 500,
		Errors: make(map[string][]string),
	}
}

// ErrorEqual compares whether two errors are equal
//
// err1 and err2 are converted to Err instances then the Error() method
// is called to create a comparable string, the two error strings are
// compared and the boolean result is return.
func ErrorEqual(err1, err2 Error) bool {
	e1 := Inst(err1)
	e2 := Inst(err2)
	if e1.Error() == e2.Error() {
		return true
	}
	return false
}
