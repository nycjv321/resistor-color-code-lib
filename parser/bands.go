package parser

import (
  "javierlvelasquez.com/resistor-color-code-calculator/bands"
  "javierlvelasquez.com/resistor-color-code-calculator/metric"
  "math"
  "errors"
  "fmt"
)

func DetermineBands(integers []int, fractional []int, prefix string, tolerance float32) ([]bands.Band, error) {
  toleranceBand := bands.FindBandByTolerance(tolerance)

  fractionalLength := len(fractional)
  metricPrefix := metric.FindMetricPrefixBySymbol(prefix)
  decimal := metricPrefix.Decimal
  var digits []int
  if fractionalLength > 0 {
    value := math.Pow(10, float64(fractionalLength))
    decimal /= value
    digits = append(integers, fractional...)
  } else {
    digits = integers
  }

  if len(fractional) > 0 && prefix == "" {
    decimal = math.Pow(10, float64(-1 * len(fractional)))
  } else if len(digits) == 2 && prefix == "" {
    decimal = 1
  }
  multiplierBand := bands.FindBandByMultiplier(decimal)


  bds := findBandsByDigits(digits)
  if len(bds) == 1 && len(bds[0].Name) == 0 {
    return []bands.Band{}, errors.New(fmt.Sprintf("%v not considered a valid band", bds[0].Name))
  }
  var value []bands.Band
  if len(multiplierBand.Name) > 0 {
    value = append(bds, []bands.Band{multiplierBand, toleranceBand}...)
  } else {
    value = append(bds, toleranceBand)
  }
  return value, nil

}

func findBandsByDigits(integers []int) []bands.Band {
  length := len(integers)
   var bds = make([]bands.Band, length, length)
   for i := 0; i < length; i++ {
     integer := integers[i]
     if i > 0 && integer == 0 {
       integer = 1
     }
     bds[i] = bands.FindBandByDigit(integer)
   }
 	return bds
}

func findLength(number float64, max int) int {
  position := 1
  for position < max {
    multiple := math.Pow(float64(10), float64(position))
    if math.Mod(number, multiple) == number {
      return position
    }
    position += 1
  }
  return -1
}
