// Decoder jwt
// Декодировщик jwt

package main

import (
	"bytes"
	"encoding/base64"

	"github.com/pkg/errors"
)

func (d *TokenDecoder) Decode(input []byte) (header []byte, body []byte, err error) {
	parts := bytes.Split(input, []byte("."))
	if len(parts) != 3 {
		return nil, nil, errors.Errorf("Expected 3 parts to the JWT, got %d", len(parts))
	}
	header, err = d.b64decode(parts[0])
	if err != nil {
		return nil, nil, err
	}
	body, err = d.b64decode(parts[1])
	if err != nil {
		return nil, nil, err
	}
	return header, body, nil
}

// TokenDecoder decodes tokens
type TokenDecoder struct{}

func (d *TokenDecoder) b64decode(data []byte) ([]byte, error) {
	var encodings = []*base64.Encoding{
		base64.URLEncoding,
		base64.StdEncoding,
		base64.RawURLEncoding,
		base64.RawStdEncoding,
	}
	for i, enc := range encodings {
		dst := make([]byte, enc.DecodedLen(len(data)))
		l, err := enc.Decode(dst, data)
		if err != nil && i == len(encodings)-1 {
			return nil, errors.Wrap(err, "No encodings passed")
		}
		if len(dst) > 0 && l > 0 {
			if dst[l-1] != byte('}') {
				continue
			}
			return dst[:l], nil
		}
	}
	return nil, errors.New("Could not decode input")
}
