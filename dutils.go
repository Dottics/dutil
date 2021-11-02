package dutil

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
)

func MarshalReader(v interface{}) (io.Reader, Error) {
	xb, err := json.Marshal(v)
	if err != nil {
		log.Println("unexpected marshal error: ", err)
		e := NewErr(500, "marshal", []string{err.Error()})
		return nil, e
	}
	p := bytes.NewReader(xb)
	return p, nil
}
