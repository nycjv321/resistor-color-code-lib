package resistance

import (
	"bytes"
	"fmt"
	"strconv"
)

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

func Calculate(colors []string) string {
	return calculate(findBands(colors))
}

func calculate(bands []Band) string {

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
