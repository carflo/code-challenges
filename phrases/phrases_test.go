package main

import (
	"testing"
)

func TestHelp(t *testing.T) {
	actual := help()
	expected := "usage: ./main pathToFile... OR cat pathToFile | ./main"

	if actual != expected {
		t.Errorf("got: %v, want: %v", actual, expected)
	}
}

func Example_printResults() {
	finalCount := map[string]int{}

	finalCount["i love sandwiches"] = 50
	finalCount["i hate sandwiches"] = 26
	finalCount["sandwiches are ok"] = 100

	printResults(finalCount)

	// Output:
	//sandwiches are ok: 100
	//i love sandwiches: 50
	//i hate sandwiches: 26
}
