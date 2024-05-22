package main

import (
	"ascii/funcs"
	"fmt"
	"os"
	"strings"
)

func main() {
	fileName := "standard.txt"
	if _, err := os.Stat(fileName); err != nil {
		fmt.Println("Banner file does not exist")
		return
	}

	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Usage: go run . <text>")
		return
	}

	inputText := args[0]

	count := funcs.News(inputText)
	if count != 0 {
		for i := 0; i < count; i++ {
			fmt.Println("")
		}
		return
	}
	inputText = strings.Replace(inputText, "\\n", "\n", -1)
	lines := strings.Split(inputText, "\n")

	if err := funcs.HashCheck(fileName); err != nil {
		fmt.Println(err.Error())
		return
	}

	if err := funcs.AsciiCheck(inputText); err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, line := range lines {
		result := funcs.Ascii_Printer(fileName, line)
		fmt.Print(result)
	}
}
