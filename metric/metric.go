package metric

import (
  	"math"
    "strconv"
)

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


var EPSILON float64 = 0.00000001
func floatEquals(a, b float64) bool {
	if ((a - b) < EPSILON && (b - a) < EPSILON) {
		return true
	}
	return false
}

func FindMetricPrefixBySymbol(symbol string) MetricPrefix {
  prefixes := metricPrefixs
  if len(symbol) == 0 {
    return MetricPrefix{}
  } else {
	   for _, prefix := range prefixes {
      if prefix.Symbol == symbol {
        return prefix
      }
    }
    return MetricPrefix{}
  }
}

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
