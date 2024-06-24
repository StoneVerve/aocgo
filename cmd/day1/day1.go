package main

import (
	"aocgo/internal/readers"
	"errors"
	"fmt"
	"os"
)

/*
 * A map with names of all decimal digits from 1 to 9 and their corresponding values
 */
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

/*
 * Given an ASCII text, finds sum of all the calibrations
 */
func getSumCalibrations(text []string) int {
	sum := 0
	for _, line := range text {
		calibration, err := getCalibration(line)
		if err != nil {
			fmt.Println(err)
		} else {
			sum += calibration
		}
	}
	return sum
}

/*
 * Given an ASCII text line, finds the corresponding calibration
 * Returns an error if there is no calibration value in the text line
 */
func getCalibration(line string) (int, error) {
	leftDigit, errLeft := getLeftDigit([]byte(line))
	if errLeft != nil {
		return 0, fmt.Errorf("There is an error with line: \"%s\" \n"+
			"Error: %w", line, errLeft)
	}
	rightDigit, errRight := getRightDigit([]byte(line))
	if errRight != nil {
		return 0, fmt.Errorf("There is an error with line: \"%s\" \n"+
			"Error: %w", line, errRight)
	}
	return (leftDigit * 10) + rightDigit, nil
}

/*
 * Given a text line in ASCII, finds the first decimal digit from left to right
 * Returns an error if the line does not contain any decimal digits
 */
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
	return 0, errors.New("the line does no contain any digit")
}

/*
 * Given a text line in ASCII, finds the first decimal digit from right to left
 * Returns an error if the line does not contain any decimal digits
 */
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
	return 0, errors.New("the line does no contain any digit")
}

/*
 * Checks if a slice of bytes encodes the name of decimal digit (1-9) in ASCII
 */
func isWordDigit(digit []byte) bool {
	word := string(digit)
	_, ok := digits[word]
	if ok {
		return true
	}
	return false
}

/*
 * Converts a slice of bytes (corresponding to name of a digit 1-9) from its ASCII encoding to the corresponding ASCII digit
 * Returns an error if the slice of bytes does not correspond to the name of a digit 1-9 in ASCII encoding
 */
func convertWord(digit []byte) (int, error) {
	if !isWordDigit(digit) {
		return 0, errors.New("the word is not the name of a single digit number")
	}
	val, _ := digits[string(digit)]
	return val, nil
}

/*
 * Checks if a byte encodes a digit in ASCII
 */
func isDigit(digit byte) bool {
	if digit >= 48 && digit <= 57 {
		return true
	}
	return false
}

/*
 * Converts a byte from its ASCII encoding to the corresponding ASCII digit
 * Returns an error if the byte does not correspond to a digit 1-9 in ASCII encoding
 */
func convertByte(digit byte) (int, error) {
	if isDigit(digit) {
		return int(digit - 48), nil
	}
	return 0, errors.New("the byte is not a single digit number in ascii")
}
