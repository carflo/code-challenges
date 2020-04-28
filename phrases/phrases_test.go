package main

import (
	"io/ioutil"
	"strings"
	"sync"
	"testing"
	"unicode"
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

func TestPhrases(t *testing.T) {
	file := "./oos.txt"

	content, _ := ioutil.ReadFile(file)

	wg := &sync.WaitGroup{}
	results := make([]map[string]int, 3)
	for i := 0; i < 3; i++ {
		wg.Add(1)
		p := &phraseCounter{
			phraseCount:  make(map[string]int),
			originalText: content,
			offset:       i,
		}
		go func(counter *phraseCounter, index int) {
			defer wg.Done()
			results[index] = counter.run()
		}(p, i)
	}
	wg.Wait()

	tests := []struct {
		testFunc func([]map[string]int) string
		testName string
	}{
		{testName: "NoSpecialCharactersInPhrase", testFunc: NoSpecialCharactersInPhrase},
		{testName: "NoExtraSpaceInPhrase", testFunc: NoExtraSpaceInPhrase},
		{testName: "PhraseHasThreeWords", testFunc: PhraseHasThreeWords},
	}

	for _, test := range tests {
		test := test
		t.Run("Phrases are valid", func(t *testing.T) {
			wg.Add(1)
			go func() {
				defer wg.Done()
				invalidPhrase := test.testFunc(results)
				if invalidPhrase != "" {
					t.Errorf(test.testName + "failed with" + invalidPhrase)
				}
			}()
			wg.Wait()
		})
	}

}

func NoSpecialCharactersInPhrase(results []map[string]int) string {
	for _, phrases := range results {
		for phrase := range phrases {
			for _, char := range phrase {
				if !unicode.IsLetter(char) {
					return phrase
				}
			}
		}
	}

	return ""
}

func NoExtraSpaceInPhrase(results []map[string]int) string {
	var lastChar rune
	for _, phrases := range results {
		for phrase := range phrases {
			for _, char := range phrase {
				if lastChar == ' ' && char == ' ' {
					return phrase
				}
				lastChar = char
			}
		}
	}

	return ""
}

func PhraseHasThreeWords(results []map[string]int) string {
	for _, phrases := range results {
		for phrase := range phrases {
			words := strings.Split(phrase, " ")
			if len(words) != 3 {
				return phrase
			}
		}
	}

	return ""
}
