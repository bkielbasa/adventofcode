package day7

import "testing"

func TestPartA_Small(t *testing.T) {
	// given
	expected := 37

	// when
	calc := partA(inputSmall)

	// then
	if calc != expected {
		t.Errorf("expected %d but %d calculated", expected, calc)
	}
}

func TestPartA_Big(t *testing.T) {
	// given
	expected := 37

	// when
	calc := partA(input)

	// then
	if calc != expected {
		t.Errorf("expected %d but %d calculated", expected, calc)
	}
}

func TestPartB_Small(t *testing.T) {
	// given
	expected := 168
	results = map[int]int{}

	// when
	calc := partB(inputSmall)

	// then
	if calc != expected {
		t.Errorf("expected %d but %d calculated", expected, calc)
	}
}

func TestPartB_Big(t *testing.T) {
	// given
	expected := 104822130
	results = map[int]int{}

	// when
	calc := partB(input)

	// then
	if calc != expected {
		t.Errorf("expected %d but %d calculated", expected, calc)
	}
}

var c int

func BenchmarkPartB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		results = map[int]int{}
		c = partB(input)
	}
}
