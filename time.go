package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	fmt.Println(now.Nanosecond())

	utc := time.Date(2022, 8, 14, 20, 30, 10, 10, time.UTC)
	fmt.Println(utc)

	layout := "2006-01-02"
	parse, _ := time.Parse(layout, "2022-08-14")
	fmt.Println(parse)
}
