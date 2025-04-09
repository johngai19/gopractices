package main

import (
	"fmt"
	"time"
)

func datetimes() {
	var now time.Time = time.Now()
	var year int = now.Year()
	var month int = int(now.Month())
	var day int = now.Day()
	var hour int = now.Hour()
	var minute int = now.Minute()
	var second int = now.Second()
	var nanosecond int = now.Nanosecond()
	fmt.Println("Current date and time:")
	fmt.Printf("%04d-%02d-%02d %02d:%02d:%02d.%09d\n", year, month, day, hour, minute, second, nanosecond)
	fmt.Println("Current date and time in RFC3339 format:", now.Format(time.RFC3339))
	fmt.Println("Current date and time in Unix format:", now.Unix())
}
