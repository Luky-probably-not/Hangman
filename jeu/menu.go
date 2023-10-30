package pendu

import (
	"fmt"
	"os"
	"time"
)

func Menu() (string, string){ //function that displays the different options at the beginning of the game
	fmt.Printf("\033[1;35mWelcome in the Hangman games ! \U0001F60A\033[0m\n \n \n")
	time.Sleep(3 * time.Second)
	fmt.Printf("\033[1;35m(in any time, you can type + to exit the program)\033[0m\n \n \n")
	time.Sleep(3 * time.Second)
	fmt.Printf("\033[1;35mfirst choose the difficulty\033[0m\n")
	time.Sleep(3 * time.Second)
	fmt.Print("\033[H\033[2J")

	var input string
	var level string
	fmt.Println()
	fmt.Printf("\033[1;34mEASY_MODE\033[0m(type e)        ")
	fmt.Printf("\033[1;34mHARD_MODE\033[0m(type h)        ")
	fmt.Printf("Choice : ")
	_, err := fmt.Scan(&level)
	fmt.Printf("\n \n \n")

	if level == "+" {
		os.Exit(0)
	}

	if err != nil {
		fmt.Println("Error")
		os.Exit(0)
	}

	for level != "e" && level != "h" {	
		fmt.Println("\nBad input, type e or h")
		fmt.Printf("Choice : ")
		_, err  = fmt.Scan(&input)
	}

	if level == "e" || level == "h" {
		fmt.Print("\033[H\033[2J")
		fmt.Printf("\033[1;35mNow Choose your path \033[1;33m\u2605\033[0m \n")
		time.Sleep(3 * time.Second)
		fmt.Print("\033[H\033[2J")
		fmt.Println()
		fmt.Printf("\033[1;34mCLASSIC_HANGMAN\033[0m(type c)        ")
		fmt.Printf("\033[1;34mFRENCH_VERSION\033[0m(type f)         ")
		fmt.Printf("\n")
		fmt.Printf("Choice : ")
		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Println("Error")
			os.Exit(0)
		}
		if input == "+" {
			os.Exit(0)
		}
		for input != "c" && input != "f" {	
			fmt.Println("\nBad input, type c or f")
			fmt.Printf("Your Choice : ")
			_, err  = fmt.Scan(&input)
		}
		fmt.Print("\033[H\033[2J")
	}
	return input, level
}
