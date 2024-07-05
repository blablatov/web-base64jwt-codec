// Decoder jwt
// Декодировщик jwt

package main

import (
	"bytes"
	"testing"

	"github.com/pkg/errors"
)

func TestDecode(t *testing.T) {

	var tests = []struct {
		jwt []byte
	}{
		{[]byte("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.XbPfbIHMI6arZ3Y922BhjWgQzWXcXNrz0ogtVhfEd2o")},
	}

	var input []byte
	for _, test := range tests {
		if bytes.Equal(test.jwt, input) {
			t.Logf("\n%s\n", test.jwt)
			input = test.jwt
		}

		parts := bytes.Split(test.jwt, []byte("."))

		if len(parts) != 3 {
			errors.Errorf("Expected 3 parts to the JWT, got %d", len(parts))
		}
		d := new(TokenDecoder)
		header, err := d.b64decode(parts[0])
		if err != nil {
			t.Error(err)
		} else {
			t.Log(string(header))
		}
		body, err := d.b64decode(parts[1])
		if err != nil {
			t.Error(err)
		} else {
			t.Log(string(body))
		}
	}
}
