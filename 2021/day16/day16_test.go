package day16

import "testing"

func TestPartA(t *testing.T) {
	// given
	expected := int64(7)

	// when
	calc := partA("EE00D40C823060")

	// then
	if calc != expected {
		t.Errorf("expected %d but %d calculated", expected, calc)
	}
}

func TestPartA_Small(t *testing.T) {
	// given
	expected := int64(1)

	// when
	calc := partA("38006F45291200")

	// then
	if calc != expected {
		t.Errorf("expected %d but %d calculated", expected, calc)
	}
}

func TestPartA_Small2(t *testing.T) {
	// given
	expected := int64(16)

	// when
	calc := partA("8A004A801A8002F478")

	// then
	if calc != expected {
		t.Errorf("expected %d but %d calculated", expected, calc)
	}
}

func TestPartA_Big(t *testing.T) {
	// given
	expected := int64(971)

	// when
	calc := partA(input)

	// then
	if calc != expected {
		t.Errorf("expected %d but %d calculated", expected, calc)
	}
}

func TestPartB_Big(t *testing.T) {
	// given
	expected := int64(831996589851)

	// when
	calc := partB(input)

	// then
	if calc != expected {
		t.Errorf("expected %d but %d calculated", expected, calc)
	}
}
