package main

import (
	"fmt"
	"os"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	sanitizedPhrases := make(chan string)
	result := make(chan map[string]int)

	go func() {
		tabulate(sanitizedPhrases, result)
	}()

	for _, input := range getInput() {
		for i := 0; i < 3; i++ {
			wg.Add(1)
			go func(in *os.File, offset int, wg *sync.WaitGroup) {
				read(in, offset, sanitizedPhrases, wg)
			}(input, i, wg)
		}
	}

	go func() {
		wg.Wait()
		close(sanitizedPhrases)
	}()

	// Print sorted results
	printResults(<-result)
}

func help() string {
	return fmt.Sprintf("usage: ./main pathToFile... OR cat pathToFile | ./main")
}

// Get a sorted slice of keys based off the values in the map
// Use new slice to print the map in order (limit to 100)
func printResults(finalCount map[string]int) {
	for index, phrase := range sortedKeys(finalCount) {
		if index == 1 {
			break
		}
		fmt.Printf("%v: %v\n", phrase, finalCount[phrase])
	}
}

func tabulate(phrases chan string, result chan map[string]int) {
	counter := struct {
		sync.RWMutex
		p map[string]int
	}{p: make(map[string]int)}

	for phrase := range phrases {
		counter.Lock()
		counter.p[phrase]++
		counter.Unlock()
	}

	result <- counter.p
}
