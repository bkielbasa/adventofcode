package main

import "testing"

func TestPartA_Small(t *testing.T) {
	// given
	expected := 5934

	// when
	calc := partA(inputSmall, 80)

	// then
	if calc != expected {
		t.Errorf("expected %d but %d calculated", expected, calc)
	}
}

func TestPartA_Big(t *testing.T) {
	// given
	expected := 5934

	// when
	calc := partA(input1, 80)

	// then
	if calc != expected {
		t.Errorf("expected %d but %d calculated", expected, calc)
	}
}

func TestPartB_Small(t *testing.T) {
	// given
	expected := 26984457539

	// when
	calc := partB(inputSmall, 256)

	// then
	if calc != expected {
		t.Errorf("expected %d but %d calculated", expected, calc)
	}
}

func TestPartB_Big(t *testing.T) {
	// given
	expected := 26984457539

	// when
	calc := partB(input1, 256)

	// then
	if calc != expected {
		t.Errorf("expected %d but %d calculated", expected, calc)
	}
}
