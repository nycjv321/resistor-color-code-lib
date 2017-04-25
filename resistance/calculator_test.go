package resistance

import (
	"javierlvelasquez.com/resistor-color-code-lib/metric"
	"testing"
)

func TestFindMetricPrefixBySymbol(t *testing.T) {

	checkResistance(
		t,
		[]string{"White", "Violet", "Yellow", "Black", "Gray"},
		&metric.Representation{9.74, metric.FindMetricPrefixBySymbol("h")},
		.05)

}

func checkResistance(t *testing.T, bands []string, representation *metric.Representation, tolerance float32) {
	r := GetResistance(bands)
	actualRepresentation := *r.Value
	if actualRepresentation != *representation {
		t.Fatalf("Incorrect resistance. Expected %s, Actual: %s", representation, actualRepresentation)
	}
	actualTolerance := r.Tolerance
	if actualTolerance != tolerance {
		t.Fatalf("Incorrect Tolerance. Expected %s, Actual: %s", tolerance, actualTolerance)
	}
}
