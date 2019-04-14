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
	computerChoices := [3]string{"r", "p", "s"}
	wins := 0
	losses := 0
	ties := 0
	fmt.Println("Rock-Paper-Scissors - GO!")
	fmt.Println("Type r, p, or s to play...")
	fmt.Println("--------------------------")
	for {
		fmt.Print("-> ")
		userChoice, _ := reader.ReadString('\n')
		// trimmed := strings.Replace(userChoice, "\n", "", -1) // doesn't work on Windows
		userChoice = strings.Replace(userChoice, "\r\n", "", -1)
		userChoice = strings.ToLower(userChoice)
		fmt.Println("You played: ", userChoice)
		if userChoice == "r" || userChoice == "p" || userChoice == "s" {
			computerChoice := computerChoices[(rand.Intn(3-0) + 0)]
			fmt.Println("Computer played: ", computerChoice)
			fmt.Println("-------------")
			if (userChoice == "r" && computerChoice == "s") || (userChoice == "p" && computerChoice == "r") || (userChoice == "s" && computerChoice == "p") {
				wins++
				fmt.Println("You won!")
			} else if userChoice == computerChoice {
				ties++
				fmt.Println("It's a tie!")
			} else {
				losses++
				fmt.Println("You LOSE")
			}
			fmt.Printf("Wins: %6d | Losses: %6d | Ties: %6d \n", wins, losses, ties)
		} else {
			fmt.Println("No go. You need to type r, p, or s.")
		}
	}
}
