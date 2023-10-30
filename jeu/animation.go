package pendu

import (
	"fmt"
	"io/ioutil"
	"time"
)

func frame(r []byte) []string { //function that separates the different frames and puts them into an array of strings
	word := ""
	line := 0
	result := []string{}

	for _, i := range r {
		if line == 8 {
			line = 0
			result = append(result, word)
			word = ""
			continue
		}
		if string(i) == "\n" {
			line++
		}
		word += string(i)
	}
	return append(result, word)
}

func HangmanLoseAnimation() { // Animation for defeat (Classic Hangman)
	text, _ := ioutil.ReadFile("jeu/hangman_frame.txt")
	frames := frame(text)
	frames = append(frames, frames[1])
	fmt.Print(frames[0])
	index := 0

	for i := 0; i < 20; i++ {
		fmt.Printf("\033[8A") //It is a specific ANSI escape sequence that means move the cursor up by 8 lines
		fmt.Printf("%s", frames[index])
		time.Sleep(500 * time.Millisecond)
		index++
		if index == 4 {
			index = 0
		}
	}
	fmt.Println("\033[1;31mGAME OVER\033[0m")
}

func HangmanWinAnimation() { // Animation for victory (Classic Hangman)
	text, _ := ioutil.ReadFile("jeu/hangman_frame2.txt")
	frames := frame(text)
	frames = append(frames, frames[1])
	fmt.Print(frames[0])
	index := 0

	for i := 0; i < 20; i++ {
		fmt.Printf("\033[8A")
		fmt.Printf("%s", frames[index])
		time.Sleep(500 * time.Millisecond)
		index++
		if index == 4 {
			index = 0
		}
	}
	fmt.Println("\033[1;32mYou saved Jose ! \U0001F60A\033[0m")
	time.Sleep(3 * time.Second)
	fmt.Print("\033[H\033[2J")
}

func GuilloLoseAnimation() { // Animation for defeat (french version)
	text, _ := ioutil.ReadFile("jeu/guillotineAnime.txt")
	frames := frame(text)
	frames = append(frames, frames[1])
	fmt.Print(frames[0])
	index := 0

	for i := 0; i < 3; i++ {
		fmt.Printf("\033[8A")
		fmt.Printf("%s", frames[index])
		time.Sleep(800 * time.Millisecond)
		index++
		if index == 4 {
			index = 0
		}
	}
	fmt.Println("\033[1;31mGAME OVER\033[0m")
}

func GuilloWinAnimation() { // Animation for victory (french version)
	text, _ := ioutil.ReadFile("jeu/guillotineAnime2.txt")
	frames := frame(text)
	frames = append(frames, frames[1])
	fmt.Print(frames[0])
	index := 0

	for i := 0; i < 20; i++ {
		fmt.Printf("\033[8A")
		fmt.Printf("%s", frames[index])
		time.Sleep(500 * time.Millisecond)
		index++
		if index == 4 {
			index = 0
		}
	}
	fmt.Println("\033[1;32mYou saved Jose ! \U0001F60A\033[0m")
	time.Sleep(3 * time.Second)
	fmt.Print("\033[H\033[2J")
}

