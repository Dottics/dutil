package dutil

import (
	"os"
	"testing"
)

func TestEnv_Load(t *testing.T) {
	e := Env{}
	if e.Vars != nil {
		t.Errorf("expected nil got %v", e.Vars)
	}

	e.Load("./testdata/.env.testfile")

	myEnvVar := os.Getenv("MY_ENV_VAR")
	myInt := os.Getenv("MY_INT")
	if myEnvVar != "this is my env var" {
		t.Errorf("expected '%v' got '%v'", "this is my env var", myEnvVar)
	}
	if myInt != "1234" {
		t.Errorf("expected '%v' got '%v'", "1234", myInt)
	}

	if e.Vars["MY_ENV_VAR"] != "this is my env var" {
		t.Errorf("expected '%v' got '%v'", "this is my env var", e.Vars["MY_ENV_VAR"])
	}
	if e.Vars["MY_INT"] != "1234" {
		t.Errorf("expected '%v' got '%v'", "1234", e.Vars["MY_INT"])
	}
}
