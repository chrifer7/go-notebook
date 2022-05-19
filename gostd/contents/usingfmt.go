package contents

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

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
