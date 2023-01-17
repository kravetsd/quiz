// this is a main.go for quiz game package.
// this will be na,ed as a main package.
// "main" package name means that this is going to be an exacutable program.
// This particular program is going to be a command line program.
// main package is the entry point of the program.
package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Hello, World!")

	//Opening a csv file from local path
	csvFile, err := os.Open("problems.csv")
	if err != nil {
		log.Fatal(err)
	}

	data := csv.NewReader(csvFile)
	records, err := data.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	// iterating over the records
	for _, record := range records {
		fmt.Println("question:", record[0], "\t answer:", record[1])
	}
}
