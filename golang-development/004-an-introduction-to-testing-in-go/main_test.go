package main

import "testing"

func TestCalculate(t *testing.T) {
	if Calculate(2) != 4 {
		t.Error("Expected 2 + 2 to exual 4")
	}
}

func TestTableCalculate(t *testing.T) {
	var tests = []struct {
		input    int
		expected int
	}{
		{2, 4},
		{-1, 1},
		{0, 2},
		{99999, 100001},
	}

	for _, test := range tests {
		if output := Calculate(test.input); output != test.expected {
			t.Error("Test Failed: {} inputted, {} expected, received: {}", test.input, test.expected, output)
		}
	}
}
