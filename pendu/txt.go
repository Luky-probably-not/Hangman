package pendu

import (
	"fmt"
	"math/rand"
	"os"
)

func SelectWord(s string) []string { // function that selects a random word from the library
	content, err := os.ReadFile(s)
	if err != nil {
		fmt.Println("file error")
		return []string(nil)
	}
	list := byteToString(content)
	mot := list[rand.Intn(len(list)-1)]
	run := []rune(mot)
	str := []string{}
	for i := 0; i < len(mot)-1; i++ {
		str = append(str, string(run[i]))
	}
	return str
}

func byteToString(b []byte) []string {//function that converts the content of a game library into a string array
	word := ""
	result := []string{}
	for _, i := range b {
		if string(i) == "\n" {
			result = append(result, word)
			word = ""
		} else {
			word += string(i)
		}
	}
	return result
}
