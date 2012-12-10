package main

import (
	"fmt"
	"flag"
	"os"
	"bufio"
)

func main() {
	n := flag.Int("n", 10, "Number of lines to read")
	printLineNumbers := flag.Bool("l", false, "Print line numbers")
	flag.Parse()

	for _, fileName := range flag.Args() {
		if file, err := os.Open(fileName); err == nil {
			reader := bufio.NewReader(file)
			for i := 0; i < *n; i++ {
				line, err := reader.ReadString('\n');
				if (err != nil) {
					break
				}
				if (*printLineNumbers) {
					fmt.Print(i, ": ", line)
				} else {
					fmt.Print(line)
				}
			}
		}
	}
}
