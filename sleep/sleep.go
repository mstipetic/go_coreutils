package main

import (
	"os"
	"strings"
	"strconv"
	"time"
)

func parseNumericPart(s, suffix string) float64 {
	toConvert := s[:len(s) - len(suffix)]
	num, err := strconv.ParseFloat(toConvert, 64)
	if err != nil {
		return 0
	}
	return num
}

func calculatePeriod(p string) (period float64) {
	if strings.HasSuffix(p, "s") {
		period = parseNumericPart(p, "s")
	} else if strings.HasSuffix(p, "m") {
		period = parseNumericPart(p, "m") * 60
	} else if strings.HasSuffix(p, "h") {
		period = parseNumericPart(p, "h") * 3600
	} else if strings.HasSuffix(p, "d") {
		period = parseNumericPart(p, "d") * 3600 * 24
	} else {
		period = parseNumericPart(p, "")
	}
	return
}

func main() {
	var period float64 
	for index, p := range os.Args {
		if index == 0 {
			continue
		}
		period += calculatePeriod(p)
	}
	
	<-time.After(time.Duration(period) * time.Second)

}
