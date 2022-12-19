package main

import (
	"fmt"
	"math/rand"
	"time"
)

func gamblingSimulation(stake int, goal int, trials int) {
	totalWins := 0
	totalBets := 0

	for i := 0; i < trials; i++ {
		bets := 0
		currentStake := stake

		for currentStake > 0 && currentStake < goal {
			bets++
			rand.Seed(time.Now().UnixNano())
			if rand.Intn(2) == 1 {
				currentStake++
			} else {
				currentStake--
			}
		}

		if currentStake == goal {
			totalWins++
		}

		totalBets += bets
	}

	percentWin := float64(totalWins) / float64(trials) * 100
	avgBets := float64(totalBets) / float64(trials)

	fmt.Println("Number of times won: ", totalWins)
	fmt.Println("Percent win: ", percentWin)
	fmt.Println("Average number of bets made: ", avgBets)
}

func main() {
	var stack, goal, trials int

	fmt.Print("Enter initiak Stack: ")
	fmt.Scanln(&stack)
	fmt.Print("Enter Goal: ")
	fmt.Scanln(&goal)
	fmt.Print("Enter Trials: ")
	fmt.Scanln(&trials)

	gamblingSimulation(stack, goal, trials)
}
