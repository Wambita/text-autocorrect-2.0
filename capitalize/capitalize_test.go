package goreloaded

import "testing"

func TestCapitalize(t *testing.T) {
	tests := [] struct{

		input string
		expected string
	}{
		{"mine was a great day", "Mine was a great day"},
		{"amazing", "Amazing"},
	}
	
	for _,test := range tests {
		output := Capitalize(test.input)
		if output != test.expected {
			t.Errorf("Error \ninput: %s \nExpected:%s \nGot: %s", test.input, test.expected, output)
		}

	}
}