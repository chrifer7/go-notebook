package contents

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
)

func Showmestuff() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("What is your name?")
	text, _ := reader.ReadString('\n')

	fmt.Printf("Hello %v", text)

	fmt.Printf("Our current version of Go is %v, and it's running on %v  \n", runtime.Version(), runtime.GOOS)

}

func Dinnertotal() {

	args := os.Args[1:]

	if len(args) == 1 && args[0] == "/help" {
		fmt.Println("Usage: dinnertotal <Total Meal Amount> <Tip Percentage>")
		fmt.Println("Example: dinnertotal 20.5 10")
	} else {
		if len(args) != 2 {
			fmt.Println("You must enter 2 arguments!\nType /help for help.")
		} else {
			mealTotal, _ := strconv.ParseFloat(args[0], 32)
			tipAmount, _ := strconv.ParseFloat(args[1], 32)

			total := calculateTotal(mealTotal, tipAmount)

			fmt.Printf("Your meal total will be %.2f", total)
		}
	}
}

func calculateTotal(mealTotal float64, tipAmount float64) float64 {
	totalPrice := mealTotal + (mealTotal * (tipAmount / 100))

	return totalPrice
}

func Flagtest() {
	archPtr := flag.String("arch", "x86", "CPU Type")

	flag.Parse()

	switch *archPtr {
	case "x86":
		fmt.Println("Running in 32 bit mode.")
	case "AMD64":
		fmt.Println("Running in 64 bit mode.")
	case "IA64":
		fmt.Println("Remember IA64?")
	}

	fmt.Println("Process complete.")
}

func Scanfdemo() {
	var name string
	fmt.Println("What is your name?")
	//count, _ := fmt.Scanf("%s", &name)
	count, _ := fmt.Scanf("%q", &name) //quoted string: "Luis Christian"

	switch count {
	case 0:
		fmt.Printf("You must enter a name. Total inputs: %d", count)
	case 1:
		fmt.Printf("Hello %s!\nNice to meet you.", name)
	}

}

func Bufiodemo() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		if scanner.Text() == "/q" || scanner.Text() == "/quit" {
			fmt.Println("Quitting...")
			os.Exit(3)
		} else {
			fmt.Println("You typed " + scanner.Text())
		}
	}

	/*After Scan returns false, the Err method will return any error that occurred
	during scanning, except that if it was io.EOF, 	Err will return nil. */
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

}

func BufiodemoText() {
	file, err := os.Open("./inputs/test.txt")

	if err != nil {
		fmt.Println(err)
	}
	defer file.Close() //deferred: executed at the end of the program.

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	/*After Scan returns false, the Err method will return any error that occurred
	during scanning, except that if it was io.EOF, 	Err will return nil. */
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

}

func Filemaker() {
	args := os.Args[1:]
	if len(args) == 0 || args[0] == "/help" {
		fmt.Println("Usage: filemaker <input file>")
	} else {

		fmt.Println("How would you like to see the text?")
		fmt.Println("1: ALL CAPS")
		fmt.Println("2: Title Case")
		fmt.Println("3: lower case")

		var option int
		_, err := fmt.Scanf("%d", &option)
		if err != nil {
			fmt.Println(err)
		}

		//open up the file
		file, err := os.Open(args[0])
		if err != nil {
			fmt.Println(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			switch option {
			case 1:
				fmt.Println(strings.ToUpper(scanner.Text()))
			case 2:
				fmt.Println(strings.ToTitle(scanner.Text()))
			case 3:
				fmt.Println(strings.ToLower(scanner.Text()))
			}
		}
	}
}
