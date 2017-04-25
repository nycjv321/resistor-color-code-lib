package resistance

import (
	"bytes"
	"fmt"
	bands "javierlvelasquez.com/resistor-color-code-lib/bands"
	metric "javierlvelasquez.com/resistor-color-code-lib/metric"
	"strconv"
)

func extractMultipler(bnds []bands.Band) ([]bands.Band, bands.Band) {
	length := len(bnds)
	if length > 2 && length < 5 {
		multipler, digits := bnds[length-1], bnds[:length-1]
		return digits, multipler
	} else {
		return []bands.Band{}, bands.Band{}
	}
}

func extractToleranceBand(bnds []bands.Band) ([]bands.Band, bands.Band) {
	length := len(bnds)
	if length > 3 && length < 6 {
		tolerance, digits := bnds[length-1], bnds[:length-1]
		return digits, tolerance
	} else {
		return []bands.Band{}, bands.Band{}
	}
}

func extractDigits(bnds []bands.Band) int {
	var buffer bytes.Buffer
	for _, element := range bnds {
		buffer.WriteString(strconv.Itoa(element.Digit))
	}
	digits, error := strconv.Atoi(buffer.String())
	if error != nil {
		return 0
	} else {
		return digits
	}
}

func Calculate(colors []string) string {
	return GetResistance(colors).String()
}

func GetResistance(colors []string) *Resistance {
	return calculate(bands.FindBandsByColors(colors))
}

func calculate(bnds []bands.Band) *Resistance {

	digits, toleranceBand := extractToleranceBand(bnds)
	var multiplerBand bands.Band
	digits, multiplerBand = extractMultipler(digits)

	tolerance := toleranceBand.Tolerance * 100
	multipler := multiplerBand.Multiplier

	resistance := metric.ToMetric(float64(extractDigits(digits)) * multipler)
	return &Resistance{resistance, tolerance}
}

type Resistance struct {
	Value     *metric.Representation
	Tolerance float32
}

func (r *Resistance) String() string {
	t := r.Tolerance
	v := r.Value
	if t == 0 {
		return fmt.Sprintf("%v Ω", v)
	} else if t < 1 {
		return fmt.Sprintf("%v Ω ± %.2f%% tolerance", v, t)
	} else {
		return fmt.Sprintf("%v Ω ± %.f%% tolerance", v, t)
	}
}
