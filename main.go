package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	// Check if we have the correct number of arguments
	if len(os.Args) != 3 {
		fmt.Println("Usage: diceroller <number of dice> <number of faces>")
		fmt.Println("Example: diceroller 3 6")
		os.Exit(1)
	}

	// Parse number of dice
	numDice, err := strconv.Atoi(os.Args[1])
	if err != nil || numDice < 1 {
		fmt.Println("Error: Number of dice must be a positive integer")
		os.Exit(1)
	}

	// Parse number of faces
	faces, err := strconv.Atoi(os.Args[2])
	if err != nil || faces < 2 {
		fmt.Println("Error: Number of faces must be an integer greater than 1")
		os.Exit(1)
	}

	// Initialize random seed
	rand.Seed(time.Now().UnixNano())

	// Roll the dice and keep track of sum
	sum := 0
	fmt.Printf("Rolling %d dice with %d faces each:\n", numDice, faces)
	
	for i := 1; i <= numDice; i++ {
		roll := rand.Intn(faces) + 1 // Generate random number between 1 and faces
		sum += roll
		fmt.Printf("Die %d: %d\n", i, roll)
	}

	// Print summary
	fmt.Printf("\nTotal number of dice rolled: %d\n", numDice)
	fmt.Printf("Sum of all dice: %d\n", sum)
}