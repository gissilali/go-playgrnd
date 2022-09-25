package main

import (
	"fmt"
	"time"
)

// tsConvert convert time stamp in "YYYY-MM-DDTHH:MM" format from one time zone to another
func tsConvert(ts, from, to string) (string, error) {
	initialTimezone, err := time.LoadLocation(from)
	if err != nil {
		fmt.Printf("error: %s", err)
	}

	finalTimezone, err := time.LoadLocation(to)

	if err != nil {
		fmt.Printf("error: %s", err)
	}

	parsedTime, err := time.ParseInLocation("2006-01-02T15:04", ts, initialTimezone)

	if err != nil {
		fmt.Printf("error: %s", err)
	}

	convertedTime := parsedTime.In(finalTimezone)

	return convertedTime.Format("2006-01-02T15:04"), nil // FIXME
}

func main() {
	ts := "2021-03-08T19:12"
	out, err := tsConvert(ts, "America/Los_Angeles", "Asia/Jerusalem")
	if err != nil {
		fmt.Printf("error: %s", err)
		return
	}

	fmt.Println(out) // 2021-03-09T05:12
}
