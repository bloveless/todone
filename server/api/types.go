package api

import (
	"fmt"
	"time"
)

type JSONTime time.Time

func (t JSONTime) MarshalJSON() ([]byte, error) {
	if time.Time(t).IsZero() {
		return []byte("null"), nil
	}
	fTime := fmt.Sprintf("\"%s\"", time.Time(t).Format(time.RFC3339))
	return []byte(fTime), nil
}
