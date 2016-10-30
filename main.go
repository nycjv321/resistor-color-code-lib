package main

import (
	"fmt"
	"strings"
	"strconv"
	"flag"
	"bytes"
	"log"
	r "javierlvelasquez.com/resistor-color-code-calculator/resistance"
	p "javierlvelasquez.com/resistor-color-code-calculator/parser"
	"javierlvelasquez.com/resistor-color-code-calculator/bands"
)

func main() {
	flags := parseFlags()
	if flags["mode"] == "calculate" {
		bands, _ := flags["bands"].(string)
		fmt.Println(r.Calculate(strings.Split(bands, ",")))
	} else {
		integers, _ := flags["integers"].(string)
		fractionals, _ := flags["fractional"].(string)
		prefix, _ := flags["prefix"].(string)
		floatTolerance, _ := strconv.ParseFloat(flags["tolerance"].(string), 32)
		tolerance :=  float32(floatTolerance)
		bands, err := p.DetermineBands(stringToInts(integers),stringToInts(fractionals),prefix, tolerance)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(bandsToString(bands))
		}
	}
}

func bandsToString(bds []bands.Band) string {
	var buffer bytes.Buffer
	length := len(bds)
	lastElement := length - 1
	for i := 0; i < length; i++ {
		name := bds[i].Name
		if len(name) == 0 {
			log.Fatal("Invalid Input")
		}
		buffer.WriteString(name)
		if i < lastElement {
			buffer.WriteString(", ")
		}
	}

	return buffer.String()
}

func stringToInts(s string) []int {
	if len(s) == 0 {
		return []int{}
	}
	split := strings.Split(s, ",")
	length := len(split)
	var output = make([]int, length, length)
	var err error
	for i := 0; i < length; i++ {
		output[i], err = strconv.Atoi(split[i])
		if err != nil {
			panic(err)
		}
	}
	return output
}

func parseFlags() map[string]interface{} {
	var flags = make(map[string]interface{})
	var mode, tolerance, fractional,integers,  prefix, bands string
	flag.StringVar(&mode, "mode", "calculate", "the mode. \"calculate\" or \"parse\".")
	flag.StringVar(&tolerance, "tolerance", "", "The tolerance.")
	flag.StringVar(&integers, "integers", "", "Integer component of resistance.")
 	flag.StringVar(&fractional, "fractional", "", "Fractional component of resistance.")
	flag.StringVar(&prefix, "prefix", "", "Resistance metric prefix.")
	flag.StringVar(&bands, "bands", "Yellow,Violet,Orange,Gold", "A comma delimited list of colors")

	flag.Parse()

	flags["mode"] = mode
	flags["tolerance"] = tolerance
	flags["integers"] = integers
	flags["fractional"] = fractional
	flags["prefix"] = prefix
	flags["bands"] = bands

	return flags
}
