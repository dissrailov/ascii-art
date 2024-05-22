package test

import (
	"ascii/funcs"
	"fmt"
	"testing"
)

func TestAsciiCheck(t *testing.T) {
	TestArr := []struct {
		input    string
		expected error
	}{
		{"HELLO", nil},
		{"♥", fmt.Errorf("Ascii error")}, // Expecting an ASCII error for "nil"
		{"ValidASCII123", nil},
		{"NoValidascii♥", fmt.Errorf("Ascii error")},
		{"hello", nil},
		{"¶", fmt.Errorf("Ascii error")}, // Add more test cases as needed
		{"¥", fmt.Errorf("Ascii error")}, // Add more test cases as needed
	}

	for _, test := range TestArr {
		err := funcs.AsciiCheck(test.input)

		// Check if the error matches the expected result
		if (err == nil && test.expected != nil) || (err != nil && test.expected == nil) || (err != nil && err.Error() != test.expected.Error()) {
			t.Errorf("AsciiCheck(%s) = %v, expected %v", test.input, err, test.expected)
		}
	}
}

func TestHashCheck(t *testing.T) {
	TestArr := []struct {
		input    string
		expected error
	}{
		{"standard.txt", nil},                           // Expecting an ASCII error for "nil"
		{"standard1.txt", fmt.Errorf("incorrect hash")}, // Expecting an ASCII error for "nil"
		{"standard2.txt", fmt.Errorf("incorrect hash")}, // Expecting an ASCII error for "nil"
	}
	for _, test := range TestArr {
		err := funcs.HashCheck(test.input)
		if (err == nil && test.expected != nil) || (err != nil && test.expected == nil) || (err != nil && err.Error() != test.expected.Error()) {
			t.Errorf("AsciiCheck(%s) = %v, expected %v", test.input, err, test.expected)
		}
	}
}

type NewlinesCheckTest struct {
	input    string
	expected int
}

func TestNewlinesCheck(t *testing.T) {
	tests := []NewlinesCheckTest{
		{"hello\n", 0},
		{"hello\nLALALSLDA", 0},
		{"hello\n\n\n\n\n\n", 0},
		{"hello\\n\n\\n\n\\nn", 3},
		{"hello\ndsadan", 0},
		{"hello\nnet", 0},
		{"\\n", 1},
		{"\\n\\n", 2},
		{"\\n\\n\\n", 3},
		{"\\n\\n\\n\\n", 4},
		{"\\n\\n\\n\\n\\n", 5},
	}

	for _, test := range tests {
		result := funcs.NewlinesCheck(test.input)
		if result != test.expected {
			t.Errorf("For input \"%s\", expected %d newline(s), but got %d", test.input, test.expected, result)
		}
	}
}

func TestNews(t *testing.T) {
	tests := []NewlinesCheckTest{
		{"hello\n", 0},
		{"hello\nLALALSLDA", 0},
		{"hello\n\n\n\n\n\n", 0},
		{"hello\\n\n\\n\n\\nn", 0},
		{"hello\ndsadan", 0},
		{"hello\nnet", 0},
		{"\\n", 1},
		{"\\n\\n", 2},
		{"\\n\\n\\n", 3},
		{"\\n\\n\\n\\n", 4},
		{"\\n\\n\\n\\n\\n", 5},
		{"\\n\\n\\n\\n\\n SALAM", 0},
	}

	for _, test := range tests {
		result := funcs.News(test.input)
		if result != test.expected {
			t.Errorf("For input \"%s\", expected %d newline(s), but got %d", test.input, test.expected, result)
		}
	}
}
