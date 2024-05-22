package funcs

import (
	"io/ioutil"
	"log"
	"strings"
)

func Ascii_Printer(fileName, text string) string {
	asciiContent, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	contentLines := strings.Split(string(asciiContent), "\n")
	charCodes := GetCharCodes(text)

	var result string
	for i := 0; i < 8; i++ {
		if len(charCodes) == 0 {
			break
		}
		for j := 0; j < len(charCodes); j++ {
			if charCodes[j] == 10 {
				charCodes = charCodes[1:]
				continue
			}
			count := 1 + 9*(charCodes[j]-32)
			if count+i >= 0 && count+i < len(contentLines) {
				result += contentLines[count+i]
			}
		}
		if i < 7 {
			result += "\n"
		}
	}
	lines := strings.Split(result, "\n")
	emptyLineCount := CountEmptyLines(lines)

	if emptyLineCount == len(lines)-1 {
		return result
	} else {
		return result + "\n"
	}
}

func GetCharCodes(text string) []int {
	var charCodes []int
	for _, char := range text {
		charCodes = append(charCodes, int(char))
	}
	return charCodes
}

func CountEmptyLines(lines []string) int {
	emptyLineCount := 0
	for _, line := range lines {
		if line == "" {
			emptyLineCount++
		}
	}
	return emptyLineCount
}
