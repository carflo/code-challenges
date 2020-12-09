package main

import (
	"regexp"
	"strings"
)

var (
	newregex                = regexp.MustCompile(`[^\w]+`)
	nonAlphabeticRegex      = regexp.MustCompile("[^a-zA-Z ]+") // Remove non alphabetic chars
	newLineRegex            = regexp.MustCompile(`\r?\n`)       // Remove CR and LF
	doubleNewLineRegex      = regexp.MustCompile(`\r?\n\r?\n`)  // Remove LineBreaks
	extraSpaceStrippedRegex = regexp.MustCompile(`\s+`)         // Remove extra spaces
)

// phraseCounter is one of the counters
// that will tally up phrase frequency
type phraseCounter struct {
	phraseCount  map[string]int
	offset       int
	originalText []byte
}

func (p *phraseCounter) run() map[string]int {
	sanitized := p.sanitize2()
	parsed := p.parse(sanitized)
	p.count(parsed)

	return p.phraseCount
}

func (p *phraseCounter) parse(sanitized string) <-chan string {
	result := make(chan string)
	words := strings.Split(sanitized, string(" "))

	go func() {
		for i := p.offset; i < len(words)-2; i = i + 3 {
			phrase := strings.Join([]string{words[i], words[i+1], words[i+2]}, " ")
			result <- phrase
		}
		close(result)
	}()
	return result
}

func (p *phraseCounter) parse2(sanitized string) <-chan string {
	result := make(chan string)
	words := strings.Split(sanitized, string(" "))

	go func() {
		for i := p.offset; i < len(words)-2; i = i + 3 {
			phrase := strings.Join([]string{words[i], words[i+1], words[i+2]}, " ")
			result <- phrase
		}
		close(result)
	}()
	return result
}

func (p *phraseCounter) sanitize2() string {
	lowercase := strings.ToLower(string(p.originalText))
	sanitized := newregex.ReplaceAllString(lowercase, " ")

	return sanitized
}

func (p *phraseCounter) sanitize() string {
	lowercase := strings.ToLower(string(p.originalText))
	nonAlphabetic := nonAlphabeticRegex.ReplaceAllString(lowercase, "")
	blankLinesStripped := doubleNewLineRegex.ReplaceAllString(nonAlphabetic, " ")
	newLineStripped := newLineRegex.ReplaceAllString(blankLinesStripped, " ")
	sanitized := strings.TrimSpace(extraSpaceStrippedRegex.ReplaceAllString(newLineStripped, " "))

	return sanitized
}

func (p *phraseCounter) count(phrases <-chan string) {
	for phrase := range phrases {
		p.phraseCount[phrase]++
	}
}
