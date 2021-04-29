package rssi

import "errors"

type RSSI int

func New(raw int) (RSSI, error) {
	if raw >= 0 {
		return 0, errors.New("rssi must be more than zero.")
	}

	return RSSI(raw), nil
}
