package migration

import (
	"encoding/base64"
	"errors"
	"fmt"
	"net/url"

	"google.golang.org/protobuf/proto"
)

// ErrUnkown scheme or host
var ErrUnkown = errors.New("unknown")

// Data extracts data part from URL string
func Data(link string) ([]byte, error) {
	u, err := url.Parse(link)
	if err != nil {
		return nil, err
	}
	if u.Scheme != "otpauth-migration" {
		return nil, fmt.Errorf("scheme %s: %w", u.Scheme, ErrUnkown)
	}
	if u.Host != "offline" {
		return nil, fmt.Errorf("host %s: %w", u.Host, ErrUnkown)
	}
	data := u.Query().Get("data")
	return base64.StdEncoding.DecodeString(data)
}

// Unmarshal otpauth-migration data
func Unmarshal(data []byte) (*Payload, error) {
	var p Payload
	if err := proto.Unmarshal(data, &p); err != nil {
		return nil, err
	}
	return &p, nil
}

// UnmarshalURL decodes otpauth-migration from URL
func UnmarshalURL(link string) (*Payload, error) {
	data, err := Data(link)
	if err != nil {
		return nil, err
	}
	return Unmarshal(data)
}
