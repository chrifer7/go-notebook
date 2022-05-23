package contents

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"
)

func Strings() {
	ourString := "\x48\x65\x78\x61\x64\x65\x63\x69\x6d\x61\x6c\x20\x6e\x6f\x74\x61\x74\x69\x6f\x6e\x20\x69\x73\x20\x75\x73\x65\x64\x20\x61\x73\x20\x61\x20\x68\x75\x6d\x61\x6e\x2d\x66\x72\x69\x65\x6e\x64\x6c\x79\x20\x72\x65\x70\x72\x65\x73\x65\x6e\x74\x61\x74\x69\x6f\x6e\x20"

	fmt.Println(ourString)

	for i := 0; i < len(ourString); i++ {
		fmt.Printf("%x", ourString[i])
	}

	for i := 0; i < len(ourString); i++ {
		fmt.Printf("%q\n", ourString[i])
	}

	newString := "This is a string!"

	fmt.Print(newString[3])
	println()
	fmt.Print(newString[0:5])
}

func Compare() {
	string1 := "this is a string!"
	string2 := "this is a string!" //"this is another string!"

	if string1 == string2 { //case sensitive
		fmt.Println("The strings are indentical!")
	} else {
		fmt.Println("The strings do not match!")
	}

	println()

	if strings.Compare(string1, string2) == 0 { //case sensitive
		fmt.Println("The strings are indentical!")
	} else {
		fmt.Println("The strings do not match!")
	}

	println()

	stooges := []string{"Larry", "Curly", "Moe"}
	for _, stooge := range stooges {
		fmt.Println(strings.Compare("Larry", stooge))
	}

}

func CompareTest() {
	string1 := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Suspendisse in consectetur odio. Suspendisse condimentum ac quam quis placerat. Proin non augue et erat pellentesque rhoncus quis ut magna. Aenean vulputate eget nisl sed rhoncus. Suspendisse tempus et quam sed consequat. Duis eget viverra nisl, eget venenatis lectus. Vestibulum sagittis sodales augue, at facilisis nunc fermentum pretium. Phasellus efficitur at arcu id pulvinar. Aliquam dapibus facilisis risus, pellentesque bibendum metus sollicitudin ut. Vivamus finibus luctus ligula, non iaculis nibh dapibus sit amet. In hac habitasse platea dictumst. Cras sagittis porttitor risus, et viverra orci sagittis id. Aliquam imperdiet urna et blandit rhoncus. Aenean tincidunt facilisis felis, sed luctus est sollicitudin in. Proin sed dapibus augue, non sollicitudin nibh. Sed ut est ex. Proin arcu odio, lobortis ac interdum vel, aliquam condimentum risus. Aliquam at egestas felis, sed aliquet lorem. Pellentesque sed rhoncus ex."

	string2 := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Suspendisse in consectetur odio. Suspendisse condimentum ac quam quis placerat. Proin non augue et erat pellentesque rhoncus quis ut magna. Aenean vulputate eget nisl sed rhoncus. Suspendisse tempus et quam sed consequat. Duis eget viverra nisl, eget venenatis lectus. Vestibulum sagittis sodales augue, at facilisis nunc fermentum pretium. Phasellus efficitur at arcu id pulvinar. Aliquam dapibus facilisis risus, pellentesque bibendum metus sollicitudin ut. Vivamus finibus luctus ligula, non iaculis nibh dapibus sit amet. In hac habitasse platea dictumst. Cras sagittis porttitor risus, et viverra orci sagittis id. Aliquam imperdiet urna et blandit rhoncus. Aenean tincidunt facilisis felis, sed luctus est sollicitudin in. Proin sed dapibus augue, non sollicitudin nibh. Sed ut est ex. Proin arcu odio, lobortis ac interdum vel, aliquam condimentum risus. Aliquam at egestas felis, sad aliquet lorem. Pellentesque sed rhoncus ex."

	//basic compare
	start1 := time.Now()
	if string1 == string2 { //case sensitive
		fmt.Println("The strings are indentical!")
	} else {
		fmt.Println("The strings do not match!")
	}
	writeTime(start1, "Basic compare")

	println()

	start2 := time.Now()
	if strings.Compare(string1, string2) == 0 { //case sensitive
		fmt.Println("The strings are indentical!")
	} else {
		fmt.Println("The strings do not match!")
	}
	writeTime(start2, "Strings compare")
}

func CaseTest() {
	string1 := "i like turtles" //"Hey this is a string!"
	string2 := "I like turtles"

	fmt.Println(compareCaseIns(string1, string2))
}

func compareCaseIns(a, b string) bool {
	if len(a) == len(b) {
		return strings.EqualFold(a, b)
		/* if strings.ToLower(a) == strings.ToLower(b) {
			return true
		} */
	}

	return false
}

func SplitTest() {
	ourString := "This is a string!"
	sep := " "

	ourString += "| This is another one | I like turtles"
	sep = "|"

	stringCollection := strings.Split(ourString, sep)

	for i := range stringCollection {
		fmt.Printf("%d) %s\n", i, stringCollection[i])
	}

	stringCollection = strings.SplitAfter(ourString, sep)
	for i := range stringCollection {
		fmt.Printf("%d) %s\n", i, stringCollection[i])
	}

	println()

	//Split N

	ourString = "This is a string!"
	sep = " "

	stringCollection = strings.SplitN(ourString, sep, 2)

	for i := range stringCollection {
		fmt.Printf("%d) %s\n", i, stringCollection[i])
	}

	println()

	//Split NewLine

	ourString = "This is a string!\nThis is a new line!"
	sep = "\n"

	stringCollection = strings.Split(ourString, sep)

	for i := range stringCollection {
		fmt.Printf("%d) %s\n", i, stringCollection[i])
	}
}

func SplitFile() {
	file, _ := os.Open("./inputs/customerlist.csv")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		items := strings.Split(line, ",")

		fmt.Println("-- new record --")
		for i := range items {
			fmt.Println(items[i])
		}
		fmt.Println()
	}
}

func FindText() {
	sampleString := "I really enjoy the Go language"
	searchTerm := "Go"

	if len(os.Args) > 1 {
		searchTerm = os.Args[1]
	} else {
		fmt.Println("You must enter a search term, default: " + searchTerm)
	}

	result := strings.Contains(sampleString, searchTerm)

	fmt.Printf("The sample text includes %s?\n%t\n", searchTerm, result)

	result = strings.HasPrefix(sampleString, searchTerm)

	fmt.Printf("The sample text starts with %s?\n%t\n", searchTerm, result)

	result = strings.HasSuffix(sampleString, searchTerm)

	fmt.Printf("The sample text ends with %s?\n%t\n", searchTerm, result)

}

func LogParser() {
	if len(os.Args) > 1 {
		searchTerm := os.Args[1]

		file, _ := os.Open("./inputs/log.txt")
		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			line := scanner.Text()
			result := strings.HasPrefix(scanner.Text(), searchTerm)

			if result {
				fmt.Println(line)
			}
		}
	} else {
		fmt.Println("You must enter a search term.")
	}

}

func FindReplace() {
	sampleString := "I really enjoy the Go language"
	searchTerm := "Go"
	replaceWith := "Python"

	sampleString = strings.Replace(sampleString, searchTerm, replaceWith, -1)

	fmt.Println(sampleString)
}

func Regex() {
	sampleString := "Go is a statically typed, compiled programming language designed at Google by Robert Griesemer, Rob Pike, and Ken Thompson. It is syntactically similar to C, but with memory safety, garbage collection, structural typing, and CSP-style concurrency. Wikipedia"

	println(sampleString)

	rgx := `s([a-z]+)y`
	println(rgx)

	r, _ := regexp.Compile(rgx)

	fmt.Println(r.MatchString(sampleString))

	fmt.Println(r.FindAllString(sampleString, -1))

	rgx = `c([a-z]+)n`
	println(rgx)

	r, _ = regexp.Compile(rgx)

	fmt.Println(r.MatchString(sampleString))

	fmt.Println(r.FindAllString(sampleString, -1))

	rgx = `s(\w[a-z]+)y\b` //word boundary
	println(rgx)

	r, _ = regexp.Compile(rgx)

	fmt.Println(r.MatchString(sampleString))

	fmt.Println(r.FindAllString(sampleString, -1))

	fmt.Println(r.FindStringIndex(sampleString))

	newText := r.ReplaceAllString(sampleString, "funny")
	fmt.Println(newText)
}

func Trim() {
	sampleString := "       This is our text         "
	fmt.Printf("%q\n", sampleString)

	newString := strings.TrimSpace(sampleString)
	fmt.Printf("%q\n", newString)

	newString = strings.TrimLeft(sampleString, " ")
	fmt.Printf("%q\n", newString)

	sampleString = "https://app.pluralsight.com/"
	domainName := strings.Trim(strings.TrimPrefix(sampleString, "https://"), "/")

	fmt.Println(domainName)

}

func Casing() {
	sampleString := "Never trust a programmer who carries a screwdriver\n"

	fmt.Println("Before: " + sampleString)

	strLowerCase := strings.ToLower(sampleString)

	fmt.Println("Lower case: " + strLowerCase)

	strUpperCase := strings.ToUpper(sampleString)

	fmt.Println("Upper case: " + strUpperCase)

	strTitleCase := strings.ToTitle(sampleString)

	fmt.Println("Title case: " + strTitleCase)

	properTitleCase := properTitle(sampleString)

	fmt.Println("Proper Title case: " + properTitleCase)

}

func properTitle(input string) string {
	words := strings.Fields(input) //tokenizer
	smallwords := " a an on the to "

	for index, word := range words {
		if strings.Contains(smallwords, " "+word+" ") {
			words[index] = word
		} else {
			words[index] = strings.Title(word)
		}
	}

	return strings.Join(words, " ")
}
