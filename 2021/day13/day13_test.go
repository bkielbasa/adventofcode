package day13

import "testing"

func TestPartA(t *testing.T) {
	// given
	expected := `█████
█   █
█   █
█   █
█████`

	// when
	calc := partA(inputSmall, 11, 15)

	// then
	if calc != expected {
		t.Errorf("expected \n%s \n\nbut \n%s calculated", expected, calc)
	}
}

func TestPartA_Big(t *testing.T) {
	// given
	// ARHZPCUH
	expected := ` ██  ███  █  █ ████ ███   ██  █  █ █  █ 
█  █ █  █ █  █    █ █  █ █  █ █  █ █  █ 
█  █ █  █ ████   █  █  █ █    █  █ ████ 
████ ███  █  █  █   ███  █    █  █ █  █ 
█  █ █ █  █  █ █    █    █  █ █  █ █  █ 
█  █ █  █ █  █ ████ █     ██   ██  █  █
`

	// when
	calc := partA(input, 1311, 895)

	// then
	if calc != expected {
		t.Errorf("expected \n%s \n\nbut \n%s \ncalculated", expected, calc)
	}
}
