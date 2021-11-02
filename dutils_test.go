package dutil

import "testing"

func TestMarshalReader(t *testing.T) {
	type payload struct {
		Name string `json:"name"`
	}

	tt := []struct{
		name string
		v interface{}
		EErr Error
	}{
		{
			name: "successful marshal",
			v: payload{
				Name: "james",
			},
			EErr: nil,
		},
		{
			name: "marshal error",
			v: "{'}",
			EErr: nil,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			_, e := MarshalReader(tc.v)

			if tc.EErr == nil {
				if e != nil {
					t.Errorf("expected nil got '%v'", e)
				}
			} else {
				if e == nil {
					t.Errorf("expected not to be nil got '%v'", e)
				}
			}
		})
	}
}
