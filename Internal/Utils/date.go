package utils

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type DateSlash struct {
	time.Time
}

const dateRef = "2006/01/02"

func (d *DateSlash) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	t, err := time.Parse(dateRef, s)
	if err != nil {
		return fmt.Errorf("invalid date format , expect YYYY/MM/DD: %w", err)
	}
	d.Time = t
	return nil
}

func (d DateSlash) MarshalJSON() ([]byte, error) {
	formmated := d.Format(dateRef)
	return json.Marshal(formmated)
}
