package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Rock-Paper-Scissors - GO!")
	fmt.Println("Type r, p, or s to play.")
	fmt.Println("-------------------------")
	for {
		fmt.Print("-> ")
		computerChoices := [3]string{"r", "p", "s"}
		letter, _ := reader.ReadString('\n')
		// trimmed := strings.Replace(userChoice, "\n", "", -1) // doesn't work on windows
		userChoice := strings.Replace(letter, "\r\n", "", -1)
		fmt.Println("You played: ", userChoice)
		if userChoice == "r" || userChoice == "p" || userChoice == "s" {
			computerChoice := computerChoices[(rand.Intn(3-0) + 0)]
			fmt.Println("Computer played: ", computerChoice)
			fmt.Println("-------------")
			if (userChoice == "r" && computerChoice == "s") ||
				(userChoice == "p" && computerChoice == "r") ||
				(userChoice == "s" && computerChoice == "p") {
				fmt.Println("You won!")
			} else if userChoice == computerChoice {
				fmt.Println("It's a tie!")
			} else {
				fmt.Println("You LOSE")
			}
		} else {
			fmt.Println("No go. You need to type r, p, or s.")
		}

	}
}
