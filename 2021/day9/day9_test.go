package day9

import "testing"

func TestPartA(t *testing.T) {
	// given
	expected := 15

	// when
	calc := partA(inputSmall, 5, 10)

	// then
	if calc != expected {
		t.Errorf("expected %d but %d calculated", expected, calc)
	}
}

func TestPartA_Big(t *testing.T) {
	// given
	expected := 541

	// when
	calc := partA(input, 100, 100)

	// then
	if calc != expected {
		t.Errorf("expected %d but %d calculated", expected, calc)
	}
}

func TestPartB(t *testing.T) {
	// given
	expected := 1134

	// when
	calc := partB(inputSmall, 5, 10)

	// then
	if calc != expected {
		t.Errorf("expected %d but %d calculated", expected, calc)
	}
}

func TestPartB_Big(t *testing.T) {
	// given
	expected := 541

	// when
	calc := partB(input, 100, 100)

	// then
	if calc != expected {
		t.Errorf("expected %d but %d calculated", expected, calc)
	}
}
