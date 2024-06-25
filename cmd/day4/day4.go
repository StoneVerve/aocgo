package main

import (
	"aocgo/internal/readers"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var text []string
	text, _ = readers.ReadFile(os.Args[1])
	fmt.Println(countScratchCards(text))
}

func countScratchCards(text []string) int {
	amountCards := len(text)
	totalCards := 0
	cardsRepetitions := initializeSliceOnes(amountCards)
	for cardNumber, card := range text {
		copies := countWinningNumbers(card)
		for i := 1; i <= copies; i++ {
			if cardNumber+i < amountCards {
				cardsRepetitions[cardNumber+i] += cardsRepetitions[cardNumber]
			}
		}
		totalCards += cardsRepetitions[cardNumber]
	}
	return totalCards
}

/*
 * Initializes a new slice of integers with default value 1
 */
func initializeSliceOnes(size int) []int {
	slice := make([]int, size)
	for i := range slice {
		slice[i] = 1
	}
	return slice
}

func countWinningNumbers(text string) int {
	parts := strings.Split(text, "|")
	amountWinningNumbers := 0
	winningNumbers := getWinningNumbers(parts[0])
	cardNumbers := getNumbers(parts[1])
	for _, number := range cardNumbers {
		if _, isInWinningNumbers := winningNumbers[number]; isInWinningNumbers {
			amountWinningNumbers += 1
		}
	}
	return amountWinningNumbers
}

func getWinningNumbers(text string) map[int]bool {
	winningNumbers := make(map[int]bool)
	card := strings.Split(text, ":")
	numbers := getNumbers(card[1])
	for _, number := range numbers {
		winningNumbers[number] = true
	}
	return winningNumbers
}

func getNumbers(text string) []int {
	textNumbers := strings.Fields(strings.TrimSpace(text))
	numbers := make([]int, len(textNumbers))
	for i, number := range textNumbers {
		num, err := strconv.Atoi(strings.TrimSpace(number))
		if err != nil {
			fmt.Println("there was an error converting ", number, " to int")
		} else {
			numbers[i] = num
		}
	}
	return numbers
}
