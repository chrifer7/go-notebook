package contents

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

/*Input scanning*/

func Scanfdemo2() {
	stdin := bufio.NewReader(os.Stdin)

	var firstname string
	var lastname string

	fmt.Println("What is your name?")

	fmt.Scanf("%s %s", &firstname, &lastname)

	fmt.Printf("Hello %s %s!\nNice to meet you.", firstname, lastname)

	stdin.ReadString('\n') //flush input

	var firstnumber int
	var lastnumber int

	fmt.Println("\nWhat two numbers would you like to add?")

	fmt.Scanf("%d %d", &firstnumber, &lastnumber)

	fmt.Printf("Total value is: %d", firstnumber+lastnumber)
}

func Sscanfdemo() {

	file, err := os.Open("./inputs/family.csv")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var age int
		var name string

		n, err := fmt.Sscanf(scanner.Text(), "%s is %d years old", &name, &age)

		if err != nil {
			panic(err)
		}

		if n == 2 {
			fmt.Printf("%s, %d\n", name, age)
		}

	}
}

func Scanlndemo() {
	fmt.Print("What is your name?\n")
	var firstName string
	var lastName string
	fmt.Scanln(&firstName, &lastName)
	fmt.Printf("Hello %s %s nice to meet you!\n", firstName, lastName)
}

/*Formatting output*/

func Printdemo() {
	var age = 42
	out, _ := fmt.Print("Jeremy is ", age, " years old \n")
	print("Bytes written: ", out)

}

func Printfdemo() {
	var age = 42
	var name = "Jeremy"

	/*
		%v	the value in a default format
		%t	the word true or false

		Integer:
		%b	base 2
		%c	the character represented by the corresponding Unicode code point
		%d	base 10
		%o	base 8
		%q	a single-quoted character literal safely escaped with Go syntax.
		%X	base 16, with upper-case letters for A-F
		%U	Unicode format: U+1234; same as "U+%04X"

		Floating-point and complex constituents:
		%f	decimal point but no exponent, e.g. 123.456

		String and slice of bytes (treated equivalently with these verbs):
		%s	the uninterpreted bytes of the string or slice
		%q	a double-quoted string safely escaped with Go syntax
	*/

	out, _ := fmt.Printf("My name is %v I am %v years old \n", name, age)
	fmt.Printf("My name is %s I am %d years old \n", name, age)

	print("Bytes written: ", out, "\n")

	var pi float32 = 3.141592

	fmt.Printf("Pi is %f\n", pi)
	fmt.Printf("Pi is %.2f\n", pi)

	println()

	fmt.Printf("|%f|%f|%f|\n", 23.3214, 456.45, 1235.65)
	fmt.Printf("|%f|%f|%f|\n", 98.214, 56.5445, 12.05)

	println()

	fmt.Printf("|%4.2f|%4.2f|%4.2f|\n", 23.3214, 456.45, 1235.65)
	fmt.Printf("|%4.2f|%4.2f|%4.2f|\n", 98.214, 56.5445, 12.05)

	println()

	fmt.Printf("|%4.2f|%4.2f|%4.2f|\n", 23.3214, 456.45, 1235.65)
	fmt.Printf("|%4.2f|%4.2f|%4.2f|\n", 98.214, 56.5445, 12.05)

	println()

	//7: total places including decimal point
	//2: only decimal places
	fmt.Printf("|%7.2f|%7.2f|%7.2f|\n", 23.3214, 456.45, 1235.65)
	fmt.Printf("|%7.2f|%7.2f|%7.2f|\n", 98.214, 56.5445, 12.05)

	println()

	fmt.Printf("|%-7.2f|%-7.2f|%-7.2f|\n", 23.3214, 456.45, 1235.65)
	fmt.Printf("|%-7.2f|%-7.2f|%-7.2f|\n", 98.214, 56.5445, 12.05)

	println()

	fmt.Printf("|%-7s|%-7s|%-7s|\n", "qwre", "erty", "poiu")
	fmt.Printf("|%-7s|%-7s|%-7s|\n", "qwerty", "asdf", "zxcv")
	fmt.Printf("|%-7q|%-7q|%-7q|\n", "qwre", "erty", "poiu")
	fmt.Printf("|%-7s|%-7q|%-7d|\n", "qwerty", "asdf", 100)

	println()

	output := fmt.Sprintf("|%-7s|%-7s|%-7s|\n", "foo", "bar", "go")
	print(output)
}

/*Manipulating Strings*/

//Already declared in usinglogo.go file
/* type messageType int

const (
	INFO messageType = 0 + iota
	WARNING
	ERROR
)
*/

const (
	InfoColor    = "\033[1;34m%s\033[0m"
	WarningColor = "\033[1;33m%s\033[0m"
	ErrorColor   = "\033[1;31m%s\033[0m"
)

func ErrorMessages() {
	fileName := "./inputs/test.txt"
	ShowMessage(INFO, fmt.Sprintf("About to open %s", fileName))

	ShowMessage(WARNING, "If the file is not present, this application will fail.")

	file, err := os.Open(fileName)
	if err != nil {
		ShowMessage(ERROR, err.Error())
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

}

func ShowMessage(messagetype messageType, message string) {
	switch messagetype {
	case INFO:
		printMessage := fmt.Sprintf("\nInformation: \n%s\n", message)
		fmt.Printf(InfoColor, printMessage)
	case WARNING:
		printMessage := fmt.Sprintf("\nWarning: \n%s\n", message)
		fmt.Printf(WarningColor, printMessage)
	case ERROR:
		printMessage := fmt.Sprintf("\nError: \n%s\n", message)
		fmt.Printf(ErrorColor, printMessage)
	}
}

/*Formatting other data types*/

type point struct {
	x, y int
}

type Person struct {
	firstName string
	lastName  string
	age       int
}

func PrintStuff() {
	p := point{2, 3}

	fmt.Printf("%v\n", p)

	newPerson := Person{"Jeremy", "Morgan", 42}
	fmt.Printf("%v\n", newPerson)
	fmt.Printf("%T\n", newPerson)

	fmt.Printf("Value is %t\n", true)
	fmt.Printf("%d\n", 45678)
	fmt.Printf("%b\n", 45678)
	fmt.Printf("%c\n", 64)
	fmt.Printf("%x\n", 45678)
}
