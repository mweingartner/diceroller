package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func getValidInput(prompt string, minValue int) int {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input. Please try again.")
			continue
		}

		// Trim whitespace and convert to integer
		input = strings.TrimSpace(input)
		num, err := strconv.Atoi(input)
		if err != nil || num < minValue {
			fmt.Printf("Please enter a valid number (%d or greater)\n", minValue)
			continue
		}

		return num
	}
}

func diceroller() {
	// Get number of dice from user
	numDice := getValidInput("Enter the number of dice to roll: ", 1)

	// Get number of faces from user
	faces := getValidInput("Enter the number of faces on each die: ", 2)

	// Initialize random seed
	rand.Seed(time.Now().UnixNano())

	// Roll the dice and keep track of sum
	sum := 0
	fmt.Printf("\nRolling %d dice with %d faces each:\n", numDice, faces)
	
	for i := 1; i <= numDice; i++ {
		roll := rand.Intn(faces) + 1 // Generate random number between 1 and faces
		sum += roll
		fmt.Printf("Die %d: %d\n", i, roll)
	}

	// Print summary
	fmt.Printf("\nTotal number of dice rolled: %d\n", numDice)
	fmt.Printf("Sum of all dice: %d\n", sum)
}

func main() {
	diceroller()
}