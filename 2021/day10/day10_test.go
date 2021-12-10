package day10

import "testing"

func TestPartA(t *testing.T) {
	// given
	expected := 26397

	// when
	calc := partA(inputSmall)

	// then
	if calc != expected {
		t.Errorf("expected %d but %d calculated", expected, calc)
	}
}

func TestPartA_Big(t *testing.T) {
	// given
	expected := 847504

	// when
	calc := partA(input)

	// then
	if calc != expected {
		t.Errorf("expected %d but %d calculated", expected, calc)
	}
}

func TestPartB(t *testing.T) {
	// given
	expected := 288957

	// when
	calc := partB(inputSmall)

	// then
	if calc != expected {
		t.Errorf("expected %d but %d calculated", expected, calc)
	}
}

func TestPartB_Big(t *testing.T) {
	// given
	expected := 288957

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
	for i := 0; i < b.N; i++ {
		c = partB(input)
	}
}
