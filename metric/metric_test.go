package metric

import (
  "testing"
)

func TestFindMetricPrefixBySymbol(t *testing.T)  {
  checkPrefix("k", "kilo", t)
}

func checkPrefix(symbol string, name string, t *testing.T) {
  actual := FindMetricPrefixBySymbol(symbol)
  if actual.Name != name || symbol != actual.Symbol {
    t.Error(`FindMetricPrefixBySymbol(` + symbol +`) == false. Actual == "` + actual.Symbol + `"`)
  }
}
