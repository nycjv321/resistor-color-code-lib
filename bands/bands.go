package bands

import (
	"fmt"
	"math"
	"os"
	"strings"
)

type Band struct {
	Name       string
	Digit      int
	Multiplier float64
	Tolerance  float32
}

var bands = []Band{Band{Name: "Black", Digit: 0, Multiplier: 1},
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

func FindBandByDigit(digit int) Band {
	for _, element := range bands {
		if element.Digit == digit {
			return element
		}
	}
	return Band{}
}

func FindBandByTolerance(tolerance float32) Band {
	for _, element := range bands {
		if element.Tolerance == tolerance {
			return element
		}
	}
	return Band{}
}

func FindBandByMultiplier(multipler float64) Band {
	for _, element := range bands {
		if floatEquals(element.Multiplier, multipler) {
			return element
		}
	}
	return Band{}
}

var EPSILON float64 = 0.00000001

func floatEquals(a, b float64) bool {
	if (a-b) < EPSILON && (b-a) < EPSILON {
		return true
	}
	return false
}

func findBandByName(band string) Band {
	for _, element := range bands {
		if strings.EqualFold(element.Name, strings.Trim(band, " ")) {
			return element
		}
	}
	return Band{}
}

func FindBandsByColors(strings []string) []Band {
	var bds = make([]Band, len(strings), len(strings))
	for index, element := range strings {
		Band := findBandByName(element)
		if Band.Name == "" {
			fmt.Fprintf(os.Stderr, "\"%v\" was not a valid Band.\n", element)
			os.Exit(1)
		} else {
			bds[index] = Band
		}
	}
	return bds
}
