package main

import (
	"os"
	"fmt"
	"flag"
	"io"
	"bufio"
	"strings"
)

type Folder struct {
	width int
	foldOnSpaces bool
	countBytes bool
}

func (f Folder) Write(in []byte) (n int, err error) {
//	fmt.Print("XXXXX")
	reader := bufio.NewReader(strings.NewReader(string(in)))
	buffer := ""
	toPrint := ""
	for {
		line, err := reader.ReadString('\n')
		fullLine := buffer + line
		if len(fullLine) > f.width {
			toPrint = fullLine[:f.width]
			buffer = fullLine[f.width:]
		} else {
			toPrint = fullLine
			buffer = ""
		}

		newLineLocation := strings.Index(toPrint, "\n")

		if newLineLocation != -1 && newLineLocation != len(toPrint) -1 && newLineLocation != 0 {
//			fmt.Println("jebiga", newLineLocation)
			excess := toPrint[newLineLocation:]
			toPrint = toPrint[:newLineLocation]
			buffer = excess + buffer
		}

		if strings.HasSuffix(toPrint, "\n") {
			fmt.Print(toPrint)
//			fmt.Print("YYYYY")
		} else {
			fmt.Println(toPrint)
		}

		if err != nil && len(buffer) == 0 {
			break
		}
	}
	n = len(in)
	return
}
// This doesn't work properly. No desire to fiddle with it now.
func main() {
	var width int
	flag.IntVar(&width, "w", 80, "Column width on which to fold")
	var foldOnSpaces bool
	flag.BoolVar(&foldOnSpaces, "s", false, "Fold on spaces")
	var countBytes bool
	flag.BoolVar(&countBytes, "b", false, "Count bytes rather than columns")

	flag.Parse()
	writer := Folder{width, foldOnSpaces, countBytes}

	if len(flag.Args()) == 0 {
		io.Copy(writer, os.Stdin)
	}

}
