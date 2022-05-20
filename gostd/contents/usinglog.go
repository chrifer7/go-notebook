package contents

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime/trace"
	"time"
)

/*Formatting log output*/

type messageType int

const (
	INFO messageType = 0 + iota
	WARNING
	ERROR
	FATAL
)

func OurFirstLog() {
	file, err := os.OpenFile("./inputs/log.txt",
		os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
	defer file.Close()

	message := "First log"
	log.Println(message)
}

func PrintLogs() {
	WriteLog(INFO, "this is an information message!")
	WriteLog(WARNING, "this is a warning")
	WriteLog(ERROR, "this is an error")
	WriteLog(FATAL, "we crashed")
}

func WriteLog(messagetype messageType, message string) {
	file, err := os.OpenFile("./inputs/log.txt",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	//log.SetOutput(file)
	defer file.Close()

	switch messagetype {
	case INFO:
		logger := log.New(file, fmt.Sprintf("%-9s", "INFO:"), log.Ldate|log.Ltime|log.Lshortfile)
		logger.Println(message)
	case WARNING:
		logger := log.New(file, fmt.Sprintf("%-9s", "WARNING:"), log.Ldate|log.Ltime|log.Lshortfile)
		logger.Println(message)
	case ERROR:
		logger := log.New(file, fmt.Sprintf("%-9s", "ERROR:"), log.Ldate|log.Ltime|log.Lshortfile)
		logger.Println(message)
	case FATAL:
		logger := log.New(file, fmt.Sprintf("%-9s", "FATAL:"), log.Ldate|log.Ltime|log.Lshortfile)
		logger.Fatal(message) //kill the program after execution
	}

}

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
	FatalLogger   *log.Logger
)

//And finally means finally: init is called after all the variable declarations in the package have evaluated their initializers,
//and those are evaluated only after all the imported packages have been initialized.
//import --> const --> var --> init()
func init() {
	file, err := os.OpenFile("./inputs/log.txt",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	//defer file.Close()

	//InfoLogger = log.New(file, fmt.Sprintf("%-9s", "INFO:"), log.Ldate|log.Ltime|log.Lshortfile)
	InfoLogger = log.New(file, fmt.Sprintf("%-9s", "INFO:"), log.LUTC|log.Lmicroseconds|log.Llongfile)
	WarningLogger = log.New(file, fmt.Sprintf("%-9s", "WARNING:"), log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(file, fmt.Sprintf("%-9s", "ERROR:"), log.Ldate|log.Ltime|log.Lshortfile)
	FatalLogger = log.New(file, fmt.Sprintf("%-9s", "FATAL:"), log.Ldate|log.Ltime|log.Lshortfile)

}

func PrintUsingLoggers() {
	InfoLogger.Println("this is an information message!")
	WarningLogger.Println("this is a warning")
	ErrorLogger.Println("this is an error")
	FatalLogger.Fatal("we crashed")
}

/*Creating useful log files*/
//nil

/*Utilizing the Trace Logger*/
func TraceThis() {
	f, err := os.Create("./inputs/trace.out")

	if err != nil {
		log.Fatalf("We did not create a trace file! %v\n", err)
	}

	/*
		The first expression RHS is a function value. In the second version the RHS
		is the value returned by the function - i.e. a function call.

		So is the semantics of:

		defer f
		vs

		defer f()
	*/
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalf("Failed to close trace file %v\n", err)
		}
	}() //expression in defer must be function call

	if err := trace.Start(f); err != nil {
		log.Fatalf("We failed to start a trace: %d\n", err)
	}
	defer trace.Stop()

	//go tool trace .\inputs\trace.out

	addRandomNumbers()
}

func addRandomNumbers() {
	firstNumber := rand.Intn(100)
	secondNumber := rand.Intn(100)

	time.Sleep(2 * time.Second)

	var result = firstNumber * secondNumber

	fmt.Printf("Result is %d\n", result)
}
