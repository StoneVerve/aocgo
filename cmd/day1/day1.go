package main

import (
	"aocgo/internal/readers"
	"errors"
	"fmt"
	"os"
)

var digits = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func main() {
	var text []string
	text, _ = readers.ReadFile(os.Args[1])
	fmt.Println(getSumCalibrations(text))
}

func getSumCalibrations(text []string) int {
	sum := 0
	for _, line := range text {
		sum = sum + getCalibration(line)
	}
	return sum
}

func getCalibration(line string) int {
	leftDigit, errLeft := getLeftDigit([]byte(line))
	rightDigit, errRight := getRightDigit([]byte(line))
	if errLeft != nil || errRight != nil {
		fmt.Println("There was en error with line")
	}
	return (leftDigit * 10) + rightDigit
}

func getLeftDigit(line []byte) (int, error) {
	number := 0
	lineLength := len(line)
	for i := 0; i < lineLength; i++ {
		if isDigit(line[i]) {
			number, _ = convertByte(line[i])
			return number, nil
		} else if i+3 < lineLength && isWordDigit(line[i:i+3]) {
			number, _ = convertWord(line[i : i+3])
			return number, nil
		} else if i+4 < lineLength && isWordDigit(line[i:i+4]) {
			number, _ = convertWord(line[i : i+4])
			return number, nil
		} else if i+5 < lineLength && isWordDigit(line[i:i+5]) {
			number, _ = convertWord(line[i : i+5])
			return number, nil
		}
	}
	return 0, errors.New("The line does no contain any digit")
}

func getRightDigit(line []byte) (int, error) {
	number := 0
	lineLength := len(line)
	for i := lineLength; i > 0; i-- {
		if isDigit(line[i-1]) {
			number, _ = convertByte(line[i-1])
			return number, nil
		} else if i-3 >= 0 && isWordDigit(line[i-3:i]) {
			number, _ = convertWord(line[i-3 : i])
			return number, nil
		} else if i-4 >= 0 && isWordDigit(line[i-4:i]) {
			number, _ = convertWord(line[i-4 : i])
			return number, nil
		} else if i-5 >= 0 && isWordDigit(line[i-5:i]) {
			number, _ = convertWord(line[i-5 : i])
			return number, nil
		}
	}
	return 0, errors.New("The line does no contain any digit")
}

func isDigit(digit byte) bool {
	if digit >= 48 && digit <= 57 {
		return true
	}
	return false
}

func convertWord(digit []byte) (int, error) {
	if !isWordDigit(digit) {
		return 0, errors.New("The word is not a digit")
	}
	val, _ := digits[string(digit)]
	return val, nil
}

func isWordDigit(digit []byte) bool {
	word := string(digit)
	_, ok := digits[word]
	if ok {
		return true
	}
	return false
}

func convertByte(digit byte) (int, error) {
	if isDigit(digit) {
		return int(digit - 48), nil
	}
	return 0, errors.New("Not a number")
}
