package contents

import (
	"fmt"
	"time"
)

func ShowTime() {
	t := time.Now()

	fmt.Print(t, "\n")

	fmt.Print(t.Year(), "\n")
	fmt.Print(t.Month(), "\n")
	fmt.Print(t.Day(), "\n")

	year := t.Year()
	month := t.Month()
	day := t.Day()
	fmt.Printf("Hoy es: %d/%d/%d", day, month, year)

}

func TimeFormat() {
	t := time.Now()
	fmt.Println("ANSIC: \t\t", t.Format(time.ANSIC))
	fmt.Println("RFC3339: \t", t.Format(time.RFC3339))
	fmt.Println("UnixDate: \t", t.Format(time.UnixDate))
	fmt.Println("Kitchen: \t", t.Format(time.Kitchen))

	println()

	// Mon Jan 2 15:04:05 MST 2006
	fmt.Println("CUSTOM1: \t", t.Format("Mon Jan 2 15:04:05 MST 2006"))
	fmt.Println("CUSTOM2: \t", t.Format("MST Jan 2 Mon 2006 15:04:05"))
	fmt.Println("CUSTOM3: \t", t.Format("Monday, January 2 in the year 2006"))

	println()

	startDate := time.Date(2018, 12, 25, 15, 35, 00, 00, time.UTC)
	fmt.Println("DEFAULT: \t", startDate)
	fmt.Println("CUSTOM: \t", startDate.Format("Mon Jan 2 15:04:05 MST 2006"))

}
