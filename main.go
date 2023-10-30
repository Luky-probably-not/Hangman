package main

import (
	"fmt"
	"math/rand"
	pendu "pendu/jeu"
	"strconv"
	"slices"
	"os"
)

func main() {
	score := 0
	player := []string{}

	for { // infinite game loop
		fmt.Print("\033[H\033[2J") // code to clear the terminal

		death, level := pendu.Menu()
		fmt.Print("\033[1;35mNew Game\033[0m \n \n \n")

		for { // infinite score loop
			listLetter := []string{}
			text := "jeu/words" + strconv.Itoa(rand.Intn(4)) + ".txt" // choosing a random word library
			answer := pendu.SelectWord(text)
			if level == "e" {
				player = pendu.StartEasy(answer)
			} else if level == "h" {
				player = pendu.StartHard(answer)
			}
			pendu.Display(player)
			countdown := 10
			for pendu.Victory(player, answer) == false && countdown != 0 { // countdown loop
				fmt.Printf("\033[1;35mRemaining Trials : \033[0m%d             ", countdown)
				fmt.Printf("\033[1;35mletters already tried :\033[0m ")
				pendu.DisplayLetterUsed(listLetter)
				mot, score, lettre := pendu.Turn(player, answer)
				if score == -1 && !slices.Contains(listLetter, lettre) {
					countdown += score
					listLetter = append(listLetter, lettre)
				}
				if death == "c" {
					pendu.DisplayHangman(countdown)
				} else if death == "f" {
					pendu.DisplayGuillo(countdown)
				}
				pendu.Display(mot)
			}
			if countdown == 0 { //condition for defeat
				fmt.Print("\033[H\033[2J")
				fmt.Println("\nThe answer was : ")
				pendu.Display(answer)
				fmt.Printf("\033[1;31mJose is Dead, Sorry \033[1;37m\u2620\033[0m\n")
				if death == "c" {
					pendu.HangmanLoseAnimation()
				}
				if death == "f" {
					pendu.GuilloLoseAnimation()
				}
				os.Exit(0)
			} else { //condition for victory
				if death == "c" {
					pendu.HangmanWinAnimation()
				} else if death == "f" {
					pendu.GuilloWinAnimation()
				}
				score++
				fmt.Printf("Score : %d \n", score)
				fmt.Println("\033[1;35mNew Game\033[0m")
				fmt.Printf("\n \n \n \n")
			}
		}
	}
}