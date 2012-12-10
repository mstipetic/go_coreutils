package main

import (
	"fmt"
	"os"
)

func main() {
	if cwd, err := os.Getwd(); err == nil {
		fmt.Println(cwd)
	} else {
		fmt.Println(err)
	}

}
