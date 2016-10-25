# Resistor Color Code Calculator

## About

 Little Go app to calculate electrical resistance given a list of colors.
 It supports 4 and 5 band schemes.

## Installation

 Clone this project to "$GOPATH/src/javierlvelasquez.com/resistor-color-code-calculator".
 The package structure refers to javierlvelasquez.com and not github.com

## Examples

    ./resistor-color-code-calculator --bands White,Violet,Yellow,Black,gray
    9.74 h Ω ± 0.05% tolerance

    ./resistor-color-code-calculator --bands Orange,Blue,Black,Silver
    3.60 da Ω ± 10% tolerance

    ./resistor-color-code-calculator --bands Gray,Orange,Orange,Gold
    83 k Ω ± 5% tolerance
