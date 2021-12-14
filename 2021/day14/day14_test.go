package day14

import "testing"

func TestPartA(t *testing.T) {
	// given
	expected := 1588

	// when
	calc := partA(inputSmall, 10)

	// then
	if calc != expected {
		t.Errorf("expected %d but %d calculated", expected, calc)
	}
}

func TestPartB(t *testing.T) {
	// given
	expected := 2188189693529

	// when
	calc := partB(inputSmall, 40)

	// then
	if calc != expected {
		t.Errorf("expected %d but %d calculated", expected, calc)
	}
}

func TestPartA_Big(t *testing.T) {
	// given
	expected := 4792

	// when
	calc := partA(input, 40)

	// then
	if calc != expected {
		t.Errorf("expected %d but %d calculated", expected, calc)
	}
}

func TestPartB_Big(t *testing.T) {
	// given
	expected := 2926813379532

	// when
	calc := partB(input, 40)

	// then
	if calc != expected {
		t.Errorf("expected %d but %d calculated", expected, calc)
	}
}

var c int

func BenchmarkPartB(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		c = partB(input, 40)
	}
}
