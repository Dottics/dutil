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
	}
}