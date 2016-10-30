# Resistor Color Code Calculator

Go library to calculate resistance ± tolerance given a list of bands as well as vice versa. 

## Installation
Clone this project to "$GOPATH/src/javierlvelasquez.com/resistor-color-code-calculator".

## Example Usage

### Calculating resistance ± tolerance

    resistance.Calculate("White,Violet,Yellow,Black,gray") // returns 9.74 h Ω ± 0.05% tolerance

### Calculating bands

    bds, err = parser.DetermineBands([]int{2, 7}, []int{}, "k", .01) # // returns "Red", "Violet", "Orange", "Brown" (as a list of colors)

## Example Implementation

  See a cli implementation [here](https://github.com/nycjv321/resistor-color-code-calculator-cli).
