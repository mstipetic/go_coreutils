package main

import (
	"os"
	"io"
)

func main() {
	for index, fileName := range os.Args {
		if index == 0 {
			// program name
			continue
		}
		if file, err := os.Open(fileName); err == nil {
			io.Copy(os.Stdout, file)
		}
	}
}
