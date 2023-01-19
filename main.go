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
)

func main() {
	fmt.Println("Starting quiz game...")

	// setting a flags to choose a file with quiz questions
	flag.String("problem-file", "problems.csv", "path to csv file with quiz questions and answers")

	// setting help flag and meesage
	flag.String("h", "", "print this help message")

	// parsing the flags
	flag.Parse()

	// getting the value of the flag
	csvFilePath := flag.Lookup("problem-file").Value.String()

	fmt.Println("Openning quiz from file:", csvFilePath)

	// Opening a csv file from local path
	csvFile, err := os.Open(csvFilePath)
	if err != nil {
		log.Fatal(err)
	}
	// create a receiver for csv file
	csvReader := csv.NewReader(csvFile)

	// ask if user is ready to start the quiz
	var ready string
	fmt.Println("Are you ready to start the quiz? (y/n)")
	fmt.Scanln(&ready) // getting answer from user via console
	if ready != "y" {
		fmt.Println("Ok, bye!")
		return
	}

	//create a variable to store the score
	var score int

	// read all the records from csv file
	recs, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for _, rec := range recs {
		fmt.Println("Question:", rec[0], "?")
		// getting answer from user via console
		var answer string
		fmt.Println("Answer: ")
		fmt.Scanln(&answer)
		if answer == rec[1] {
			score++
		}
	}

	// print the score
	fmt.Println("Your score is:", score, "from", len(recs))

}
