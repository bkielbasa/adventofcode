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
	expected := 357353

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
	results = make([]int, 2000)

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
	results = make([]int, 2000)

	// when
	calc := partB(input)

	// then
	if calc != expected {
		t.Errorf("expected %d but %d calculated", expected, calc)
	}
}

var c int

func BenchmarkPartB(b *testing.B) {
	b.ReportAllocs()
	results = make([]int, 2000)
	for i := 0; i < b.N; i++ {
		c = partB(input)
	}
}
