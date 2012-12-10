package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

func main() {
	for index, fileName := range os.Args {
		if index == 0 {
			continue
		}
		file, err := os.Open(fileName)
		if err != nil {
			fmt.Println(err)
			continue
		}
		reader := bufio.NewReader(file)
		// lines, words, characters
		var l, w, c int
		for {
			if line, err := reader.ReadString('\n'); err == nil {
				l++
				w += len(strings.Fields(line))
				c += len(line)
			} else {
				break;
			}
		}
		fmt.Printf("Lines: %d\nWords: %d\nCharacters: %d\n", l, w, c);
	}
}
