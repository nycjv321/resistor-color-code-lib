package parser

import (
  "javierlvelasquez.com/resistor-color-code-calculator/bands"
  "testing"
)

func TestDetermineBands(t *testing.T)  {
  bds, _ := DetermineBands([]int{6,6},[]int{}, "k", .1)

  if len(bds) != 4 {
    t.Error("Incorrect bands")
  }

  checkBand(bds[0], "Blue", t)
  checkBand(bds[1], "Blue", t)
  checkBand(bds[2], "Orange", t)
  checkBand(bds[3], "Silver", t)

  bds, _ = DetermineBands([]int{6},[]int{6}, "M", .1)

  if len(bds) != 4 {
    t.Error("Incorrect bands")
  }

  checkBand(bds[0], "Blue", t)
  checkBand(bds[1], "Blue", t)
  checkBand(bds[2], "Green", t)
  checkBand(bds[3], "Silver", t)


  bds, _ = DetermineBands([]int{2,7},[]int{}, "k", .01)

  if len(bds) != 4 {
    t.Error("Incorrect bands: {0}", bds)
  }

  checkBand(bds[0], "Red", t)
  checkBand(bds[1], "Violet", t)
  checkBand(bds[2], "Orange", t)
  checkBand(bds[3], "Brown", t)

  bds, _ = DetermineBands([]int{7,2,0},[]int{}, "", .05)

  if len(bds) != 4 {
    t.Error("Incorrect bands: {0}", bds)
  }

  checkBand(bds[0], "Violet", t)
  checkBand(bds[1], "Red", t)
  checkBand(bds[2], "Brown", t)
  checkBand(bds[3], "Gold", t)

  bds, _ = DetermineBands([]int{7,9},[]int{}, "", .02)

  if len(bds) != 4 {
    t.Error("Incorrect bands: {0}", bds)
  }

  checkBand(bds[0], "Violet", t)
  checkBand(bds[1], "White", t)
  checkBand(bds[2], "Black", t)
  checkBand(bds[3], "Red", t)

}

func TestFractional(t *testing.T)  {
  bds, _ := DetermineBands([]int{7},[]int{2}, "", .10)

  if len(bds) != 4 {
    t.Error("Incorrect bands: {0}", bds)
  }

  checkBand(bds[0], "Violet", t)
  checkBand(bds[1], "Red", t)
  checkBand(bds[2], "Gold", t)
  checkBand(bds[3], "Silver", t)
}

func TestFractionalAndMultipler(t *testing.T)  {
  bds, _ := DetermineBands([]int{7},[]int{2}, "k", .10)

  if len(bds) != 4 {
    t.Error("Incorrect bands: {0}", bds)
  }

  checkBand(bds[0], "Violet", t)
  checkBand(bds[1], "Red", t)
  checkBand(bds[2], "Red", t)
  checkBand(bds[3], "Silver", t)
}

func TestNegativeMultipler(t *testing.T)  {
  bds, _ := DetermineBands([]int{},[]int{4,5}, "", .10)

  if len(bds) != 4 {
    t.Error("Incorrect bands: {0}", bds)
  }

  checkBand(bds[0], "Yellow", t)
  checkBand(bds[1], "Green", t)
  checkBand(bds[2], "Silver", t)
  checkBand(bds[3], "Silver", t)
}

func TestInvalidBand(t *testing.T)   {

    bds, err:= DetermineBands([]int{10},[]int{}, "", .05)

    if err == nil {
      t.Error("Expected for \"10\" to be considered invalid input. Got: {0}", bds)
    }

}

func checkBand(band bands.Band, name string, t *testing.T) {
  if band.Name != name {
    t.Error("Expected: \""+name+"\", Actual: \""+ band.Name+"\"")
  }
}
