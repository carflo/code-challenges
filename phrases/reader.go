package main

import (
	"bufio"
	"os"
	"regexp"
	"strings"
	"sync"
)

var (
	nonAlphabetic = regexp.MustCompile("[^a-zA-Z]+")
)

func getInput() []*os.File {
	files := []*os.File{}
	input, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	if input.Mode()&os.ModeNamedPipe == os.ModeNamedPipe {
		files = append(files, os.Stdin)
	} else {
		for _, arg := range os.Args[1:] {
			file, err := os.Open(arg)
			files = append(files, file)
			if err != nil {
				panic(err)
			}
		}
	}

	return files
}

func read(file *os.File, offset int, stream chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	phraseSlice := []string{}
	count := -1
	for scanner.Scan() {
		count++
		if count < offset {
			continue
		}

		word := scanner.Text()
		sanitized := nonAlphabetic.ReplaceAll([]byte(word), []byte(""))
		if len(sanitized) > 0 {
			phraseSlice = append(phraseSlice, strings.ToLower(string(sanitized)))
		}

		if len(phraseSlice) == 3 {
			stream <- strings.Join(phraseSlice, " ")
			phraseSlice = []string{}
		}
	}
}
