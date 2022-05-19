package examples

import "github.com/pluralsight/webservice/models"

// Program flow examples
// Capital letter to exporting function
func ProgramFlow() {
	//loopsCondition()
	//loopsPostClause()
	//infiniteLoops()
	//loopingCollection()
	//panicFunction()

	//ifStatements()
	switchStatements()
}

func switchStatements() {
	type HTTPRequest struct {
		Method string
	}

	r := HTTPRequest{Method: "HEAD"}

	//implicit break
	switch r.Method {
	case "GET":
		println("Get request")
		//fallthrough //to avoid the implicit break
	case "DELETE":
		println("Delete request")
	case "POST":
		println("Post request")
	case "PUT":
		println("Put request")
	default:
		println("Unhandled method")
	}
}

func ifStatements() {
	u1 := models.User{
		ID:        1,
		FirstName: "Camila",
		LastName:  "Sincabello",
	}

	u2 := models.User{
		ID:        2,
		FirstName: "Camilo",
		LastName:  "Sincabello",
	}

	if u1.ID == u2.ID {
		println("Same user!")
	} else if u1.FirstName == u2.FirstName || u1.LastName == u2.LastName {
		println("Similar user.")
	} else {
		println("Different user!")
	}
}

func loopsCondition() {
	var i int
	for i < 5 {
		println(i)
		i++

		if i == 3 {
			//break
			continue
		}
		println("continuing...")
	}
}

func loopsPostClause() {
	for i := 0; i < 5; i++ {
		println(i)
	}
}

func infiniteLoops() {
	var i int
	for {
		if i == 5 {
			break
		}
		println(i)
		i++
	}
}

func loopingCollection() {
	slice := []int{1, 2, 3}

	for i := 0; i < len(slice); i++ {
		println(slice[i])
	}

	for index, value := range slice {
		println(index, ": ", value)
	}

	wellKnownPorts := map[string]int{"http": 80, "https": 443}
	for index, value := range wellKnownPorts {
		println(index, ": ", value)
	}

	for index := range wellKnownPorts {
		println(index)
	}

	for _, value := range wellKnownPorts {
		println(value)
	}
}

func panicFunction() {
	println("Starting server...")
	//
	panic("Something bad just happened")

	println("Starting started.")
}
