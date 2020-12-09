package main

import (
	"io/ioutil"
	"reflect"
	"testing"
)

func TestSanitize(t *testing.T) {
	text, err := ioutil.ReadFile("text.example")
	if err != nil {
		t.Error(err)
	}

	expected := "i love sandwiches"
	if result != "i love sandwiches" {
		t.Errorf("got: %v, want: %v", result, expected)
	}
}

func TestCount(t *testing.T) {
	phrases := make(chan string)

	go func() {
		phrases <- "i love sandwiches"
		phrases <- "i love sandwiches"
		phrases <- "sandwiches are ok"
		close(phrases)
	}()

	p := &phraseCounter{
		phraseCount: make(map[string]int),
	}

	p.count(phrases)

	expected := make(map[string]int)
	expected["i love sandwiches"] = 2
	expected["sandwiches are ok"] = 1
	if !reflect.DeepEqual(p.phraseCount, expected) {
		t.Errorf("got: %v, want: %v", p.phraseCount, expected)
	}
}

func TestParse(t *testing.T) {
	p := phraseCounter{}

	result := p.parse("i love sandwiches i love")
	expected := "i love sandwiches"
	actual := <-result

	if actual != expected {
		t.Errorf("got: %v, want: %v", actual, expected)
	}
}

func TestRun(t *testing.T) {
	text, err := ioutil.ReadFile("text.example")
	if err != nil {
		t.Error(err)
	}

	p := &phraseCounter{
		originalText: text,
		offset:       0,
		phraseCount:  make(map[string]int),
	}

	p.run()

	expectedCount := make(map[string]int)
	expectedCount["i love sandwiches"] = 1

	expected := &phraseCounter{
		originalText: text,
		offset:       0,
		phraseCount:  expectedCount,
	}

	if !reflect.DeepEqual(p, expected) {
		t.Errorf("got: %v, want: %v", p, expected)
	}
}
