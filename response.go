package dutil

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
)

/* ENCODING and DECODING */

// Resp is a structure for the basis of all responses to clients.
// This enforces a consistent structure for all responses.
type Resp struct {
	Status  int         `json:"-"`
	Header  http.Header `json:"-"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Errors  Errors      `json:"errors"`
}

// Respond encodes to JSON and writes the response structure to the client
func (resp *Resp) Respond(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Accept", "*/*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-Token")
	w.Header().Set("Access-Control-Expose-Headers", "X-Token")
	w.Header().Set("Content-Type", "application/json")
	for key, values := range resp.Header {
		w.Header().Set(key, values[0])
	}

	w.WriteHeader(resp.Status)
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		log.Println("Respond encoding error", err)
		http.Error(w, "encoding error", 500)
	}
	log.Printf("-> %v [ %v %v ] <- %d -\n", r.Host, r.Method, r.RequestURI, resp.Status)
}

// Decode decodes the http.Request body to the value pointed to by v.
func Decode(w http.ResponseWriter, r *http.Request, v interface{}) Error {
	err := json.NewDecoder(r.Body).Decode(v)
	if err != nil {
		e := NewErr(400, "decode", []string{"unable to decode data", err.Error()})
		return e
	}
	return nil
}

// QueryParams parses the URL query parameters to a map
func QueryParams(w http.ResponseWriter, r *http.Request) (map[string][]string, error) {
	qp, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		log.Println("queryParams", err)
	}
	return qp, err
}
