package main

import "testing"

func TestPartA_Small(t *testing.T) {
	// given
	expected := 5
	matrix := make([][]int, 10)
	for i := 0; i < 10; i++ {
		matrix[i] = make([]int, 10)
	}

	// when
	calc := partA(input2, matrix)

	// then
	if expected != calc {
		t.Errorf("expected %d but %d calculated", expected, calc)
	}
}

func TestPartA_Big(t *testing.T) {
	// given
	expected := 7085
	matrix := make([][]int, 1000)
	for i := 0; i < 1000; i++ {
		matrix[i] = make([]int, 1000)
	}

	// when
	calc := partA(input1, matrix)

	// then
	if expected != calc {
		t.Errorf("expected %d but %d calculated", expected, calc)
	}
}

func TestPartB_Small(t *testing.T) {
	// given
	expected := 12
	matrix := make([][]int, 10)
	for i := 0; i < 10; i++ {
		matrix[i] = make([]int, 10)
	}

	// when
	calc := partB(input2, matrix)

	// then
	if expected != calc {
		t.Errorf("expected %d but %d calculated", expected, calc)
	}
}

func TestPartB_Big(t *testing.T) {
	// given
	expected := 20271
	matrix := make([][]int, 1000)
	for i := 0; i < 1000; i++ {
		matrix[i] = make([]int, 1000)
	}

	// when
	calc := partB(input1, matrix)

	// then
	if expected != calc {
		t.Errorf("expected %d but %d calculated", expected, calc)
	}
}

var calc int

func BenchmarkPartB_Big(b *testing.B) {
	matrix := make([][]int, 1000)
	for i := 0; i < 1000; i++ {
		matrix[i] = make([]int, 1000)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		calc = partB(input1, matrix)
	}
}

func BenchmarkPartA_Big(b *testing.B) {
	matrix := make([][]int, 1000)
	for i := 0; i < 1000; i++ {
		matrix[i] = make([]int, 1000)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		calc = partA(input1, matrix)
	}
}
