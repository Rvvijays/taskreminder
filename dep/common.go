package dep

import (
	"fmt"
	"time"
)

func IntTOAPITime(unixTimestamp int64) string {
	t := time.Unix(unixTimestamp, 0)

	rfc3339Time := t.Format(time.RFC3339)

	return rfc3339Time
}

func APIToIntTime(timestamp string) (int64, error) {
	dueDateTime, err := time.Parse(time.RFC3339, timestamp)
	if err != nil {
		fmt.Println("Err", err)
		return 0, err
	}

	return dueDateTime.Unix(), nil
}
