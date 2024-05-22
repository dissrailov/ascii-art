package funcs

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"strings"
)

func AsciiCheck(s string) error {
	for _, ch := range s {
		if ch != 10 && (ch < 32 || ch > 126) {
			return fmt.Errorf("Ascii error")
		}
	}
	return nil
}

func HashCheck(fileName string) error {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}
	hasher := sha256.New()
	hasher.Write(data)
	hash := hasher.Sum(nil)
	hashString := hex.EncodeToString(hash)
	if hashString != "e194f1033442617ab8a78e1ca63a2061f5cc07a3f05ac226ed32eb9dfd22a6bf" {
		return fmt.Errorf("incorrect hash")
	}
	return nil
}

func NewlinesCheck(text string) int {
	count := 0
	for i := 0; i < len(text)-1; i++ {
		if text[i] == 92 && text[i+1] == 'n' {
			count++
		}
	}
	return count
}

func News(inputText string) int {
	clon := strings.Replace(inputText, "\\n", "", -1)
	if len(clon) > 0 {
		return 0
	}
	return NewlinesCheck(inputText)
}
