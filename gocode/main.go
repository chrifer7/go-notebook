package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/pluralsight/webservice/controllers"
	"github.com/pluralsight/webservice/examples"
	"github.com/pluralsight/webservice/models"
)

func main() {
	//examplesData()
	//demoWebServer()
	//examplesProgramFlow()

	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)

}

func demoWebServer() {
	port := 3000
	_, err := startWebServer(port, 2)
	fmt.Println(err)
}

func startWebServer(port int, numberOfRetries int) (int, error) {
	fmt.Println("Starting server...")
	//
	fmt.Println("Starting started on port:", port)
	fmt.Println("Number of retries:", numberOfRetries)

	return port, errors.New("something went wrong")
}

func examplesProgramFlow() {
	examples.ProgramFlow()
}

func examplesData() {
	examples.PrimitiveData()
	examples.Collections()
}

func examplesData2() {
	u := models.User{
		ID:        2,
		FirstName: "Gordon",
		LastName:  "Freeman",
	}

	fmt.Println(u)
}
