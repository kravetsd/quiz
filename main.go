// this is a main.go for quiz game package.
// this will be na,ed as a main package.
// "main" package name means that this is going to be an exacutable program.
// This particular program is going to be a command line program.
// main package is the entry point of the program.
package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	fmt.Println("Starting quiz game...")

	// setting a flags to choose a file with quiz questions
	flag.String("problem-file", "problems.csv", "path to csv file with quiz questions and answers")

	// setting help flag and meesage
	flag.String("h", "", "print this help message")

	// default timer set to 30s
	flag.String("timer", "10s", "set the timer for the quiz")

	// parsing the flags
	flag.Parse()

	// getting the value of the flag
	csvFilePath := flag.Lookup("problem-file").Value.String()

	// getting the value of the timer flag
	timer, err := time.ParseDuration(flag.Lookup("timer").Value.String())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Openning quiz from file:", csvFilePath)

	// Opening a csv file from local path
	csvFile, err := os.Open(csvFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer csvFile.Close()
	// create a receiver for csv file
	csvReader := csv.NewReader(csvFile)

	// ask if user is ready to start the quiz
	var ready string
	fmt.Println("Are you ready to start the quiz and timer? (y/n)")
	fmt.Scanln(&ready) // getting ENTER from user
	if ready != "y" {
		fmt.Println("Ok, bye!")
		os.Exit(0)
	}

	// read all the records from csv file
	recs, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// preparing channels for goroutines to signal that they are done
	done := make(chan bool)

	go func() {
		time.Sleep(timer)
		fmt.Println("Time is up!")
		done <- true
		close(done)

	}()

	counter := 0
	go func() {
		for _, rec := range recs {
			fmt.Println("Question:", rec[0], "?")
			// getting answer from user via console
			var answer string
			fmt.Println("Answer: ")
			fmt.Scanln(&answer)
			if answer == rec[1] {
				counter++
			}
		}
		done <- true
		close(done)
	}()

	if <-done {
		fmt.Println("Your score is:", counter, "from", len(recs))
	}
}
