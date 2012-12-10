package main

import (
	"fmt"
	"strconv"
	"flag"
)

func main() {
	var separator string
	flag.StringVar(&separator, "s", "\n", "Element separator")
	flag.Parse()

	var first, last int
	var err error

	if len(flag.Args()) == 1 {
		first = 1
		if last, err = strconv.Atoi(flag.Args()[0]); err != nil {
			fmt.Println(err)
			return
		}
	} else if len(flag.Args()) == 2 {
		if first, err = strconv.Atoi(flag.Args()[0]); err != nil {
			fmt.Println(err)
			return
		}
		if last, err = strconv.Atoi(flag.Args()[1]); err != nil {
			fmt.Println(err)
			return
		}
	}

	for i := first; i <= last; i++ {
		fmt.Printf("%d%s", i, separator)
	}
}
