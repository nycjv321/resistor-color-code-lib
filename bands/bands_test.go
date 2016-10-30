package bands

import (
  "testing"
  "strconv"
)

func TestFindBandByTolerance(t *testing.T)  {
  checkTolerance(.01, "Brown", t)
  checkTolerance(.005, "Green", t)
  checkTolerance(.05, "Gold", t)
  checkTolerance(.20, "", t)
}

func TestFindBandByMultiplier(t *testing.T)  {
  checkMultipler(10, "Brown", t)
  checkMultipler(100000, "Green", t)
  checkMultipler(.1, "Gold", t)
  checkMultipler(.20, "", t)
}

func TestFindBandByDigit(t *testing.T)  {
  checkDigit(1, "Brown", t)
  checkDigit(5, "Green", t)
  checkDigit(9, "White", t)
  checkDigit(10, "", t)
}

func TestFindBandByName(t *testing.T)  {
  checkColor("Brown", 1, t)
  checkColor("Green", 5, t)
  checkColor("Gold", 0, t)
  checkColor("", 0, t)
  checkColor("Magenta", 0, t)
}

func checkDigit(digit int, color string, t *testing.T) {
  actual := FindBandByDigit(digit).Name
  if actual != color {
    t.Error(`FindBandByDigit(` + strconv.Itoa(digit) +`) == false. Expected == "`+ color +`",  Actual == "` + actual + `"`)
  }
}

func checkMultipler(multipler float64, color string, t *testing.T) {
  actual := FindBandByMultiplier(multipler).Name
  if actual != color {
    t.Error(`FindBandByMultiplier(` + strconv.FormatFloat(multipler, 'f', 2, 32) +`) == false. Expected == "`+ color +`",  Actual == "` + actual + `"`)
  }
}

func checkColor(color string, digit int, t *testing.T) {
  actual := findBandByName(color).Digit
  if actual != digit {
    t.Error(`findBandByName(` + color +`) == false. Actual == "` + strconv.Itoa(actual) + `"`)
  }
}

func checkTolerance(tolerance float32, expected string, t *testing.T) {
  actual := FindBandByTolerance(tolerance).Name
  if actual != expected {
    t.Error(`findBandByTolerance(` + strconv.FormatFloat(float64(tolerance), 'f', 2, 32) +`) == false. Actual == "` + actual + `"`)
  }
}
