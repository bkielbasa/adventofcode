package day03

import "testing"

func TestPartA(t *testing.T) {
	// given
	expected := 157

	// when
	calc := partA(inputSmall)

	// then
	if calc != expected {
		t.Errorf("expected %d but %d calculated", expected, calc)
	}
}

func TestPartA_Big(t *testing.T) {
	// given
	expected := 7795

	// when
	calc := partA(input)

	// then
	if calc != expected {
		t.Errorf("expected %d but %d calculated", expected, calc)
	}
}

func TestPartB(t *testing.T) {
	// given
	expected := 70

	// when
	calc := partB(inputSmall)

	// then
	if calc != expected {
		t.Errorf("expected %d but %d calculated", expected, calc)
	}
}

func TestPartB_Big(t *testing.T) {
	// given
	expected := 10238

	// when
	calc := partB(input)

	// then
	if calc != expected {
		t.Errorf("expected %d but %d calculated", expected, calc)
	}
}
