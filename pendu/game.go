package pendu

import (
	"fmt"
	"math/rand"
	"os"
)

func StartHard(s []string) []string {  //function that puts no letters at the beginning of the game (Hard mode)
	var index int
	word := []string{}
	for _, i := range s {
		word = append(word, i)
	}
	for i := 0; i < len(word); i++ {
		index = rand.Intn(len(word))
		for word[index] == "_" {
			index = rand.Intn(len(s))
		}
		word[index] = "_"
	}
	return word
}

func StartEasy(s []string) []string { //function that puts (len(word)/2)-1 letters at the beginning of the game (easy mode)
	var index int
    word := []string{}
    for _, i := range s {
        word = append(word, i)
    }
    for i := 0; i < ((len(word) / 2) + 1); i++ {
        index = rand.Intn(len(word))
        for word[index] == "_" {
            index = rand.Intn(len(s))
        }
        word[index] = "_"
    }
    return word
}

func Turn(s, answer []string) ([]string, int, string) { // function that handles each game turn
	var input string
	var cd int
	fmt.Printf("LETTER : ")
	_, err := fmt.Scan(&input)
	fmt.Print("\033[H\033[2J")
	if err != nil {
		return []string(nil), 0, ""
	}
	if len(input) != 1 {
		return []string(nil), 0, ""
	}
	if input == "+" {
		os.Exit(0)
	}
	l := check(input, answer)
	if len(l) == 0 {
		fmt.Printf("You type %s, \033[1;31mWrong !\033[0m\n \n", input)
		cd = -1
	} else {
		fmt.Printf("You type %s, \033[1;32mGreat !\033[0m\n \n", input)
		cd = 0
	}
	for _, i := range l {
		s[i] = answer[i]
	}
	return s, cd, string(input)
}

func check(s string, answer []string) []int { //function that checks the presence of the entered letter in the word
	l := []int{}
	for i := 0; i < len(answer); i++ {
		if s == answer[i] {
			l = append(l, i)
		}
	}
	return l
}

func Victory(s, answer []string) bool { //function that checks if the player has guessed the entire word
	for f, i := range answer {
		if s[f] != i {
			return false
		}
	}
	fmt.Print("\033[H\033[2J")
	fmt.Printf("\033[1;32m%s\033[0m\n \n", s)
	fmt.Println("\033[1;32mYou got this one Congrats !\033[0m")
	return true
}
