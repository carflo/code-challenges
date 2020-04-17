package main

import (
	"fmt"
	"os"
	"sync"
)

func main() {
	// Get input text from pipe or os.args
	originalText, err := getText()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Check for zero length input
	if len(originalText) == 0 {
		fmt.Println("Nothing to read")
		fmt.Println(help())
		os.Exit(1)
	}

	// Spin up 3 workers to parse the phrases.
	// The offset allows them to read every possible
	// combo of 3 word phrases in the text.
	wg := sync.WaitGroup{}
	results := make([]map[string]int, 3)
	for i := 0; i < 3; i++ {
		wg.Add(1)
		p := &phraseCounter{
			phraseCount:  make(map[string]int),
			originalText: originalText,
			offset:       i,
		}
		go func(counter *phraseCounter, index int) {
			defer wg.Done()
			results[index] = counter.run()
		}(p, i)
	}
	wg.Wait()

	// Tabulate results at the end.
	finalCount := make(map[string]int)
	for _, counter := range results {
		for phrase := range counter {
			finalCount[phrase] = finalCount[phrase] + counter[phrase]
		}
	}

	// Print sorted results
	printResults(finalCount)
}

func help() string {
	return fmt.Sprintf("usage: ./main pathToFile... OR cat pathToFile | ./main")
}

// Get a sorted slice of keys based off the values in the map
// Use new slice to print the map in order (limit to 100)
func printResults(finalCount map[string]int) {
	for index, phrase := range sortedKeys(finalCount) {
		if index == 100 {
			break
		}
		fmt.Printf("%v: %v\n", phrase, finalCount[phrase])
	}
}
