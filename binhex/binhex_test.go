package goreloaded

import "testing"

func TestBinToDec(t *testing.T) {
	tests := []struct {
		testStr  string
		expected string
	}{
		{"0000", "0"},
		{"0010", "2"},
	}
	for _, test := range tests {
		output := BinToDec(test.testStr)
		if output != test.expected {
			t.Errorf("Error(%s), /nExpected(%s) /nGot: (%s)", test.testStr, test.expected, output)
		}
	}
}

func TestHexToDec(t *testing.T) {
	tests := []struct {
		testStr  string
		expected string
	}{
		{"1DA6", "7590"},
		{"E8B", "3723"},
	}
	for _, test := range tests {
		output := HexToDec(test.testStr)
		if output != test.expected {
			t.Errorf("Error(%s), /nExpected(%s) /nGot: (%s)", test.testStr, test.expected, output)
		}
	}
}
