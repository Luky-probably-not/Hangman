package pendu

import (
	"fmt"
	"os"
)

func Display(s []string) { //function that displays the word to guess at the start of the game
	final := ""
	for _, i := range s {
		final += i
	}
	fmt.Println(final)
}

func DisplayHangman(hp int) { // function that displays the progress of the hangman
	b,_ := os.ReadFile("jeu/hangman.txt")
	word := ""
	line := 0
	result := []string{}
	for _, i := range b {
		if line == 8 {
			line = 0
			result = append(result, word)
			word = ""
		}
		if string(i) == "\n" {
			line++
		}
		word += string(i)
	}
	
	print(append(result, word), hp)
}

func DisplayGuillo(hp int) { //function that displays the progress of the guillotine
	b,_ := os.ReadFile("jeu/guillotine.txt")
	word := ""
	ligne := 0
	result := []string{}
	for _, i := range b {
		if ligne == 9 {
			ligne = 0
			result = append(result, word)
			word = ""
		}
		if string(i) == "\n" {
			ligne++
		}
		word += string(i)
	}
	
	print(append(result, word), hp)	
}
func print(s []string, hp int) {
	if hp < 10 {
		fmt.Println(s[9-hp])
	}
}

func DisplayLetterUsed(s []string) { // function that displays the letters already used
	lettre := ""
	for _, i := range s {
		lettre += " " + i
	}
	fmt.Println(lettre)
}