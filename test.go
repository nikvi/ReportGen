package main

import (
	"fmt"
	"github.com/vly/go-gadata"
)

func main() {
	// initialise GAData
	gaTest := new(gadata.GAData)

	// initialise instance incl. authentication
	gaTest.Init()

	// build a basic GA query, replace your ga ID
	testRequest := gadata.GaRequest{"ga:43047246", // GA id
		"2014-01-01", // start date
		"2014-12-31", // end date
		"ga:visits",  // metrics
		"ga:month",   // dimensions
		"ga:eventlabel==hr.unimelb.edu.au/__data/assets/word_doc/0003/143949/HR4_updated.doc", // filters
		"",  // segments
		"",  // sort
		100} // results no.

	// launch data
	result := gaTest.GetData(1, &testRequest)
	fmt.Printf("results: %s\n", result)

}
