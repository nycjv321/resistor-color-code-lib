package main

import (
	"bytes"
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Band struct {
	Name       string
	Digit      int
	Multiplier float64
	Tolerance  float32
}

type MetricPrefix struct {
	Name       string
	ShortScale string
	Symbol     string
	Decimal    float64
}

var metricPrefixs = []MetricPrefix{MetricPrefix{Name: "yotta", ShortScale: "septillion", Symbol: "Y", Decimal: math.Pow(10, 24)},
	MetricPrefix{Name: "zetta", ShortScale: "sextillion", Symbol: "Z", Decimal: math.Pow(10, 21)},
	MetricPrefix{Name: "exa", ShortScale: "quintillion", Symbol: "E", Decimal: math.Pow(10, 18)},
	MetricPrefix{Name: "peta", ShortScale: "quadrillion", Symbol: "P", Decimal: math.Pow(10, 15)},
	MetricPrefix{Name: "tera", ShortScale: "trillion", Symbol: "T", Decimal: math.Pow(10, 12)},
	MetricPrefix{Name: "giga", ShortScale: "billion", Symbol: "G", Decimal: math.Pow(10, 9)},
	MetricPrefix{Name: "mega", ShortScale: "million", Symbol: "M", Decimal: math.Pow(10, 6)},
	MetricPrefix{Name: "kilo", ShortScale: "thousand", Symbol: "k", Decimal: math.Pow(10, 3)},
	MetricPrefix{Name: "hecto", ShortScale: "hundred", Symbol: "h", Decimal: math.Pow(10, 2)},
	MetricPrefix{Name: "deca", ShortScale: "ten", Symbol: "da", Decimal: math.Pow(10, 1)},
	MetricPrefix{ShortScale: "One", Decimal: math.Pow(10, 0)},
	MetricPrefix{Name: "deci", ShortScale: "tenth", Symbol: "n", Decimal: math.Pow(10, -1)},
	MetricPrefix{Name: "centi", ShortScale: "hundredth", Symbol: "n", Decimal: math.Pow(10, -2)},
	MetricPrefix{Name: "milli", ShortScale: "thousandth", Symbol: "m", Decimal: math.Pow(10, -3)},
	MetricPrefix{Name: "micro", ShortScale: "millionth", Symbol: "μ", Decimal: math.Pow(10, -6)},
	MetricPrefix{Name: "nano", ShortScale: "billionth", Symbol: "n", Decimal: math.Pow(10, -9)},
	MetricPrefix{Name: "pica", ShortScale: "trillionth", Symbol: "p", Decimal: math.Pow(10, -12)},
	MetricPrefix{Name: "femto", ShortScale: "quadrillionth", Symbol: "f", Decimal: math.Pow(10, -15)},
	MetricPrefix{Name: "atto", ShortScale: "quintillionth", Symbol: "a", Decimal: math.Pow(10, -18)},
	MetricPrefix{Name: "zepto", ShortScale: "sextillionth", Symbol: "z", Decimal: math.Pow(10, -21)},
	MetricPrefix{Name: "yocto", ShortScale: "septillionth", Symbol: "y", Decimal: math.Pow(10, -24)}, MetricPrefix{}}

var Bands = []Band{Band{Name: "Black", Digit: 0, Multiplier: 1},
	Band{Name: "Brown", Digit: 1, Multiplier: 10, Tolerance: .01},
	Band{Name: "Red", Digit: 2, Multiplier: math.Pow(10, 2), Tolerance: .02},
	Band{Name: "Orange", Digit: 3, Multiplier: math.Pow(10, 3), Tolerance: .03},
	Band{Name: "Yellow", Digit: 4, Multiplier: math.Pow(10, 4), Tolerance: .04},
	Band{Name: "Green", Digit: 5, Multiplier: math.Pow(10, 5), Tolerance: .005},
	Band{Name: "Blue", Digit: 6, Multiplier: math.Pow(10, 6), Tolerance: .0025},
	Band{Name: "Violet", Digit: 7, Multiplier: math.Pow(10, 7), Tolerance: .001},
	Band{Name: "Gray", Digit: 8, Multiplier: math.Pow(10, 8), Tolerance: .0005},
	Band{Name: "White", Digit: 9, Multiplier: math.Pow(10, 9)},
	Band{Name: "Gold", Multiplier: math.Pow(10, -1), Tolerance: .05},
	Band{Name: "Silver", Multiplier: math.Pow(10, -2), Tolerance: .1},
	Band{Tolerance: .2}}

func ToMetric(number float64) string {
	prefixes := metricPrefixs
	for _, prefix := range prefixes {
		decimal := prefix.Decimal
		if decimal < number {
			quotient := number / decimal
			if math.Mod(quotient, 1) > 0 {
				return strconv.FormatFloat(quotient, 'f', 2, 64) + " " + prefix.Symbol
			} else {
				return strconv.FormatFloat(quotient, 'f', 0, 64) + " " + prefix.Symbol
			}
		}
	}
	return ""
}

func findBand(Band string) Band {
	for _, element := range Bands {
		if strings.EqualFold(element.Name, strings.Trim(Band, " ")) {
			return element
		}
	}
	return Bands[len(Bands)-1]
}

func extractMultipler(bands []Band) ([]Band, Band) {
	length := len(bands)
	if length > 2 && length < 5 {
		multipler, digits := bands[length-1], bands[:length-1]
		return digits, multipler
	} else {
		return []Band{}, Band{}
	}
}

func extractToleranceBand(bands []Band) ([]Band, Band) {
	length := len(bands)
	if length > 3 && length < 6 {
		tolerance, digits := bands[length-1], bands[:length-1]
		return digits, tolerance
	} else {
		return []Band{}, Band{}
	}
}

func extractDigits(bands []Band) int {
	var buffer bytes.Buffer
	for _, element := range bands {
		buffer.WriteString(strconv.Itoa(element.Digit))
	}
	digits, error := strconv.Atoi(buffer.String())
	if error != nil {
		return 0
	} else {
		return digits
	}
}

func findBands(strings []string) []Band {
	var Bands = make([]Band, len(strings), len(strings))
	for index, element := range strings {
		Band := findBand(element)
		if Band.Name == "" {
			fmt.Fprintf(os.Stderr, "\"%v\" was not a valid Band.\n", element)
			os.Exit(1)
		} else {
			Bands[index] = Band
		}
	}
	return Bands
}

func parseFlags() string {
	var rawBands string
	flag.StringVar(&rawBands, "bands", "Yellow,Violet,Orange,Gold", "A comma delimited list of Bands representing resistor bands")
	flag.Parse()
	return rawBands
}

func CalculateResistance(bands []Band) string {

	digits, toleranceBand := extractToleranceBand(bands)
	var multiplerBand Band
	digits, multiplerBand = extractMultipler(digits)

	tolerance := toleranceBand.Tolerance * 100
	multipler := multiplerBand.Multiplier

	resistance := ToMetric(float64(extractDigits(digits)) * multipler)
	if tolerance == 0 {
		return fmt.Sprintf("%v Ω", resistance)
	} else if tolerance < 1 {
		return fmt.Sprintf("%v Ω ± %.2f%% tolerance", resistance, tolerance)
	} else {
		return fmt.Sprintf("%v Ω ± %.f%% tolerance", resistance, tolerance)
	}
}

func main() {
	fmt.Println(CalculateResistance(findBands(strings.Split(parseFlags(), ","))))
}
