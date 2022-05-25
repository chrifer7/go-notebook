package contents

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"time"

	"alten.it/cf/gostd/media"
)

func CustomType() {
	fmt.Println("My Favorite Movie")

	/*myFav := media.Movie{}

	myFav.Title = "Top Gun"
	myFav.Rating = media.R
	myFav.BoxOffice = 43.2 */

	//myFav := media.NewMovie("Top Gun", media.R, 43.2)
	var myFav media.Catalogable = &media.Movie{}
	myFav.NewMovie("Top Gun", media.R, 43.2)

	fmt.Printf("My favorite movie is %s\n", myFav.GetTitle())
	fmt.Printf("It was rated %v\n", myFav.GetRating())
	fmt.Printf("It made %f in the box office.\n", myFav.GetBoxOffice())

	//It generates a copy of the structure
	/* myFav.Title = "Dumb and Dumber"
	fmt.Printf("My favorite movie is now  %s\n", myFav.getTitle) */

	myFav.SetTitle("Dumb and Dumber")
	fmt.Printf("My favorite movie is now %s\n", myFav.GetTitle())

}

func ShowTypes() {
	type person struct {
		personId  int
		firstName string
		lastName  string
	}

	newPerson := person{0, "Pedro", "Picapiedra"}

	fmt.Printf("Our person is %s %s with an id of %d\n", newPerson.firstName,
		newPerson.lastName, newPerson.personId)

	fmt.Printf("The type is %v \n", reflect.TypeOf(newPerson))

	fmt.Printf("The value is %v \n", reflect.ValueOf(newPerson))

	fmt.Printf("The kind of type is %v \n", reflect.ValueOf(newPerson).Kind())

	type employee struct {
		personId  int
		firstName string
		lastName  string
	}

	type customer struct {
		personId  int
		firstName string
		lastName  string
		company   string
	}

	newEmployee := employee{0, "Pedro", "Picapiedra"}

	newCustomer := customer{234, "Barney", "Dinosaur", "Slate Rock and Gravel"}

	addPerson(newEmployee)
	addPerson(newCustomer)

}

func addPerson(p interface{}) bool {
	if reflect.ValueOf(p).Kind() == reflect.Struct {
		var sqlStr string
		v := reflect.ValueOf(p)

		switch reflect.TypeOf(p).Name() {
		case "employee":
			sqlStr = "INSERT INTO employees (personId, firstName, lastName) VALUES (?, ?, ?)"

		case "customer":
			sqlStr = "INSERT INTO customers (personId, firstName, lastName, company) VALUES (?, ?, ?, ?)"

		}

		fmt.Printf("SQL: %s\n", sqlStr)
		fmt.Printf("Added %v\n", v.Field(1))

		return true
	} else {
		return false
	}
}

func TypeDemo() {

	println("----------- Reflection struct -----------")

	type employee struct {
		personId  int
		firstName string
		lastName  string
	}

	newEmployee := employee{0, "Pedro", "Picapiedra"}

	//name
	fmt.Printf("The name is %v \n", reflect.TypeOf(newEmployee).Name())
	// type
	fmt.Printf("The type is %v \n", reflect.TypeOf(newEmployee))
	// kind
	fmt.Printf("The kind of type is %v \n", reflect.ValueOf(newEmployee).Kind())
	// value
	fmt.Printf("The value is %v \n", reflect.ValueOf(newEmployee))

	println("----------- Reflection slice -----------")

	employees := make([]employee, 3)

	employees = append(employees, employee{0, "Harry", "Oxford"})
	employees = append(employees, employee{1, "Hina", "Ishika"})
	employees = append(employees, employee{2, "Jorge", "Fernández"})

	//name
	fmt.Printf("The name is %v \n", reflect.TypeOf(employees).Name())
	// type
	fmt.Printf("The type is %v \n", reflect.TypeOf(employees))
	// kind
	fmt.Printf("The kind of type is %v \n", reflect.ValueOf(employees).Kind())
	// value
	fmt.Printf("The value is %v \n", reflect.ValueOf(employees))

	println("----------- Grab type with reflection -----------")

	eType := reflect.TypeOf(employees)

	newEmployeeList := reflect.MakeSlice(eType, 0, 0)

	newEmployeeList = reflect.Append(newEmployeeList, reflect.ValueOf(employee{3, "Mary", "Popping"}))

	fmt.Printf("First list of employees: %v\n\n", employees)

	fmt.Printf("List created by reflection: %v\n", newEmployeeList)
}

func Primitive() {
	ourTitle := "the go standard library"

	newTitle := properTitle(ourTitle)
	fmt.Println(newTitle)

	fmt.Println(doubleOurNumber(3))

}

func properTitle(input string) string {
	// from: http://golangcookbook.com/chapters/strings/title/
	words := strings.Fields(input)
	smallwords := " a an on the to "

	for index, word := range words {
		if strings.Contains(smallwords, " "+word+" ") {
			words[index] = word
		} else {
			//executes: go get golang.org/x/text/cases cases.Title(word) //
			words[index] = strings.Title(word)
		}
	}
	return strings.Join(words, " ")
}

func doubleOurNumber(a int) int {
	time.Sleep(1 * time.Second)
	return a * 2
}

func CreateFunctions() {
	ourTitle := "the go standard library"
	timed := MakeTimedFunction(properTitle).(func(string) string)
	newTitle := timed(ourTitle)
	fmt.Println(newTitle)

	println()

	timedToo := MakeTimedFunction(doubleOurNumber).(func(int) int)
	fmt.Println(timedToo(2))
}

func MakeTimedFunction(f interface{}) interface{} {
	rf := reflect.TypeOf(f)

	if rf.Kind() != reflect.Func {
		panic("expects a function")
	}

	vf := reflect.ValueOf(f)
	wrapperF := reflect.MakeFunc(rf, func(in []reflect.Value) []reflect.Value {
		start := time.Now()
		out := vf.Call(in)
		end := time.Now()

		fmt.Printf("calling %s took %v\n", runtime.FuncForPC(vf.Pointer()).Name(), end.Sub(start))

		return out
	})

	return wrapperF.Interface()

}
