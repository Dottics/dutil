package dutil

import (
	"log"
	"os"
	"strings"
)

/* DOTENV */

type Env struct {
	Vars map[string]string
}

// Load takes a filename s and loads the contents of the file as environmental
// variables
func (e *Env) Load(s string) {
	if e.Vars == nil {
		e.Vars = make(map[string]string)
	}

	bs, _ := os.ReadFile(s)
	rows := strings.Split(string(bs), "\n")
	for _, ln := range rows {
		if ln == "" {
			continue
		}

		xln := strings.Split(ln, "=")
		// set struct map
		e.Vars[xln[0]] = xln[1]
		// set env var
		err := os.Setenv(xln[0], xln[1])
		if err != nil {
			log.Panic(err.Error())
		}
	}
}
