# Resistor Color Code Calculator

## About

 Little Go app to calculate electrical resistance given a list of colors.
 It supports 4 and 5 band schemes.

## Installation

 Clone this project to "$GOPATH/src/javierlvelasquez.com/resistor-color-code-calculator".
 The package structure refers to javierlvelasquez.com and not github.com

## Examples

### Calculating Resistance

    ./resistor-color-code-calculator -mode=calculate --bands White,Violet,Yellow,Black,gray
    9.74 h Ω ± 0.05% tolerance

    ./resistor-color-code-calculator -mode=calculate --bands Orange,Blue,Black,Silver
    3.60 da Ω ± 10% tolerance

    ./resistor-color-code-calculator -mode=calculate --bands Gray,Orange,Orange,Gold
    83 k Ω ± 5% tolerance

    ./resistor-color-code-calculator -integers=5,5 -tolerance=.10 -prefix=M -mode=parse
    Green, Green, Blue, Silver

    ./resistor-color-code-calculator -integers=234121 -fractional=23412 -tolerance=121 -mode=parse
    2016/10/30 01:44:40 Invalid Input

### Determine Bands

### Simple Four Band Examples

    ./resistor-color-code-calculator -integers=7,5 -tolerance=.01 -prefix=k -mode=parse
    Violet, Green, Orange, Brown

    ./resistor-color-code-calculator -integers=7,2,0 -tolerance=.05 -mode=parse
    Violet, Red, Brown, Gold

    ./resistor-color-code-calculator -integers=4,5 -tolerance=.01 -prefix=k -mode=parse
    Yellow, Green, Orange, Brown

#### Floating Point Resistor

    ./resistor-color-code-calculator -integers=7 -fractional=2 -tolerance=.1 -mode=parse
    Violet, Red, Gold, Silver

    ./resistor-color-code-calculator -integers=2 -fractional=7 -tolerance=.1 -mode=parse
    Red, Violet, Gold, Silver

    ./resistor-color-code-calculator -integers=1 -fractional=9 -tolerance=.1 -mode=parse
    Brown, White, Gold, Silver


#### Five Bands

    ./resistor-color-code-calculator -integers=2,1 -fractional=1 -tolerance=.02 -prefix=k -mode=parse
    Red, Brown, Brown, Red, Red
