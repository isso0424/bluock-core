package mac_address

import (
	"errors"
	"regexp"
)

type MacAddress string

var r = regexp.MustCompile("	^([0-9A-F]{2}[:-]){5}([0-9A-F]{2})$")

func New(raw string) (MacAddress, error) {
	if !r.Match([]byte(raw)) {
		return "", errors.New("invalid format for mac address.")
	}

	return MacAddress(raw), nil
}
