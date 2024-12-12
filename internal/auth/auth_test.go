package auth

import (
	"net/http"
	"reflect"
	"testing"
)

func TestGetAPIKey(t *testing.T) {

	validHeader := http.Header{}
	validHeader.Set("Authorization", "ApiKey someapikeytext")
	invalidHeader := http.Header{}
	invalidHeader.Set("Authorization", "ApiKey othertext andsomemore")

	type returnValue struct {
		key string
		err error
	}

	tests := map[string]struct {
		input http.Header
		want  returnValue
	}{
		"Valid Key":      {input: validHeader, want: returnValue{key: "someapikeytext", err: nil}},
		"Missing Key":    {input: http.Header{}, want: returnValue{key: "", err: ErrNoAuthHeaderIncluded}},
		"Invalid Header": {input: invalidHeader, want: returnValue{key: "", err: ErrMalformedHeader}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {

			key, err := GetAPIKey(tc.input)
			got := returnValue{key: key, err: err}

			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("%s expected: %v, got: %v", name, tc.want, got)
			}

		})

	}

}
