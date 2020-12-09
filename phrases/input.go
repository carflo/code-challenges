package main

import (
	"bufio"
	"io/ioutil"
	"os"
	"strings"
)

func getText() ([]byte, error) {
	var originalText []byte
	// attempt to read from stdin
	if len(os.Args) < 2 {
		return getPipeInput()
	}

	// attempt to read files from command line arguments
	for _, arg := range os.Args[1:] {
		content, err := ioutil.ReadFile(arg)
		if err != nil {
			return nil, err
		}
		originalText = append(originalText, content...)
	}

	return originalText, nil
}

func getPipeInput() ([]byte, error) {
	var originalText []byte
	input, err := os.Stdin.Stat()
	if err != nil {
		return nil, err
	}
	if input.Mode()&os.ModeNamedPipe == os.ModeNamedPipe {
		originalText, err = ioutil.ReadAll(os.Stdin)
		if err != nil {
			return nil, err
		}
	}

	return originalText, nil
}

func getText2() ([]byte, error) {
	var originalText []byte
	// attempt to read from stdin
	if len(os.Args) < 2 {
		return getPipeInput()
	}

	// attempt to read files from command line arguments
	for _, arg := range os.Args[1:] {
		content, err := ioutil.ReadFile(arg)
		if err != nil {
			return nil, err
		}
		originalText = append(originalText, content...)
	}

	return originalText, nil
}

func getPipeInput2() ([]byte, error) {
	var originalText []byte
	input, err := os.Stdin.Stat()
	if err != nil {
		return nil, err
	}
	if input.Mode()&os.ModeNamedPipe == os.ModeNamedPipe {
		originalText, err = ioutil.ReadAll(os.Stdin)
		if err != nil {
			return nil, err
		}
	}

	return originalText, nil
}

func read(stream chan string) {
	input, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	if input.Mode()&os.ModeNamedPipe == os.ModeNamedPipe {
		f := os.Stdin

		scanner := bufio.NewScanner(f)
		scanner.Split(bufio.ScanWords)
		phraseSlice := []string{}
		for scanner.Scan() {
			if len(phraseSlice) == 3 {
				stream <- strings.Join(phraseSlice, " ")
				phraseSlice = []string{}
			}
			phraseSlice = append(phraseSlice, scanner.Text())
		}

	} else {
		for _, arg := range os.Args[1:] {
			f, err := os.Open(arg)
			if err != nil {
				panic(err)
			}

			scanner := bufio.NewScanner(f)
			scanner.Split(bufio.ScanWords)
			phraseSlice := []string{}
			for scanner.Scan() {
				if len(phraseSlice) == 3 {
					stream <- strings.Join(phraseSlice, " ")
					phraseSlice = []string{}
				}
				phraseSlice = append(phraseSlice, scanner.Text())
			}

			defer f.Close()
		}
	}
}
