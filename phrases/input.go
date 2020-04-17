package main

import (
	"io/ioutil"
	"os"
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
