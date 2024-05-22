package test

import (
	"ascii/funcs"
	"fmt"
	"os"
	"reflect"
	"testing"
)

type GetCharCodesTest struct {
	input    string
	expected []int
}

func TestGetCharCodes(t *testing.T) {
	tests := []GetCharCodesTest{
		{"Hello, World!", []int{72, 101, 108, 108, 111, 44, 32, 87, 111, 114, 108, 100, 33}},
		{"Привет, мир!", []int{1055, 1088, 1080, 1074, 1077, 1090, 44, 32, 1084, 1080, 1088, 33}},
		{"@@@@@@@@@@@@@@@", []int{64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64}},
		{"privet", []int{112, 114, 105, 118, 101, 116}},
		{"LDASDLASLDALDASDLAASDLADDDD", []int{76, 68, 65, 83, 68, 76, 65, 83, 76, 68, 65, 76, 68, 65, 83, 68, 76, 65, 65, 83, 68, 76, 65, 68, 68, 68, 68}},
	}

	for _, test := range tests {
		result := funcs.GetCharCodes(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For input \"%s\", expected %v, but got %v", test.input, test.expected, result)
		}
	}
}

func TestAsciiPrint(t *testing.T) {
	texts := []string{"Hello", "\n\n", "Hello\n\nThere", "ABCDEFGHIJKLMNOPQRSTUVWXYZ"}
	file := "standard.txt"
	for i := 0; i < 4; i++ {
		path := fmt.Sprintf("cases/testcase%d.txt", i+1)
		want, err := os.ReadFile(path)
		if err != nil {
			return
		}
		text := texts[i]
		have := funcs.Ascii_Printer(file, text)
		fmt.Println(have)
		if string(want) != have {
			t.Fatalf("\nWhat you need: \n%s \nWhat you have: \n%s", want, have)
		}
	}
}
