package main

import (
	"fmt"
	"strings"
	"flag"
	r "javierlvelasquez.com/resistor-color-code-calculator/resistance"
)

func main() {
	fmt.Println(r.Calculate(strings.Split(parseFlags(), ",")))
}

func parseFlags() string {
	var rawBands string
	flag.StringVar(&rawBands, "bands", "Yellow,Violet,Orange,Gold", "A comma delimited list of Bands representing resistor bands")
	flag.Parse()
	return rawBands
}
