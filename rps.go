package main

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var reader = bufio.NewReader(os.Stdin)

var GameChoices = [3]string{"r", "p", "s"}

var GameChoiceNameMap = map[string]string{
	"r": "Rock",
	"s": "Scissors",
	"p": "Paper",
}

var GameWinRules = map[string]string {
	"r": "s",
	"p": "r",
	"s": "p",
}

var wins, losses, ties int

func main() {
	// Initialize game.
	initGame()

	for {
		// Get User's Choice
		userChoice, err := readInUserChoice()
		if err != nil {
			logError(err)
			continue
		}

		// Get Computer's choice.
		computerChoice := generateComputerChoice()

		// Determine Winner.
		if userChoice == computerChoice {
			ties++
			fmt.Println("It's a tie!")
		} else if GameWinRules[userChoice] == computerChoice {
			wins++
			fmt.Println("You won!")
		} else {
			losses++
			fmt.Println("You LOSE")
		}

		logStandings()
	}
}

func initGame() {
	rand.Seed(time.Now().UnixNano())
	wins = 0
	losses = 0
	ties = 0
	logGameIntroduction()
}

func readInUserChoice() (string, error) {
	fmt.Print("-> ")
	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)
	choice = strings.ToLower(choice)
	fmt.Println("You played: ", GameChoiceNameMap[choice])

	if !validateInput(choice) {
		return "", errors.New("no go. You need to type 'r', 'p', or 's'")
	}
	return choice, nil
}

func generateComputerChoice() string {
	choice := GameChoices[(rand.Intn(3-0) + 0)]
	fmt.Println("Computer played: ", GameChoiceNameMap[choice])
	fmt.Println("----------------")
	return choice
}

func logGameIntroduction() {
	fmt.Println("Rock-Paper-Scissors - GO!")
	fmt.Println("Type 'r', 'p', or 's' to play...")
	fmt.Println("--------------------------------")
}

func logError(err error) {
	fmt.Println("ERROR: ", err)
}

func validateInput(choice string) bool {
	return choice == "r" || choice == "p" || choice == "s"
}

func logStandings() {
	fmt.Printf("Wins: %6d | Losses: %6d | Ties: %6d \n", wins, losses, ties)
}
