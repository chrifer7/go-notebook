package contents

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
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

const (
	layoutISO = "2006-01-02"
	layoutUS  = "January 2, 2006"
	layoutEU  = "2 January, 2006"
)

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

	println()

	fmt.Println("CUST. ISO: \t", startDate.Format(layoutISO))
	fmt.Println("CUST. US: \t", startDate.Format(layoutUS))
	fmt.Println("CUST. EU: \t", startDate.Format(layoutEU))

}

func TimeSpans() {
	first := time.Now()
	fmt.Printf("It is currently %v\n", first.Format("15:04:05"))
	//time.Sleep(1000000000)
	time.Sleep(time.Second)
	second := time.Now()
	fmt.Printf("It is now %v\n", second.Format("15:04:05"))

	println()

	today := time.Now()
	fmt.Printf("It is currently %v\n", today.Format("Moday, Jan 2 2006"))
	startDate := time.Date(2018, 07, 04, 9, 0, 0, 0, time.UTC)
	elapsed := time.Since(startDate)

	fmt.Printf("%v has passed since %v\n", elapsed, startDate.Format("Moday, Jan 2 2006"))
	fmt.Printf("Hours: %.0f Minutes: %.0f Seconds: %.0f\n", elapsed.Hours(), elapsed.Minutes(), elapsed.Seconds())

	println()

	today = time.Now() //already declared
	future := today.AddDate(0, 6, 0)
	fmt.Printf("In six months it will be %v\n", future.Format("Monday, January 2 2006"))
	past := today.AddDate(0, -6, 0)
	fmt.Printf("Six months ago it was %v\n", past.Format("Monday, January 2 2006"))

	println()

	today = time.Now() //already declared
	future2 := today.Add(6 * time.Hour)
	fmt.Printf("In six hours it will be %v\n", future2.Format("15:04"))

	println()

	bedtime := time.Date(2022, 05, 23, 23, 0, 0, 0, time.Local)
	fmt.Printf("There is %.0f hours until bed time\n", time.Until(bedtime).Hours())

}

func Simpleapp() {
	start := time.Now()

	args := os.Args
	//open the customer list
	custlist, err := os.Open(string(args[1]))
	check(err)
	defer custlist.Close()

	writeTime(start, "Opened Customer List")

	//create an output file
	outfile, err := os.Create("./inputs/outfile.csv")
	check(err)
	defer outfile.Close()

	writeTime(start, "Created outfile")

	//scan the customer list
	scanner := bufio.NewScanner(custlist)
	for scanner.Scan() {
		names := strings.Split(scanner.Text(), ",")
		outfile.WriteString(names[1] + "," + names[2] + "\n")
	}

	check(scanner.Err())
	writeTime(start, "Wrote data outfile")
	defer writeTime(start, "Application Finished")
}

func writeTime(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %v", name, elapsed)
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
