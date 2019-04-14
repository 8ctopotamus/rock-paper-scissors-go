package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	userChoice := flag.String("choice", "foo", "--choose= r, p, or s to play.")
	flag.Parse()

	if *userChoice == "r" || *userChoice == "p" || *userChoice == "s" {
		rand.Seed(time.Now().UnixNano())
		computerChoices := [3]string{"r", "p", "s"}
		computerChoice := computerChoices[(rand.Intn(3-0) + 0)]
		fmt.Println("User played: ", *userChoice)
		fmt.Println("Computer played: ", computerChoice)
		fmt.Println("-------------")
		if (*userChoice == "r" && computerChoice == "s") ||
			(*userChoice == "p" && computerChoice == "r") ||
			(*userChoice == "s" && computerChoice == "p") {
			fmt.Println("You won!")
		} else if *userChoice == computerChoice {
			fmt.Println("It's a tie!")
		} else {
			fmt.Println("You LOSE")
		}
	} else {
		fmt.Println("No go. You need to use the flag --choice=r, p, or s.")
	}
}
